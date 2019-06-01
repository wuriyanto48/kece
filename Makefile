TEST_OPTS=-covermode=atomic -v -cover -race -coverprofile=coverage.txt

APP_NAME=github.com/wuriyanto48/kece
APP_RELEASE_VERSION=v0.0.0

# Testing
test:
	@go test $(TEST_OPTS) ./...

unittest:
	@go test -short $(TEST_OPTS) ./...

# Linter
lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.13.2

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

release:
	./scripts/build.sh $(APP_NAME) $(APP_RELEASE_VERSION)

.PHONY: lint lint-prepare clean build unittest test release
