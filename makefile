build: dep
	GOOS=linux go build -o main
dep:
	dep ensure
test:
	go test