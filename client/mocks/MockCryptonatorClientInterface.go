// Code generated by MockGen. DO NOT EDIT.
// Source: ./client/cryptonatorClientInterface.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	models "Golang-assignment-srikomm/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptonatorClientInterface is a mock of CryptonatorClientInterface interface.
type MockCryptonatorClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCryptonatorClientInterfaceMockRecorder
}

// MockCryptonatorClientInterfaceMockRecorder is the mock recorder for MockCryptonatorClientInterface.
type MockCryptonatorClientInterfaceMockRecorder struct {
	mock *MockCryptonatorClientInterface
}

// NewMockCryptonatorClientInterface creates a new mock instance.
func NewMockCryptonatorClientInterface(ctrl *gomock.Controller) *MockCryptonatorClientInterface {
	mock := &MockCryptonatorClientInterface{ctrl: ctrl}
	mock.recorder = &MockCryptonatorClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptonatorClientInterface) EXPECT() *MockCryptonatorClientInterfaceMockRecorder {
	return m.recorder
}

// GetETHCurrentPrice mocks base method.
func (m *MockCryptonatorClientInterface) GetETHCurrentPrice() (models.Crypto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetETHCurrentPrice")
	ret0, _ := ret[0].(models.Crypto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetETHCurrentPrice indicates an expected call of GetETHCurrentPrice.
func (mr *MockCryptonatorClientInterfaceMockRecorder) GetETHCurrentPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetETHCurrentPrice", reflect.TypeOf((*MockCryptonatorClientInterface)(nil).GetETHCurrentPrice))
}