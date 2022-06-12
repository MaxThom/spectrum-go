.PHONY: run clear build

exec:
	sudo ./bin/spectrum

clear:
	sudo ./bin/spectrum clear

run: build
	sudo ./bin/spectrum

build:
	go build -o ./bin/spectrum ./cmd/spectrum/

docker-build:
	docker build --tag maxthom/spectrum-go:latest -f build/Dockerfile .

docker-run:
	docker run --rm --privileged maxthom/spectrum-go:latest

docker-exec:
	docker run -it --rm --privileged maxthom/spectrum-go:latest /bin/bash

docker-push:
	docker push maxthom/spectrum-go:latest

docker-builx:
	docker buildx build --push --platform linux/arm/v7,linux/arm64 -t maxthom/spectrum-go:latest -f build/Dockerfile .

deploy:
	docker-compose -f ./build/docker-compose/docker-compose.yaml down
	docker-compose -f ./build/docker-compose/docker-compose.yaml rm -f
	docker-compose -f ./build/docker-compose/docker-compose.yaml pull
	docker-compose -f ./build/docker-compose/docker-compose.yaml up --build -d
	docker-compose -f ./build/docker-compose/docker-compose.yaml logs -f
