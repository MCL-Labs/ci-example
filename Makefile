build:
	go mod tidy
	go vet ./...
	go build -o bin/ci_example main.go


test:
	go test -v ./...