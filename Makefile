REPO ?= danberg/hello-app
TAG ?= $(shell git rev-parse --short HEAD)
IMAGE := $(REPO):$(TAG)
LATEST_IMAGE := $(REPO):latest

all: docker-build docker-push timestamp

.PHONY: buildgo
buildgo:
	# Disable CGO to avoid requiring the Go libs in the container
	CGO_ENABLED=0 go build -v .

docker-build:
	docker build -t $(IMAGE) -f Dockerfile .
	docker tag $(IMAGE) $(LATEST_IMAGE)
docker-push:
	docker push $(IMAGE)
	docker push $(LATEST_IMAGE)
timestamp:
	date
