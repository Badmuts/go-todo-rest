all: run

build:
	docker build -t badmuts/$$JOB_NAME:$$GIT_BRANCH-$$BUILD_NUMBER -f operations/docker/Dockerfile

test:
	docker build -t go-todo-rest:test -f operations/docker/Dockerfile.test .
	docker run --rm -it go-todo-rest:test

coverage:
	echo "No coverage for you!"

push: build
	docker push badmuts/$$JOB_NAME:$$GIT_BRANCH-$$BUILD_NUMBER

cleanup:
	docker rmi go-todo-rest:test
	docker rmi badmuts/$$JOB_NAME:$$GIT_BRANCH-$$BUILD_NUMBER

run:
	docker-compose up -d