mocks:
	mockery --config=./resource/.mockery.yml
lint:
	golangci-lint run --config ./resource/.golangci.yml ./...