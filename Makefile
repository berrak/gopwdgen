#!/usr/bin/make
# WARN: gmake syntax

all: 
	@./build.sh
release:
	@NOCOPY=1 ./build.sh package
test:
	@go test ./...
clean-all:
	@rm -f gopwdgen
	@rm -fr packages	
clean:
	@rm -f gopwdgen
install: all
	@cp gopwdgen $(GOPATH)/bin
uninstall: 
	@rm -f $(GOPATH)/bin/gopwdgen
