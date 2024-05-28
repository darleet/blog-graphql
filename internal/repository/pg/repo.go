package pg

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

const ArticleLimit = 5
const CommentLimit = 5

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Close() {
	r.pool.Close()
}
