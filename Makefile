.PHONY: run clear build

run: build
	sudo ./bin/spectrum

clear: build
	sudo ./bin/spectrum clear

build:
	go build -o ./bin/spectrum ./cmd/spectrum/