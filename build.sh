GOOS=darwin GOARCH=arm64 go build -o server-m1 server.go
GOOS=darwin GOARCH=amd64 go build -o server server.go
GOOS=linux GOARCH=amd64 go build -o server-linux server.go
GOOS=windows GOARCH=amd64 go build -o server.exe server.go