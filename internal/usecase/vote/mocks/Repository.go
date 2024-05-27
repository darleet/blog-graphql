// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/darleet/blog-graphql/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserVote provides a mock function with given fields: ctx, userID, articleID
func (_m *Repository) GetUserVote(ctx context.Context, userID string, articleID string) (model.VoteValue, error) {
	ret := _m.Called(ctx, userID, articleID)

	var r0 model.VoteValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (model.VoteValue, error)); ok {
		return rf(ctx, userID, articleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) model.VoteValue); ok {
		r0 = rf(ctx, userID, articleID)
	} else {
		r0 = ret.Get(0).(model.VoteValue)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVotes provides a mock function with given fields: ctx, articleID
func (_m *Repository) GetVotes(ctx context.Context, articleID string) (int, error) {
	ret := _m.Called(ctx, articleID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, articleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, articleID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertVote provides a mock function with given fields: ctx, userID, value
func (_m *Repository) InsertVote(ctx context.Context, userID string, value model.Vote) error {
	ret := _m.Called(ctx, userID, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Vote) error); ok {
		r0 = rf(ctx, userID, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetVote provides a mock function with given fields: ctx, userID, value
func (_m *Repository) SetVote(ctx context.Context, userID string, value model.Vote) error {
	ret := _m.Called(ctx, userID, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Vote) error); ok {
		r0 = rf(ctx, userID, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
