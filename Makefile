VERSION?=1.1.0

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}
BINARY=pepp

VET_REPORT=vet.report
LINT_REPORT=lint.report
TEST_REPORT=test.report
TEST_XUNIT_REPORT=test.report.xml

OS := $(shell uname -s)
ifeq ($(OS),Darwin)
	DYLIB=.dylib
	INSTALL=install
	LDCONFIG=
	NEBBINARY=$(BINARY)
	BUUILDLOG=
else
	DYLIB=.so
	INSTALL=sudo install
	LDCONFIG=sudo /sbin/ldconfig
	NEBBINARY=$(BINARY)-$(COMMIT)
	BUUILDLOG=-rm -f $(BINARY); ln -s $(BINARY)-$(COMMIT) $(BINARY)
endif

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.compileAt=`date +%s`"

# Build the project
.PHONY: build build-linux clean dep lint run test vet link-libs

all: clean vet fmt lint build test

dep:
	dep ensure -v

deploy-v8:
	$(INSTALL) nf/nvm/native-lib/*$(DYLIB) /usr/local/lib/
	$(LDCONFIG)

deploy-libs:
	$(INSTALL) nf/nvm/native-lib/*$(DYLIB) /usr/local/lib/
	$(INSTALL) native-lib/*$(DYLIB) /usr/local/lib/
	$(LDCONFIG)

build:
	cd cmd/pepp; go build -gcflags "-N -l" $(LDFLAGS) -o ../../$(BINARY)-$(COMMIT)
	cd cmd/crashreporter; go build -gcflags "-N -l" $(LDFLAGS) -o ../../pepp-crashreporter
	rm -f $(BINARY)
	cp -f $(BINARY)-$(COMMIT) $(BINARY)

build-linux:
	cd cmd/pepp; GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ../../$(BINARY)-linux

test:
	env GOCACHE=off go test ./... 2>&1 | tee $(TEST_REPORT); go2xunit -fail -input $(TEST_REPORT) -output $(TEST_XUNIT_REPORT)

vet:
	go vet $$(go list ./...) 2>&1 | tee $(VET_REPORT)

fmt:
	goimports -w $$(go list -f "{{.Dir}}" ./... | grep -v /vendor/)

lint:
	golint $$(go list ./...) | sed "s:^$(BUILD_DIR)/::" | tee $(LINT_REPORT)

clean:
	-rm -f $(VET_REPORT)
	-rm -f $(LINT_REPORT)
	-rm -f $(TEST_REPORT)
	-rm -f $(TEST_XUNIT_REPORT)
	-rm -f $(BINARY)
	-rm -f $(BINARY)-$(COMMIT)

