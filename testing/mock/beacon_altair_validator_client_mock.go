// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Kevionte/prysm_beacon/proto/prysm/v1alpha1 (interfaces: BeaconNodeValidatorAltair_StreamBlocksClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v2 "github.com/Kevionte/prysm_beacon/v5/proto/prysm/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// BeaconNodeValidatorAltair_StreamBlocksClient is a mock of BeaconNodeValidatorAltair_StreamBlocksClient interface
type BeaconNodeValidatorAltair_StreamBlocksClient struct {
	ctrl     *gomock.Controller
	recorder *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder
}

// BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder is the mock recorder for MockBeaconNodeValidatorAltair_StreamBlocksClient
type BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder struct {
	mock *BeaconNodeValidatorAltair_StreamBlocksClient
}

// NewMockBeaconNodeValidatorAltair_StreamBlocksClient creates a new mock instance
func NewMockBeaconNodeValidatorAltair_StreamBlocksClient(ctrl *gomock.Controller) *BeaconNodeValidatorAltair_StreamBlocksClient {
	mock := &BeaconNodeValidatorAltair_StreamBlocksClient{ctrl: ctrl}
	mock.recorder = &BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) EXPECT() *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).CloseSend))
}

// Context mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).Context))
}

// Header mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).Header))
}

// Recv mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) Recv() (*v2.StreamBlocksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v2.StreamBlocksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *BeaconNodeValidatorAltair_StreamBlocksClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *BeaconNodeValidatorAltair_StreamBlocksClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*BeaconNodeValidatorAltair_StreamBlocksClient)(nil).Trailer))
}
