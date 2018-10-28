clean:
	rm -rfv main
build: clean test
	GOOS=linux GOARCH=amd64 go build -o main
test:
	go test
zippy: build
	zip deployment.zip main
