package pg

import (
	"context"
	"fmt"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
)

func parseValue(value model.VoteValue) int {
	switch value {
	case model.VoteValueUp:
		return 1
	case model.VoteValueDown:
		return -1
	default:
		return 0
	}
}

func (r *Repository) SetArticleVote(ctx context.Context, userID string, input model.VoteArticle) (int, error) {
	q := `
		INSERT INTO articles_votes (article_id, author_id, value) VALUES ($1, $2, $3)  
		ON CONFLICT (article_id, author_id) DO UPDATE SET value = $3 
		RETURNING (SELECT COALESCE(SUM(value), 0) FROM articles_votes WHERE article_id = $1)
	`

	value := parseValue(input.Value)

	var votes int
	err := r.pool.QueryRow(ctx, q, input.ArticleID, userID, value).Scan(&votes)

	if err != nil {
		return 0, errors.NewInternalServerError(fmt.Errorf("VoteRepository.SetArticleVote: %w", err))
	}
	return votes + value, nil
}

func (r *Repository) SetCommentVote(ctx context.Context, userID string, input model.VoteComment) (int, error) {
	q := `
		INSERT INTO comments_votes (comment_id, author_id, value) VALUES ($1, $2, $3)  
		ON CONFLICT (comment_id, author_id) DO UPDATE SET value = $3 
		RETURNING (SELECT COALESCE(SUM(value), 0) FROM comments_votes WHERE comment_id = $1)
	`

	value := parseValue(input.Value)

	var votes int
	err := r.pool.QueryRow(ctx, q, input.CommentID, userID, value).Scan(&votes)

	if err != nil {
		return 0, errors.NewInternalServerError(fmt.Errorf(
			"VoteRepository.SetCommentVote: %w", err))
	}
	return votes + value, nil
}
