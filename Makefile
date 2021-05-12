TEST=./test/**/*.go
GOCMD=$(if $(shell which richgo),richgo,go)

test:
	$(GOCMD) test $(TEST)

test-coverage:
	$(GOCMD) test -v $(TEST) -covermode=atomic -coverpkg=./... -coverprofile=coverage.out -json > report.json
