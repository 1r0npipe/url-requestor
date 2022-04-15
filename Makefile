.PHONY: build test local deploy

build:
	env GOOS=linux go build -o bin/url-requestor cmd/main.go

test:
	go test ./...

deploy:
	kubectl apply -f k8s/url-requestor-deployment.yaml
	kubectl apply -f k8s/url-requestor-svc.yaml

run-mac:
	env GOOS=darwin go run cmd/main.go

run:
	go fmt ./...
	go run cmd/main.go
