# Run Stage
FROM alpine:latest
COPY informer /
ENV path $path
EXPOSE 8080
CMD ["./informer", "-path=$path"]
