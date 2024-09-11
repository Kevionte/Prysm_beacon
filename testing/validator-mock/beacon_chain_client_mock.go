// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Kevionte/prysm_beacon/v1/validator/client/iface (interfaces: BeaconChainClient)
//
// Generated by this command:
//
//	mockgen -package=validator_mock -destination=testing/validator-mock/beacon_chain_client_mock.go github.com/Kevionte/prysm_beacon/v1/validator/client/iface BeaconChainClient
//

// Package validator_mock is a generated GoMock package.
package validator_mock

import (
	context "context"
	reflect "reflect"

	eth "github.com/Kevionte/prysm_beacon/v1/proto/prysm/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockBeaconChainClient is a mock of BeaconChainClient interface.
type MockBeaconChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconChainClientMockRecorder
}

// MockBeaconChainClientMockRecorder is the mock recorder for MockBeaconChainClient.
type MockBeaconChainClientMockRecorder struct {
	mock *MockBeaconChainClient
}

// NewMockBeaconChainClient creates a new mock instance.
func NewMockBeaconChainClient(ctrl *gomock.Controller) *MockBeaconChainClient {
	mock := &MockBeaconChainClient{ctrl: ctrl}
	mock.recorder = &MockBeaconChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeaconChainClient) EXPECT() *MockBeaconChainClientMockRecorder {
	return m.recorder
}

// GetChainHead mocks base method.
func (m *MockBeaconChainClient) GetChainHead(arg0 context.Context, arg1 *emptypb.Empty) (*eth.ChainHead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainHead", arg0, arg1)
	ret0, _ := ret[0].(*eth.ChainHead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainHead indicates an expected call of GetChainHead.
func (mr *MockBeaconChainClientMockRecorder) GetChainHead(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainHead", reflect.TypeOf((*MockBeaconChainClient)(nil).GetChainHead), arg0, arg1)
}

// GetValidatorParticipation mocks base method.
func (m *MockBeaconChainClient) GetValidatorParticipation(arg0 context.Context, arg1 *eth.GetValidatorParticipationRequest) (*eth.ValidatorParticipationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorParticipation", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorParticipationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorParticipation indicates an expected call of GetValidatorParticipation.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorParticipation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorParticipation", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorParticipation), arg0, arg1)
}

// GetValidatorPerformance mocks base method.
func (m *MockBeaconChainClient) GetValidatorPerformance(arg0 context.Context, arg1 *eth.ValidatorPerformanceRequest) (*eth.ValidatorPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorPerformance", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorPerformance indicates an expected call of GetValidatorPerformance.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorPerformance(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorPerformance", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorPerformance), arg0, arg1)
}

// GetValidatorQueue mocks base method.
func (m *MockBeaconChainClient) GetValidatorQueue(arg0 context.Context, arg1 *emptypb.Empty) (*eth.ValidatorQueue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorQueue", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorQueue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorQueue indicates an expected call of GetValidatorQueue.
func (mr *MockBeaconChainClientMockRecorder) GetValidatorQueue(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorQueue", reflect.TypeOf((*MockBeaconChainClient)(nil).GetValidatorQueue), arg0, arg1)
}

// ListValidatorBalances mocks base method.
func (m *MockBeaconChainClient) ListValidatorBalances(arg0 context.Context, arg1 *eth.ListValidatorBalancesRequest) (*eth.ValidatorBalances, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListValidatorBalances", arg0, arg1)
	ret0, _ := ret[0].(*eth.ValidatorBalances)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValidatorBalances indicates an expected call of ListValidatorBalances.
func (mr *MockBeaconChainClientMockRecorder) ListValidatorBalances(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValidatorBalances", reflect.TypeOf((*MockBeaconChainClient)(nil).ListValidatorBalances), arg0, arg1)
}

// ListValidators mocks base method.
func (m *MockBeaconChainClient) ListValidators(arg0 context.Context, arg1 *eth.ListValidatorsRequest) (*eth.Validators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListValidators", arg0, arg1)
	ret0, _ := ret[0].(*eth.Validators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValidators indicates an expected call of ListValidators.
func (mr *MockBeaconChainClientMockRecorder) ListValidators(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValidators", reflect.TypeOf((*MockBeaconChainClient)(nil).ListValidators), arg0, arg1)
}
