BINARY_NAME=examples
build:
	echo "Compiling for every OS and Platform"
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
    GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
    GOARCH=amd64 GOOS=window go build -o bin/${BINARY_NAME}-windows main.go

run:
	echo "Running..."
	go run main.go

clean:
	go clean
	rm bin/${BINARY_NAME}-darwin
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

lint:
	golangci-lint run --enable-all

all: test test_coverage run