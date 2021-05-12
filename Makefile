TEST=./test/**/*_test.go
GOCMD=$(if $(shell which richgo),richgo,go)

test:
	$(GOCMD) test $(TEST)

test-coverage:
	$(GOCMD) test -v $(TEST) -covermode=atomic -coverpkg=./... -coverprofile=coverage.out -json > report.json

release:
	@bash -c "$$(curl -s https://raw.githubusercontent.com/escaletech/releaser/master/tag-and-push.sh)"
