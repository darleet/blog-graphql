package comment

import (
	"context"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

func (uc *Usecase) IsAuthor(ctx context.Context, commentID string) (bool, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return false, errors.NewUnauthorizedError("CommentUsecase.IsAuthor: unauthenticated, userID is empty")
	}
	authorID, err := uc.repo.GetAuthorID(ctx, commentID)
	if err != nil {
		return false, err
	}
	return userID == authorID, nil
}
