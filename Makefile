include contrib/common.mk
export

.PHONY: fmt
fmt:
ifeq ($(tools),)
	go install golang.org/x/tools/...@latest
endif
	goimports -local $(MODULE) -w . || true
	gofmt -l -w . || true

.PHONY: gomod
gomod:
	go mod tidy -compat=1.18
	go mod download

.PHONY: lint
lint: fmt gomod
ifeq ($(linter),)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.45.2
endif
	$(linter) run -c golangci.yml

.PHONY: test
test:
	go test ./... -cover

.PHONY: docs
docs:
ifeq ($(goreadme),)
	go install github.com/posener/goreadme/cmd/goreadme@latest
endif
ifeq ($(gomarkdoc),)
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
endif
	$(GOREADME)
	$(GOMARKDOC) && git add .