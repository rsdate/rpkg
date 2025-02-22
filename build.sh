#!/bin/bash

# Linux
archs=(amd64 arm arm64 ppc64le ppc64 s390x)

for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/linux/rpkg_linux_${arch}
done

# MacOS
archs=(amd64 arm64)

for arch in ${archs[@]}
do
	env GOOS=darwin GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/darwin/rpkg_darwin_${arch}
done

# Windows
archs=(amd64 arm64 386)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/windows/rpkg_windows_${arch}
done

# FreeBSD
archs=(amd64 arm64 386 arm riscv64)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/freebsd/rpkg_freebsd_${arch}
done

# NetBSD
archs=(amd64 arm64 386 arm)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/netbsd/rpkg_netbsd_${arch}
done

# OpenBSD
archs=(amd64 arm64 386 arm ppc64 riscv64)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o out/openbsd/rpkg_openbsd_${arch}
done