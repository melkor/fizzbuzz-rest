 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build install

install:
	$(GOINSTALL) -v ./...

build: 
	$(GOBUILD) -v ./...

test: 
	$(GOTEST) -v -cover ./...

clean: 
	$(GOCLEAN)

