// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/smartcontractkit/chainlink/core/store (interfaces: TxManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	go_ethereum "github.com/ethereum/go-ethereum"
	accounts "github.com/ethereum/go-ethereum/accounts"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	store "github.com/smartcontractkit/chainlink/core/store"
	assets "github.com/smartcontractkit/chainlink/core/store/assets"
	models "github.com/smartcontractkit/chainlink/core/store/models"
	null_v3 "gopkg.in/guregu/null.v3"
	big "math/big"
	reflect "reflect"
)

// MockTxManager is a mock of TxManager interface
type MockTxManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxManagerMockRecorder
}

// MockTxManagerMockRecorder is the mock recorder for MockTxManager
type MockTxManagerMockRecorder struct {
	mock *MockTxManager
}

// NewMockTxManager creates a new mock instance
func NewMockTxManager(ctrl *gomock.Controller) *MockTxManager {
	mock := &MockTxManager{ctrl: ctrl}
	mock.recorder = &MockTxManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxManager) EXPECT() *MockTxManagerMockRecorder {
	return m.recorder
}

// BumpGasUntilSafe mocks base method
func (m *MockTxManager) BumpGasUntilSafe(arg0 common.Hash) (*models.TxReceipt, store.AttemptState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BumpGasUntilSafe", arg0)
	ret0, _ := ret[0].(*models.TxReceipt)
	ret1, _ := ret[1].(store.AttemptState)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BumpGasUntilSafe indicates an expected call of BumpGasUntilSafe
func (mr *MockTxManagerMockRecorder) BumpGasUntilSafe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BumpGasUntilSafe", reflect.TypeOf((*MockTxManager)(nil).BumpGasUntilSafe), arg0)
}

// CheckAttempt mocks base method
func (m *MockTxManager) CheckAttempt(arg0 *models.TxAttempt, arg1 uint64) (*models.TxReceipt, store.AttemptState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAttempt", arg0, arg1)
	ret0, _ := ret[0].(*models.TxReceipt)
	ret1, _ := ret[1].(store.AttemptState)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckAttempt indicates an expected call of CheckAttempt
func (mr *MockTxManagerMockRecorder) CheckAttempt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAttempt", reflect.TypeOf((*MockTxManager)(nil).CheckAttempt), arg0, arg1)
}

// Connect mocks base method
func (m *MockTxManager) Connect(arg0 *models.Head) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockTxManagerMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockTxManager)(nil).Connect), arg0)
}

// Connected mocks base method
func (m *MockTxManager) Connected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Connected indicates an expected call of Connected
func (mr *MockTxManagerMockRecorder) Connected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockTxManager)(nil).Connected))
}

// ContractLINKBalance mocks base method
func (m *MockTxManager) ContractLINKBalance(arg0 models.WithdrawalRequest) (assets.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContractLINKBalance", arg0)
	ret0, _ := ret[0].(assets.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContractLINKBalance indicates an expected call of ContractLINKBalance
func (mr *MockTxManagerMockRecorder) ContractLINKBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContractLINKBalance", reflect.TypeOf((*MockTxManager)(nil).ContractLINKBalance), arg0)
}

// CreateTx mocks base method
func (m *MockTxManager) CreateTx(arg0 common.Address, arg1 []byte) (*models.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", arg0, arg1)
	ret0, _ := ret[0].(*models.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx
func (mr *MockTxManagerMockRecorder) CreateTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockTxManager)(nil).CreateTx), arg0, arg1)
}

// CreateTxWithEth mocks base method
func (m *MockTxManager) CreateTxWithEth(arg0, arg1 common.Address, arg2 *assets.Eth) (*models.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTxWithEth", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTxWithEth indicates an expected call of CreateTxWithEth
func (mr *MockTxManagerMockRecorder) CreateTxWithEth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTxWithEth", reflect.TypeOf((*MockTxManager)(nil).CreateTxWithEth), arg0, arg1, arg2)
}

// CreateTxWithGas mocks base method
func (m *MockTxManager) CreateTxWithGas(arg0 null_v3.String, arg1 common.Address, arg2 []byte, arg3 *big.Int, arg4 uint64) (*models.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTxWithGas", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*models.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTxWithGas indicates an expected call of CreateTxWithGas
func (mr *MockTxManagerMockRecorder) CreateTxWithGas(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTxWithGas", reflect.TypeOf((*MockTxManager)(nil).CreateTxWithGas), arg0, arg1, arg2, arg3, arg4)
}

// Disconnect mocks base method
func (m *MockTxManager) Disconnect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect")
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockTxManagerMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockTxManager)(nil).Disconnect))
}

