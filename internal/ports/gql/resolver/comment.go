package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
)

// Author is the resolver for the author field.
func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Replies - replies"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, input *model.UpdateComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: UpdateComment - updateComment"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteComment - deleteComment"))
}

// ListenComments is the resolver for the listenComments field.
func (r *subscriptionResolver) ListenComments(ctx context.Context, articleID string) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: ListenComments - listenComments"))
}

// Comment returns runtime.CommentResolver implementation.
func (r *Resolver) Comment() runtime.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
