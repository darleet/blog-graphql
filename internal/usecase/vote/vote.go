package vote

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

//go:generate mockery --name=Repository
type Repository interface {
	GetArticleVotes(ctx context.Context, articleID string) (int, error)
	GetCommentVotes(ctx context.Context, commentID string) (int, error)
	SetArticleVote(ctx context.Context, userID string, value model.VoteArticle) error
	SetCommentVote(ctx context.Context, userID string, value model.VoteComment) error
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) GetArticleVotes(ctx context.Context, articleID string) (int, error) {
	return uc.repo.GetArticleVotes(ctx, articleID)
}

func (uc *Usecase) GetCommentVotes(ctx context.Context, commentID string) (int, error) {
	return uc.repo.GetCommentVotes(ctx, commentID)
}

func (uc *Usecase) VoteArticle(ctx context.Context, input model.VoteArticle) (int, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return 0, errors.NewUnauthorizedError("VoteUsecase.VoteArticle: unauthenticated, userID is empty")
	}
	err := uc.repo.SetArticleVote(ctx, userID, input)
	if err != nil {
		return 0, err
	}
	return uc.repo.GetArticleVotes(ctx, input.ArticleID)
}

func (uc *Usecase) VoteComment(ctx context.Context, input model.VoteComment) (int, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return 0, errors.NewUnauthorizedError("VoteUsecase.VoteComment: unauthenticated, userID is empty")
	}
	err := uc.repo.SetCommentVote(ctx, userID, input)
	if err != nil {
		return 0, err
	}
	return uc.repo.GetCommentVotes(ctx, input.CommentID)
}
