mocks:
	mockery --config=./resource/.mockery.yml
generate:
	go generate ./...
lint:
	golangci-lint run --config ./resource/.golangci.yml ./...