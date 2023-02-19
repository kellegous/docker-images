#!/bin/bash

set -euo pipefail

PROTOC_VERSION="22.0"
PROTOC_BASE_URL="https://github.com/protocolbuffers/protobuf/releases/download"
GO_VERSION="1.20.1"

install_protoc() {
	local VERSION=$1
	local ARCH
	case "$2" in
		"x86_64")
		ARCH="x86_64"
		;;
		"aarch64")
		ARCH="aarch_64"
		;;
		*)
		return 1
		;;
	esac
	local DST="protoc-${VERSION}-linux-${ARCH}.zip"
	curl -OL "${PROTOC_BASE_URL}/v${VERSION}/${DST}"
	unzip $DST -d /usr/local
	rm $DST
}

install_go() {
	local VERSION=$1
	local ARCH
	case "$2" in
		"x86_64")
		ARCH="amd64"
		;;
		"aarch64")
		ARCH="arm64"
		;;
		*)
	esac
	curl "https://dl.google.com/go/go${VERSION}.linux-${ARCH}.tar.gz" | tar xz -C /usr/local
}

# Check for supported architectures
ARCH=$(uname -m)
case "$ARCH" in
	"x86_64"|"aarch64")
	;;
	*)
	echo "Unsupported machine arch $(uname -m)"
	exit 1
	;;
esac

apt-get update

apt-get install -y curl

curl -sL https://deb.nodesource.com/setup_18.x | bash - 

apt-get update

apt-get install -y nodejs git build-essential closure-compiler zip

apt-get clean

install_protoc $PROTOC_VERSION $ARCH
install_go $GO_VERSION $ARCH
rm $0