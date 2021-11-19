all: fetch reload

fetch:
	git pull

up:
	docker-compose up

reload: build
	docker-compose up

build: down
	docker-compose build

down:
	docker-compose down
