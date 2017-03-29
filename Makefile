all: compile

deps:
	ls -lah
	go get -d -v
	go install -v

compile: deps
	go build -v
	go test -v

build:
	@docker run --rm -v $$(pwd):/usr/local/go/src/github.com/badmuts/go-todo-rest -w /usr/local/go/src/github.com/badmuts/go-todo-rest golang:1.8 bash -c make
