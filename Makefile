VERSION = $(shell git describe --tags)
FLAG = "-X main.Version=$(VERSION)"

build:
	go build -ldflags $(FLAG)

install:
	go install -ldflags $(FLAG)

uninstall:
	rm $(shell go env GOPATH)/bin/getdl

clean:
	rm $(shell go env GOPATH)/bin/getdl
	rm getdl

