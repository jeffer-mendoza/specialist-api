package storage

import (
	"app/internal/domain/model"
	"app/internal/domain/port"
	"app/internal/infrastructure/container"
	"github.com/jeffer-mendoza/go-common-lib/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

const specialistTable = "specialist"

func NewSpecialistImpl(container *container.Dependency) port.Specialist {
	return &specialistImpl{
		container: container,
	}
}

type specialistImpl struct {
	container *container.Dependency
}

func (repo *specialistImpl) Create(ctx *gin.Context, specialist model.Specialist) error {
	item, err := dynamodbattribute.MarshalMap(specialist)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(specialistTable),
		Item:      item,
	}

	_, err = repo.container.DynamoDbClient.PutItemWithContext(ctx, input)
	return err
}

func (repo *specialistImpl) Find(ID string) (model.Specialist, error) {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(ID),
		},
	}
	input := &dynamodb.GetItemInput{
		Key: key,
		TableName: aws.String(specialistTable),
	}

	response, err := repo.container.DynamoDbClient.GetItem(input)

	var specialist model.Specialist

	if response.Item == nil {
		err = errors.NewNotFoundApiError("no found specialist: " + ID)
		return specialist, err
	}

	if err != nil {
		err = errors.NewInternalServerApiError("communication", err)
		return specialist, err
	}

	err = dynamodbattribute.UnmarshalMap(response.Item, &specialist)
	if err != nil {
		err = errors.NewInternalServerApiError("unmarshall", err)
		return specialist, err
	}

	return specialist, nil
}

func (repo *specialistImpl) FindAll() ([]model.Specialist, error) {
	result, err := repo.container.DynamoDbClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String(specialistTable),
	})

	if err != nil {
		return nil, err
	}

	items := make([]model.Specialist, 0)
	if len(result.Items) == 0 {
		return items, nil
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *specialistImpl) Delete(ID string) error {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(ID),
		},
	}

	input:=  &dynamodb.DeleteItemInput{
		TableName: aws.String(specialistTable),
		Key: key,
	}

	_, err := repo.container.DynamoDbClient.DeleteItem(input)

	return err
}
