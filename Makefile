.PHONY: run clear build

run: build
	sudo ./bin/spectrum rainbow

wipe: build
	sudo ./bin/spectrum wipe

maze: build
	sudo ./bin/spectrum maze

clear: build
	sudo ./bin/spectrum clear

build:
	go build -o ./bin/spectrum ./cmd/spectrum/