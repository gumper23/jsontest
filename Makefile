# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=jsontest

all: test tidy build 
build: 
		CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v
tidy:
		$(GOMOD) tidy
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run: build
		./$(BINARY_NAME)
