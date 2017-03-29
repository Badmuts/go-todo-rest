REPO=badmuts

# Name of the image
IMAGE=go-todo-rest

# Current branch-commit (example: master-ab01c1z)
CURRENT=`echo $$GIT_BRANCH | cut -d'/' -f 2-`-$$(git rev-parse HEAD | cut -c1-7)

all: run

build:
	docker build -t $(REPO)/$(IMAGE):$(CURRENT) -f operations/docker/Dockerfile .

test:
	docker build -t go-todo-rest:test -f operations/docker/Dockerfile.test .
	docker run --rm go-todo-rest:test

coverage:
	echo "No coverage for you!"

push: build
	docker push $(REPO)/$(IMAGE):$(CURRENT)

cleanup:
	docker rmi go-todo-rest:test
	docker rmi $(REPO)/$(IMAGE):$(CURRENT)

run:
	docker-compose up -d