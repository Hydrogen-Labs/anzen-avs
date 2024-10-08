// Code generated by MockGen. DO NOT EDIT.
// Source: anzen-avs/operator (interfaces: AggregatorRpcClienter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/rpc_client.go -package=mocks anzen-avs/operator AggregatorRpcClienter
//
// Package mocks is a generated GoMock package.
package mocks

import (
	aggregator "anzen-avs/aggregator"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAggregatorRpcClienter is a mock of AggregatorRpcClienter interface.
type MockAggregatorRpcClienter struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorRpcClienterMockRecorder
}

// MockAggregatorRpcClienterMockRecorder is the mock recorder for MockAggregatorRpcClienter.
type MockAggregatorRpcClienterMockRecorder struct {
	mock *MockAggregatorRpcClienter
}

// NewMockAggregatorRpcClienter creates a new mock instance.
func NewMockAggregatorRpcClienter(ctrl *gomock.Controller) *MockAggregatorRpcClienter {
	mock := &MockAggregatorRpcClienter{ctrl: ctrl}
	mock.recorder = &MockAggregatorRpcClienterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregatorRpcClienter) EXPECT() *MockAggregatorRpcClienterMockRecorder {
	return m.recorder
}

// SendSignedOraclePullTaskReponseToAggregator mocks base method.
func (m *MockAggregatorRpcClienter) SendSignedOraclePullTaskReponseToAggregator(arg0 *aggregator.SignedOraclePullTaskResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendSignedOraclePullTaskReponseToAggregator", arg0)
}

// SendSignedOraclePullTaskReponseToAggregator indicates an expected call of SendSignedOraclePullTaskReponseToAggregator.
func (mr *MockAggregatorRpcClienterMockRecorder) SendSignedOraclePullTaskReponseToAggregator(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSignedOraclePullTaskReponseToAggregator", reflect.TypeOf((*MockAggregatorRpcClienter)(nil).SendSignedOraclePullTaskReponseToAggregator), arg0)
}
