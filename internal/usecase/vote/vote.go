package vote

import (
	"context"
	errs "errors"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

//go:generate mockery --name=Repository
type Repository interface {
	GetUserVote(ctx context.Context, userID, articleID string) (model.VoteValue, error)
	GetVotes(ctx context.Context, articleID string) (int, error)
	SetVote(ctx context.Context, userID string, value model.Vote) error
	InsertVote(ctx context.Context, userID string, value model.Vote) error
}

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) GetVotes(ctx context.Context, articleID string) (int, error) {
	return uc.repo.GetVotes(ctx, articleID)
}

func (uc *Usecase) ProcessVote(ctx context.Context, input model.Vote) (int, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return 0, errors.NewUnauthorizedError("VoteUsecase.ProcessVote: unauthenticated, userID is empty")
	}

	vote, err := uc.repo.GetUserVote(ctx, userID, input.ArticleID)
	if err != nil && errs.Is(err, errors.NotFound) {
		err = uc.repo.InsertVote(ctx, userID, input)
	} else if err != nil {
		return 0, err
	} else if vote != input.Value {
		err = uc.repo.SetVote(ctx, userID, input)
	}

	if err != nil {
		return 0, err
	}
	return uc.repo.GetVotes(ctx, input.ArticleID)
}
