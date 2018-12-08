build: 
	docker build --rm -t kartz/ratings:latest . -f Dockerfile.prod

build-dev:
	docker build --rm -t kartz/ratings.dev:latest . -f Dockerfile.dev

up:
	docker-compose up

down:
	docker-compose down