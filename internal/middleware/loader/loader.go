package loader

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/vikstrous/dataloadgen"
	"net/http"
	"time"
)

type UserRepository interface {
	GetUsers(ctx context.Context, userIDs []string) ([]*model.User, []error)
}

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// userReader reads Users from a database
type userReader struct {
	repo UserRepository
}

// getUsers returns a batch function that can retrieve many users by ID for use in a dataloader
func (r *userReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	return r.repo.GetUsers(ctx, userIDs)
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(repo UserRepository) *Loaders {
	r := &userReader{repo: repo}
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(r.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(repo UserRepository, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(repo)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns multiple users by id efficiently
func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
