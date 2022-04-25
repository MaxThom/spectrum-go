.PHONY: run clear build

run: build
	sudo ./bin/spectrum rainbow

wipe: build
	sudo ./bin/spectrum wipe

clear: build
	sudo ./bin/spectrum clear

build:
	go build -v -o ./bin/spectrum ./cmd/spectrum/