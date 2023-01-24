SHELL=/bin/sh
OUTPUT=callhistory
SRC=$(shell go list $(GOFLAGS) -f "{{.Dir}}:{{.GoFiles}}" . | tr -d '[]' | awk 'BEGIN{FS=":"}{n=split($$2,files," "); for (i=1; i<=n; i++) { printf ("%s/%s ",$$1, files[i]); } ; };' )
GO=go

BUILD=$(shell git describe --tags)

.PHONY: all clean test
.NOTPARALLEL: clean

all: $(OUTPUT)


$(OUTPUT): $(SRC)
	$(GO) build -ldflags "-s -w -X main.build=$(BUILD)" -o $@

test:
	$(GO) test ./... -cover

clean:
	-rm $(OUTPUT)

goreleaser:
	goreleaser check
	goreleaser release --snapshot --rm-dist
