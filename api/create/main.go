package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/maurogestoso/serverless-stack-notes/api/create/handler"
	"github.com/maurogestoso/serverless-stack-notes/note"
)

func main() {
	conf, err := getConfig()
	if err != nil {
		log.Fatal(err.Error())

	}

	ns := note.NewStore(conf.TableName, conf.Region)
	h := handler.New(ns.PutNote)

	lambda.Start(h.HandleRequest)
}

func getConfig() (configuration, error) {
	tn, found := os.LookupEnv("NOTES_TABLE_NAME")
	if !found {
		return configuration{}, fmt.Errorf("could not find environment variable NOTES_TABLE_NAME")
	}
	r, found := os.LookupEnv("AWS_REGION")
	if !found {
		return configuration{}, fmt.Errorf("could not find environment variable AWS_REGION")
	}
	return configuration{
		TableName: tn,
		Region:    r,
	}, nil
}

type configuration struct {
	TableName, Region string
}
