package api

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dev3mike/go-serverless-boilerplate/src/errors"
	"github.com/dev3mike/go-serverless-boilerplate/src/helpers"
	"github.com/dev3mike/go-serverless-boilerplate/src/services"
	"github.com/dev3mike/go-serverless-boilerplate/src/types"
)

type ApiHandler struct {
	userService *services.UserService
}

func NewApiHandler(userService *services.UserService) *ApiHandler {
	return &ApiHandler{
		userService: userService,
	}
}

func(api *ApiHandler) CreateUser(userDto *types.UserDto) (events.APIGatewayProxyResponse, error) {

	inputErr := helpers.ValidateInput(
		&helpers.Input{Name: "email", Value: userDto.Email},
		&helpers.Input{Name: "firstName", Value: userDto.FirstName},
		&helpers.Input{Name: "lastName", Value: userDto.LastName},
		&helpers.Input{Name: "password", Value: userDto.Password},
	)

	if inputErr != nil {
		return errors.NewApiErrorResponse(inputErr.Error(), http.StatusBadRequest), nil
	}

	userEntity, err := userDto.GetMappedEntity();
	
	if err != nil{
		return errors.NewApiErrorResponse("Invalid input", http.StatusBadRequest), err
	}

	err = api.userService.CreateUser(userEntity);

	if err != nil{
		return errors.NewInternalServerError(), err
	}

	return helpers.NewApiMessageResponse("User successfully created"), nil
}

func(api *ApiHandler) GetUser(email string) (events.APIGatewayProxyResponse, error) {
	user, err := api.userService.GetUser(email)

	if err != nil {
		return errors.NewInternalServerError(), err
	}

	if user == nil {
		return errors.NewApiErrorResponse("User could not be found", http.StatusNotFound), nil
	}

	dto, err := user.GetMappedResponseDto();

	if err != nil {
		return errors.NewInternalServerError(), err
	}

	return helpers.NewApiResponse(dto), nil
}