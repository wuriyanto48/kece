TEST_OPTS=-covermode=atomic -v -cover -race -coverprofile=coverage.txt

# Testing
.PHONY: unittest
unittest:
	@go test -short $(TEST_OPTS) ./...

.PHONY: test
test:
	@go test $(TEST_OPTS) ./...

# Linter
lint-prepare: 
	@echo "Installing golangci-lint"
	# @go get -u github.com/golangci/golangci-lint/cmd/golangci-lint 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.13.2

lint: 
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: lint lint-prepare clean build unittest 