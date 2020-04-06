.PHONY: all
all: help

.PHONY: build

setup: ## install project dependences
	go mod tidy && go mod download

start-docker:
	docker-compose up -d

serve-watch:
	go get -u github.com/cosmtrek/air
	make start-docker
	air
