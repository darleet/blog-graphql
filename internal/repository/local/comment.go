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

func (r *Repository) CreateComment(ctx context.Context, userID string, input model.NewComment) (*model.Comment, error) {
	u, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("CommentRepository.CreateComment: invalid userID")
	}

	a, err := strconv.ParseUint(input.ArticleID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("CommentRepository.CreateComment: invalid articleID")
	}

	var p uint64
	if input.ParentID != nil {
		p, err = strconv.ParseUint(*input.ParentID, 10, 64)
		if err != nil {
			return nil, errors.NewBadRequestError("CommentRepository.CreateComment: invalid parentID")
		}
	}

	comment := &schema.Comment{
		ID:        r.serialComments,
		Content:   input.Content,
		Parent:    p,
		UserID:    u,
		Votes:     make(map[uint64]int),
		CreatedAt: time.Now(),
	}

	m := &model.Comment{
		ID:        strconv.FormatUint(r.serialComments, 10),
		Content:   comment.Content,
		UserID:    strconv.FormatUint(comment.UserID, 10),
		CreatedAt: comment.CreatedAt,
	}

	r.comments[r.serialComments] = comment
	r.articles[a].Comments = append(r.articles[a].Comments, r.serialComments)
	r.serialComments++
	return m, nil
}

func (r *Repository) UpdateComment(ctx context.Context, input model.UpdateComment) (*model.Comment, error) {
	u, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("CommentRepository.UpdateComment: invalid commentID")
	}
	comment, ok := r.comments[u]
	if !ok {
		return nil, errors.NewNotFoundError("CommentRepository.UpdateComment: comment not found")
	}
	if input.Content != nil {
		comment.Content = *input.Content
	}
	return &model.Comment{
		ID:      input.ID,
		Content: comment.Content,
	}, nil
}

func (r *Repository) DeleteComment(ctx context.Context, id string) (bool, error) {
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return false, errors.NewBadRequestError("CommentRepository.DeleteComment: invalid commentID")
	}
	_, ok := r.comments[u]
	if !ok {
		return false, errors.NewNotFoundError("CommentRepository.DeleteComment: comment not found")
	}
	delete(r.comments, u)
	return true, nil
}

func (r *Repository) GetReplies(ctx context.Context, commentID string, after *string,
	sort *model.Sort) ([]*model.Comment, error) {
	u, err := strconv.ParseUint(commentID, 10, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("CommentRepository.GetReplies: invalid commentID")
	}
	_, ok := r.comments[u]
	if !ok {
		return nil, errors.NewNotFoundError("CommentRepository.GetReplies: comment not found")
	}

	var cursor uint64
	if after != nil {
		cursor, err = strconv.ParseUint(*after, 10, 64)
		if err != nil {
			return nil, errors.NewBadRequestError("CommentRepository.GetReplies: invalid after")
		}
	}

	comments := make([]*schema.Comment, 0)
	for k, v := range r.comments {
		if v.Parent == u && k > cursor {
			comments = append(comments, v)
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

func (r *Repository) GetCommentAuthorID(ctx context.Context, id string) (string, error) {
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return "", errors.NewBadRequestError("CommentRepository.GetCommentAuthorID: invalid commentID")
	}
	comment, ok := r.comments[u]
	if !ok {
		return "", errors.NewNotFoundError("CommentRepository.GetCommentAuthorID: comment not found")
	}
	return strconv.FormatUint(comment.UserID, 10), nil
}
