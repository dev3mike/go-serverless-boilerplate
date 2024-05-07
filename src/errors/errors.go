package errors

import "fmt"

const (
	InternalServerError = "InternalServerError"
	EmailAlreadyExist = "EmailAlreadyExist"
	InvalidInput = "InvalidInput"
	UserCouldNotBeFound = "UserCouldNotBeFound"
)

var errorsMap = map[string]string{
	InternalServerError: "something went wrong, try again later",
	EmailAlreadyExist: "email already exist in the database",
	InvalidInput: "the input is invalid",
	UserCouldNotBeFound: "UserCouldNotBeFound",
}

func NewError(errorKey string) error {
	if errorMessage, exists := errorsMap[errorKey]; exists {
		return fmt.Errorf(errorMessage);
	}

	return fmt.Errorf("undefined error key: %s", errorKey)
}

func NewErrorWithError(errorKey string, err error) error {
	if errorMessage, exists := errorsMap[errorKey]; exists {
		return fmt.Errorf(errorMessage + ": %s", errorMessage);
	}

	return fmt.Errorf("undefined error key: %s, message: %s", errorKey, err)
}