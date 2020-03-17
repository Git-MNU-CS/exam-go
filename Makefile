test:
	go test ./...

linux:
	GOOS=linux GOARCH=amd64 go build -o exam entry/main.go

