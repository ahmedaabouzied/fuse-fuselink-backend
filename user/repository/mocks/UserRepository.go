// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) Create(ctx context.Context, _a1 *entities.User) (*entities.User, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *entities.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, userID
func (_m *UserRepository) Delete(ctx context.Context, userID string) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByCognitoID provides a mock function with given fields: ctx, cognitoID
func (_m *UserRepository) GetByCognitoID(ctx context.Context, cognitoID string) (*entities.User, error) {
	ret := _m.Called(ctx, cognitoID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, cognitoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cognitoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *UserRepository) GetByID(ctx context.Context, ID string) (*entities.User, error) {
	ret := _m.Called(ctx, ID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *UserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) Update(ctx context.Context, _a1 *entities.User) (*entities.User, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *entities.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
