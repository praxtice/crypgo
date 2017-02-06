default: build

.PHONY: clean

build:
	go build

run: 
	go build
	./go-iex
