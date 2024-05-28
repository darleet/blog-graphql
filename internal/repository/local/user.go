package local

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/pkg/errors"
	"strconv"
)

func (r *Repository) GetUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	users := make([]*model.User, 0, len(userIDs))
	errs := make([]error, 0, len(userIDs))

	for _, id := range userIDs {
		u, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			errs = append(errs, errors.NewBadRequestError("UserRepository.GetUsers: invalid userID"))
			users = append(users, nil)
			continue
		}

		user, ok := r.users[u]
		if !ok {
			errs = append(errs, errors.NewNotFoundError("UserRepository.GetUsers: user not found"))
			users = append(users, nil)
			continue
		}

		url := model.URL(user.AvatarURL.String())

		errs = append(errs, nil)
		users = append(users, &model.User{
			ID:        strconv.FormatUint(user.ID, 10),
			Username:  user.Username,
			AvatarURL: &url,
		})
	}

	return users, errs
}
