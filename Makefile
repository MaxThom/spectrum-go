.PHONY: run clear build

exec:
	sudo ./bin/spectrum

clear:
	sudo ./bin/spectrum clear

run: build
	sudo ./bin/spectrum

build:
	go build -o ./bin/spectrum ./cmd/spectrum/
