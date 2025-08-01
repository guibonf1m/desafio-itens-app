// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	item "desafio-itens-app/internal/domain/item"

	mock "github.com/stretchr/testify/mock"
)

// ItemRepository is an autogenerated mock type for the ItemRepository type
type ItemRepository struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: _a0
func (_m *ItemRepository) AddItem(_a0 item.Item) (item.Item, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddItem")
	}

	var r0 item.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(item.Item) (item.Item, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(item.Item) item.Item); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(item.Item)
	}

	if rf, ok := ret.Get(1).(func(item.Item) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CodeExists provides a mock function with given fields: code
func (_m *ItemRepository) CodeExists(code string) (bool, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for CodeExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountItens provides a mock function with given fields: status
func (_m *ItemRepository) CountItens(status *item.Status) (int, error) {
	ret := _m.Called(status)

	if len(ret) == 0 {
		panic("no return value specified for CountItens")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*item.Status) (int, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(*item.Status) int); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*item.Status) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteItem provides a mock function with given fields: id
func (_m *ItemRepository) DeleteItem(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetItem provides a mock function with given fields: id
func (_m *ItemRepository) GetItem(id int) (*item.Item, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *item.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*item.Item, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *item.Item); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItens provides a mock function with no fields
func (_m *ItemRepository) GetItens() ([]item.Item, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetItens")
	}

	var r0 []item.Item
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]item.Item, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []item.Item); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItensFiltrados provides a mock function with given fields: status, limit
func (_m *ItemRepository) GetItensFiltrados(status *item.Status, limit int) ([]item.Item, error) {
	ret := _m.Called(status, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetItensFiltrados")
	}

	var r0 []item.Item
	var r1 error
	if rf, ok := ret.Get(0).(func(*item.Status, int) ([]item.Item, error)); ok {
		return rf(status, limit)
	}
	if rf, ok := ret.Get(0).(func(*item.Status, int) []item.Item); ok {
		r0 = rf(status, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(*item.Status, int) error); ok {
		r1 = rf(status, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItensFiltradosPaginados provides a mock function with given fields: status, page, pageSize
func (_m *ItemRepository) GetItensFiltradosPaginados(status *item.Status, page int, pageSize int) ([]item.Item, int, error) {
	ret := _m.Called(status, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetItensFiltradosPaginados")
	}

	var r0 []item.Item
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(*item.Status, int, int) ([]item.Item, int, error)); ok {
		return rf(status, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(*item.Status, int, int) []item.Item); ok {
		r0 = rf(status, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(*item.Status, int, int) int); ok {
		r1 = rf(status, page, pageSize)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(*item.Status, int, int) error); ok {
		r2 = rf(status, page, pageSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetItensPaginados provides a mock function with given fields: ofsset, limit
func (_m *ItemRepository) GetItensPaginados(ofsset int, limit int) ([]item.Item, int, error) {
	ret := _m.Called(ofsset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetItensPaginados")
	}

	var r0 []item.Item
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]item.Item, int, error)); ok {
		return rf(ofsset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []item.Item); ok {
		r0 = rf(ofsset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]item.Item)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(ofsset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(ofsset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateItem provides a mock function with given fields: _a0
func (_m *ItemRepository) UpdateItem(_a0 item.Item) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(item.Item) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewItemRepository creates a new instance of ItemRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewItemRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ItemRepository {
	mock := &ItemRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
