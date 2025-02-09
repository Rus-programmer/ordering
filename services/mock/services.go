// Code generated by MockGen. DO NOT EDIT.
// Source: ordering/services (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -package mockService -destination services/mock/services.go ordering/services Service
//

// Package mockService is a generated GoMock package.
package mockService

import (
	context "context"
	dto "ordering/dto"
	auth "ordering/services/auth"
	customers "ordering/services/customers"
	order "ordering/services/orders"
	products "ordering/services/products"
	token "ordering/token"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockService) CreateCustomer(ctx context.Context, req customers.CreateCustomer) (dto.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", ctx, req)
	ret0, _ := ret[0].(dto.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockServiceMockRecorder) CreateCustomer(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockService)(nil).CreateCustomer), ctx, req)
}

// CreateProduct mocks base method.
func (m *MockService) CreateProduct(ctx context.Context, body products.CreateProduct) (dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, body)
	ret0, _ := ret[0].(dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockServiceMockRecorder) CreateProduct(ctx, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockService)(nil).CreateProduct), ctx, body)
}

// DeleteOrder mocks base method.
func (m *MockService) DeleteOrder(ctx context.Context, req order.DeleteOrderParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockServiceMockRecorder) DeleteOrder(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockService)(nil).DeleteOrder), ctx, req)
}

// DeleteProduct mocks base method.
func (m *MockService) DeleteProduct(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockServiceMockRecorder) DeleteProduct(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockService)(nil).DeleteProduct), ctx, id)
}

// GetOrder mocks base method.
func (m *MockService) GetOrder(ctx context.Context, req order.GetOrder) (dto.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, req)
	ret0, _ := ret[0].(dto.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockServiceMockRecorder) GetOrder(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockService)(nil).GetOrder), ctx, req)
}

// GetProduct mocks base method.
func (m *MockService) GetProduct(ctx context.Context, id int64) (dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, id)
	ret0, _ := ret[0].(dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockServiceMockRecorder) GetProduct(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockService)(nil).GetProduct), ctx, id)
}

// GetTokenMaker mocks base method.
func (m *MockService) GetTokenMaker() token.Maker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenMaker")
	ret0, _ := ret[0].(token.Maker)
	return ret0
}

// GetTokenMaker indicates an expected call of GetTokenMaker.
func (mr *MockServiceMockRecorder) GetTokenMaker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenMaker", reflect.TypeOf((*MockService)(nil).GetTokenMaker))
}

// ListOrders mocks base method.
func (m *MockService) ListOrders(ctx context.Context, req order.ListOrders) ([]dto.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", ctx, req)
	ret0, _ := ret[0].([]dto.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockServiceMockRecorder) ListOrders(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockService)(nil).ListOrders), ctx, req)
}

// ListProducts mocks base method.
func (m *MockService) ListProducts(ctx context.Context, req products.ListProductRequest) ([]dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", ctx, req)
	ret0, _ := ret[0].([]dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockServiceMockRecorder) ListProducts(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockService)(nil).ListProducts), ctx, req)
}

// Login mocks base method.
func (m *MockService) Login(ctx context.Context, req auth.LoginRequest, clientIp, userAgent string) (dto.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req, clientIp, userAgent)
	ret0, _ := ret[0].(dto.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(ctx, req, clientIp, userAgent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), ctx, req, clientIp, userAgent)
}

// RenewAccessToken mocks base method.
func (m *MockService) RenewAccessToken(ctx context.Context, refreshToken string) (dto.RenewAccessTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewAccessToken", ctx, refreshToken)
	ret0, _ := ret[0].(dto.RenewAccessTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewAccessToken indicates an expected call of RenewAccessToken.
func (mr *MockServiceMockRecorder) RenewAccessToken(ctx, refreshToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewAccessToken", reflect.TypeOf((*MockService)(nil).RenewAccessToken), ctx, refreshToken)
}

// UpdateProduct mocks base method.
func (m *MockService) UpdateProduct(ctx context.Context, id int64, body products.UpdateProduct) (dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, id, body)
	ret0, _ := ret[0].(dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockServiceMockRecorder) UpdateProduct(ctx, id, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockService)(nil).UpdateProduct), ctx, id, body)
}
