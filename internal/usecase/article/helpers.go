package article

import (
	"context"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
)

func (a *Usecase) IsAuthor(ctx context.Context, articleID string) (bool, error) {
	userID := utils.GetUserID(ctx)
	if userID == "" {
		return false, errors.NewUnauthorizedError("ArticleUsecase.IsAuthor: unauthenticated, userID is empty")
	}
	article, err := a.repo.GetByID(ctx, articleID)
	if err != nil {
		return false, err
	}
	return userID == article.Author.ID, nil
}
