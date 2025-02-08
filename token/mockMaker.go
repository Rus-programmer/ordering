// Code generated by MockGen. DO NOT EDIT.
// Source: ordering/token (interfaces: Maker)
//
// Generated by this command:
//
//	mockgen -package token -destination token/mockMaker.go ordering/token Maker
//

// Package token is a generated GoMock package.
package token

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockMaker is a mock of Maker interface.
type MockMaker struct {
	ctrl     *gomock.Controller
	recorder *MockMakerMockRecorder
	isgomock struct{}
}

// MockMakerMockRecorder is the mock recorder for MockMaker.
type MockMakerMockRecorder struct {
	mock *MockMaker
}

// NewMockMaker creates a new mock instance.
func NewMockMaker(ctrl *gomock.Controller) *MockMaker {
	mock := &MockMaker{ctrl: ctrl}
	mock.recorder = &MockMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaker) EXPECT() *MockMakerMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MockMaker) CreateToken(customerID int64, role string, duration time.Duration) (string, *Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", customerID, role, duration)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*Payload)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockMakerMockRecorder) CreateToken(customerID, role, duration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockMaker)(nil).CreateToken), customerID, role, duration)
}

// VerifyToken mocks base method.
func (m *MockMaker) VerifyToken(token string) (*Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", token)
	ret0, _ := ret[0].(*Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockMakerMockRecorder) VerifyToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockMaker)(nil).VerifyToken), token)
}
