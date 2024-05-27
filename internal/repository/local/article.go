package local

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/repository/local/schema"
	"github.com/darleet/blog-graphql/pkg/errors"
	sorting "sort"
	"strconv"
	"time"
)

func (r *Repository) CreateArticle(ctx context.Context, userID string,
	input model.NewArticle) (*model.Article, error) {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("ArticleRepository.CreateArticle: invalid userID")
	}

	article := &schema.Article{
		ID:        r.serialArticles,
		Title:     input.Title,
		Content:   input.Content,
		UserID:    u,
		Comments:  make([]uint64, 0),
		Votes:     make(map[uint64]int),
		CreatedAt: time.Now(),
	}

	m := &model.Article{
		ID:        strconv.FormatUint(r.serialArticles, 10),
		Title:     article.Title,
		Content:   article.Content,
		UserID:    strconv.FormatUint(article.UserID, 10),
		IsClosed:  article.IsClosed,
		CreatedAt: article.CreatedAt,
	}

	r.articles[r.serialArticles] = article
	r.serialArticles++
	return m, nil
}

func (r *Repository) UpdateArticle(ctx context.Context, input model.UpdateArticle) (*model.Article, error) {
	u, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("ArticleRepository.UpdateArticle: invalid articleID")
	}
	article, ok := r.articles[u]
	if !ok {
		return nil, errors.NewNotFoundError("ArticleRepository.UpdateArticle: article not found")
	}

	if input.Title != nil {
		article.Title = *input.Title
	}
	if input.Content != nil {
		article.Content = *input.Content
	}
	if input.IsClosed != nil {
		article.IsClosed = *input.IsClosed
	}

	m := &model.Article{
		ID:        input.ID,
		Title:     article.Title,
		Content:   article.Content,
		UserID:    strconv.FormatUint(article.UserID, 10),
		IsClosed:  article.IsClosed,
		CreatedAt: article.CreatedAt,
	}
	return m, nil
}

func (r *Repository) DeleteArticle(ctx context.Context, id string) (bool, error) {
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return false, errors.NewBadRequestError("ArticleRepository.DeleteArticle: invalid articleID")
	}
	_, ok := r.articles[u]
	if !ok {
		return false, errors.NewNotFoundError("ArticleRepository.DeleteArticle: article not found")
	}
	delete(r.articles, u)
	return true, nil
}

func (r *Repository) GetArticlesList(ctx context.Context, after *string,
	sort *model.Sort) ([]*model.Article, error) {
	var cursor uint64
	var err error

	if after != nil {
		cursor, err = strconv.ParseUint(*after, 10, 64)
		if err != nil {
			return nil, errors.NewBadRequestError("ArticleRepository.GetArticlesList: invalid after")
		}
	}

	articles := make([]*schema.Article, 0)

	for _, v := range r.articles {
		if v.ID > cursor {
			articles = append(articles, v)
		}
	}

	if sort != nil {
		switch *sort {
		case model.SortNewAsc:
			sorting.Slice(articles, func(i, j int) bool {
				return articles[i].CreatedAt.Before(articles[j].CreatedAt)
			})
		case model.SortNewDesc:
			sorting.Slice(articles, func(i, j int) bool {
				return articles[i].CreatedAt.After(articles[j].CreatedAt)
			})
		case model.SortTopAsc:
			sorting.Slice(articles, func(i, j int) bool {
				return voteSum(articles[i].Votes) < voteSum(articles[j].Votes)
			})
		case model.SortTopDesc:
			sorting.Slice(articles, func(i, j int) bool {
				return voteSum(articles[i].Votes) > voteSum(articles[j].Votes)
			})
		}
	}

	res := make([]*model.Article, 0, ArticleLimit)
	for k := 0; k < ArticleLimit && k < len(articles); k++ {
		res = append(res, &model.Article{
			ID:        strconv.FormatUint(articles[k].ID, 10),
			Title:     articles[k].Title,
			Content:   articles[k].Content,
			UserID:    strconv.FormatUint(articles[k].UserID, 10),
			IsClosed:  articles[k].IsClosed,
			CreatedAt: articles[k].CreatedAt,
		})
	}

	return res, nil
}

func (r *Repository) GetArticle(ctx context.Context, articleID string) (*model.Article, error) {
	u, err := strconv.ParseUint(articleID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("ArticleRepository.GetArticle: invalid articleID")
	}
	article, ok := r.articles[u]
	if !ok {
		return nil, errors.NewNotFoundError("ArticleRepository.GetArticle: article not found")
	}
	m := &model.Article{
		ID:        articleID,
		Title:     article.Title,
		Content:   article.Content,
		UserID:    strconv.FormatUint(article.UserID, 10),
		IsClosed:  article.IsClosed,
		CreatedAt: article.CreatedAt,
	}
	return m, nil
}

func (r *Repository) GetComments(ctx context.Context, articleID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	u, err := strconv.ParseUint(articleID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("ArticleRepository.GetComments: invalid articleID")
	}
	article, ok := r.articles[u]
	if !ok {
		return nil, errors.NewNotFoundError("ArticleRepository.GetComments: article not found")
	}

	var cursor uint64
	if after != nil {
		cursor, err = strconv.ParseUint(*after, 10, 64)
		if err != nil {
			return nil, errors.NewBadRequestError("ArticleRepository.GetComments: invalid after")
		}
	}

	comments := make([]*schema.Comment, 0)
	for _, v := range article.Comments {
		if v > cursor {
			comments = append(comments, r.comments[v])
		}
	}

	if sort != nil {
		switch *sort {
		case model.SortNewAsc:
			sorting.Slice(comments, func(i, j int) bool {
				return comments[i].CreatedAt.Before(comments[j].CreatedAt)
			})
		case model.SortNewDesc:
			sorting.Slice(comments, func(i, j int) bool {
				return comments[i].CreatedAt.After(comments[j].CreatedAt)
			})
		case model.SortTopAsc:
			sorting.Slice(comments, func(i, j int) bool {
				return voteSum(comments[i].Votes) < voteSum(comments[j].Votes)
			})
		case model.SortTopDesc:
			sorting.Slice(comments, func(i, j int) bool {
				return voteSum(comments[i].Votes) > voteSum(comments[j].Votes)
			})
		}
	}

	res := make([]*model.Comment, 0, CommentLimit)
	for k := 0; k < CommentLimit && k < len(comments); k++ {
		res = append(res, &model.Comment{
			ID:        strconv.FormatUint(comments[k].ID, 10),
			Content:   comments[k].Content,
			UserID:    strconv.FormatUint(comments[k].UserID, 10),
			CreatedAt: comments[k].CreatedAt,
		})
	}

	return res, nil
}

func (r *Repository) GetArticleAuthorID(ctx context.Context, id string) (string, error) {
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return "", errors.NewBadRequestError("ArticleRepository.GetArticleAuthorID: invalid articleID")
	}
	article, ok := r.articles[u]
	if !ok {
		return "", errors.NewNotFoundError("ArticleRepository.GetArticleAuthorID: article not found")
	}
	return strconv.FormatUint(article.UserID, 10), nil
}
