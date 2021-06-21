TEST ?= ./...
GOCMD=$(if $(shell which richgo),richgo,go)

test-unit:
	$(GOCMD) test $(TEST)

test-coverage:
	rm -rf .coverage && mkdir .coverage
	$(GOCMD) test -v $(TEST) -covermode=atomic -coverpkg=./... -coverprofile=.coverage/coverage.out -json > .coverage/report.json

release:
	@bash -c "$$(curl -s https://raw.githubusercontent.com/escaletech/releaser/master/tag-and-push.sh)"
