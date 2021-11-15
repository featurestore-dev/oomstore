// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/jinghan/source/jinghancc/oom-ai/oomstore/internal/database/metadata/store.go

// Package mock_metadata is a generated GoMock package.
package mock_metadata

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "github.com/oom-ai/oomstore/internal/database/metadata"
	types "github.com/oom-ai/oomstore/pkg/oomstore/types"
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

// CreateEntity mocks base method.
func (m *MockStore) CreateEntity(ctx context.Context, opt metadata.CreateEntityOpt) (int16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntity", ctx, opt)
	ret0, _ := ret[0].(int16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntity indicates an expected call of CreateEntity.
func (mr *MockStoreMockRecorder) CreateEntity(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntity", reflect.TypeOf((*MockStore)(nil).CreateEntity), ctx, opt)
}

// CreateFeature mocks base method.
func (m *MockStore) CreateFeature(ctx context.Context, opt metadata.CreateFeatureOpt) (int16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeature", ctx, opt)
	ret0, _ := ret[0].(int16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeature indicates an expected call of CreateFeature.
func (mr *MockStoreMockRecorder) CreateFeature(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeature", reflect.TypeOf((*MockStore)(nil).CreateFeature), ctx, opt)
}

// CreateFeatureGroup mocks base method.
func (m *MockStore) CreateFeatureGroup(ctx context.Context, opt metadata.CreateFeatureGroupOpt) (int16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeatureGroup", ctx, opt)
	ret0, _ := ret[0].(int16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeatureGroup indicates an expected call of CreateFeatureGroup.
func (mr *MockStoreMockRecorder) CreateFeatureGroup(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeatureGroup", reflect.TypeOf((*MockStore)(nil).CreateFeatureGroup), ctx, opt)
}

// CreateRevision mocks base method.
func (m *MockStore) CreateRevision(ctx context.Context, opt metadata.CreateRevisionOpt) (int32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRevision", ctx, opt)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRevision indicates an expected call of CreateRevision.
func (mr *MockStoreMockRecorder) CreateRevision(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRevision", reflect.TypeOf((*MockStore)(nil).CreateRevision), ctx, opt)
}

// GetEntity mocks base method.
func (m *MockStore) GetEntity(ctx context.Context, id int16) (*types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntity", ctx, id)
	ret0, _ := ret[0].(*types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntity indicates an expected call of GetEntity.
func (mr *MockStoreMockRecorder) GetEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntity", reflect.TypeOf((*MockStore)(nil).GetEntity), ctx, id)
}

// GetEntityByName mocks base method.
func (m *MockStore) GetEntityByName(ctx context.Context, name string) (*types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntityByName", ctx, name)
	ret0, _ := ret[0].(*types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntityByName indicates an expected call of GetEntityByName.
func (mr *MockStoreMockRecorder) GetEntityByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntityByName", reflect.TypeOf((*MockStore)(nil).GetEntityByName), ctx, name)
}

// GetFeature mocks base method.
func (m *MockStore) GetFeature(ctx context.Context, id int16) (*types.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeature", ctx, id)
	ret0, _ := ret[0].(*types.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeature indicates an expected call of GetFeature.
func (mr *MockStoreMockRecorder) GetFeature(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeature", reflect.TypeOf((*MockStore)(nil).GetFeature), ctx, id)
}

// GetFeatureByName mocks base method.
func (m *MockStore) GetFeatureByName(ctx context.Context, name string) (*types.Feature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureByName", ctx, name)
	ret0, _ := ret[0].(*types.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureByName indicates an expected call of GetFeatureByName.
func (mr *MockStoreMockRecorder) GetFeatureByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureByName", reflect.TypeOf((*MockStore)(nil).GetFeatureByName), ctx, name)
}

// GetFeatureGroup mocks base method.
func (m *MockStore) GetFeatureGroup(ctx context.Context, id int16) (*types.FeatureGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureGroup", ctx, id)
	ret0, _ := ret[0].(*types.FeatureGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureGroup indicates an expected call of GetFeatureGroup.
func (mr *MockStoreMockRecorder) GetFeatureGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureGroup", reflect.TypeOf((*MockStore)(nil).GetFeatureGroup), ctx, id)
}

// GetFeatureGroupByName mocks base method.
func (m *MockStore) GetFeatureGroupByName(ctx context.Context, name string) (*types.FeatureGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeatureGroupByName", ctx, name)
	ret0, _ := ret[0].(*types.FeatureGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeatureGroupByName indicates an expected call of GetFeatureGroupByName.
func (mr *MockStoreMockRecorder) GetFeatureGroupByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeatureGroupByName", reflect.TypeOf((*MockStore)(nil).GetFeatureGroupByName), ctx, name)
}

// GetRevision mocks base method.
func (m *MockStore) GetRevision(ctx context.Context, id int32) (*types.Revision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevision", ctx, id)
	ret0, _ := ret[0].(*types.Revision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevision indicates an expected call of GetRevision.
func (mr *MockStoreMockRecorder) GetRevision(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevision", reflect.TypeOf((*MockStore)(nil).GetRevision), ctx, id)
}

// GetRevisionBy mocks base method.
func (m *MockStore) GetRevisionBy(ctx context.Context, groupID int16, revision int64) (*types.Revision, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRevisionBy", ctx, groupID, revision)
	ret0, _ := ret[0].(*types.Revision)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRevisionBy indicates an expected call of GetRevisionBy.
func (mr *MockStoreMockRecorder) GetRevisionBy(ctx, groupID, revision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRevisionBy", reflect.TypeOf((*MockStore)(nil).GetRevisionBy), ctx, groupID, revision)
}

// ListEntity mocks base method.
func (m *MockStore) ListEntity(ctx context.Context) types.EntityList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntity", ctx)
	ret0, _ := ret[0].(types.EntityList)
	return ret0
}

// ListEntity indicates an expected call of ListEntity.
func (mr *MockStoreMockRecorder) ListEntity(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntity", reflect.TypeOf((*MockStore)(nil).ListEntity), ctx)
}

// ListFeature mocks base method.
func (m *MockStore) ListFeature(ctx context.Context, opt metadata.ListFeatureOpt) types.FeatureList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeature", ctx, opt)
	ret0, _ := ret[0].(types.FeatureList)
	return ret0
}

// ListFeature indicates an expected call of ListFeature.
func (mr *MockStoreMockRecorder) ListFeature(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeature", reflect.TypeOf((*MockStore)(nil).ListFeature), ctx, opt)
}

// ListFeatureGroup mocks base method.
func (m *MockStore) ListFeatureGroup(ctx context.Context, entityID *int16) types.FeatureGroupList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFeatureGroup", ctx, entityID)
	ret0, _ := ret[0].(types.FeatureGroupList)
	return ret0
}

// ListFeatureGroup indicates an expected call of ListFeatureGroup.
func (mr *MockStoreMockRecorder) ListFeatureGroup(ctx, entityID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeatureGroup", reflect.TypeOf((*MockStore)(nil).ListFeatureGroup), ctx, entityID)
}

// ListRevision mocks base method.
func (m *MockStore) ListRevision(ctx context.Context, opt metadata.ListRevisionOpt) types.RevisionList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRevision", ctx, opt)
	ret0, _ := ret[0].(types.RevisionList)
	return ret0
}

// ListRevision indicates an expected call of ListRevision.
func (mr *MockStoreMockRecorder) ListRevision(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRevision", reflect.TypeOf((*MockStore)(nil).ListRevision), ctx, opt)
}

// Refresh mocks base method.
func (m *MockStore) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh.
func (mr *MockStoreMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockStore)(nil).Refresh))
}

// UpdateEntity mocks base method.
func (m *MockStore) UpdateEntity(ctx context.Context, opt metadata.UpdateEntityOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntity", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntity indicates an expected call of UpdateEntity.
func (mr *MockStoreMockRecorder) UpdateEntity(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntity", reflect.TypeOf((*MockStore)(nil).UpdateEntity), ctx, opt)
}

// UpdateFeature mocks base method.
func (m *MockStore) UpdateFeature(ctx context.Context, opt metadata.UpdateFeatureOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeature", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeature indicates an expected call of UpdateFeature.
func (mr *MockStoreMockRecorder) UpdateFeature(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeature", reflect.TypeOf((*MockStore)(nil).UpdateFeature), ctx, opt)
}

// UpdateFeatureGroup mocks base method.
func (m *MockStore) UpdateFeatureGroup(ctx context.Context, opt metadata.UpdateFeatureGroupOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeatureGroup", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeatureGroup indicates an expected call of UpdateFeatureGroup.
func (mr *MockStoreMockRecorder) UpdateFeatureGroup(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeatureGroup", reflect.TypeOf((*MockStore)(nil).UpdateFeatureGroup), ctx, opt)
}

// UpdateRevision mocks base method.
func (m *MockStore) UpdateRevision(ctx context.Context, opt metadata.UpdateRevisionOpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRevision", ctx, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRevision indicates an expected call of UpdateRevision.
func (mr *MockStoreMockRecorder) UpdateRevision(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRevision", reflect.TypeOf((*MockStore)(nil).UpdateRevision), ctx, opt)
}
