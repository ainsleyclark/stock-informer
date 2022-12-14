# Run Stage
FROM alpine:latest
COPY informer /
EXPOSE 8080
ENTRYPOINT ["./informer"]
