// Code generated by MockGen. DO NOT EDIT.
// Source: gen/go/proto/users/v3/users.pb.go

// Package mock_v3pb is a generated GoMock package.
package mock_v3pb

import (
	x "."
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUsersClient is a mock of UsersClient interface
type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

// MockUsersClientMockRecorder is the mock recorder for MockUsersClient
type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

// NewMockUsersClient creates a new mock instance
func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUsersClient) Get(ctx context.Context, in *x.GetRequest, opts ...grpc.CallOption) (*x.User, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*x.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsersClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsersClient)(nil).Get), varargs...)
}

// MockUsersServer is a mock of UsersServer interface
type MockUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServerMockRecorder
}

// MockUsersServerMockRecorder is the mock recorder for MockUsersServer
type MockUsersServerMockRecorder struct {
	mock *MockUsersServer
}

// NewMockUsersServer creates a new mock instance
func NewMockUsersServer(ctrl *gomock.Controller) *MockUsersServer {
	mock := &MockUsersServer{ctrl: ctrl}
	mock.recorder = &MockUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersServer) EXPECT() *MockUsersServerMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUsersServer) Get(arg0 context.Context, arg1 *x.GetRequest) (*x.User, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*x.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsersServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsersServer)(nil).Get), arg0, arg1)
}
