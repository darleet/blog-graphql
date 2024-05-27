package local

import (
	"github.com/darleet/blog-graphql/internal/repository/local/schema"
)

const ArticleLimit = 5
const CommentLimit = 5

type Repository struct {
	articles       map[uint64]*schema.Article
	comments       map[uint64]*schema.Comment
	serialArticles uint64
	serialComments uint64
}

func NewRepository() *Repository {
	return &Repository{
		articles:       make(map[uint64]*schema.Article),
		comments:       make(map[uint64]*schema.Comment),
		serialArticles: 1,
		serialComments: 1,
	}
}
