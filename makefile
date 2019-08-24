# Variables
SEMVER := 0.1.0
COMMIT := $(shell git rev-list -1 HEAD | cut -c1-7)
VERSION := $(shell if [ "$$(git branch | grep \* | cut -d ' ' -f2)" = "master" ]; then echo $(SEMVER); else echo $(SEMVER)-$(COMMIT); fi)
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

# Default target
default: test

# Output the semantic version
semver:
	@echo $(SEMVER)

# Output the current commit hash
commit:
	@echo $(COMMIT)

# Output the full version name
version:
	@echo $(VERSION)

#
# Builders
#

build:
	$(shell if [ ! -d ./bin ]; then mkdir ./bin; fi)
	go build -o ./bin $(LDFLAGS) ./...

install:
	go install $(LDFLAGS) ./...

#
# Tests
#

test:
	go test -v -cover ./...

codecov:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
