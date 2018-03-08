# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run


install: install_jsnice
all: build_jsnice
build_jsnice: 
	$(GOBUILD) jsnice.go

clean_jsnice: 
	rm -f jsnice

install_jsnice: build_jsnice
	mv jsnice $(HOME)/bin/jsnice

