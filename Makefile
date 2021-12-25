GITCOMMIT = $(shell git rev-parse --short HEAD)
VERSION = $(shell git describe --tags)

build:
	go build -ldflags "-X getdl/cmd.Version=$(VERSION) -X getdl/cmd.Commit=$(GITCOMMIT)"

install:
	go install -ldflags "-X getdl/cmd.Version=$(VERSION) -X getdl/cmd.Commit=$(GITCOMMIT)"

uninstall:
	rm $(shell go env GOPATH)/bin/getdl

clean:
	rm $(shell go env GOPATH)/bin/getdl
	rm getdl

