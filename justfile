APP_NAME := "wko"
VERSION := `git describe --tags --abbrev=0`

all: clean osx-arm64 osx-amd64 linux-amd64 linux-arm64 windows release

build goos goarch exe:
    GOOS={{goos}} GOARCH={{goarch}} CGO_ENABLED=0 \
    go build -ldflags "-X main.Version={{VERSION}}" -o bin/{{goos}}/{{goarch}}/{{exe}}

osx-arm64:
    just build darwin arm64 {{APP_NAME}}

osx-amd64:
    just build darwin amd64 {{APP_NAME}}

linux-amd64:
    just build linux amd64 {{APP_NAME}}

linux-arm64:
    just build linux arm64 {{APP_NAME}}

windows:
    just build windows amd64 {{APP_NAME}}.exe

clean:
    rm -rf bin
    mkdir -p bin/osx
    mkdir -p bin/linux
    mkdir -p bin/windows
    

zip goos goarch:
    cd bin/{{goos}}/{{goarch}} && zip -j ../../../release/{{APP_NAME}}-{{goos}}-{{goarch}}-{{VERSION}}.zip {{APP_NAME}}*

release:
    rm -rf release && mkdir -p release
    just zip darwin arm64
    just zip darwin amd64
    just zip linux amd64
    just zip linux arm64
    just zip windows amd64
