// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	entity "rpl-sixmath/entity"

	mock "github.com/stretchr/testify/mock"

	model "rpl-sixmath/model"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: userId
func (_m *UserRepository) DeleteUser(userId int) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserAll provides a mock function with given fields:
func (_m *UserRepository) FindUserAll() ([]entity.User, error) {
	ret := _m.Called()

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserById provides a mock function with given fields: userId
func (_m *UserRepository) FindUserById(userId int) (entity.User, error) {
	ret := _m.Called(userId)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(int) entity.User); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindUserByUsername(username string) (entity.User, error) {
	ret := _m.Called(username)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataUser provides a mock function with given fields: month, year
func (_m *UserRepository) GetDataUser(month string, year string) (model.GetTotalUser, error) {
	ret := _m.Called(month, year)

	var r0 model.GetTotalUser
	if rf, ok := ret.Get(0).(func(string, string) model.GetTotalUser); ok {
		r0 = rf(month, year)
	} else {
		r0 = ret.Get(0).(model.GetTotalUser)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(month, year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: user
func (_m *UserRepository) InsertUser(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *UserRepository) UpdateUser(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewUserRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t NewUserRepositoryT) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
