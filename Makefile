REPO=badmuts

# Name of the image
IMAGE=go-todo-rest

# Current branch-commit (example: master-ab01c1z)
CURRENT=`echo $$GIT_BRANCH | cut -d'/' -f 2-`-$$(git rev-parse HEAD | cut -c1-7)

.PHONY: coverage

run: start
suite: test coverage
package: compile build

# Run linters, simple code quality check
lint:
	go tool vet $$(go list ./... | grep -v /vendor/)
	golint ./...

# Run tests
test:
	go test -v -coverprofile=coverage/c.out

# Create coverage report
coverage:
	go tool cover -html=coverage/c.out -o coverage/coverage.html

# Moves coverage reports to specific jenkins folder
mv-jenkins-reports:
	mv coverage/coverage.html /var/jenkins_home/workspace/go-todo-rest/coverage/

# Moves build binary to specific jenkins folder
mv-jenkins-build:
	mv go-todo-rest /var/jenkins_home/workspace/go-todo-rest/

# Jenkins step to run complete pipeline
ci-jenkins-tests:
	docker build -t go-todo-rest:test -f operations/docker/Dockerfile.test .
	docker run --rm --volumes-from jenkins go-todo-rest:test bash -c make suite && make mv-jenkins-reports
	docker run --rm --volumes-from jenkins go-todo-rest:test bash -c make compile && make mv-jenkins-build

# Jenkins step to run complete pipeline
ci-jenkins: ci-jenkins-tests build push cleanup

# Create binary
compile:
	go build 

# Create docker image with tag badmuts/go-todo-rest:branch-sha
build:
	docker build -t $(REPO)/$(IMAGE):$(CURRENT) -f operations/docker/Dockerfile .

# Push image to the hub, this also build the image
push: build
	docker push $(REPO)/$(IMAGE):$(CURRENT)

# Cleanup step to remove test image and build image
cleanup:
	docker rmi go-todo-rest:test
	docker rmi $(REPO)/$(IMAGE):$(CURRENT)

# Run development via docker-compose. This autoreloads/compiles on change etc.
start:
	docker-compose up -d
