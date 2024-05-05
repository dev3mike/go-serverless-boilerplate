package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dev3mike/go-serverless-boilerplate/src/app"
)

func main(){
	application := app.NewApp()
	lambda.Start(application.ApiHandler.CreateUser)
} 