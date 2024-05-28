package pg

import (
	"context"
	errs "errors"
	"fmt"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) IsArticleClosed(ctx context.Context, articleID string) (bool, error) {
	q := `SELECT is_closed FROM articles WHERE id = $1`

	var isClosed bool
	err := r.pool.QueryRow(ctx, q, articleID).Scan(&isClosed)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return false, errors.NewNotFoundError(fmt.Errorf("CommentRepository.IsArticleClosed: %w", err))
	} else if err != nil {
		return false, errors.NewInternalServerError(fmt.Errorf("CommentRepository.IsArticleClosed: %w", err))
	}

	return isClosed, nil
}

func (r *Repository) CreateComment(ctx context.Context, userID string,
	input model.NewComment) (*model.Comment, error) {
	q := `
		INSERT INTO comments (article_id, author_id, parent_id, body) 
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	var c model.Comment
	err := r.pool.QueryRow(ctx, q, input.ArticleID, userID, input.ParentID, input.Content).Scan(&c.ID, &c.CreatedAt)

	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("CommentRepository.CreateComment: %w", err))
	}
	c.Content = input.Content
	c.UserID = userID

	return &c, nil
}

func (r *Repository) UpdateComment(ctx context.Context, input model.UpdateComment) (*model.Comment, error) {
	q := `
		UPDATE comments SET body = COALESCE($1, body) 
		WHERE id = $2 RETURNING created_at, author_id, body, (
			SELECT COALESCE(SUM(value), 0) FROM comments_votes WHERE comment_id = $2
		)
	`

	var c model.Comment
	err := r.pool.QueryRow(ctx, q, input.Content, input.ID).Scan(&c.CreatedAt, &c.UserID, &c.Content, &c.Votes)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return nil, errors.NewNotFoundError(fmt.Errorf("CommentRepository.UpdateComment: %w", err))
	}
	c.ID = input.ID

	return &c, nil
}

func (r *Repository) DeleteComment(ctx context.Context, id string) (bool, error) {
	q := "DELETE FROM comments WHERE id = $1"
	_, err := r.pool.Exec(ctx, q, id)
	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return false, errors.NewNotFoundError(fmt.Errorf("CommentRepository.DeleteComment: %w", err))
	}
	if err != nil {
		return false, errors.NewInternalServerError(fmt.Errorf("CommentRepository.DeleteComment: %w", err))
	}
	return true, nil
}

func (r *Repository) GetReplies(ctx context.Context, commentID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	q := `
		SELECT c.id, c.created_at, c.author_id, c.body, 
			COALESCE(SUM(v.value), 0) AS vote_sum
		FROM comments c
		LEFT JOIN comments_votes v ON c.id = v.comment_id
		WHERE c.parent_id = $1 AND c.id > $2
		GROUP BY c.id, c.created_at
		ORDER BY CASE WHEN $3 = 'NEW_DESC' THEN c.created_at END DESC,
		         CASE WHEN $3 = 'NEW_ASC' THEN c.created_at END,
		         CASE WHEN $3 = 'VOTE_DESC' THEN COALESCE(SUM(v.value), 0) END DESC,
		         CASE WHEN $3 = 'VOTE_ASC' THEN COALESCE(SUM(v.value), 0) END
		LIMIT $4
	`

	if after == nil {
		after = new(string)
		*after = "0"
	}

	var c []*model.Comment
	rows, err := r.pool.Query(ctx, q, commentID, after, sort, CommentLimit)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return nil, errors.NewNotFoundError(fmt.Errorf("CommentRepository.GetReplies: %w", err))
	}
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("CommentRepository.GetReplies: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.ID, &comment.CreatedAt, &comment.UserID, &comment.Content, &comment.Votes)
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Errorf("CommentRepository.GetReplies: %w", err))
		}
		c = append(c, &comment)
	}

	return c, nil
}

func (r *Repository) GetCommentAuthorID(ctx context.Context, id string) (string, error) {
	q := "SELECT author_id FROM comments WHERE id = $1"
	var userID string
	err := r.pool.QueryRow(ctx, q, id).Scan(&userID)
	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return "", errors.NewNotFoundError(fmt.Errorf("CommentRepository.GetCommentAuthorID: %w", err))
	}
	if err != nil {
		return "", errors.NewInternalServerError(fmt.Errorf("CommentRepository.GetCommentAuthorID: %w", err))
	}
	return userID, nil
}
