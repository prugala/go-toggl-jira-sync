EXECUTABLE=toggl-jira-sync
BUILD_DIR=./build/
WINDOWS=$(BUILD_DIR)$(EXECUTABLE)_windows_amd64.exe
LINUX=$(BUILD_DIR)$(EXECUTABLE)_linux_amd64
DARWIN=$(BUILD_DIR)$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: all test clean

build: windows linux darwin ## Build binaries
	cp .env $(BUILD_DIR).env
	cp users.json.dist $(BUILD_DIR)users.json
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  .

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  .

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  .

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
