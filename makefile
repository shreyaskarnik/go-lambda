clean:
	rm -rfv main
build: clean dep test
	GOOS=linux GOARCH=amd64 go build -o main
dep:
	dep ensure
test: dep
	go test
zippy: build
	zip deployment.zip main