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
	docker run --rm --volumes-from jenkins -v /var/jenkins_home/workspace/go-todo-rest/coverage go-todo-rest:test go test -v -coverprofile=/var/jenkins_home/workspace/go-todo-rest/coverage/c.out

coverage:
	docker run --rm --volumes-from jenkins -v /var/jenkins_home/workspace/go-todo-rest/coverage go-todo-rest:test go tool cover -html=/var/jenkins_home/workspace/go-todo-rest/coverage/c.out -o /var/jenkins_home/workspace/go-todo-rest/coverage /var/jenkins_home/workspace/go-todo-rest/coverage/coverage.html

push: build
	docker push $(REPO)/$(IMAGE):$(CURRENT)

cleanup:
	docker rmi go-todo-rest:test
	docker rmi $(REPO)/$(IMAGE):$(CURRENT)

run:
	docker-compose up -d