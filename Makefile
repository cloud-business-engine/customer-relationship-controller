PROJECT=customer-relationship-controller
VERSION := $(shell git describe --tags `git rev-list --tags --max-count=1`)

REPORTS=.reports
IMAGE_NAME=$(PROJECT)

SSH_KEY_PATH := $(HOME)/.ssh/id_rsa
MODULE := $(shell awk 'NR==1{print $$2}' go.mod)

$(REPORTS):
	mkdir $(REPORTS)

setup: $(REPORTS) .tidy .install-tools

clean:
	rm -rf $(REPORTS)
	go clean -modcache

format: .gci .tidy
	golangci-lint run --fix ./...

lint: setup
	golangci-lint run ./... --verbose --out-format json > $(REPORTS)/golangci-lint.json

tests:
	go test -race -coverprofile=$(REPORTS)/coverage.out ./... > /dev/null
	go tool cover -func $(REPORTS)/coverage.out
	go tool cover -html=$(REPORTS)/coverage.out -o $(REPORTS)/coverage.html

.PHONY: mock
mock:
	go generate ./...

docker-image:
	DOCKER_BUILDKIT=1 docker build . \
		--ssh default=$(SSH_KEY_PATH) \
		--tag $(IMAGE_NAME) \
		--pull \
		--no-cache

echo-image-name:
	@echo $(IMAGE_NAME)

echo-image-tag:
	@echo $(IMAGE_NAME):$(VERSION)

.gci:
	gci write \
		--section Standard \
		--section Default \
		--section "Prefix($(MODULE))" \
		--skip-generated \
		$$(find . -name "*.go" -type f)

.tidy:
	go mod tidy

.install-tools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

all: format lint tests docker-image

.DEFAULT_GOAL := all
