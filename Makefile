#!/usr/bin/make
# WARN: gmake syntax

help:
	@echo
	@echo "[develop and test]---------------------------------"
	@echo "build --------- build binary"
	@echo "test ---------- run tests"
	@echo "coverage ------ run coverage tests"
	@echo "clean --------- remove build binary"
	@echo "clean-all ----- removes all build artifacts"
	@echo "install ------- compile and installs binary to GOPATH/bin"
	@echo "uninstall ----- uninstalls binary in GOPATH/bin"
	@echo "[release]------------------------------------------"
	@echo "dist ---------- creates source tar archive"
	@echo "release ------- creates archives with compiled binaries for Linux, Mac, Windows & Free BSD"
	@echo

build: 
	@./build.sh
release:
	@NOCOPY=1 ./build.sh package
test:
	@go test ./...
clean-all:
	@rm -f gopwdgen
	@rm -fr packages
	@rm -rf *.tar.gz	
clean:
	@rm -f gopwdgen
install: build
	@cp gopwdgen $(GOPATH)/bin
uninstall: 
	@rm -f $(GOPATH)/bin/gopwdgen
coverage:
	@sh scripts/coverage
dist:
	@./build.sh archive
