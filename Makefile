.DEFAULT_GOAL := help

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint fixing issues
	golangci-lint run --fix

.PHONY: tests
tests: ## Run tests
	go test ./... --tags=integration,unit -coverpkg=./...

.PHONY: mocks
mocks: ## Generate mocks
	go generate ./...

.PHONY: app
app: ## Run app
	docker-compose up -d app

.PHONY: logs
logs: ## Show app logs
	docker-compose logs -f

.PHONY: clean
clean: ## clean docker containers, images, volumes and unused networks
	-docker rm -f `docker ps -a -q`
	-docker rmi -f `docker images -q`
	-docker volume rm `docker volume ls -q`
	docker network prune -f
