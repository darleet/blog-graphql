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

// Create provides a mock function with given fields: ctx, userID, input
func (_m *Repository) Create(ctx context.Context, userID string, input model.NewArticle) (*model.Article, error) {
	ret := _m.Called(ctx, userID, input)

	var r0 *model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.NewArticle) (*model.Article, error)); ok {
		return rf(ctx, userID, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.NewArticle) *model.Article); ok {
		r0 = rf(ctx, userID, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.NewArticle) error); ok {
		r1 = rf(ctx, userID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, articleID
func (_m *Repository) Get(ctx context.Context, articleID string) (*model.Article, error) {
	ret := _m.Called(ctx, articleID)

	var r0 *model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Article, error)); ok {
		return rf(ctx, articleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Article); ok {
		r0 = rf(ctx, articleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, articleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorID provides a mock function with given fields: ctx, id
func (_m *Repository) GetAuthorID(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetComments provides a mock function with given fields: ctx, articleID, after, sort
func (_m *Repository) GetComments(ctx context.Context, articleID string, after *string, sort *model.Sort) ([]*model.Comment, error) {
	ret := _m.Called(ctx, articleID, after, sort)

	var r0 []*model.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *model.Sort) ([]*model.Comment, error)); ok {
		return rf(ctx, articleID, after, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *model.Sort) []*model.Comment); ok {
		r0 = rf(ctx, articleID, after, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *string, *model.Sort) error); ok {
		r1 = rf(ctx, articleID, after, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: ctx, after, sort
func (_m *Repository) GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error) {
	ret := _m.Called(ctx, after, sort)

	var r0 []*model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *string, *model.Sort) ([]*model.Article, error)); ok {
		return rf(ctx, after, sort)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *string, *model.Sort) []*model.Article); ok {
		r0 = rf(ctx, after, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *string, *model.Sort) error); ok {
		r1 = rf(ctx, after, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, input
func (_m *Repository) Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateArticle) (*model.Article, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateArticle) *model.Article); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateArticle) error); ok {
		r1 = rf(ctx, input)
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
