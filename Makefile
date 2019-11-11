EXECUTABLE=toggl-jira-sync
WINDOWS=$(EXECUTABLE)_windows.exe
LINUX=$(EXECUTABLE)_linux
MACOS=$(EXECUTABLE)_macos

all: test build

test:
	go test .

build: windows linux darwin

run:
    go run .

windows: $(WINDOWS)

linux: $(LINUX)

macos: $(MACOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o ./build/$(WINDOWS) .

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o ./build/$(LINUX) .

$(MACOS):
	env GOOS=darwin GOARCH=amd64 go build -o ./build/$(DARWIN) .

clean:
	rm -f $(WINDOWS) $(LINUX) $(MACOS)