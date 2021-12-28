// Code generated by MockGen. DO NOT EDIT.
// Source: internal/database/online/store.go

// Package mock_online is a generated GoMock package.
package mock_online

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dbutil "github.com/oom-ai/oomstore/internal/database/dbutil"
	online "github.com/oom-ai/oomstore/internal/database/online"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// Get mocks base method.
func (m *MockStore) Get(ctx context.Context, opt online.GetOpt) (dbutil.RowMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, opt)
	ret0, _ := ret[0].(dbutil.RowMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, opt)
}

// Import mocks base method.
func (m *MockStore) Import(ctx context.Context, opt online.ImportOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Import indicates an expected call of Import.
func (mr *MockStoreMockRecorder) Import(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockStore)(nil).Import), ctx, opt)
}

// MultiGet mocks base method.
func (m *MockStore) MultiGet(ctx context.Context, opt online.MultiGetOpt) (map[string]dbutil.RowMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGet", ctx, opt)
	ret0, _ := ret[0].(map[string]dbutil.RowMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockStoreMockRecorder) MultiGet(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockStore)(nil).MultiGet), ctx, opt)
}

// Ping mocks base method.
func (m *MockStore) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockStoreMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockStore)(nil).Ping), ctx)
}

// Purge mocks base method.
func (m *MockStore) Purge(ctx context.Context, revisionID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Purge", ctx, revisionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Purge indicates an expected call of Purge.
func (mr *MockStoreMockRecorder) Purge(ctx, revisionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Purge", reflect.TypeOf((*MockStore)(nil).Purge), ctx, revisionID)
}

// Push mocks base method.
func (m *MockStore) Push(ctx context.Context, opt online.PushOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockStoreMockRecorder) Push(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockStore)(nil).Push), ctx, opt)
}
