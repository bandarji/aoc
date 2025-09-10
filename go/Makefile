.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build executable
	@go build -o ./aoc cmd/*.go

clean: ## Remove unnecessary files
	@go clean
	@rm -f aoc coverage.out coverage.html

run: build ## Play the game
	@./aoc -y $(YEAR) -d $(DAY)
	@rm -f aoc

test: ## Execute unit tests
	@go test -v -count=1 -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

coverage: test ## Show code coverage in browser
	@go tool cover -html=coverage.out -o coverage.html
