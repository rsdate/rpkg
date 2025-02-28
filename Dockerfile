# syntax=docker/dockerfile:1

FROM golang:1.24.0 AS builder
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod \
go get -u ./... \
&& go mod download
COPY --chmod=755 ./build.sh /app/build.sh
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
/app/build.sh \
&& tar -czf rpkg.tar.gz ./out/

FROM scratch AS export
WORKDIR /app
COPY --from=builder /app/rpkg.tar.gz ./rpkg.tar.gz
ENTRYPOINT ["./rpkg.tar.gz"]