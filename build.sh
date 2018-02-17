#!/bin/bash
set -e

# Binary name
NAME=gopwdgen

# Ensure git is installed
GITBIN=$(which git)
if [ "$GITBIN" == "" ]; then
echo "Git is not installed. Install git and try again..."
fi

# VERSION file provides one canocial location
MAJOR=$(cat < VERSION | cut -f1 -d' ')
MINOR=$(cat < VERSION | cut -f2 -d' ')
PATCH=$(cat < VERSION | cut -f3 -d' ')

GITCOMMITS=$(git rev-list HEAD | wc -l 2>/dev/null)
if [ "$GITCOMMITS" -eq 0 ];then
echo "Please do 'git init'"
echo "Commit with: git add .; git commit -m 'some commit message'"
fi

# Use semantic versioning with information from git
GIT_HASH=$(git log -n 1 --format="%h")
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD | sed 's/[-_.\/]//g')
DEV_VERSION=${MAJOR}.${MINOR}.${PATCH}+${GIT_HASH}.${GITCOMMITS}
REL_VERSION=${MAJOR}.${MINOR}.${PATCH}

if [ "$1" != "package" ];then
	VERSION=$DEV_VERSION
else
	VERSION=$REL_VERSION
	if [ "$GIT_BRANCH" != "master" ];then
		echo "Must be on branch 'master' to do release. Aborting release..."
	fi
fi

PROTECTED_MODE="no"

export GO15VENDOREXPERIMENT=1

cd $(dirname "${BASH_SOURCE[0]}")
OD="$(pwd)"
WD=$OD

package(){
	echo Packaging $1 Binary
	bdir=gopwdgen-${VERSION}-$2-$3
	rm -rf packages/$bdir && mkdir -p packages/$bdir
	GOOS=$2 GOARCH=$3 ./build.sh
	if [ "$2" == "windows" ]; then
		mv gopwdgen packages/$bdir/gopwdgen.exe
	else
		mv gopwdgen packages/$bdir
	fi
	cp README.md packages/$bdir
	cd packages
	if [ "$2" == "linux" ]; then
		tar -zcf $bdir.tar.gz $bdir
	else
		zip -r -q $bdir.zip $bdir
	fi
	rm -rf $bdir
	cd ..
}

if [ "$1" == "package" ]; then
	rm -rf packages/
	package "Windows" "windows" "amd64"
	package "Mac" "darwin" "amd64"
	package "Linux" "linux" "amd64"
	package "FreeBSD" "freebsd" "amd64"
	exit
fi

PKG_NAME=${NAME}-${REL_VERSION}
PKG=${PKG_NAME}.tar.gz

if [ "$1" == "archive" ]; then
	rm -rf *.tar.gz
    git archive --output=${PKG} --prefix=${PKG_NAME}/ HEAD
	exit
fi

# temp directory for storing isolated environment.
TMP="$(mktemp -d -t sdb.XXXX)"
function rmtemp {
	rm -rf "$TMP"
}
trap rmtemp EXIT

if [ "$NOCOPY" != "1" ]; then
	# copy all files to an isloated directory.
	WD="$TMP/src/github.com/berrak/gopwdgen"
	export GOPATH="$TMP"
	for file in `find . -type f`; do
		# TODO: use .gitignore to ignore, or possibly just use git to determine the file list.
		if [[ "$file" != "." && "$file" != ./.git* && "$file" != ./gopwdgen ]]; then
			mkdir -p "$WD/$(dirname "${file}")"
			cp -P "$file" "$WD/$(dirname "${file}")"
		fi
	done
	cd $WD
fi


# build and store objects into original directory.
go build -ldflags "-X main.Version=$VERSION" -o "$OD/gopwdgen" cmd/gopwdgen/main.go