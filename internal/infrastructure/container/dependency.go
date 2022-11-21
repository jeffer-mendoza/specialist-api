package container

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Dependency struct {
	// controller
	SpecialistController interface{}

	// user cases
	SpecialistApplication interface{}

	// ports
	SpecialistDB interface{}

	// adapters
	DynamoDbClient *dynamodb.DynamoDB

}
