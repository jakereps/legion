BINDIR ?= bin
BINARY ?= legion

default: build

test:
	go test -cover -v ./...

build: $(BINDIR)/$(BINARY)

$(BINDIR)/$(BINARY):
	go build -o $(BINDIR)/$(BINARY) ./...

$(BINDIR):
	mkdir -p $(BINDIR)

.PHONY: clean
clean:
	go clean
	rm $(BINDIR)/$(BINARY)
