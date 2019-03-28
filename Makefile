 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build

build: 
	$(GOBUILD) -v ./...

test: 
	$(GOTEST) -v -cover ./...

clean: 
	$(GOCLEAN)

