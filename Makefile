.PHONY: run clear build

run: build
	sudo ./bin/spectrum rainbow

clear: build
	sudo ./bin/spectrum clear

build:
	go build -v -o ./bin/spectrum ./cmd/spectrum/