package pg

import (
	"context"
	"fmt"
	"github.com/darleet/blog-graphql/internal/model"
	"strconv"
)

func (r *Repository) GetUsers(ctx context.Context, userID []string) ([]*model.User, []error) {
	var params string
	for i := range userID {
		params += "$" + strconv.Itoa(i+1) + ","
	}
	params = params[:len(params)-1]

	query := fmt.Sprintf(`SELECT id, username, avatar_url FROM users WHERE id IN (%s)`, params)

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, []error{err}
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	errors := make([]error, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.AvatarURL); err != nil {
			errors = append(errors, err)
		}
		users = append(users, &user)
	}
	return users, nil
}
