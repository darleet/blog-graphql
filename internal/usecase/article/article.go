package article

import (
	"context"

	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, userID string, input model.NewArticle) (*model.Article, error)
	Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error)
	Get(ctx context.Context, articleID string) (*model.Article, error)
	GetComments(ctx context.Context, articleID string, after *string, sort *model.Sort) ([]*model.Comment, error)
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) Create(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return nil, errors.NewUnauthorizedError("ArticleUsecase.Create: unauthenticated, userID is empty")
	}
	return uc.repo.Create(ctx, userID, input)
}

func (uc *Usecase) Update(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	isAuthor, err := uc.IsAuthor(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, errors.NewForbiddenError("ArticleUsecase.Update: you are not the author of this article")
	}
	return uc.repo.Update(ctx, input)
}

func (uc *Usecase) Delete(ctx context.Context, id string) (bool, error) {
	isAuthor, err := uc.IsAuthor(ctx, id)
	if err != nil {
		return false, err
	}
	if !isAuthor {
		return false, errors.NewForbiddenError("ArticleUsecase.Delete: you are not the author of this article")
	}
	return uc.repo.Delete(ctx, id)
}

func (uc *Usecase) GetList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error) {
	return uc.repo.GetList(ctx, after, sort)
}

func (uc *Usecase) Get(ctx context.Context, articleID string) (*model.Article, error) {
	return uc.repo.Get(ctx, articleID)
}

func (uc *Usecase) GetComments(ctx context.Context, articleID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	return uc.repo.GetComments(ctx, articleID, after, sort)
}
