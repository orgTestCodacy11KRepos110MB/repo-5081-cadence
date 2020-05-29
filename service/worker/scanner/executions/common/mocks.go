// The MIT License (MIT)
// 
// Copyright (c) 2017-2020 Uber Technologies Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/worker/scanner/executions/common/interfaces.go

// Package common is a generated GoMock package.
package common

import (
	gomock "github.com/golang/mock/gomock"
	persistence "github.com/uber/cadence/common/persistence"
	reflect "reflect"
)

// MockPersistenceRetryer is a mock of PersistenceRetryer interface
type MockPersistenceRetryer struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceRetryerMockRecorder
}

// MockPersistenceRetryerMockRecorder is the mock recorder for MockPersistenceRetryer
type MockPersistenceRetryerMockRecorder struct {
	mock *MockPersistenceRetryer
}

// NewMockPersistenceRetryer creates a new mock instance
func NewMockPersistenceRetryer(ctrl *gomock.Controller) *MockPersistenceRetryer {
	mock := &MockPersistenceRetryer{ctrl: ctrl}
	mock.recorder = &MockPersistenceRetryerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersistenceRetryer) EXPECT() *MockPersistenceRetryerMockRecorder {
	return m.recorder
}

// ListConcreteExecutions mocks base method
func (m *MockPersistenceRetryer) ListConcreteExecutions(arg0 *persistence.ListConcreteExecutionsRequest) (*persistence.ListConcreteExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConcreteExecutions", arg0)
	ret0, _ := ret[0].(*persistence.ListConcreteExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConcreteExecutions indicates an expected call of ListConcreteExecutions
func (mr *MockPersistenceRetryerMockRecorder) ListConcreteExecutions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConcreteExecutions", reflect.TypeOf((*MockPersistenceRetryer)(nil).ListConcreteExecutions), arg0)
}

// GetWorkflowExecution mocks base method
func (m *MockPersistenceRetryer) GetWorkflowExecution(arg0 *persistence.GetWorkflowExecutionRequest) (*persistence.GetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowExecution", arg0)
	ret0, _ := ret[0].(*persistence.GetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowExecution indicates an expected call of GetWorkflowExecution
func (mr *MockPersistenceRetryerMockRecorder) GetWorkflowExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowExecution", reflect.TypeOf((*MockPersistenceRetryer)(nil).GetWorkflowExecution), arg0)
}

// GetCurrentExecution mocks base method
func (m *MockPersistenceRetryer) GetCurrentExecution(arg0 *persistence.GetCurrentExecutionRequest) (*persistence.GetCurrentExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentExecution", arg0)
	ret0, _ := ret[0].(*persistence.GetCurrentExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentExecution indicates an expected call of GetCurrentExecution
func (mr *MockPersistenceRetryerMockRecorder) GetCurrentExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentExecution", reflect.TypeOf((*MockPersistenceRetryer)(nil).GetCurrentExecution), arg0)
}

// ReadHistoryBranch mocks base method
func (m *MockPersistenceRetryer) ReadHistoryBranch(arg0 *persistence.ReadHistoryBranchRequest) (*persistence.ReadHistoryBranchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHistoryBranch", arg0)
	ret0, _ := ret[0].(*persistence.ReadHistoryBranchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHistoryBranch indicates an expected call of ReadHistoryBranch
func (mr *MockPersistenceRetryerMockRecorder) ReadHistoryBranch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHistoryBranch", reflect.TypeOf((*MockPersistenceRetryer)(nil).ReadHistoryBranch), arg0)
}

// DeleteWorkflowExecution mocks base method
func (m *MockPersistenceRetryer) DeleteWorkflowExecution(arg0 *persistence.DeleteWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecution", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecution indicates an expected call of DeleteWorkflowExecution
func (mr *MockPersistenceRetryerMockRecorder) DeleteWorkflowExecution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecution", reflect.TypeOf((*MockPersistenceRetryer)(nil).DeleteWorkflowExecution), arg0)
}

// DeleteCurrentWorkflowExecution mocks base method
func (m *MockPersistenceRetryer) DeleteCurrentWorkflowExecution(request *persistence.DeleteCurrentWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurrentWorkflowExecution", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCurrentWorkflowExecution indicates an expected call of DeleteCurrentWorkflowExecution
func (mr *MockPersistenceRetryerMockRecorder) DeleteCurrentWorkflowExecution(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentWorkflowExecution", reflect.TypeOf((*MockPersistenceRetryer)(nil).DeleteCurrentWorkflowExecution), request)
}

// MockInvariantManager is a mock of InvariantManager interface
type MockInvariantManager struct {
	ctrl     *gomock.Controller
	recorder *MockInvariantManagerMockRecorder
}

// MockInvariantManagerMockRecorder is the mock recorder for MockInvariantManager
type MockInvariantManagerMockRecorder struct {
	mock *MockInvariantManager
}

// NewMockInvariantManager creates a new mock instance
func NewMockInvariantManager(ctrl *gomock.Controller) *MockInvariantManager {
	mock := &MockInvariantManager{ctrl: ctrl}
	mock.recorder = &MockInvariantManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvariantManager) EXPECT() *MockInvariantManagerMockRecorder {
	return m.recorder
}

// RunChecks mocks base method
func (m *MockInvariantManager) RunChecks(arg0 Execution) ManagerCheckResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChecks", arg0)
	ret0, _ := ret[0].(ManagerCheckResult)
	return ret0
}

// RunChecks indicates an expected call of RunChecks
func (mr *MockInvariantManagerMockRecorder) RunChecks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChecks", reflect.TypeOf((*MockInvariantManager)(nil).RunChecks), arg0)
}

// RunFixes mocks base method
func (m *MockInvariantManager) RunFixes(arg0 Execution) ManagerFixResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunFixes", arg0)
	ret0, _ := ret[0].(ManagerFixResult)
	return ret0
}

// RunFixes indicates an expected call of RunFixes
func (mr *MockInvariantManagerMockRecorder) RunFixes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFixes", reflect.TypeOf((*MockInvariantManager)(nil).RunFixes), arg0)
}

// InvariantTypes mocks base method
func (m *MockInvariantManager) InvariantTypes() []InvariantType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvariantTypes")
	ret0, _ := ret[0].([]InvariantType)
	return ret0
}

