.PHONY: all
all: help

.PHONY: build

setup: ## install project dependences
	cp .env.dist .env
	go mod tidy && go mod download
	docker-compose pull

start-docker:
	docker-compose up -d

serve-watch:
	go get -u github.com/cosmtrek/air
	make start-docker
	air
