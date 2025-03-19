package main

import (
	"awesomeProject/internal/repository/handler/CLI"
	"awesomeProject/internal/repository/infrastucture/inMemory"
	"awesomeProject/internal/usecase"
)

func main() {
	userRepo := inMemory.NewUserInMemory()

	userUseCase := usecase.NewUserUseCase(userRepo)

	cliHandler := CLI.NewUserHandlerCLI(userUseCase)

	cliHandler.AddUser()
}
