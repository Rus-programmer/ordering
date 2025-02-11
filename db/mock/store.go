// Code generated by MockGen. DO NOT EDIT.
// Source: ordering/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go ordering/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	db "ordering/db/sqlc"
	reflect "reflect"

	uuid "github.com/google/uuid"
	pgtype "github.com/jackc/pgx/v5/pgtype"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
	isgomock struct{}
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

// CreateCustomer mocks base method.
func (m *MockStore) CreateCustomer(ctx context.Context, arg db.CreateCustomerParams) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", ctx, arg)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockStoreMockRecorder) CreateCustomer(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockStore)(nil).CreateCustomer), ctx, arg)
}

// CreateLog mocks base method.
func (m *MockStore) CreateLog(ctx context.Context, arg db.CreateLogParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLog", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLog indicates an expected call of CreateLog.
func (mr *MockStoreMockRecorder) CreateLog(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLog", reflect.TypeOf((*MockStore)(nil).CreateLog), ctx, arg)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(ctx context.Context, arg db.CreateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, arg)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), ctx, arg)
}

// CreateOrderProducts mocks base method.
func (m *MockStore) CreateOrderProducts(ctx context.Context, arg []db.CreateOrderProductsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderProducts", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderProducts indicates an expected call of CreateOrderProducts.
func (mr *MockStoreMockRecorder) CreateOrderProducts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderProducts", reflect.TypeOf((*MockStore)(nil).CreateOrderProducts), ctx, arg)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(ctx context.Context, arg db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, arg)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), ctx, arg)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, arg)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), ctx, arg)
}

// DeleteCustomer mocks base method.
func (m *MockStore) DeleteCustomer(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomer indicates an expected call of DeleteCustomer.
func (mr *MockStoreMockRecorder) DeleteCustomer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomer", reflect.TypeOf((*MockStore)(nil).DeleteCustomer), ctx, id)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), ctx, id)
}

// DeleteOrderProduct mocks base method.
func (m *MockStore) DeleteOrderProduct(ctx context.Context, arg db.DeleteOrderProductParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderProduct", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderProduct indicates an expected call of DeleteOrderProduct.
func (mr *MockStoreMockRecorder) DeleteOrderProduct(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderProduct", reflect.TypeOf((*MockStore)(nil).DeleteOrderProduct), ctx, arg)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), ctx, id)
}

// ExecTx mocks base method.
func (m *MockStore) ExecTx(ctx context.Context, fn func(*db.Queries) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockStoreMockRecorder) ExecTx(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockStore)(nil).ExecTx), ctx, fn)
}

// GetCustomerByUsername mocks base method.
func (m *MockStore) GetCustomerByUsername(ctx context.Context, username string) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerByUsername", ctx, username)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerByUsername indicates an expected call of GetCustomerByUsername.
func (mr *MockStoreMockRecorder) GetCustomerByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerByUsername", reflect.TypeOf((*MockStore)(nil).GetCustomerByUsername), ctx, username)
}

