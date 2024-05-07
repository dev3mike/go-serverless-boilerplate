package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dev3mike/go-serverless-boilerplate/src/app"
	"github.com/dev3mike/go-serverless-boilerplate/src/errors"
	"github.com/dev3mike/go-serverless-boilerplate/src/mapper"
)

func main(){
	application := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
		switch request.Path{
		case "/register":
			dto, err := mapper.GetMappedUserDto(request);

			if err != nil{
				return errors.NewApiErrorResponse("The input is invalid", 400), err
			}

			return application.ApiHandler.CreateUser(dto)
		default:
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
} 