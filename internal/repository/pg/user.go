package pg

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
)

func (r *Repository) GetUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	query := `SELECT id, username, avatar_url FROM users WHERE id = ANY ($1)`

	rows, err := r.pool.Query(ctx, query, userIDs)
	if err != nil {
		return nil, []error{err}
	}
	defer rows.Close()

	users := make([]*model.User, 0, len(userIDs))
	errors := make([]error, 0, len(userIDs))
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.AvatarURL); err != nil {
			users = append(users, nil)
			errors = append(errors, err)
		} else {
			users = append(users, &user)
			errors = append(errors, nil)
		}
	}
	return users, errors
}
