// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/utils/queryparams/queryparams.go

// Package mock_queryparams is a generated GoMock package.
package mock_queryparams

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPageParams is a mock of PageParams interface
type MockPageParams struct {
	ctrl     *gomock.Controller
	recorder *MockPageParamsMockRecorder
}

// MockPageParamsMockRecorder is the mock recorder for MockPageParams
type MockPageParamsMockRecorder struct {
	mock *MockPageParams
}

// NewMockPageParams creates a new mock instance
func NewMockPageParams(ctrl *gomock.Controller) *MockPageParams {
	mock := &MockPageParams{ctrl: ctrl}
	mock.recorder = &MockPageParamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPageParams) EXPECT() *MockPageParamsMockRecorder {
	return m.recorder
}

// Page mocks base method
func (m *MockPageParams) Page() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Page")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Page indicates an expected call of Page
func (mr *MockPageParamsMockRecorder) Page() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Page", reflect.TypeOf((*MockPageParams)(nil).Page))
}

// PerPage mocks base method
func (m *MockPageParams) PerPage() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerPage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerPage indicates an expected call of PerPage
func (mr *MockPageParamsMockRecorder) PerPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerPage", reflect.TypeOf((*MockPageParams)(nil).PerPage))
}