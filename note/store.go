package note

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewStore constructor
func NewStore(tableName, region string) *Store {
	session, _ := session.NewSession(&aws.Config{
		Region: &region,
	})
	return &Store{
		client:    dynamodb.New(session),
		tableName: tableName,
	}
}

// Store is the notes DynamoDB store
type Store struct {
	client    dynamodbiface.DynamoDBAPI
	tableName string
}

// PutNote puts a note record into the store
func (s Store) PutNote(n Note) error {
	item, err := dynamodbattribute.MarshalMap(n)
	if err != nil {
		return err
	}
	_, err = s.client.PutItem(&dynamodb.PutItemInput{
		TableName: &s.tableName,
		Item:      item,
	})
	return err
}
