.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/create api/create/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

test:
	go test ./...

run-localstack:
	docker kill localstack-notes-app-api || true
	docker run --rm -d --name localstack-notes-app-api -p 4569:4569 -p 4568:4568 -e PORT_WEB_UI=8888 -e SERVICES=dynamodb localstack/localstack