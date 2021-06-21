TEST ?= ./...
GOCMD=$(if $(shell which richgo),richgo,go)

test-unit:
	$(GOCMD) test -count=1 $(TEST)

test-coverage:
	$(GOCMD) test -count=1 -v $(TEST) -covermode=atomic -coverpkg=./... -coverprofile=coverage.out -json > report.json

release:
	@bash -c "$$(curl -s https://raw.githubusercontent.com/escaletech/releaser/master/tag-and-push.sh)"
