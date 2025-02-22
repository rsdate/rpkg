FROM golang:1.24.0 as builder
WORKDIR /app
RUN --mount=type=bind,src=./,dst=/app/ go get -u ./... && go mod download
RUN --mount=type=bind,src=./,dst=/app/ \
chmod +x build.sh \
&& ./build.sh \
&& tar -czf rpkg.tar.gz ./out/

FROM scratch as export
COPY --from=builder /app/rpkg.tar.gz /app/rpkg.tar.gz
ENTRYPOINT ["/app/rpkg.tar.gz"]