// InvariantTypes indicates an expected call of InvariantTypes
func (mr *MockInvariantManagerMockRecorder) InvariantTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvariantTypes", reflect.TypeOf((*MockInvariantManager)(nil).InvariantTypes))
}

// MockInvariant is a mock of Invariant interface
type MockInvariant struct {
	ctrl     *gomock.Controller
	recorder *MockInvariantMockRecorder
}

// MockInvariantMockRecorder is the mock recorder for MockInvariant
type MockInvariantMockRecorder struct {
	mock *MockInvariant
}

// NewMockInvariant creates a new mock instance
func NewMockInvariant(ctrl *gomock.Controller) *MockInvariant {
	mock := &MockInvariant{ctrl: ctrl}
	mock.recorder = &MockInvariantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvariant) EXPECT() *MockInvariantMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockInvariant) Check(arg0 Execution, arg1 *InvariantResourceBag) CheckResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(CheckResult)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockInvariantMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockInvariant)(nil).Check), arg0, arg1)
}

// Fix mocks base method
func (m *MockInvariant) Fix(arg0 Execution, arg1 *InvariantResourceBag) FixResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fix", arg0, arg1)
	ret0, _ := ret[0].(FixResult)
	return ret0
}

// Fix indicates an expected call of Fix
func (mr *MockInvariantMockRecorder) Fix(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fix", reflect.TypeOf((*MockInvariant)(nil).Fix), arg0, arg1)
}

// InvariantType mocks base method
func (m *MockInvariant) InvariantType() InvariantType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvariantType")
	ret0, _ := ret[0].(InvariantType)
	return ret0
}

// InvariantType indicates an expected call of InvariantType
func (mr *MockInvariantMockRecorder) InvariantType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvariantType", reflect.TypeOf((*MockInvariant)(nil).InvariantType))
}

// MockExecutionIterator is a mock of ExecutionIterator interface
type MockExecutionIterator struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionIteratorMockRecorder
}

// MockExecutionIteratorMockRecorder is the mock recorder for MockExecutionIterator
type MockExecutionIteratorMockRecorder struct {
	mock *MockExecutionIterator
}

// NewMockExecutionIterator creates a new mock instance
func NewMockExecutionIterator(ctrl *gomock.Controller) *MockExecutionIterator {
	mock := &MockExecutionIterator{ctrl: ctrl}
	mock.recorder = &MockExecutionIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutionIterator) EXPECT() *MockExecutionIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockExecutionIterator) Next() (*Execution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*Execution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockExecutionIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockExecutionIterator)(nil).Next))
}

