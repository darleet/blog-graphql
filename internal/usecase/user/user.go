package user

import (
	"context"
	"github.com/darleet/blog-graphql/internal/middleware/loader"
	"github.com/darleet/blog-graphql/internal/model"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (uc *Usecase) GetUser(ctx context.Context, userID string) (*model.User, error) {
	return loader.GetUser(ctx, userID)
}
