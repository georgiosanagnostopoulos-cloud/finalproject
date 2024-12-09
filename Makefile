# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOTEST = $(GOCMD) test
BINARY_NAME = api

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) api.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