// HasNext mocks base method
func (m *MockExecutionIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext
func (mr *MockExecutionIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockExecutionIterator)(nil).HasNext))
}

// MockScanOutputIterator is a mock of ScanOutputIterator interface
type MockScanOutputIterator struct {
	ctrl     *gomock.Controller
	recorder *MockScanOutputIteratorMockRecorder
}

// MockScanOutputIteratorMockRecorder is the mock recorder for MockScanOutputIterator
type MockScanOutputIteratorMockRecorder struct {
	mock *MockScanOutputIterator
}

// NewMockScanOutputIterator creates a new mock instance
func NewMockScanOutputIterator(ctrl *gomock.Controller) *MockScanOutputIterator {
	mock := &MockScanOutputIterator{ctrl: ctrl}
	mock.recorder = &MockScanOutputIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScanOutputIterator) EXPECT() *MockScanOutputIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockScanOutputIterator) Next() (*ScanOutputEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*ScanOutputEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockScanOutputIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockScanOutputIterator)(nil).Next))
}

// HasNext mocks base method
func (m *MockScanOutputIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext
func (mr *MockScanOutputIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockScanOutputIterator)(nil).HasNext))
}

// MockExecutionWriter is a mock of ExecutionWriter interface
type MockExecutionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionWriterMockRecorder
}

// MockExecutionWriterMockRecorder is the mock recorder for MockExecutionWriter
type MockExecutionWriterMockRecorder struct {
	mock *MockExecutionWriter
}

// NewMockExecutionWriter creates a new mock instance
func NewMockExecutionWriter(ctrl *gomock.Controller) *MockExecutionWriter {
	mock := &MockExecutionWriter{ctrl: ctrl}
	mock.recorder = &MockExecutionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutionWriter) EXPECT() *MockExecutionWriterMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockExecutionWriter) Add(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockExecutionWriterMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockExecutionWriter)(nil).Add), arg0)
}

// Flush mocks base method
func (m *MockExecutionWriter) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockExecutionWriterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockExecutionWriter)(nil).Flush))
}

// FlushedKeys mocks base method
func (m *MockExecutionWriter) FlushedKeys() *Keys {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushedKeys")
	ret0, _ := ret[0].(*Keys)
	return ret0
}

// FlushedKeys indicates an expected call of FlushedKeys
func (mr *MockExecutionWriterMockRecorder) FlushedKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushedKeys", reflect.TypeOf((*MockExecutionWriter)(nil).FlushedKeys))
}

// MockScanner is a mock of Scanner interface
type MockScanner struct {
	ctrl     *gomock.Controller
	recorder *MockScannerMockRecorder
}

// MockScannerMockRecorder is the mock recorder for MockScanner
type MockScannerMockRecorder struct {
	mock *MockScanner
}

// NewMockScanner creates a new mock instance
func NewMockScanner(ctrl *gomock.Controller) *MockScanner {
	mock := &MockScanner{ctrl: ctrl}
	mock.recorder = &MockScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScanner) EXPECT() *MockScannerMockRecorder {
	return m.recorder
}

// Scan mocks base method
func (m *MockScanner) Scan() ShardScanReport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan")
	ret0, _ := ret[0].(ShardScanReport)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockScannerMockRecorder) Scan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockScanner)(nil).Scan))
}

// MockFixer is a mock of Fixer interface
type MockFixer struct {
	ctrl     *gomock.Controller
	recorder *MockFixerMockRecorder
}

// MockFixerMockRecorder is the mock recorder for MockFixer
type MockFixerMockRecorder struct {
	mock *MockFixer
}

// NewMockFixer creates a new mock instance
func NewMockFixer(ctrl *gomock.Controller) *MockFixer {
	mock := &MockFixer{ctrl: ctrl}
	mock.recorder = &MockFixerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFixer) EXPECT() *MockFixerMockRecorder {
	return m.recorder
}

// Fix mocks base method
func (m *MockFixer) Fix() ShardFixReport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fix")
	ret0, _ := ret[0].(ShardFixReport)
	return ret0
}

// Fix indicates an expected call of Fix
func (mr *MockFixerMockRecorder) Fix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fix", reflect.TypeOf((*MockFixer)(nil).Fix))
}