################################
# Dependency related commands
################################
.PHONY: dependency
dependency: install-golangci-lint install-goose install-go-enum  install-goimports install-gocover-cobertura

.PHONY: install-go-enum
install-go-enum:
	go install github.com/abice/go-enum@latest

.PHONY: install-golangci-lint
install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

.PHONY: install-goimports
install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

install-gocover-cobertura:
	go install github.com/boumenot/gocover-cobertura

.PHONY: generate
generate:
	go generate ./...

.PHONY: unittest
unittest:
	go test -short -v ./...
################################
# CI related commands
################################

.PHONY: test
test: 
	go test -cover ./...


.PHONY: clean
clean:
	if [ -d bin ] ; then rm -rf bin ; fi

.PHONY: format
format:
	go fmt ./...

.PHONY: analyze
analyze:
	go vet ./...

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: release-staging
release-staging:
	standard-version -p beta

.PHONY: release-prod
release-prod:
	standard-version

.PHONY: fix-imports
fix-imports:
	goimports -w ./cmd
	goimports -w ./internal
	goimports -w ./pkg