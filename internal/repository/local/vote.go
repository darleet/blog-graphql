package local

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"strconv"
)

func (r *Repository) GetArticleVotes(ctx context.Context, articleID string) (int, error) {
	u, err := strconv.ParseUint(articleID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.GetArticleVotes: invalid articleID")
	}
	return voteSum(r.articles[u].Votes), nil
}

func (r *Repository) GetCommentVotes(ctx context.Context, commentID string) (int, error) {
	u, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("VoteRepository.GetCommentVotes: invalid commentID")
	}
	return voteSum(r.comments[u].Votes), nil
}

func (r *Repository) SetArticleVote(ctx context.Context, userID string, input model.VoteArticle) error {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return errors.NewBadRequestError("VoteRepository.SetArticleVote: invalid userID")
	}
	a, err := strconv.ParseUint(input.ArticleID, 10, 64)
	if err != nil {
		return errors.NewBadRequestError("VoteRepository.SetArticleVote: invalid articleID")
	}
	article, ok := r.articles[a]
	if !ok {
		return errors.NewNotFoundError("VoteRepository.SetArticleVote: vote not found")
	}
	switch input.Value {
	case model.VoteValueNone:
		delete(article.Votes, u)
	case model.VoteValueUp:
		article.Votes[u] = 1
	case model.VoteValueDown:
		article.Votes[u] = -1
	}
	return nil
}

func (r *Repository) SetCommentVote(ctx context.Context, userID string, input model.VoteComment) error {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return errors.NewBadRequestError("VoteRepository.SetCommentVote: invalid userID")
	}
	c, err := strconv.ParseUint(input.CommentID, 10, 64)
	if err != nil {
		return errors.NewBadRequestError("VoteRepository.SetCommentVote: invalid commentID")
	}
	comment, ok := r.comments[c]
	if !ok {
		return errors.NewNotFoundError("VoteRepository.SetCommentVote: vote not found")
	}
	switch input.Value {
	case model.VoteValueNone:
		delete(comment.Votes, u)
	case model.VoteValueUp:
		comment.Votes[u] = 1
	case model.VoteValueDown:
		comment.Votes[u] = -1
	}
	return nil
}
