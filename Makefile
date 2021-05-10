SHELL=bash

DEFAULT_ENV_FILE := ~/.config/gui/gui.env
# If CSDIP_ENV variable is set we use it to define environment file path
ENV_FILE := ${gui_ENV}

# If env file is not set using CSDIP_ENV we try a default location
ifeq ($(ENV_FILE),)
	ifneq ("$(wildcard $(DEFAULT_ENV_FILE))","")
		ENV_FILE := $(DEFAULT_ENV_FILE)
	endif
endif

# If env file path is set we include it
ifneq ($(ENV_FILE),)
$(info including ENV file $(ENV_FILE))
include $(ENV_FILE)
endif

export GO111MODULE=on

COVERAGE_FILE:=coverage.out
COVERAGE_HTML:=coverage.html

GUI_BIN:=dist/gui
GUI_SRC:=./cmd/gui
#LAUNCHER_BIN:=dist/hdetect_launcher
#LAUNCHER_SRC:=./cmd/launcher

DIST = dist

LAST_TAGGED=$(shell git rev-list --tags --max-count=1)
VERSION=$(shell git describe --tags $(LAST_TAGGED))
BUILDTIME=$(shell TZ=GMT date "+%Y-%m-%d_%H:%M_GMT")
GITCOMMIT=$(shell git rev-parse --short HEAD 2>/dev/null)
GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)

PKG=github.com/suchy1105/GUIcontroler

VERSION_TAG=-X $(PKG)/config.AppVersion=$(VERSION)
COMMIT_TAG=-X $(PKG)/config.GitCommit=$(GITCOMMIT)
BRANCH_TAG=-X $(PKG)/config.GitBranch=$(GITBRANCH)
BUILDTIME_TAG=-X $(PKG)/config.BuildTime=$(BUILDTIME)
ARCH_TAG=-X $(PKG)/config.Arch=x64
LDFLAGS=-s -w $(VERSION_TAG) $(COMMIT_TAG) $(BRANCH_TAG) $(BUILDTIME_TAG) $(ARCH_TAG)

.PHONY: all
all: clean prereq build
#all: clean prereq test build

.PHONY: prereq
prereq:
	go get golang.org/x/lint/golint
#	go get golang.org/x/tools/cmd/goimports

dirs: $(DIST)

$(DIST):
	mkdir -p $@

.PHONY: clean
clean:
	rm -rf dist

#.PHONY: test
#test:
#	go test --covermode=count -coverprofile=$(COVERAGE_FILE) ./... && tail -q -n +2 $(COVERAGE_FILE) | go tool cover -func=$(COVERAGE_FILE)
#	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

#.PHONY: race
#race:
#	go test -race -short ./...

#.PHONY: msan
#msan:
#	go test -msan -short ./...

#.PHONY: lint
#lint:
#	golint -set_exit_status ./...

.PHONY: fmt
fmt:
	goimports -l ./


.PHONY: build $(GUI_SRC)
build: dirs $(GUI_BIN)

$(GUI_BIN): dist
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags '$(LDFLAGS)' -o $(GUI_BIN) -v $(GUI_SRC)


