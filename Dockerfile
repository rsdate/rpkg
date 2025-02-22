FROM golang:1.24.0 as builder
WORKDIR /app
RUN --mount=type=bind,src=./,dst=/app/ go get -u ./... && go mod download
COPY . .
RUN chmod +x ./build.sh
RUN ./build.sh
RUN tar -czvf rpkg.tar.gz ./out/

FROM scratch as export
COPY --from=builder /app/rpkg.tar.gz /app/rpkg.tar.gz
ENTRYPOINT ["/app/rpkg.tar.gz"]