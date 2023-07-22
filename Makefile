NAME := app
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd/${NAME}
DEV_CONFIG_PATH := ./configs/dev.yml
STAGE_CONFIG_PATH := ./configs/stage.yml
CONFIG_TEMPLATE_PATH := ./configs/template.yml

# Docker
DOCKER_APP_FILENAME ?= deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml
VERSION ?= v0.1.2

# sed
SECRET_KEY ?= "very-secret-key"
CONFIG_PATH ?= ./configs/new.yml

define sedi
    sed --version >/dev/null 2>&1 && sed -- $(1) > ${CONFIG_PATH} || sed "" $(1) > ${CONFIG_PATH}
endef

.PHONY: run
run: gen
	go run cmd/$(NAME)/main.go ${DEV_CONFIG_PATH}

.PHONY: gen
gen: 
	go generate ./internal/ent

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes
