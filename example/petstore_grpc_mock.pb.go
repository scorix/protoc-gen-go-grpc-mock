// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: petstore.proto

package petstore

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockPetStoreClient is a mock of PetStoreClient interface.
type MockPetStoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockPetStoreClientMockRecorder
}

// MockPetStoreClientMockRecorder is the mock recorder for MockPetStoreClient.
type MockPetStoreClientMockRecorder struct {
	mock *MockPetStoreClient
}

// NewMockPetStoreClient creates a new mock instance.
func NewMockPetStoreClient(ctrl *gomock.Controller) *MockPetStoreClient {
	mock := &MockPetStoreClient{ctrl: ctrl}
	mock.recorder = &MockPetStoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetStoreClient) EXPECT() *MockPetStoreClientMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockPetStoreClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Pets, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAll", varargs...)
	ret0, _ := ret[0].(*Pets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPetStoreClientMockRecorder) GetAll(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPetStoreClient)(nil).GetAll), varargs...)
}

// GetPet mocks base method.
func (m *MockPetStoreClient) GetPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPet", varargs...)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPet indicates an expected call of GetPet.
func (mr *MockPetStoreClientMockRecorder) GetPet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPet", reflect.TypeOf((*MockPetStoreClient)(nil).GetPet), varargs...)
}

// CreatePet mocks base method.
func (m *MockPetStoreClient) CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePet", varargs...)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePet indicates an expected call of CreatePet.
func (mr *MockPetStoreClientMockRecorder) CreatePet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePet", reflect.TypeOf((*MockPetStoreClient)(nil).CreatePet), varargs...)
}

// UpdatePet mocks base method.
func (m *MockPetStoreClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePet", varargs...)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePet indicates an expected call of UpdatePet.
func (mr *MockPetStoreClientMockRecorder) UpdatePet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePet", reflect.TypeOf((*MockPetStoreClient)(nil).UpdatePet), varargs...)
}

// DeletePet mocks base method.
func (m *MockPetStoreClient) DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePet", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePet indicates an expected call of DeletePet.
func (mr *MockPetStoreClientMockRecorder) DeletePet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePet", reflect.TypeOf((*MockPetStoreClient)(nil).DeletePet), varargs...)
}

// MockPetStoreServer is a mock of PetStoreServer interface.
type MockPetStoreServer struct {
	ctrl     *gomock.Controller
	recorder *MockPetStoreServerMockRecorder
}

// MockPetStoreServerMockRecorder is the mock recorder for MockPetStoreServer.
type MockPetStoreServerMockRecorder struct {
	mock *MockPetStoreServer
}

// NewMockPetStoreServer creates a new mock instance.
func NewMockPetStoreServer(ctrl *gomock.Controller) *MockPetStoreServer {
	mock := &MockPetStoreServer{ctrl: ctrl}
	mock.recorder = &MockPetStoreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetStoreServer) EXPECT() *MockPetStoreServerMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockPetStoreServer) GetAll(arg0 context.Context, arg1 *emptypb.Empty) (*Pets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].(*Pets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPetStoreServerMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPetStoreServer)(nil).GetAll), arg0, arg1)
}

// GetPet mocks base method.
func (m *MockPetStoreServer) GetPet(arg0 context.Context, arg1 *Pet) (*Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPet", arg0, arg1)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPet indicates an expected call of GetPet.
func (mr *MockPetStoreServerMockRecorder) GetPet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPet", reflect.TypeOf((*MockPetStoreServer)(nil).GetPet), arg0, arg1)
}

// CreatePet mocks base method.
func (m *MockPetStoreServer) CreatePet(arg0 context.Context, arg1 *Pet) (*Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePet", arg0, arg1)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePet indicates an expected call of CreatePet.
func (mr *MockPetStoreServerMockRecorder) CreatePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePet", reflect.TypeOf((*MockPetStoreServer)(nil).CreatePet), arg0, arg1)
}

// UpdatePet mocks base method.
func (m *MockPetStoreServer) UpdatePet(arg0 context.Context, arg1 *Pet) (*Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePet", arg0, arg1)
	ret0, _ := ret[0].(*Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePet indicates an expected call of UpdatePet.
func (mr *MockPetStoreServerMockRecorder) UpdatePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePet", reflect.TypeOf((*MockPetStoreServer)(nil).UpdatePet), arg0, arg1)
}

// DeletePet mocks base method.
func (m *MockPetStoreServer) DeletePet(arg0 context.Context, arg1 *Pet) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePet", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePet indicates an expected call of DeletePet.
func (mr *MockPetStoreServerMockRecorder) DeletePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePet", reflect.TypeOf((*MockPetStoreServer)(nil).DeletePet), arg0, arg1)
}

// mustEmbedUnimplementedPetStoreServer mocks base method.
func (m *MockPetStoreServer) mustEmbedUnimplementedPetStoreServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPetStoreServer")
}

// mustEmbedUnimplementedPetStoreServer indicates an expected call of mustEmbedUnimplementedPetStoreServer.
func (mr *MockPetStoreServerMockRecorder) mustEmbedUnimplementedPetStoreServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPetStoreServer", reflect.TypeOf((*MockPetStoreServer)(nil).mustEmbedUnimplementedPetStoreServer))
}

// MockUnsafePetStoreServer is a mock of UnsafePetStoreServer interface.
type MockUnsafePetStoreServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePetStoreServerMockRecorder
}

// MockUnsafePetStoreServerMockRecorder is the mock recorder for MockUnsafePetStoreServer.
type MockUnsafePetStoreServerMockRecorder struct {
	mock *MockUnsafePetStoreServer
}

// NewMockUnsafePetStoreServer creates a new mock instance.
func NewMockUnsafePetStoreServer(ctrl *gomock.Controller) *MockUnsafePetStoreServer {
	mock := &MockUnsafePetStoreServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePetStoreServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePetStoreServer) EXPECT() *MockUnsafePetStoreServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPetStoreServer mocks base method.
func (m *MockUnsafePetStoreServer) mustEmbedUnimplementedPetStoreServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPetStoreServer")
}

// mustEmbedUnimplementedPetStoreServer indicates an expected call of mustEmbedUnimplementedPetStoreServer.
func (mr *MockUnsafePetStoreServerMockRecorder) mustEmbedUnimplementedPetStoreServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPetStoreServer", reflect.TypeOf((*MockUnsafePetStoreServer)(nil).mustEmbedUnimplementedPetStoreServer))
}