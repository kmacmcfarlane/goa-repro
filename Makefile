PROJECT_NAME=endorsements

all: |deps gen build test

deps:
	go get -u goa.design/goa/v3
	go get -u goa.design/goa/v3/...

gen:
	go generate ./...
