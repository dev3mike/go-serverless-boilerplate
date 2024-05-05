package api

import (
	"github.com/dev3mike/go-serverless-boilerplate/src/helpers"
	"github.com/dev3mike/go-serverless-boilerplate/src/services"
	"github.com/dev3mike/go-serverless-boilerplate/src/types"
)

type ApiHandler struct {
	userService services.UserService
}

func NewApiHandler(userService services.UserService) ApiHandler {
	return ApiHandler{
		userService: userService,
	}
}

func(api ApiHandler) CreateUser(event types.UserDto) error {

	inputErr := helpers.ValidateInput(
		helpers.Input{Name: "email", Value: event.Email},
		helpers.Input{Name: "firstName", Value: event.FirstName},
		helpers.Input{Name: "lastName", Value: event.LastName},
		helpers.Input{Name: "password", Value: event.Password},
	)

	if inputErr != nil {
		return inputErr
	}

	err := api.userService.CreateUser(&event);

	if err != nil{
		return err
	}

	return nil
}