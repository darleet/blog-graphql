package article

import (
	"context"
	"github.com/darleet/blog-graphql/pkg/utils"
)

func (a *Usecase) IsAuthor(ctx context.Context, articleID string) (bool, error) {
	userID := utils.GetUserID(ctx)
	article, err := a.repo.GetByID(ctx, articleID)
	if err != nil {
		return false, err
	}
	return userID == article.Author.ID, nil
}
