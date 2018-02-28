IMAGE_NAME = newrelic-demo
DEPENDENCIES = build

COMMONENVVAR = GOOS=$(shell uname -s | tr A-Z a-z) GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))
BUILDENVVAR = CGO_ENABLED=0

all: build

deps:
	glide install -v

# test:
# 	$(COMMONENVVAR) $(BUILDENVVAR) go test ./... -v

build:
	$(COMMONENVVAR) $(BUILDENVVAR) go build -o main *.go

.PHONY: all deps test build
