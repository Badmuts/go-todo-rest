all: get-deps compile

get-deps:
	# glide install

compile:
	@go build -v

build:
	@docker run --rm -v $$(pwd):/usr/src/api -w /usr/src/api golang:1.8 bash -c make