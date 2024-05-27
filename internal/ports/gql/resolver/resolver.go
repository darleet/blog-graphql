package resolver

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
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
