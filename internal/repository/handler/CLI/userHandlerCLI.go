package CLI

import (
	"awesomeProject/internal/repository/handler/CLI/dto"
	"awesomeProject/internal/usecase"
	"fmt"
	"log"
)

type UserHandlerCLI struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandlerCLI(u *usecase.UserUseCase) *UserHandlerCLI {
	return &UserHandlerCLI{u}
}

func (uh *UserHandlerCLI) AddUser() {
	var name, email string

	fmt.Print("Enter name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatalf("Error reading name: %v", err)
		return
	}

	fmt.Print("Enter email: ")
	_, err = fmt.Scanln(&email)
	if err != nil {
		log.Fatalf("Error reading email: %v", err)
		return
	}

	userDto := dto.NewUserDTO(name, email)

	user := userDto.Convert()

	id, err := uh.userUseCase.CreateUser(user)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("ID of successful created user: %+v\n", id)
}
