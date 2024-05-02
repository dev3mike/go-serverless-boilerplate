package app

import (
	"github.com/dev3mike/go-serverless-boilerplate/src/api"
	"github.com/dev3mike/go-serverless-boilerplate/src/database"
	"github.com/dev3mike/go-serverless-boilerplate/src/services"
)



type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {

	dbClient := database.NewDynamoDbClient()
	userService := services.NewUserService(dbClient)
	apiHandler := api.NewApiHandler(userService)

	return App{
		ApiHandler: apiHandler,
	}
}