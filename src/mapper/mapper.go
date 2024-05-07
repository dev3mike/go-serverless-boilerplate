package mapper

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dev3mike/go-serverless-boilerplate/src/types"
)


func GetMappedUserDto(request events.APIGatewayProxyRequest) (*types.UserDto, error){
	var userDto *types.UserDto

	err := json.Unmarshal([]byte(request.Body), &userDto)

	if err != nil {
		return userDto, err
	}

	return userDto, nil
}