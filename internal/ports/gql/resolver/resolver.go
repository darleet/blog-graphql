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
	GetArticle(ctx context.Context, articleID string) (*model.Article, error)
	GetComments(ctx context.Context, articleID string, after *string, sort *model.Sort) ([]*model.Comment, error)
}

type CommentUsecase interface {
	Create(ctx context.Context, input model.NewComment) (*model.Comment, error)
	Update(ctx context.Context, input model.UpdateComment) (*model.Comment, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetReplies(ctx context.Context, id string, after *string, sort *model.Sort) ([]*model.Comment, error)
	Subscribe(ctx context.Context, articleID string) (<-chan *model.Comment, error)
}

type UserUsecase interface {
	GetUser(ctx context.Context, userID string) (*model.User, error)
}

type Resolver struct {
	articles ArticleUsecase
	comments CommentUsecase
	users    UserUsecase
}

func NewRootResolvers(articles ArticleUsecase, comments CommentUsecase, users UserUsecase) runtime.Config {
	c := runtime.Config{
		Resolvers: &Resolver{
			articles: articles,
			comments: comments,
			users:    users,
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
