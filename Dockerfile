# syntax=docker/dockerfile:1

FROM golang:1.24.0 AS builder
WORKDIR /app
RUN --mount=type=bind,src=./,dst=/app/,rw go get -u ./... && go mod download
COPY --chmod=755 ./build.sh /app/build.sh
COPY . .
RUN /app/build.sh && tar -czf rpkg.tar.gz ./out/

FROM scratch AS export
WORKDIR /app
COPY --from=builder /app/rpkg.tar.gz ./rpkg.tar.gz
ENTRYPOINT ["./rpkg.tar.gz"]