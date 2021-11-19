all: fetch reload

fetch:
	git pull

up:
	docker-compose up -d

reload: build
	docker-compose up -d

build: down
	docker-compose build

down:
	docker-compose down
