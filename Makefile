TEST_OPTS=-covermode=atomic -v -cover

# Testing
.PHONY: unittest
unittest:
	@go test -short $(TEST_OPTS) ./...

.PHONY: test
test:
	@go test $(TEST_OPTS) ./...