// GetErrorRate mocks base method.
func (m *MockStore) GetErrorRate(ctx context.Context) (pgtype.Numeric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorRate", ctx)
	ret0, _ := ret[0].(pgtype.Numeric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetErrorRate indicates an expected call of GetErrorRate.
func (mr *MockStoreMockRecorder) GetErrorRate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorRate", reflect.TypeOf((*MockStore)(nil).GetErrorRate), ctx)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(ctx context.Context, arg db.GetOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, arg)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), ctx, arg)
}

// GetOrderProducts mocks base method.
func (m *MockStore) GetOrderProducts(ctx context.Context, orderID int64) ([]db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderProducts", ctx, orderID)
	ret0, _ := ret[0].([]db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderProducts indicates an expected call of GetOrderProducts.
func (mr *MockStoreMockRecorder) GetOrderProducts(ctx, orderID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderProducts", reflect.TypeOf((*MockStore)(nil).GetOrderProducts), ctx, orderID)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(ctx context.Context, id int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, id)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), ctx, id)
}

// GetProductForUpdate mocks base method.
func (m *MockStore) GetProductForUpdate(ctx context.Context, id int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductForUpdate", ctx, id)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductForUpdate indicates an expected call of GetProductForUpdate.
func (mr *MockStoreMockRecorder) GetProductForUpdate(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductForUpdate", reflect.TypeOf((*MockStore)(nil).GetProductForUpdate), ctx, id)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(ctx context.Context, id uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, id)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), ctx, id)
}

// GetTotalPrice mocks base method.
func (m *MockStore) GetTotalPrice(ctx context.Context, orderID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPrice", ctx, orderID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPrice indicates an expected call of GetTotalPrice.
func (mr *MockStoreMockRecorder) GetTotalPrice(ctx, orderID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPrice", reflect.TypeOf((*MockStore)(nil).GetTotalPrice), ctx, orderID)
}

// GetTotalRequests mocks base method.
func (m *MockStore) GetTotalRequests(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequests", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalRequests indicates an expected call of GetTotalRequests.
func (mr *MockStoreMockRecorder) GetTotalRequests(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequests", reflect.TypeOf((*MockStore)(nil).GetTotalRequests), ctx)
}

// GetTotalRequestsByMethod mocks base method.
func (m *MockStore) GetTotalRequestsByMethod(ctx context.Context) ([]db.GetTotalRequestsByMethodRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequestsByMethod", ctx)
	ret0, _ := ret[0].([]db.GetTotalRequestsByMethodRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalRequestsByMethod indicates an expected call of GetTotalRequestsByMethod.
func (mr *MockStoreMockRecorder) GetTotalRequestsByMethod(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequestsByMethod", reflect.TypeOf((*MockStore)(nil).GetTotalRequestsByMethod), ctx)
}

// GetTotalRequestsByPath mocks base method.
func (m *MockStore) GetTotalRequestsByPath(ctx context.Context) ([]db.GetTotalRequestsByPathRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequestsByPath", ctx)
	ret0, _ := ret[0].([]db.GetTotalRequestsByPathRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalRequestsByPath indicates an expected call of GetTotalRequestsByPath.
func (mr *MockStoreMockRecorder) GetTotalRequestsByPath(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequestsByPath", reflect.TypeOf((*MockStore)(nil).GetTotalRequestsByPath), ctx)
}

// GetTotalRequestsByStatusCode mocks base method.
func (m *MockStore) GetTotalRequestsByStatusCode(ctx context.Context) ([]db.GetTotalRequestsByStatusCodeRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalRequestsByStatusCode", ctx)
	ret0, _ := ret[0].([]db.GetTotalRequestsByStatusCodeRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalRequestsByStatusCode indicates an expected call of GetTotalRequestsByStatusCode.
func (mr *MockStoreMockRecorder) GetTotalRequestsByStatusCode(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalRequestsByStatusCode", reflect.TypeOf((*MockStore)(nil).GetTotalRequestsByStatusCode), ctx)
}

// ListCustomers mocks base method.
func (m *MockStore) ListCustomers(ctx context.Context, arg db.ListCustomersParams) ([]db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCustomers", ctx, arg)
	ret0, _ := ret[0].([]db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomers indicates an expected call of ListCustomers.
func (mr *MockStoreMockRecorder) ListCustomers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomers", reflect.TypeOf((*MockStore)(nil).ListCustomers), ctx, arg)
}

// ListOrders mocks base method.
func (m *MockStore) ListOrders(ctx context.Context, arg db.ListOrdersParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", ctx, arg)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockStoreMockRecorder) ListOrders(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), ctx, arg)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(ctx context.Context, arg db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", ctx, arg)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), ctx, arg)
}

// SoftDeleteOrder mocks base method.
func (m *MockStore) SoftDeleteOrder(ctx context.Context, arg db.SoftDeleteOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDeleteOrder", ctx, arg)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftDeleteOrder indicates an expected call of SoftDeleteOrder.
func (mr *MockStoreMockRecorder) SoftDeleteOrder(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDeleteOrder", reflect.TypeOf((*MockStore)(nil).SoftDeleteOrder), ctx, arg)
}

// UpdateOrder mocks base method.
func (m *MockStore) UpdateOrder(ctx context.Context, arg db.UpdateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, arg)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockStoreMockRecorder) UpdateOrder(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockStore)(nil).UpdateOrder), ctx, arg)
}

// UpdateOrderProduct mocks base method.
func (m *MockStore) UpdateOrderProduct(ctx context.Context, arg db.UpdateOrderProductParams) (db.OrderProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderProduct", ctx, arg)
	ret0, _ := ret[0].(db.OrderProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderProduct indicates an expected call of UpdateOrderProduct.
func (mr *MockStoreMockRecorder) UpdateOrderProduct(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderProduct", reflect.TypeOf((*MockStore)(nil).UpdateOrderProduct), ctx, arg)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(ctx context.Context, arg db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, arg)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), ctx, arg)
}
