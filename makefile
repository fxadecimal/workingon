all: clean osx linux windows

APP_NAME := wko
VERSION := $(shell git describe --tags --abbrev=0)

osx:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/osx/$(APP_NAME) .
	cp bin/osx/$(APP_NAME) .
	
linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/linux/$(APP_NAME) .

windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/windows/$(APP_NAME).exe .

clean:
	rm -rf bin
	mkdir -p bin/osx
	mkdir -p bin/linux
	mkdir -p bin/windows

zip:
	rm -rf release && mkdir -p release
	zip -r release/$(APP_NAME)-osx-$(VERSION).zip bin/osx/$(APP_NAME) 
	zip -r release/$(APP_NAME)-linux-$(VERSION).zip bin/linux/$(APP_NAME) 
	zip -r release/$(APP_NAME)-windows-$(VERSION).zip bin/windows/$(APP_NAME).exe

.phony: osx clean linux windows
