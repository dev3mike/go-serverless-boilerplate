package database

import "github.com/dev3mike/go-serverless-boilerplate/src/types"

type DataStore interface {
	DoesUserExist(email string) (bool, error)
	CreateUser(userDto *types.UserDto) error
}