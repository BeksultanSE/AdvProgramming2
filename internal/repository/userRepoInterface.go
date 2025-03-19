package repository

import (
	"awesomeProject/internal/domain"
)

type UserRepoInterface interface {
	Save(user *domain.User) (uint64, error)
	FindById(id uint64) (*domain.User, error)
}
