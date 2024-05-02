package services

import (
	"github.com/dev3mike/go-serverless-boilerplate/src/database"
	"github.com/dev3mike/go-serverless-boilerplate/src/errors"
	"github.com/dev3mike/go-serverless-boilerplate/src/types"
)

type UserService struct {
	dbClient database.DynamoDbClient
}

func NewUserService(dbClient database.DynamoDbClient) UserService{
	return UserService{
		dbClient: dbClient,
	}
}

func(u UserService) CreateUser(user *types.UserDto) error{
	userExists, err := u.dbClient.DoesUserExist(user.Email);

	if err != nil{
		return errors.NewErrorWithError(errors.InternalServerError, err)
	}

	if userExists != false{
		return errors.NewError(errors.EmailAlreadyExist)
	}

	err = u.dbClient.CreateUser(user)

	if err != nil{
		return errors.NewErrorWithError(errors.InternalServerError, err)
	}

	return nil
}