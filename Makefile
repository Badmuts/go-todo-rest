REPO=badmuts

# Name of the image
IMAGE=go-todo-rest

# Current branch-commit (example: master-ab01c1z)
CURRENT=`echo $$GIT_BRANCH | cut -d'/' -f 2-`-$$(git rev-parse HEAD | cut -c1-7)

.PHONY: coverage

all: run

build:
	docker build -t $(REPO)/$(IMAGE):$(CURRENT) -f operations/docker/Dockerfile .

test:
	docker build -t go-todo-rest:test -f operations/docker/Dockerfile.test .
	docker run --rm --volumes-from jenkins -v $$(pwd)/coverage:/coverage go-todo-rest:test go test -v -coverprofile=/coverage/c.out

coverage:
	docker run --rm --volumes-from jenkins -v $$(pwd)/coverage:/coverage go-todo-rest:test go tool cover -html=/coverage/c.out -o /coverage/coverage.html

push: build
	docker push $(REPO)/$(IMAGE):$(CURRENT)

cleanup:
	docker rmi go-todo-rest:test
	docker rmi $(REPO)/$(IMAGE):$(CURRENT)

run:
	docker-compose up -d