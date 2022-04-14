.PHONY: build test local deploy

build:
	env GOOS=linux go build -o bin/url-requestor cmd/main.go

test:
	go test -coverprofile cp.out ./...

deploy:
	kubectl apply -f k8s/url-requestor-deployment.yaml

run-mac:
	env GOOS=darwin go run cmd/main.go
	kubectl apply -f k8s/url-requestor-svc.yaml

run:
	go fmt ./...
	go run cmd/main.go