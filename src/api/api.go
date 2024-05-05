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

func(api ApiHandler) CreateUser(userDto types.UserDto) error {

	inputErr := helpers.ValidateInput(
		helpers.Input{Name: "email", Value: userDto.Email},
		helpers.Input{Name: "firstName", Value: userDto.FirstName},
		helpers.Input{Name: "lastName", Value: userDto.LastName},
		helpers.Input{Name: "password", Value: userDto.Password},
	)

	if inputErr != nil {
		return inputErr
	}

	userEntity, err := userDto.GetMappedEntity();
	
	if err != nil{
		return err
	}

	err = api.userService.CreateUser(userEntity);

	if err != nil{
		return err
	}

	return nil
}