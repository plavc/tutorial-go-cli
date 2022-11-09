
deps:
	go mod tidy
	go mod vendor

build:
	go build -o out/tutorial-go-cli cmd/cli/main.go

run:
	out/tutorial-go-cli

docker-build: