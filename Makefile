.PHONY: check-fmt test test-ci

check-fmt:
	DIFF=$$(gofmt -d .);echo "$${DIFF}";test -z "$${DIFF}"

test:
	go test -covermode=set -failfast

test-ci:
	go test -coverprofile=coverage.out -covermode=set
