package types

import (
	"github.com/dev3mike/go-serverless-boilerplate/src/helpers"
)

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserEntity struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func(dto UserDto) GetEntity() (UserEntity, error){
	hashedPassword, err := helpers.HashPassoword(dto.Password)

	if err != nil {
		return UserEntity{}, err
	}

	return UserEntity{
		FirstName: dto.FirstName,
		LastName: dto.LastName,
		Email: dto.Email,
		Password: hashedPassword,
	}, nil
}