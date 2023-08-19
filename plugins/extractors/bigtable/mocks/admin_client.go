// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	bigtable "cloud.google.com/go/bigtable"

	mock "github.com/stretchr/testify/mock"
)

// AdminClient is an autogenerated mock type for the AdminClient type
type AdminClient struct {
	mock.Mock
}

type AdminClient_Expecter struct {
	mock *mock.Mock
}

func (_m *AdminClient) EXPECT() *AdminClient_Expecter {
	return &AdminClient_Expecter{mock: &_m.Mock}
}

// TableInfo provides a mock function with given fields: ctx, table
func (_m *AdminClient) TableInfo(ctx context.Context, table string) (*bigtable.TableInfo, error) {
	ret := _m.Called(ctx, table)

	var r0 *bigtable.TableInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*bigtable.TableInfo, error)); ok {
		return rf(ctx, table)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *bigtable.TableInfo); ok {
		r0 = rf(ctx, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bigtable.TableInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AdminClient_TableInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TableInfo'
type AdminClient_TableInfo_Call struct {
	*mock.Call
}

// TableInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - table string
func (_e *AdminClient_Expecter) TableInfo(ctx interface{}, table interface{}) *AdminClient_TableInfo_Call {
	return &AdminClient_TableInfo_Call{Call: _e.mock.On("TableInfo", ctx, table)}
}

func (_c *AdminClient_TableInfo_Call) Run(run func(ctx context.Context, table string)) *AdminClient_TableInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AdminClient_TableInfo_Call) Return(_a0 *bigtable.TableInfo, _a1 error) *AdminClient_TableInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AdminClient_TableInfo_Call) RunAndReturn(run func(context.Context, string) (*bigtable.TableInfo, error)) *AdminClient_TableInfo_Call {
	_c.Call.Return(run)
	return _c
}

// Tables provides a mock function with given fields: ctx
func (_m *AdminClient) Tables(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AdminClient_Tables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tables'
type AdminClient_Tables_Call struct {
	*mock.Call
}

// Tables is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AdminClient_Expecter) Tables(ctx interface{}) *AdminClient_Tables_Call {
	return &AdminClient_Tables_Call{Call: _e.mock.On("Tables", ctx)}
}

func (_c *AdminClient_Tables_Call) Run(run func(ctx context.Context)) *AdminClient_Tables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AdminClient_Tables_Call) Return(_a0 []string, _a1 error) *AdminClient_Tables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AdminClient_Tables_Call) RunAndReturn(run func(context.Context) ([]string, error)) *AdminClient_Tables_Call {
	_c.Call.Return(run)
	return _c
}

// NewAdminClient creates a new instance of AdminClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminClient {
	mock := &AdminClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
