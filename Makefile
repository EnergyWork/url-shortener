.PHONY: build
build:
	go build -v ./cmd/url.shortener.service.go

.DEFAULT_GOAL := build