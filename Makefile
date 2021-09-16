.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o /app/dist/requestcan main.go
