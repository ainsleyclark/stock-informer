# Run Stage
FROM alpine:latest
COPY informer /
EXPOSE 8080
CMD ["./krang", "start"]
