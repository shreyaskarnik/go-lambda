clean:
	rm -rfv main
build: clean dep test
	GOOS=linux go build -o main
dep:
	dep ensure
test:
	go test
zippy: build
	zip deployment.zip main