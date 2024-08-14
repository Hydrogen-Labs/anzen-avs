// Code generated by MockGen. DO NOT EDIT.
// Source: anzen-avs/safety-factor (interfaces: SafetyFactorServicer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/safety_factor_servicer.go -package=mocks anzen-avs/safety-factor SafetyFactorServicer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	safety_factor_base "anzen-avs/safety-factor/safety-factor-base"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSafetyFactorServicer is a mock of SafetyFactorServicer interface.
type MockSafetyFactorServicer struct {
	ctrl     *gomock.Controller
	recorder *MockSafetyFactorServicerMockRecorder
}

// MockSafetyFactorServicerMockRecorder is the mock recorder for MockSafetyFactorServicer.
type MockSafetyFactorServicerMockRecorder struct {
	mock *MockSafetyFactorServicer
}

// NewMockSafetyFactorServicer creates a new mock instance.
func NewMockSafetyFactorServicer(ctrl *gomock.Controller) *MockSafetyFactorServicer {
	mock := &MockSafetyFactorServicer{ctrl: ctrl}
	mock.recorder = &MockSafetyFactorServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSafetyFactorServicer) EXPECT() *MockSafetyFactorServicerMockRecorder {
	return m.recorder
}

// GetModuleByOracleIndex mocks base method.
func (m *MockSafetyFactorServicer) GetModuleByOracleIndex(arg0 int) (safety_factor_base.SFModule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleByOracleIndex", arg0)
	ret0, _ := ret[0].(safety_factor_base.SFModule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModuleByOracleIndex indicates an expected call of GetModuleByOracleIndex.
func (mr *MockSafetyFactorServicerMockRecorder) GetModuleByOracleIndex(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleByOracleIndex", reflect.TypeOf((*MockSafetyFactorServicer)(nil).GetModuleByOracleIndex), arg0)
}

// GetSafetyFactorInfoByOracleIndex mocks base method.
func (m *MockSafetyFactorServicer) GetSafetyFactorInfoByOracleIndex(arg0 int) (*safety_factor_base.SFModuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSafetyFactorInfoByOracleIndex", arg0)
	ret0, _ := ret[0].(*safety_factor_base.SFModuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSafetyFactorInfoByOracleIndex indicates an expected call of GetSafetyFactorInfoByOracleIndex.
func (mr *MockSafetyFactorServicerMockRecorder) GetSafetyFactorInfoByOracleIndex(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSafetyFactorInfoByOracleIndex", reflect.TypeOf((*MockSafetyFactorServicer)(nil).GetSafetyFactorInfoByOracleIndex), arg0)
}

// IsSafetyFactorInfoStale mocks base method.
func (m *MockSafetyFactorServicer) IsSafetyFactorInfoStale(arg0 int32) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSafetyFactorInfoStale", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSafetyFactorInfoStale indicates an expected call of IsSafetyFactorInfoStale.
func (mr *MockSafetyFactorServicerMockRecorder) IsSafetyFactorInfoStale(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSafetyFactorInfoStale", reflect.TypeOf((*MockSafetyFactorServicer)(nil).IsSafetyFactorInfoStale), arg0)
}
