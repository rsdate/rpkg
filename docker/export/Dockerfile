# syntax=docker/dockerfile:1

FROM golang:1.24.1 AS builder
WORKDIR /app
RUN go env -w GOMODCACHE=/root/.cache/go-build
RUN \ 
--mount=type=bind,src=./,dst=/app/,rw \
--mount=type=cache,target=/root/.cache/go-build \
go get -u ./... && go mod download
COPY --chmod=755 ./build.sh /app/build.sh
COPY . .
RUN \
--mount=type=cache,target=/root/.cache/go-build \
/app/build.sh && tar -czf rpkg.tar.gz ./out/

FROM scratch AS export
WORKDIR /app
COPY --from=builder /app/rpkg.tar.gz ./rpkg.tar.gz
ENTRYPOINT ["./rpkg.tar.gz"]