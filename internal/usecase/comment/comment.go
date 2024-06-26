package comment

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
	"sync"
)

//go:generate mockery --name=Repository
type Repository interface {
	IsArticleClosed(ctx context.Context, id string) (bool, error)
	CreateComment(ctx context.Context, userID string, input model.NewComment) (*model.Comment, error)
	UpdateComment(ctx context.Context, input model.UpdateComment) (*model.Comment, error)
	DeleteComment(ctx context.Context, id string) (bool, error)
	GetReplies(ctx context.Context, commentID string, after *string, sort *model.Sort) ([]*model.Comment, error)
	GetCommentAuthorID(ctx context.Context, id string) (string, error)
}

type Usecase struct {
	repo Repository
	mu   sync.RWMutex
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
		return nil, errors.NewUnauthorizedError("ArticleUsecase.CreateArticle: " +
			"unauthenticated, userID is empty")
	}

	isClosed, err := uc.repo.IsArticleClosed(ctx, input.ArticleID)
	if err != nil {
		return nil, err
	}
	if isClosed {
		return nil, errors.NewForbiddenError("CommentUsecase.CreateComment: article is closed")
	}

	comment, err := uc.repo.CreateComment(ctx, userID, input)
	if err != nil {
		return nil, err
	}

	go func() {
		uc.mu.RLock()
		defer uc.mu.RUnlock()
		for _, c := range uc.sub[input.ArticleID] {
			c <- comment
		}
	}()

	return comment, nil
}

func (uc *Usecase) Update(ctx context.Context, input model.UpdateComment) (*model.Comment, error) {
	isAuthor, err := uc.IsAuthor(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, errors.NewForbiddenError("CommentUsecase.UpdateArticle: " +
			"you are not the author of this comment")
	}
	return uc.repo.UpdateComment(ctx, input)
}

func (uc *Usecase) Delete(ctx context.Context, id string) (bool, error) {
	isAuthor, err := uc.IsAuthor(ctx, id)
	if err != nil {
		return false, err
	}
	if !isAuthor {
		return false, errors.NewForbiddenError("CommentUsecase.DeleteArticle: " +
			"you are not the author of this comment")
	}
	return uc.repo.DeleteComment(ctx, id)
}

func (uc *Usecase) GetReplies(ctx context.Context, commentID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	return uc.repo.GetReplies(ctx, commentID, after, sort)
}

func (uc *Usecase) Subscribe(ctx context.Context, articleID string) (<-chan *model.Comment, error) {
	c := make(chan *model.Comment, 1)
	uc.mu.Lock()
	defer uc.mu.Unlock()
	if _, ok := uc.sub[articleID]; !ok {
		uc.sub[articleID] = make([]chan *model.Comment, 0)
	}
	uc.sub[articleID] = append(uc.sub[articleID], c)
	return c, nil
}
