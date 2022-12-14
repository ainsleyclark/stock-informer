# Run Stage
FROM alpine:latest
COPY informer /
ARG path
EXPOSE 8080
CMD ["./informer", "-path=$path"]
