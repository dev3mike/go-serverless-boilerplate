package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassoword(plainPassword string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsPasswordValid(hashedPassword, plainPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}