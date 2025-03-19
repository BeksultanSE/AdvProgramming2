package usecase

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/repository"
)

type UserUseCaseInterface interface {
	CreateUser(user *domain.User) (uint64, error)
	FindUserById(userId uint64) (domain.User, error)
}

type UserUseCase struct {
	userRepo repository.UserRepoInterface
}

func NewUserUseCase(userRepo repository.UserRepoInterface) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}
func (uc *UserUseCase) CreateUser(user *domain.User) (uint64, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}

	id, err := uc.userRepo.Save(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (uc *UserUseCase) FindUserById(id uint64) (*domain.User, error) {
	user, err := uc.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
