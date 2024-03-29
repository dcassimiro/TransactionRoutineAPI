// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/account/account.go
//
// Generated by this command:
//
//	mockgen -source=./app/account/account.go -destination=./mocks/account_app_mock.go -package=mocks -mock_names=App=MockAccountApp
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/pismo/TransactionRoutineAPI/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountApp is a mock of App interface.
type MockAccountApp struct {
	ctrl     *gomock.Controller
	recorder *MockAccountAppMockRecorder
}

// MockAccountAppMockRecorder is the mock recorder for MockAccountApp.
type MockAccountAppMockRecorder struct {
	mock *MockAccountApp
}

// NewMockAccountApp creates a new mock instance.
func NewMockAccountApp(ctrl *gomock.Controller) *MockAccountApp {
	mock := &MockAccountApp{ctrl: ctrl}
	mock.recorder = &MockAccountAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountApp) EXPECT() *MockAccountAppMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountApp) Create(ctx context.Context, account model.AccountRequest) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, account)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountAppMockRecorder) Create(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountApp)(nil).Create), ctx, account)
}

// ReadOne mocks base method.
func (m *MockAccountApp) ReadOne(ctx context.Context, accountID string) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOne", ctx, accountID)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOne indicates an expected call of ReadOne.
func (mr *MockAccountAppMockRecorder) ReadOne(ctx, accountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOne", reflect.TypeOf((*MockAccountApp)(nil).ReadOne), ctx, accountID)
}