// GetBlockByNumber mocks base method
func (m *MockTxManager) GetBlockByNumber(arg0 string) (models.BlockHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0)
	ret0, _ := ret[0].(models.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockTxManagerMockRecorder) GetBlockByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockTxManager)(nil).GetBlockByNumber), arg0)
}

// GetEthBalance mocks base method
func (m *MockTxManager) GetEthBalance(arg0 common.Address) (*assets.Eth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEthBalance", arg0)
	ret0, _ := ret[0].(*assets.Eth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEthBalance indicates an expected call of GetEthBalance
func (mr *MockTxManagerMockRecorder) GetEthBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEthBalance", reflect.TypeOf((*MockTxManager)(nil).GetEthBalance), arg0)
}

// GetLINKBalance mocks base method
func (m *MockTxManager) GetLINKBalance(arg0 common.Address) (*assets.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLINKBalance", arg0)
	ret0, _ := ret[0].(*assets.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLINKBalance indicates an expected call of GetLINKBalance
func (mr *MockTxManagerMockRecorder) GetLINKBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLINKBalance", reflect.TypeOf((*MockTxManager)(nil).GetLINKBalance), arg0)
}

// GetLogs mocks base method
func (m *MockTxManager) GetLogs(arg0 go_ethereum.FilterQuery) ([]models.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0)
	ret0, _ := ret[0].([]models.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockTxManagerMockRecorder) GetLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockTxManager)(nil).GetLogs), arg0)
}

// GetTxReceipt mocks base method
func (m *MockTxManager) GetTxReceipt(arg0 common.Hash) (*models.TxReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxReceipt", arg0)
	ret0, _ := ret[0].(*models.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxReceipt indicates an expected call of GetTxReceipt
func (mr *MockTxManagerMockRecorder) GetTxReceipt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxReceipt", reflect.TypeOf((*MockTxManager)(nil).GetTxReceipt), arg0)
}

// NextActiveAccount mocks base method
func (m *MockTxManager) NextActiveAccount() *store.ManagedAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextActiveAccount")
	ret0, _ := ret[0].(*store.ManagedAccount)
	return ret0
}

// NextActiveAccount indicates an expected call of NextActiveAccount
func (mr *MockTxManagerMockRecorder) NextActiveAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextActiveAccount", reflect.TypeOf((*MockTxManager)(nil).NextActiveAccount))
}

// OnNewHead mocks base method
func (m *MockTxManager) OnNewHead(arg0 *models.Head) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNewHead", arg0)
}

// OnNewHead indicates an expected call of OnNewHead
func (mr *MockTxManagerMockRecorder) OnNewHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNewHead", reflect.TypeOf((*MockTxManager)(nil).OnNewHead), arg0)
}

// Register mocks base method
func (m *MockTxManager) Register(arg0 []accounts.Account) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register
func (mr *MockTxManagerMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockTxManager)(nil).Register), arg0)
}

// SubscribeToLogs mocks base method
func (m *MockTxManager) SubscribeToLogs(arg0 chan<- models.Log, arg1 go_ethereum.FilterQuery) (models.EthSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToLogs", arg0, arg1)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToLogs indicates an expected call of SubscribeToLogs
func (mr *MockTxManagerMockRecorder) SubscribeToLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLogs", reflect.TypeOf((*MockTxManager)(nil).SubscribeToLogs), arg0, arg1)
}

// SubscribeToNewHeads mocks base method
func (m *MockTxManager) SubscribeToNewHeads(arg0 chan<- models.BlockHeader) (models.EthSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewHeads", arg0)
	ret0, _ := ret[0].(models.EthSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToNewHeads indicates an expected call of SubscribeToNewHeads
func (mr *MockTxManagerMockRecorder) SubscribeToNewHeads(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewHeads", reflect.TypeOf((*MockTxManager)(nil).SubscribeToNewHeads), arg0)
}

// WithdrawLINK mocks base method
func (m *MockTxManager) WithdrawLINK(arg0 models.WithdrawalRequest) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawLINK", arg0)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawLINK indicates an expected call of WithdrawLINK
func (mr *MockTxManagerMockRecorder) WithdrawLINK(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawLINK", reflect.TypeOf((*MockTxManager)(nil).WithdrawLINK), arg0)
}
