package note

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
)

func TestPutNote(t *testing.T) {
	client, tableName, err := createTestTable()
	if err != nil {
		t.Error("failed to create test table")
	}
	store := Store{
		client:    client,
		tableName: tableName,
	}

	testNote := Note{
		UserID:     "user1",
		NoteID:     "note1",
		Content:    "This is a test note",
		Attachment: "hello.jpg",
	}
	err = store.PutNote(testNote)
	if err != nil {
		t.Error(err)
	}

	out, err := store.client.GetItem(&dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"userId": &dynamodb.AttributeValue{
				S: aws.String("user1"),
			},
			"noteId": &dynamodb.AttributeValue{
				S: aws.String("note1"),
			},
		},
	})
	var dbNote Note
	dynamodbattribute.UnmarshalMap(out.Item, &dbNote)

	assert.Equal(t, testNote, dbNote)
}
