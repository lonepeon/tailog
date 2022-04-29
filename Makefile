BUILD_FOLDER = target

OS := darwin
ARCH := amd64
BINARY_NAME := "tlog"
FULL_BINARY_NAME := $(BINARY_NAME)-$(OS)-$(ARCH)

PROJECT_USERNAME := lonepeon
PROJECT_REPOSITORY := tlog

GIT_BIN := git

GIT_COMMIT := $(shell $(GIT_BIN) rev-parse HEAD)
GIT_BRANCH := $(shell $(GIT_BIN) branch --no-color | awk '/^\* / { print $$2 }')
GIT_STATE := $(shell if [ -z "$(shell $(GIT_BIN) status --short)" ]; then echo clean; else echo dirty; fi)

GO_BIN := go

.PHONY: release
release:
	@$(MAKE) compile OS=darwin ARCH=amd64
	@$(MAKE) compile OS=linux ARCH=amd64

.PHONY: compile
compile:
	@echo "+$@ $(OS) $(ARCH)"
	@touch internal/cmd/version.go
	@GOOS=$(OS) GOARCH=$(ARCH) go build $(BUILD_OPTIONS) \
		-ldflags \
			 "-X github.com/lonepeon/tailog/internal/cmd.gitBranch=$(GIT_BRANCH) \
		 	  -X github.com/lonepeon/tailog/internal/cmd.gitCommit=$(GIT_COMMIT) \
			  -X github.com/lonepeon/tailog/internal/cmd.gitState=$(GIT_STATE) \
			  -X github.com/lonepeon/tailog/internal/cmd.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')" \
		-o $(BUILD_FOLDER)/$(FULL_BINARY_NAME)

.PHONY: test
test: test-unit test-format test-lint test-security

.PHONY: test-unit
test-unit:
	@echo "+ $@"
	@go test ./...

.PHONY: test-format
test-format:
	@echo "+ $@"
	@data=$$(gofmt -l main.go internal);\
		 if [ -n "$${data}" ]; then \
			>&2 echo "format is broken:"; \
			>&2 echo "$${data}"; \
			exit 1; \
		 fi

.PHONY: test-generate
test-generate:
	@echo "+ $@"
	@./scripts/assert-generated-files-updated.sh

.PHONY: test-lint
test-lint:
	@echo "+ $@"
	@$(GO_BIN) run ./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint run

.PHONY: test-security
test-security:
	@echo "+ $@"
	@$(GO_BIN) run ./vendor/honnef.co/go/tools/cmd/staticcheck/staticcheck.go ./...
