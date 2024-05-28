package local

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"strconv"
)

func (r *Repository) SetArticleVote(ctx context.Context, userID string, input model.VoteArticle) (int, error) {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.SetArticleVote: invalid userID")
	}
	a, err := strconv.ParseUint(input.ArticleID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.SetArticleVote: invalid articleID")
	}
	article, ok := r.articles[a]
	if !ok {
		return 0, errors.NewNotFoundError("VoteRepository.SetArticleVote: vote not found")
	}
	switch input.Value {
	case model.VoteValueNone:
		delete(article.Votes, u)
	case model.VoteValueUp:
		article.Votes[u] = 1
	case model.VoteValueDown:
		article.Votes[u] = -1
	}
	return article.Votes[u], nil
}

func (r *Repository) SetCommentVote(ctx context.Context, userID string, input model.VoteComment) (int, error) {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.SetCommentVote: invalid userID")
	}
	c, err := strconv.ParseUint(input.CommentID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.SetCommentVote: invalid commentID")
	}
	comment, ok := r.comments[c]
	if !ok {
		return 0, errors.NewNotFoundError("VoteRepository.SetCommentVote: vote not found")
	}
	switch input.Value {
	case model.VoteValueNone:
		delete(comment.Votes, u)
	case model.VoteValueUp:
		comment.Votes[u] = 1
	case model.VoteValueDown:
		comment.Votes[u] = -1
	}
	return comment.Votes[u], nil
}
