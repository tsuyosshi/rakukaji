DOCKER_REPOSITORY := rakukaji-api
BACKEND_DOCKERFILE := backend/Dockerfile
DOCKER_COMPOSE := docker-compose.yml

.PHONY: build
build:
	docker build \
		-t $(DOCKER_REPOSITORY) \
		-f $(BACKEND_DOCKERFILE) \
		.

.PHONY: cup
cup:
	docker-compose \
		-f ./$(DOCKER_COMPOSE) \
		up -d

.PHONY: cdown
cdown:
	docker-compose \
		-f ./$(DOCKER_COMPOSE) \
		down
