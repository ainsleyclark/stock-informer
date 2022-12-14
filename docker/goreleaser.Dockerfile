# Run Stage
FROM alpine:latest
COPY informer /
ENTRYPOINT ["./informer"]
