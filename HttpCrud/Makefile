build-docker:
	docker build . -t meetup:latest

run-docker:
	docker run -p 8000:8000 meetup:latest

size-docker:
	docker image ls | egrep "meetup|alpine"

all: build-docker run-docker