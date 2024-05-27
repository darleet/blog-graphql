package article

import (
	"context"

	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

type Repository interface {
	Create(ctx context.Context, userID string, input model.NewArticle) (*model.Article, error)
	Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error)
	GetByID(ctx context.Context, articleID string) (*model.Article, error)
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (a *Usecase) Create(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	userID := utils.GetUserID(ctx)
	return a.repo.Create(ctx, userID, input)
}

func (a *Usecase) Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	isAuthor, err := a.IsAuthor(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, errors.NewForbiddenError("you are not the author of this article")
	}
	return a.repo.Update(ctx, input)
}

func (a *Usecase) Delete(ctx context.Context, id string) (bool, error) {
	isAuthor, err := a.IsAuthor(ctx, id)
	if err != nil {
		return false, err
	}
	if !isAuthor {
		return false, errors.NewForbiddenError("you are not the author of this article")
	}
	return a.repo.Delete(ctx, id)
}

func (a *Usecase) GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error) {
	return a.repo.GetList(ctx, after, sort)
}

func (a *Usecase) GetByID(ctx context.Context, articleID string) (*model.Article, error) {
	return a.repo.GetByID(ctx, articleID)
}
