# gopwdgen

[![Build Status](https://travis-ci.org/berrak/gopwdgen.svg?branch=master)](https://travis-ci.org/berrak/gopwdgen)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat)](./LICENSE)
[![release](https://img.shields.io/badge/release-v0.1.0-blue.svg)]()
[![homebrew](https://img.shields.io/badge/homebrew-v0.1.0-orange.svg)]()

## Generate random passwords or hashes for automatic script retrieval

## Features

- [x] Generates random fixed size 32 characters long passwords
- [x] Supports generation of common (md5, sha1, sha256) file hashes in one command
- [x] Use scrypt library for environment hashes
- [x] Generates 40 characters hash with salts given in environment varible GOPWDGEN_SALT
- [x] In Oracle database environments use ORACLE_SID variable for hash generation 
- [x] Implemented in pure Golang

gopwdgen implements generation of random passwords with provided
requirements as described by [AgileBits
1Password](https://discussions.agilebits.com/discussion/23842/how-random-are-the-generated-passwords)

## TODO

- [ ] Implement interactive flag (-i) to input password at command line
- [ ] Add more tests
- [ ] Add more user control on generated passwords and scrypt hashes

## Linux Install

First, install [Go](https://golang.org), and then update `GOPATH`, `GOBIN` and `PATH` like so:

```bash
export GOPATH=$HOME/go
export GOBIN=$HOME/$GOPATH/bin
export PATH=$GOBIN:/usr/local/go/bin:$PATH
```
Next download the project and build the binary file.

```bash
## download source
$ go get -u github.com/berrak/gopwdgen/cmd/gopwdgen
$ cd $GOPATH/src/github.com/berrak/gopwdgen/cmd/gopwdgen
## create the binary
$ go build
## first gopwdgen run
$ ./gopwdgen --version
## install binary
$ go install
```

## Usage

### Generate random passwords
```bash
$ gopwdgen 
  x)70b0476))8Tj~X@fw.Hl`36e))q)4a
```
It is probably best to store this password in a secure password manager.

### Generate file hashes
```bash
$ gopwdgen LICENSE 
  md5sum: LICENSE 86d3f3a95c324c9479bd8986968f4327
    sha1: LICENSE 7df059597099bb7dcf25d2a9aedfaf4465f72d8d
  sha256: LICENSE c71d239df91726fc519c6eb72d318ec65820627232b2f796219e87dcf35d0ab4
```
### Generate scrypt hashes with salt in environment
```bash
$ export GOPWDGEN_SALT="some_salty_text_string"
$ gopwdgen secret 
  01bf2ff0eb5b41be60d4d0677182534770526688726fcbd47e217674034fb1f2
```
In the same way, the value of ORACLE_SID is accepted as salt in an Oracle database environment.

### Help

```bash 
$ gopwdgen --help
```

### Supported environment variables
```bash 
export ORACLE_SID
export GOPWDGEN_SALT
export GOPWDGEN_UP
```
The latter set to true will force generated scrypt hashes to upper case letters.

## Referenses and credits
[(1) github.com/sethvargo](https://github.com/sethvargo/password) Included but modified source to exclude some troublesome symbols.

## License
This project is under the Apache License. See the LICENSE file for the full license text.
