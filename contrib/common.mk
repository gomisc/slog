UNAME_S = $(shell uname -s)
PROJECT_NAME = $(shell basename `git rev-parse --show-toplevel`)
MODULE = $(shell go list -m | xargs)

ifeq ($(UNAME_S),Linux)
	SHELL=/bin/bash
	.SHELLFLAGS = -o pipefail -c
endif

goreadme = $(shell command -v goreadme 2> /dev/null)
gomarkdoc = $(shell command -v gomarkdoc 2> /dev/null)
linter = $(shell command -v golangci-lint 2> /dev/null)
tools = $(shell command -v stringer 2> /dev/null)

GOREADME = goreadme -title="$(PROJECT_NAME)" -credit=false -skip-sub-packages > README.md
GOMARKDOC = gomarkdoc --output '{{.Dir}}/README.md' ./zaplogger/...

