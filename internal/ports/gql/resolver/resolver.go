package resolver

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ArticleUsecase interface {
	Create(ctx context.Context, input model.NewArticle) (*model.Article, error)
	Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error)
	GetByID(ctx context.Context, articleID string) (*model.Article, error)
}

type Resolver struct {
	articles ArticleUsecase
}

func NewRootResolvers(articles ArticleUsecase) runtime.Config {
	c := runtime.Config{
		Resolvers: &Resolver{
			articles: articles,
		},
	}

	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{},
		next graphql.Resolver) (res interface{}, err error) {
		userID := utils.GetUserID(ctx)
		if userID != "" {
			return next(ctx)
		} else {
			return nil, errors.NewUnauthorizedError("you are unauthorized to perform this action")
		}
	}
	return c
}
