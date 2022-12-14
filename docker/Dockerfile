# Build Stage
FROM golang:1.19 AS build

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go generate ./...
ENV CGO_ENABLED=0
RUN go build -o informer

# Run Stage
FROM alpine:latest AS runtime
COPY --from=build app/informer ./
ENTRYPOINT ["./informer"]
