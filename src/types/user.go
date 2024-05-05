package types

import "golang.org/x/crypto/bcrypt"

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)

	if err != nil {
		return UserEntity{}, err
	}

	return UserEntity{
		FirstName: dto.FirstName,
		LastName: dto.LastName,
		Email: dto.Email,
		Password: string(hashedPassword),
	}, nil
}