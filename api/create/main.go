package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/maurogestoso/serverless-stack-notes/note"
)

func main() {
	tableName := os.Getenv("NOTES_TABLE_NAME")
	region := os.Getenv("AWS_REGION")

	ns := note.NewStore(tableName, region)
	h := handler{
		putNote: ns.PutNote,
	}
	lambda.Start(h.handleRequest)
}

type handler struct {
	putNote note.Putter
}

func (h handler) handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var b requestBody
	err := json.Unmarshal([]byte(req.Body), &b)
	if err != nil {
		return createResponse(http.StatusBadRequest, err.Error()), err
	}

	userID := req.RequestContext.Identity.CognitoIdentityID
	n, err := note.New(userID, b.Content, b.Attachment)
	if err != nil {
		return createResponse(http.StatusInternalServerError, err.Error()), err
	}

	err = h.putNote(n)
	if err != nil {
		return createResponse(http.StatusInternalServerError, err.Error()), err
	}

	return createResponse(http.StatusCreated, ""), nil
}

type requestBody struct {
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

func createResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
