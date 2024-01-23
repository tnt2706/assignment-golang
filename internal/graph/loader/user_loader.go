package loader

import (
	"assignment/internal/model"
	repo "assignment/internal/repository"
	"context"
)

type UserReader struct {
	userRepo repo.UserRepo
}

func (u *UserReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	users, err := u.userRepo.GetUsersByIds(userIDs)
	if err != nil {
		return nil, []error{err}
	}

	lenUser := len(userIDs)

	d := make(map[string]*model.User, lenUser)

	for _, user := range users {
		d[user.ID] = user
	}

	result := make([]*model.User, lenUser)

	for index, id := range userIDs {
		result[index] = d[id]
	}

	return result, nil
}
