package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

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

func NewApiErrorResponse(message string, httpCode int) events.APIGatewayProxyResponse{
	errorMessage, _ := json.Marshal(map[string]string{
		"error": message,
	})

    return events.APIGatewayProxyResponse{
        StatusCode: httpCode,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
        Body: string(errorMessage),
    }
}

func NewInternalServerError() events.APIGatewayProxyResponse{
	return NewApiErrorResponse("Internal server error", http.StatusInternalServerError)
}