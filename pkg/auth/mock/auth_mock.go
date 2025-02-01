// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/moc/pkg/auth (interfaces: Authorizer)

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	credentials "google.golang.org/grpc/credentials"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// WithRPCAuthorization mocks base method.
func (m *MockAuthorizer) WithRPCAuthorization() credentials.PerRPCCredentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithRPCAuthorization")
	ret0, _ := ret[0].(credentials.PerRPCCredentials)
	return ret0
}

// WithRPCAuthorization indicates an expected call of WithRPCAuthorization.
func (mr *MockAuthorizerMockRecorder) WithRPCAuthorization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithRPCAuthorization", reflect.TypeOf((*MockAuthorizer)(nil).WithRPCAuthorization))
}

// WithTransportAuthorization mocks base method.
func (m *MockAuthorizer) WithTransportAuthorization() credentials.TransportCredentials {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransportAuthorization")
	ret0, _ := ret[0].(credentials.TransportCredentials)
	return ret0
}

// WithTransportAuthorization indicates an expected call of WithTransportAuthorization.
func (mr *MockAuthorizerMockRecorder) WithTransportAuthorization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransportAuthorization", reflect.TypeOf((*MockAuthorizer)(nil).WithTransportAuthorization))
}
