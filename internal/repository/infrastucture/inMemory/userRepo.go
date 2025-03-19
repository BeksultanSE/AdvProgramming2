package inMemory

import (
	"awesomeProject/internal/domain"
	"errors"
)

type UserInMemory struct {
	users []*domain.User
}

func (r *UserInMemory) Save(user *domain.User) (uint64, error) {
	user.ID = uint64(len(r.users) + 1)
	r.users = append(r.users, user)
	return user.ID, nil
}

func (r *UserInMemory) FindById(id uint64) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewUserInMemory() *UserInMemory {
	return &UserInMemory{
		users: make([]*domain.User, 0),
	}
}
