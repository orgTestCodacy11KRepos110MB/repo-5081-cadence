// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package matching

import (
	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/common/types"
)

type (
	// genericTaskInfo contains the info for an activity or decision task
	genericTaskInfo struct {
		*persistence.TaskInfo
		completionFunc func(*persistence.TaskInfo, error)
	}
	// queryTaskInfo contains the info for a query task
	queryTaskInfo struct {
		taskID  string
		request *types.MatchingQueryWorkflowRequest
	}
	// startedTaskInfo contains info for any task received from
	// another matching host. This type of task is already marked as started
	startedTaskInfo struct {
		decisionTaskInfo *types.MatchingPollForDecisionTaskResponse
		activityTaskInfo *types.PollForActivityTaskResponse
	}
	// InternalTask represents an activity, decision, query or started (received from another host).
	// this struct is more like a union and only one of [ query, event, forwarded ] is
	// non-nil for any given task
	InternalTask struct {
		event                    *genericTaskInfo // non-nil for activity or decision task that's locally generated
		query                    *queryTaskInfo   // non-nil for a query task that's locally sync matched
		started                  *startedTaskInfo // non-nil for a task received from a parent partition which is already started
		domainName               string
		source                   types.TaskSource
		forwardedFrom            string     // name of the child partition this task is forwarded from (empty if not forwarded)
		responseC                chan error // non-nil only where there is a caller waiting for response (sync-match)
		backlogCountHint         int64
		activityTaskDispatchInfo *types.ActivityTaskDispatchInfo
	}
)

func newInternalTask(
	info *persistence.TaskInfo,
	completionFunc func(*persistence.TaskInfo, error),
	source types.TaskSource,
	forwardedFrom string,
	forSyncMatch bool,
	activityTaskDispatchInfo *types.ActivityTaskDispatchInfo,
) *InternalTask {
	task := &InternalTask{
		event: &genericTaskInfo{
			TaskInfo:       info,
			completionFunc: completionFunc,
		},
		source:                   source,
		forwardedFrom:            forwardedFrom,
		activityTaskDispatchInfo: activityTaskDispatchInfo,
	}
	if forSyncMatch {
		task.responseC = make(chan error, 1)
	}
	return task
}

func newInternalQueryTask(
	taskID string,
	request *types.MatchingQueryWorkflowRequest,
) *InternalTask {
	return &InternalTask{
		query: &queryTaskInfo{
			taskID:  taskID,
			request: request,
		},
		forwardedFrom: request.GetForwardedFrom(),
		responseC:     make(chan error, 1),
	}
}

func newInternalStartedTask(info *startedTaskInfo) *InternalTask {
	return &InternalTask{started: info}
}

// isQuery returns true if the underlying task is a query task
func (task *InternalTask) isQuery() bool {
	return task.query != nil
}

// isStarted is true when this task is already marked as started
func (task *InternalTask) isStarted() bool {
	return task.started != nil
}

// isForwarded returns true if the underlying task is forwarded by a remote matching host
// forwarded tasks are already marked as started in history
func (task *InternalTask) isForwarded() bool {
	return task.forwardedFrom != ""
}

func (task *InternalTask) workflowExecution() *types.WorkflowExecution {
	switch {
	case task.event != nil:
		return &types.WorkflowExecution{WorkflowID: task.event.WorkflowID, RunID: task.event.RunID}
	case task.query != nil:
		return task.query.request.GetQueryRequest().GetExecution()
	case task.started != nil && task.started.decisionTaskInfo != nil:
		return task.started.decisionTaskInfo.WorkflowExecution
	case task.started != nil && task.started.activityTaskInfo != nil:
		return task.started.activityTaskInfo.WorkflowExecution
	}
	return &types.WorkflowExecution{}
}

// pollForDecisionResponse returns the poll response for a decision task that is
// already marked as started. This method should only be called when isStarted() is true
func (task *InternalTask) pollForDecisionResponse() *types.MatchingPollForDecisionTaskResponse {
	if task.isStarted() {
		return task.started.decisionTaskInfo
	}
	return nil
}

// pollForActivityResponse returns the poll response for an activity task that is
// already marked as started. This method should only be called when isStarted() is true
func (task *InternalTask) pollForActivityResponse() *types.PollForActivityTaskResponse {
	if task.isStarted() {
		return task.started.activityTaskInfo
	}
	return nil
}

// finish marks a task as finished. Should be called after a poller picks up a task
// and marks it as started. If the task is unable to marked as started, then this
// method should be called with a non-nil error argument.
func (task *InternalTask) finish(err error) {
	switch {
	case task.responseC != nil:
		task.responseC <- err
	case task.event.completionFunc != nil:
		task.event.completionFunc(task.event.TaskInfo, err)
	}
}
