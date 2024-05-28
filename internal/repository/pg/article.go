package pg

import (
	"context"
	errs "errors"
	"fmt"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateArticle(ctx context.Context, userID string,
	input model.NewArticle) (*model.Article, error) {

	q := "INSERT INTO articles (author_id, title, body, is_closed) VALUES ($1, $2, $3, $4) RETURNING id, created_at"
	var a model.Article

	err := r.pool.QueryRow(ctx, q, userID, input.Title, input.Content, input.IsClosed).Scan(&a.ID, &a.CreatedAt)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.CreateArticle: %w", err))
	}

	a.Title = input.Title
	a.Content = input.Content
	a.IsClosed = input.IsClosed
	a.UserID = userID

	return &a, nil
}

func (r *Repository) UpdateArticle(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	q := `
		UPDATE articles SET title = COALESCE($1, title), 
            body = COALESCE($2, body), 
            is_closed = COALESCE($3, is_closed) 
        WHERE id = $4 RETURNING title, body, is_closed, created_at, author_id, (
            SELECT COALESCE(SUM(value), 0) FROM articles_votes WHERE article_id = $4
        )
	`

	var a model.Article
	err := r.pool.QueryRow(ctx, q, input.Title, input.Content, input.IsClosed).
		Scan(&a.Title, &a.Content, &a.IsClosed, &a.CreatedAt, &a.Votes, &a.UserID)
	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return nil, errors.NewNotFoundError(fmt.Errorf("ArticleRepository.UpdateArticle: %w", err))
	} else if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.UpdateArticle: %w", err))
	}

	a.ID = input.ID

	return &a, nil
}

func (r *Repository) DeleteArticle(ctx context.Context, id string) (bool, error) {
	q := `DELETE FROM articles WHERE id = $1`

	_, err := r.pool.Exec(ctx, q, id)
	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return false, errors.NewNotFoundError(fmt.Errorf("ArticleRepository.DeleteArticle: %w", err))
	} else if err != nil {
		return false, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.DeleteArticle: %w", err))
	}

	return true, nil
}

func (r *Repository) GetArticlesList(ctx context.Context, after *string, sort *model.Sort) ([]*model.Article, error) {
	query := `
		SELECT a.id, a.title, a.body, a.author_id, a.is_closed, a.created_at,
			COALESCE(SUM(c.value), 0) AS vote_sum
		FROM articles a
		LEFT JOIN articles_votes c ON a.id = c.article_id
		WHERE a.id > $1
		GROUP BY a.id, a.created_at
		ORDER BY CASE WHEN $2 = 'NEW_DESC' THEN a.created_at END DESC,
				 CASE WHEN $2 = 'NEW_ASC' THEN a.created_at END,
				 CASE WHEN $2 = 'TOP_ASC' THEN vote_sum END DESC,
				 CASE WHEN $2 = 'TOP_DESC' THEN vote_sum END
		LIMIT $3
	`

	var articles []*model.Article
	rows, err := r.pool.Query(ctx, query, after, sort, ArticleLimit)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetArticlesList: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var article model.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.UserID,
			&article.IsClosed,
			&article.CreatedAt,
			&article.Votes,
		)
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetArticlesList: %w", err))
		}
		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetArticlesList: %w", err))
	}

	return articles, nil
}

func (r *Repository) GetArticle(ctx context.Context, articleID string) (*model.Article, error) {
	q := `
		SELECT *, 
		       (SELECT COALESCE(SUM(value), 0) AS vote_sum FROM articles_votes WHERE article_id = $1) 
		FROM articles WHERE id = $1
	`

	var article model.Article
	err := r.pool.QueryRow(ctx, q, articleID).Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.UserID,
		&article.IsClosed,
		&article.CreatedAt,
		&article.Votes,
	)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return nil, errors.NewNotFoundError(fmt.Errorf("ArticleRepository.GetArticle: %w", err))
	}
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetArticle: %w", err))
	}

	return &article, nil
}

func (r *Repository) GetComments(ctx context.Context, articleID string, after *string, sort *model.Sort) ([]*model.Comment, error) {
	q := `
		SELECT c.id, c.body, c.author_id, c.created_at,
			COALESCE(SUM(v.value), 0) AS vote_sum
		FROM comments c
		LEFT JOIN comments_votes v ON c.id = v.comment_id
		WHERE c.article_id = $1
		GROUP BY c.id, c.created_at
		ORDER BY CASE WHEN $2 = 'NEW_DESC' THEN c.created_at END DESC,
				 CASE WHEN $2 = 'NEW_ASC' THEN c.created_at END,
				 CASE WHEN $2 = 'TOP_ASC' THEN vote_sum END DESC,
				 CASE WHEN $2 = 'TOP_DESC' THEN vote_sum END
		LIMIT $3
	`

	var comments []*model.Comment
	rows, err := r.pool.Query(ctx, q, articleID, sort, CommentLimit)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetComments: %w", err))
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.UserID,
			&comment.CreatedAt,
			&comment.Votes,
		)
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetComments: %w", err))
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetComments: %w", err))
	}

	return comments, nil
}

func (r *Repository) GetArticleAuthorID(ctx context.Context, id string) (string, error) {
	q := "SELECT author_id FROM articles WHERE id = $1"

	var authorID string
	err := r.pool.QueryRow(ctx, q, id).Scan(&authorID)

	if err != nil && errs.Is(err, pgx.ErrNoRows) {
		return "", errors.NewNotFoundError(fmt.Errorf("ArticleRepository.GetArticleAuthorID: %w", err))
	} else if err != nil {
		return "", errors.NewInternalServerError(fmt.Errorf("ArticleRepository.GetArticleAuthorID: %w", err))
	}

	return authorID, nil
}
