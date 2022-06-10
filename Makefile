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