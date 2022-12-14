# Variables
DIR=$(PWD)

setup: # Setup dependencies
	go mod tidy
	go generate ./...
.PHONY: setup

run: # Run the app
	go run main.go
.PHONY: run

format: # Run gofmt
	go fmt ./...
.PHONY: format

lint: # Run linter
	golangci-lint run ./...
.PHONY: lint

test: # Test uses race and coverage
	go clean -testcache && go test -race $$(go list ./... | grep -v tests | grep -v res | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

test-v: # Test with -v
	go clean -testcache && go test -race -v $$(go list ./... | grep -v tests | grep -v res | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

cover: test # Run all the tests and opens the coverage report
	go tool cover -html=coverage.out
.PHONY: cover

dist: # Creates and build dist folder
	goreleaser check
	goreleaser release --rm-dist --snapshot
.PHONY: dist

docker-clean: # Removes the docker image
	docker image rm stock-informer
.PHONY: docker-clean

docker-build: # Build the docker image
	docker build -f ./docker/Dockerfile -t stock-informer .
.PHONY: docker-build

docker-run: # Run the docker image
	docker run -it --rm \
		-v $(DIR)/config.yml:/mnt/config.yml \
		-p 8010:8080 stock-informer \
		-path=/mnt/config.yml
.PHONY: docker-run

mock: # Make mocks keeping directory tree
	rm -rf gen/mocks \
	&& mockery --all --keeptree --exported=true --output=./mocks
.PHONY: mock

bench: # Runs benchmarks
	go test -benchmem -bench .
.PHONY: bench

doc: # Run go doc
	godoc -http localhost:8080
.PHONY: doc

all: # Make format, lint and test
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
.PHONY: all

todo: # Show to-do items per file
	$(Q) grep \
		--exclude=Makefile.util \
		--exclude-dir=vendor \
		--exclude-dir=.idea \
		--text \
		--color \
		-nRo \
		-E '\S*[^\.]TODO.*' \
		.
.PHONY: todo

help: # Display this help
	$(Q) awk 'BEGIN {FS = ":.*#"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?#/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help
