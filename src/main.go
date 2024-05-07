package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dev3mike/go-serverless-boilerplate/src/app"
	"github.com/dev3mike/go-serverless-boilerplate/src/errors"
	"github.com/dev3mike/go-serverless-boilerplate/src/mapper"
)

type AppHandler struct {
	Application *app.App
}

func (h *AppHandler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/register":
		dto, err := mapper.GetMappedUserDto(request)
		if err != nil {
			return errors.NewApiErrorResponse("The input is invalid", 400), err
		}
		return h.Application.ApiHandler.CreateUser(dto)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil
	}
}

func main() {
	application := app.NewApp()
	handler := &AppHandler{Application: &application}
	
	lambda.Start(handler.HandleRequest)
}
