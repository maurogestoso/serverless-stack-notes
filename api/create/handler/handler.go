package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/maurogestoso/serverless-stack-notes/note"
)

// Handler struct
type Handler struct {
	PutNote note.Putter
}

// New creates a new Handler
func New(np note.Putter) Handler {
	return Handler{
		PutNote: np,
	}
}

// HandleRequest handles an API Gateway request
func (h Handler) HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	err = h.PutNote(n)
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
