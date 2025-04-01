// Code generated by mockery v2.52.2. DO NOT EDIT.

package fees

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	mock "github.com/stretchr/testify/mock"

	rateregistry "github.com/xmtp/xmtpd/pkg/abi/rateregistry"
)

// MockRatesContract is an autogenerated mock type for the RatesContract type
type MockRatesContract struct {
	mock.Mock
}

type MockRatesContract_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRatesContract) EXPECT() *MockRatesContract_Expecter {
	return &MockRatesContract_Expecter{mock: &_m.Mock}
}

// GetRates provides a mock function with given fields: opts, fromIndex
func (_m *MockRatesContract) GetRates(opts *bind.CallOpts, fromIndex *big.Int) (struct {
	Rates   []rateregistry.RateRegistryRates
	HasMore bool
}, error) {
	ret := _m.Called(opts, fromIndex)

	if len(ret) == 0 {
		panic("no return value specified for GetRates")
	}

	var r0 struct {
		Rates   []rateregistry.RateRegistryRates
		HasMore bool
	}
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) (struct {
		Rates   []rateregistry.RateRegistryRates
		HasMore bool
	}, error)); ok {
		return rf(opts, fromIndex)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) struct {
		Rates   []rateregistry.RateRegistryRates
		HasMore bool
	}); ok {
		r0 = rf(opts, fromIndex)
	} else {
		r0 = ret.Get(0).(struct {
			Rates   []rateregistry.RateRegistryRates
			HasMore bool
		})
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, fromIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRatesContract_GetRates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRates'
type MockRatesContract_GetRates_Call struct {
	*mock.Call
}

// GetRates is a helper method to define mock.On call
//   - opts *bind.CallOpts
//   - fromIndex *big.Int
func (_e *MockRatesContract_Expecter) GetRates(opts interface{}, fromIndex interface{}) *MockRatesContract_GetRates_Call {
	return &MockRatesContract_GetRates_Call{Call: _e.mock.On("GetRates", opts, fromIndex)}
}

func (_c *MockRatesContract_GetRates_Call) Run(run func(opts *bind.CallOpts, fromIndex *big.Int)) *MockRatesContract_GetRates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bind.CallOpts), args[1].(*big.Int))
	})
	return _c
}

func (_c *MockRatesContract_GetRates_Call) Return(_a0 struct {
	Rates   []rateregistry.RateRegistryRates
	HasMore bool
}, _a1 error) *MockRatesContract_GetRates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRatesContract_GetRates_Call) RunAndReturn(run func(*bind.CallOpts, *big.Int) (struct {
	Rates   []rateregistry.RateRegistryRates
	HasMore bool
}, error)) *MockRatesContract_GetRates_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRatesContract creates a new instance of MockRatesContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRatesContract(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRatesContract {
	mock := &MockRatesContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
