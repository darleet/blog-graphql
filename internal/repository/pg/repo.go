package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const ArticleLimit = 5
const CommentLimit = 5

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
