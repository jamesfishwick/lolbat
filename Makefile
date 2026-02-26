.PHONY: setup lint test

setup:
	cp scripts/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
	@echo "pre-commit hook installed"

lint:
	go vet ./...
	golangci-lint run

test:
	go test -race ./...
