package comment

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

//go:generate mockery --name=Repository
type Repository interface {
	CreateComment(ctx context.Context, userID string, input model.NewComment) (*model.Comment, error)
	UpdateComment(ctx context.Context, input model.UpdateComment) (*model.Comment, error)
	DeleteComment(ctx context.Context, id string) (bool, error)
	GetReplies(ctx context.Context, articleID string, after *string, sort *model.Sort) ([]*model.Comment, error)
	GetCommentAuthorID(ctx context.Context, id string) (string, error)
}

type Usecase struct {
	repo Repository
	sub  map[string][]chan *model.Comment
}

func NewUsecase(repo Repository) *Usecase {
	sub := make(map[string][]chan *model.Comment)
	return &Usecase{
		repo: repo,
		sub:  sub,
	}
}

func (uc *Usecase) Create(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return nil, errors.NewUnauthorizedError("ArticleUsecase.CreateArticle: unauthenticated, userID is empty")
	}
	comment, err := uc.repo.CreateComment(ctx, userID, input)
	if err != nil {
		return nil, err
	}
	for _, c := range uc.sub[input.ArticleID] {
		c <- comment
	}
	return comment, nil
}

func (uc *Usecase) Update(ctx context.Context, input model.UpdateComment) (*model.Comment, error) {
	isAuthor, err := uc.IsAuthor(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, errors.NewForbiddenError("CommentUsecase.UpdateArticle: you are not the author of this comment")
	}
	return uc.repo.UpdateComment(ctx, input)
}

func (uc *Usecase) Delete(ctx context.Context, id string) (bool, error) {
	isAuthor, err := uc.IsAuthor(ctx, id)
	if err != nil {
		return false, err
	}
	if !isAuthor {
		return false, errors.NewForbiddenError("CommentUsecase.DeleteArticle: you are not the author of this comment")
	}
	return uc.repo.DeleteComment(ctx, id)
}

func (uc *Usecase) GetReplies(ctx context.Context, articleID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	return uc.repo.GetReplies(ctx, articleID, after, sort)
}

func (uc *Usecase) Subscribe(ctx context.Context, articleID string) (<-chan *model.Comment, error) {
	c := make(chan *model.Comment, 1)
	if _, ok := uc.sub[articleID]; !ok {
		uc.sub[articleID] = make([]chan *model.Comment, 0)
	}
	uc.sub[articleID] = append(uc.sub[articleID], c)
	return c, nil
}
