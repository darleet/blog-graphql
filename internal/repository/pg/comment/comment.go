package comment

import (
	"context"
	errs "errors"
	"fmt"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Close(ctx context.Context) error {
	return r.conn.Close(ctx)
}

func (r *Repository) CreateComment(ctx context.Context, userID string,
	input model.NewComment) (*model.Comment, error) {
	q := `
		INSERT INTO comments (article_id, author_id, parent_id, body) 
		VALUES ($1, $2, $3) RETURNING id, created_at
	`

	var c model.Comment
	err := r.conn.QueryRow(ctx, q, input.ArticleID, userID, input.ParentID, input.Content).Scan(&c.ID, &c.CreatedAt)

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
	err := r.conn.QueryRow(ctx, q, input.Content, input.ID).Scan(&c.CreatedAt, &c.UserID, &c.Content, &c.Votes)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return nil, errors.NewNotFoundError(fmt.Errorf("CommentRepository.UpdateComment: %w", err))
	}
	c.ID = input.ID

	return &c, nil
}

func (r *Repository) DeleteComment(ctx context.Context, id string) (bool, error) {
	q := "DELETE FROM comments WHERE id = $1"
	_, err := r.conn.Exec(ctx, q, id)
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
		WHERE c.parent_id = $1
		GROUP BY c.id, c.created_at
		ORDER BY CASE WHEN $2 = 'NEW_DESC' THEN c.created_at END DESC,
		         CASE WHEN $2 = 'NEW_ASC' THEN c.created_at END,
		         CASE WHEN $2 = 'VOTE_DESC' THEN vote_sum END DESC,
		         CASE WHEN $2 = 'VOTE_ASC' THEN vote_sum END
	`

	var c []*model.Comment
	rows, err := r.conn.Query(ctx, q, commentID)

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
	err := r.conn.QueryRow(ctx, q, id).Scan(&userID)
	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return "", errors.NewNotFoundError(fmt.Errorf("CommentRepository.GetCommentAuthorID: %w", err))
	}
	if err != nil {
		return "", errors.NewInternalServerError(fmt.Errorf("CommentRepository.GetCommentAuthorID: %w", err))
	}
	return userID, nil
}
