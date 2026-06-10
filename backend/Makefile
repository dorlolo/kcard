.PHONY: fmt lint vet test staticcheck
fmt:
	gofmt -w ./cmd ./internal
lint: vet
vet:
	go vet ./...
test:
	go test ./...
staticcheck:
	go test ./...
