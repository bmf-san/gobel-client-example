.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.PHONY: docker-compose-build
docker-compose-build: ## Build containers by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose.ci.yml build
else
	docker-compose -f docker-compose-local.yml build
endif

.PHONY: docker-compose-up
docker-compose-up: ## Run containers by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose.ci.yml up
else
	docker-compose -f docker-compose-local.yml up
endif

.PHONY: docker-compose-up-d
docker-compose-up-d: ## Run containers in the background by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose.ci.yml up -d
else
	docker-compose -f docker-compose-local.yml up -d
endif

.PHONY: docker-compose-pull
docker-compose-pull: ## Pull images by docker-compose.
ifeq ($(env), ci)
	docker-compose -f docker-compose.ci.yml pull
else
	docker-compose -f docker-compose-local.yml pull
endif

.PHONY: setup-buildx
setup-buildx: ## Set up buildx builder.
	docker buildx create --name buildx-builder
	docker buildx use buildx-builder

.PHONY: build-and-push
build-and-push: ## Build and push image to dockerhub.
	docker buildx build --no-cache --push --platform linux/amd64,linux/arm64 --file app/Dockerfile --tag bmfsan/gobel-client-example app/

.PHONY: lint
lint: ## Run golint.
	docker exec -it gobel-client-example golint ./...

.PHONY: test
test: ## Run tests.
	docker exec -it gobel-client-example go test -v ./...

.PHONY: build
build: ## Run go build
	cd app && GOOS=linux GOARCH=amd64 go build -o app
