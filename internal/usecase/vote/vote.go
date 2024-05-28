package vote

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

//go:generate mockery --name=Repository
type Repository interface {
	SetArticleVote(ctx context.Context, userID string, input model.VoteArticle) (int, error)
	SetCommentVote(ctx context.Context, userID string, input model.VoteComment) (int, error)
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) VoteArticle(ctx context.Context, input model.VoteArticle) (int, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return 0, errors.NewUnauthorizedError("VoteUsecase.VoteArticle: unauthenticated, userID is empty")
	}
	return uc.repo.SetArticleVote(ctx, userID, input)
}

func (uc *Usecase) VoteComment(ctx context.Context, input model.VoteComment) (int, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return 0, errors.NewUnauthorizedError("VoteUsecase.VoteComment: unauthenticated, userID is empty")
	}
	return uc.repo.SetCommentVote(ctx, userID, input)
}
