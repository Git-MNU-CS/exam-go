test:
	go test ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o exam main/main.go


