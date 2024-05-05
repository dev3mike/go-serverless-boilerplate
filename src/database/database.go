package database

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dev3mike/go-serverless-boilerplate/src/types"
)

const (
	TABLE_NAME = "usersTable"
)

type DynamoDbClient struct {
	dbClient *dynamodb.DynamoDB
}


func NewDynamoDbClient() DynamoDbClient {

	dbSession := session.Must(session.NewSession())
	dbClient := dynamodb.New(dbSession)

 return DynamoDbClient{
	dbClient: dbClient,
 }
}

func(db DynamoDbClient) DoesUserExist(email string) (bool, error){

	validatedEmail, err := mail.ParseAddress(email);

	if err != nil {
		return true, err;
	}

	result, err := db.dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(validatedEmail.Address),
			},
		},
	})

	if err != nil{
		return true, err
	}

	if result.Item != nil {
		return true, nil
	}

	return false, nil;
}

func(db DynamoDbClient) CreateUser(userDto *types.UserDto) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(userDto.Email),
			},
			"firstName": {
				S: aws.String(userDto.FirstName),
			},
			"lastName": {
				S: aws.String(userDto.LastName),
			},
			"passwordHash": {
				S: aws.String(userDto.Password),
			},
		},
	}

	_, err := db.dbClient.PutItem(item)

	if err != nil{
		return err
	}

	return nil
}