package dto

import "awesomeProject/internal/domain"

type UserDTO struct {
	Name  string
	Email string
}

func NewUserDTO(name, email string) *UserDTO {
	return &UserDTO{
		Name:  name,
		Email: email,
	}
}

func (u *UserDTO) Convert() *domain.User {
	return &domain.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
