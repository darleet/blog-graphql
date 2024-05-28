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

// SetArticleVote provides a mock function with given fields: ctx, userID, input
func (_m *Repository) SetArticleVote(ctx context.Context, userID string, input model.VoteArticle) (int, error) {
	ret := _m.Called(ctx, userID, input)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.VoteArticle) (int, error)); ok {
		return rf(ctx, userID, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.VoteArticle) int); ok {
		r0 = rf(ctx, userID, input)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.VoteArticle) error); ok {
		r1 = rf(ctx, userID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCommentVote provides a mock function with given fields: ctx, userID, input
func (_m *Repository) SetCommentVote(ctx context.Context, userID string, input model.VoteComment) (int, error) {
	ret := _m.Called(ctx, userID, input)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.VoteComment) (int, error)); ok {
		return rf(ctx, userID, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.VoteComment) int); ok {
		r0 = rf(ctx, userID, input)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.VoteComment) error); ok {
		r1 = rf(ctx, userID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
