clean:
	@go mod tidy
	@rm -rfv main *.zip
build: clean test
	@GOOS=linux GOARCH=amd64 go build -o main
test:
	@go test
zippy: clean build
	@zip deployment.zip main
