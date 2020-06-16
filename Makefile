#--include .env
PROJECTNAME = "unode"
# Go related variables.
GOBASE = `pwd`
GOBIN = $(GOBASE)/bin
GOFILE = $(exporter.go)

# Use linker flags to provide version/build settings
#LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

go-clean: hello
	@echo	"===> Cleaning build cache"
	@go clean
	@rm -f $(PROJECTNAME)

## compile: Compile the binary.
hello:
	@echo "***********************************************************"
	@echo "*                                                         *"
	@echo "*                                                         *"
	@echo "*  Welcome to unode exporter for AIX Version 7.1/2        *"
	@echo "*                                                         *"
	@echo "*                                                         *"
	@echo "***********************************************************"

clean: go-clean

build: hello
	@echo "===> Building binary..." $(PROJECTNAME)
	@go build -o $(PROJECTNAME) $(GOFILE)
	

go-compile: go-build

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
 	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)
compile: hello
	@echo "===> Building binary..." $(PROJECTNAME)
	@go build -o $(PROJECTNAME) $(GOFILE)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo " clean"
	@echo " build"
