// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aliyun/wuying-agentbay-sdk/golang/pkg/agentbay/interface (interfaces: SessionInterface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	client "github.com/aliyun/wuying-agentbay-sdk/golang/api/client"
	agentbay "github.com/aliyun/wuying-agentbay-sdk/golang/pkg/agentbay"
	gomock "github.com/golang/mock/gomock"
)

// MockSessionInterface is a mock of SessionInterface interface.
type MockSessionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSessionInterfaceMockRecorder
}

// MockSessionInterfaceMockRecorder is the mock recorder for MockSessionInterface.
type MockSessionInterfaceMockRecorder struct {
	mock *MockSessionInterface
}

// NewMockSessionInterface creates a new mock instance.
func NewMockSessionInterface(ctrl *gomock.Controller) *MockSessionInterface {
	mock := &MockSessionInterface{ctrl: ctrl}
	mock.recorder = &MockSessionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionInterface) EXPECT() *MockSessionInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSessionInterface) Delete() (*agentbay.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(*agentbay.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSessionInterfaceMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionInterface)(nil).Delete))
}

// GetAPIKey mocks base method.
func (m *MockSessionInterface) GetAPIKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAPIKey indicates an expected call of GetAPIKey.
func (mr *MockSessionInterfaceMockRecorder) GetAPIKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKey", reflect.TypeOf((*MockSessionInterface)(nil).GetAPIKey))
}

// GetClient mocks base method.
func (m *MockSessionInterface) GetClient() *client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(*client.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockSessionInterfaceMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockSessionInterface)(nil).GetClient))
}

// GetLabels mocks base method.
func (m *MockSessionInterface) GetLabels() (*agentbay.LabelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(*agentbay.LabelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockSessionInterfaceMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockSessionInterface)(nil).GetLabels))
}

// GetLink mocks base method.
func (m *MockSessionInterface) GetLink(arg0 *string, arg1 *int32) (*agentbay.LinkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", arg0, arg1)
	ret0, _ := ret[0].(*agentbay.LinkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLink indicates an expected call of GetLink.
func (mr *MockSessionInterfaceMockRecorder) GetLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*MockSessionInterface)(nil).GetLink), arg0, arg1)
}

// GetSessionId mocks base method.
func (m *MockSessionInterface) GetSessionId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSessionId indicates an expected call of GetSessionId.
func (mr *MockSessionInterfaceMockRecorder) GetSessionId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionId", reflect.TypeOf((*MockSessionInterface)(nil).GetSessionId))
}

// Info mocks base method.
func (m *MockSessionInterface) Info() (*agentbay.InfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(*agentbay.InfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockSessionInterfaceMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockSessionInterface)(nil).Info))
}

// SetLabels mocks base method.
func (m *MockSessionInterface) SetLabels(arg0 string) (*agentbay.LabelResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLabels", arg0)
	ret0, _ := ret[0].(*agentbay.LabelResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockSessionInterfaceMockRecorder) SetLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockSessionInterface)(nil).SetLabels), arg0)
}
