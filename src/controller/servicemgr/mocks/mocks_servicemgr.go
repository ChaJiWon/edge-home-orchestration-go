// Code generated by MockGen. DO NOT EDIT.
// Source: ./servicemgr.go

// Package mocks is a generated GoMock package.
package mocks

import (
	executor "controller/servicemgr/executor"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	client "restinterface/client"
)

// MockServiceMgr is a mock of ServiceMgr interface
type MockServiceMgr struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMgrMockRecorder
}

// MockServiceMgrMockRecorder is the mock recorder for MockServiceMgr
type MockServiceMgrMockRecorder struct {
	mock *MockServiceMgr
}

// NewMockServiceMgr creates a new mock instance
func NewMockServiceMgr(ctrl *gomock.Controller) *MockServiceMgr {
	mock := &MockServiceMgr{ctrl: ctrl}
	mock.recorder = &MockServiceMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceMgr) EXPECT() *MockServiceMgrMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockServiceMgr) Execute(target, name string, args []interface{}, notiChan chan string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", target, name, args, notiChan)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockServiceMgrMockRecorder) Execute(target, name, args, notiChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockServiceMgr)(nil).Execute), target, name, args, notiChan)
}

// SetLocalServiceExecutor mocks base method
func (m *MockServiceMgr) SetLocalServiceExecutor(s executor.ServiceExecutor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLocalServiceExecutor", s)
}

// SetLocalServiceExecutor indicates an expected call of SetLocalServiceExecutor
func (mr *MockServiceMgrMockRecorder) SetLocalServiceExecutor(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLocalServiceExecutor", reflect.TypeOf((*MockServiceMgr)(nil).SetLocalServiceExecutor), s)
}

// ExecuteAppOnLocal mocks base method
func (m *MockServiceMgr) ExecuteAppOnLocal(appInfo map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAppOnLocal", appInfo)
}

// ExecuteAppOnLocal indicates an expected call of ExecuteAppOnLocal
func (mr *MockServiceMgrMockRecorder) ExecuteAppOnLocal(appInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAppOnLocal", reflect.TypeOf((*MockServiceMgr)(nil).ExecuteAppOnLocal), appInfo)
}

// SetClient mocks base method
func (m *MockServiceMgr) SetClient(clientAPI client.Clienter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", clientAPI)
}

// SetClient indicates an expected call of SetClient
func (mr *MockServiceMgrMockRecorder) SetClient(clientAPI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockServiceMgr)(nil).SetClient), clientAPI)
}
