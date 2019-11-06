package note

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createTestTable() (client *dynamodb.DynamoDB, name string, err error) {
	conf := aws.NewConfig().
		WithEndpoint("http://localhost:4569").
		WithRegion("eu-west-2").
		WithCredentials(credentials.NewStaticCredentials("dummy", "dummy", ""))

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: *conf,
	})
	if err != nil {
		return
	}
	client = dynamodb.New(sess)
	name = fmt.Sprintf("notes-api_test_%d", time.Now().UnixNano())
	_, err = client.CreateTable(&dynamodb.CreateTableInput{
		TableName:   &name,
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		KeySchema: []*dynamodb.KeySchemaElement{
			&dynamodb.KeySchemaElement{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("HASH"),
			},
			&dynamodb.KeySchemaElement{
				AttributeName: aws.String("noteId"),
				KeyType:       aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			&dynamodb.AttributeDefinition{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
			&dynamodb.AttributeDefinition{
				AttributeName: aws.String("noteId"),
				AttributeType: aws.String("S"),
			},
		},
	})
	if err != nil {
		return
	}
	_, err = client.UpdateTimeToLive(&dynamodb.UpdateTimeToLiveInput{
		TableName: aws.String(name),
		TimeToLiveSpecification: &dynamodb.TimeToLiveSpecification{
			AttributeName: aws.String("ttl"),
			Enabled:       aws.Bool(true),
		},
	})
	return
}

func dropTable(client *dynamodb.DynamoDB, name string) error {
	_, err := client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(name),
	})
	return err
}
