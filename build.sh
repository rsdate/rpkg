#!/bin/bash

# Linux
archs=(amd64 arm arm64 ppc64le ppc64 s390x)

for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -o rpkg_linux_${arch}
done

# MacOS
archs=(amd64 arm64)

for arch in ${archs[@]}
do
	env GOOS=darwin GOARCH=${arch} go build -o rpkg_darwin_${arch}
done

# Windows
archs=(amd64 arm64 386)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -o rpkg_windows_${arch}
done

# FreeBSD
archs=(amd64 arm64 386 arm riscv64)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -o rpkg_windows_${arch}
done

# NetBSD
archs=(amd64 arm64 386 arm)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -o rpkg_windows_${arch}
done

# OpenBSD
archs=(amd64 arm64 386 arm ppc64 riscv64)

for arch in ${archs[@]}
do
	env GOOS=windows GOARCH=${arch} go build -o rpkg_windows_${arch}
done