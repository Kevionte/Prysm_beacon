// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/genesis.go
//
// Generated by this command:
//
//	mockgen -package=mock -source=validator/client/beacon-api/genesis.go -destination=validator/client/beacon-api/mock/genesis_mock.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	structs "github.com/Kevionte/prysm_beacon/v5/api/server/structs"
	gomock "go.uber.org/mock/gomock"
)

// MockGenesisProvider is a mock of GenesisProvider interface.
type MockGenesisProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGenesisProviderMockRecorder
}

// MockGenesisProviderMockRecorder is the mock recorder for MockGenesisProvider.
type MockGenesisProviderMockRecorder struct {
	mock *MockGenesisProvider
}

// NewMockGenesisProvider creates a new mock instance.
func NewMockGenesisProvider(ctrl *gomock.Controller) *MockGenesisProvider {
	mock := &MockGenesisProvider{ctrl: ctrl}
	mock.recorder = &MockGenesisProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenesisProvider) EXPECT() *MockGenesisProviderMockRecorder {
	return m.recorder
}

// GetGenesis mocks base method.
func (m *MockGenesisProvider) GetGenesis(ctx context.Context) (*structs.Genesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesis", ctx)
	ret0, _ := ret[0].(*structs.Genesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesis indicates an expected call of GetGenesis.
func (mr *MockGenesisProviderMockRecorder) GetGenesis(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockGenesisProvider)(nil).GetGenesis), ctx)
}
