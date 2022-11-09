
default: clean build test

deps:
	go mod tidy

build:
	go build -o out/tutorgo main.go

test:
	go test ...

run:
	out/tutorgo

clean:
	rm -r out || true