// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *Transaction) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: count
func (_m *Transaction) Count(count *int64) error {
	ret := _m.Called(count)

	var r0 error
	if rf, ok := ret.Get(0).(func(*int64) error); ok {
		r0 = rf(count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *Transaction) Create(value interface{}) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: value, conds
func (_m *Transaction) Delete(value interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(value, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Distinct provides a mock function with given fields: args
func (_m *Transaction) Distinct(args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...interface{}) orm.Query); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Driver provides a mock function with given fields:
func (_m *Transaction) Driver() orm.Driver {
	ret := _m.Called()

	var r0 orm.Driver
	if rf, ok := ret.Get(0).(func() orm.Driver); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(orm.Driver)
	}

	return r0
}

// Exec provides a mock function with given fields: sql, values
func (_m *Transaction) Exec(sql string, values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, sql)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(sql, values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: dest, conds
func (_m *Transaction) Find(dest interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(dest, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// First provides a mock function with given fields: dest
func (_m *Transaction) First(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: dest, conds
func (_m *Transaction) FirstOrCreate(dest interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(dest, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDelete provides a mock function with given fields: value, conds
func (_m *Transaction) ForceDelete(value interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(value, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: dest
func (_m *Transaction) Get(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Group provides a mock function with given fields: name
func (_m *Transaction) Group(name string) orm.Query {
	ret := _m.Called(name)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string) orm.Query); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Having provides a mock function with given fields: query, args
func (_m *Transaction) Having(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Join provides a mock function with given fields: query, args
func (_m *Transaction) Join(query string, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit
func (_m *Transaction) Limit(limit int) orm.Query {
	ret := _m.Called(limit)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *Transaction) Model(value interface{}) orm.Query {
	ret := _m.Called(value)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}) orm.Query); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: offset
func (_m *Transaction) Offset(offset int) orm.Query {
	ret := _m.Called(offset)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// OrWhere provides a mock function with given fields: query, args
func (_m *Transaction) OrWhere(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value
func (_m *Transaction) Order(value interface{}) orm.Query {
	ret := _m.Called(value)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}) orm.Query); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Pluck provides a mock function with given fields: column, dest
func (_m *Transaction) Pluck(column string, dest interface{}) error {
	ret := _m.Called(column, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(column, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Raw provides a mock function with given fields: sql, values
func (_m *Transaction) Raw(sql string, values ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, sql)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(sql, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *Transaction) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: value
func (_m *Transaction) Save(value interface{}) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *Transaction) Scan(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scopes provides a mock function with given fields: funcs
func (_m *Transaction) Scopes(funcs ...func(orm.Query) orm.Query) orm.Query {
	_va := make([]interface{}, len(funcs))
	for _i := range funcs {
		_va[_i] = funcs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...func(orm.Query) orm.Query) orm.Query); ok {
		r0 = rf(funcs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *Transaction) Select(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Table provides a mock function with given fields: name, args
func (_m *Transaction) Table(name string, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Update provides a mock function with given fields: column, value
func (_m *Transaction) Update(column string, value interface{}) error {
	ret := _m.Called(column, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(column, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updates provides a mock function with given fields: values
func (_m *Transaction) Updates(values interface{}) error {
	ret := _m.Called(values)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *Transaction) Where(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WithTrashed provides a mock function with given fields:
func (_m *Transaction) WithTrashed() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

type NewTransactionT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransaction creates a new instance of Transaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction(t NewTransactionT) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
