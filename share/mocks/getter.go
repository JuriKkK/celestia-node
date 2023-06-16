// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/share (interfaces: Getter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	da "github.com/celestiaorg/celestia-app/pkg/da"
	share "github.com/celestiaorg/celestia-node/share"
	namespace "github.com/celestiaorg/nmt/namespace"
	rsmt2d "github.com/celestiaorg/rsmt2d"
	gomock "github.com/golang/mock/gomock"
)

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// GetEDS mocks base method.
func (m *MockGetter) GetEDS(arg0 context.Context, arg1 *da.DataAvailabilityHeader) (*rsmt2d.ExtendedDataSquare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEDS", arg0, arg1)
	ret0, _ := ret[0].(*rsmt2d.ExtendedDataSquare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEDS indicates an expected call of GetEDS.
func (mr *MockGetterMockRecorder) GetEDS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEDS", reflect.TypeOf((*MockGetter)(nil).GetEDS), arg0, arg1)
}

// GetShare mocks base method.
func (m *MockGetter) GetShare(arg0 context.Context, arg1 *da.DataAvailabilityHeader, arg2, arg3 int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShare", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShare indicates an expected call of GetShare.
func (mr *MockGetterMockRecorder) GetShare(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShare", reflect.TypeOf((*MockGetter)(nil).GetShare), arg0, arg1, arg2, arg3)
}

// GetSharesByNamespace mocks base method.
func (m *MockGetter) GetSharesByNamespace(arg0 context.Context, arg1 *da.DataAvailabilityHeader, arg2 namespace.ID) (share.NamespacedShares, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharesByNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].(share.NamespacedShares)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharesByNamespace indicates an expected call of GetSharesByNamespace.
func (mr *MockGetterMockRecorder) GetSharesByNamespace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharesByNamespace", reflect.TypeOf((*MockGetter)(nil).GetSharesByNamespace), arg0, arg1, arg2)
}
