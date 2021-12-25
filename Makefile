GITTAG = $(shell git describe --tags)

build:
	go build -ldflags "-X getdl/cmd.Version=$(GITTAG)"

install:
	go install -ldflags "-X getdl/cmd.Version=$(GITTAG)"

uninstall:
	rm $(shell go env GOPATH)/bin/getdl

clean:
	rm $(shell go env GOPATH)/bin/getdl
	rm getdl

