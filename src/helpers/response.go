package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func NewApiResponse(object interface{}) events.APIGatewayProxyResponse {
    // Marshal the object into a JSON string
    jsonData, err := json.Marshal(object)
    if err != nil {
        // Handle JSON marshaling errors by returning a server error response
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Headers: map[string]string{
                "Content-Type": "application/json",
            },
            Body: `{"error": "Internal Server Error"}`,
        }
    }

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
        Body: string(jsonData),
    }
}

func NewApiMessageResponse(message string) events.APIGatewayProxyResponse {
	return NewApiResponse(&MessageResponse{
		Message: message,
	})
}