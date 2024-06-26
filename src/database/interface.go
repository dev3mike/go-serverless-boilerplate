package database

import "github.com/dev3mike/go-serverless-boilerplate/src/types"


type DataStore interface {
	DoesUserExist(email string) (bool, error)
	CreateUser(userDto *types.UserEntity) error
	GetUser(email string) (*types.UserEntity, error)
}