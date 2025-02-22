#!/bin/bash

# Linux
echo Linux
archs1=(amd64 arm arm64 ppc64le ppc64 s390x)

for arch in ${archs1[@]}
do
	env GOOS=linux GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/linux/rpkg_linux_${arch} .
done

# MacOS
echo MacOS
archs2=(amd64 arm64)

for arch in ${archs2[@]}
do
	env GOOS=darwin GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/darwin/rpkg_darwin_${arch} .
done

# Windows
echo Windows
archs3=(amd64 arm64 386)

for arch in ${archs3[@]}
do
	env GOOS=windows GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/windows/rpkg_windows_${arch} .
done

# FreeBSD
echo FreeBSD
archs4=(amd64 arm64 386 arm riscv64)

for arch in ${archs4[@]}
do
	env GOOS=freebsd GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/freebsd/rpkg_freebsd_${arch} .
done

# NetBSD
echo NetBSD
archs5=(amd64 arm64 386 arm)

for arch in ${archs5[@]}
do
	env GOOS=netbsd GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/netbsd/rpkg_netbsd_${arch} .
done

# OpenBSD
echo OpenBSD
archs6=(amd64 arm64 386 arm ppc64 riscv64)

for arch in ${archs6[@]}
do
	env GOOS=openbsd GOARCH=${arch} go build -ldflags="-X main.version=$(git describe --always --long --dirty)" -o ./out/openbsd/rpkg_openbsd_${arch} .
done