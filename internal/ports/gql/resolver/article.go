package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
)

// Author is the resolver for the author field.
func (r *articleResolver) Author(ctx context.Context, obj *model.Article) (*model.User, error) {
	return r.users.GetUser(ctx, obj.UserID)
}

// Votes is the resolver for the votes field.
func (r *articleResolver) Votes(ctx context.Context, obj *model.Article) (int, error) {
	return r.votes.GetVotes(ctx, obj.ID)
}

// Comments is the resolver for the comments field.
func (r *articleResolver) Comments(ctx context.Context, obj *model.Article, after *string, sort *model.Sort) ([]*model.Comment, error) {
	return r.articles.GetComments(ctx, obj.ID, after, sort)
}

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	return r.articles.Create(ctx, input)
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	return r.articles.Update(ctx, input)
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, id string) (bool, error) {
	return r.articles.Delete(ctx, id)
}

// ArticlesList is the resolver for the articlesList field.
func (r *queryResolver) ArticlesList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error) {
	return r.articles.GetList(ctx, after, sort)
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, articleID string) (*model.Article, error) {
	return r.articles.GetArticle(ctx, articleID)
}

// Article returns runtime.ArticleResolver implementation.
func (r *Resolver) Article() runtime.ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }
