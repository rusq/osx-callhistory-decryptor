SHELL=/bin/sh
OUTPUT=callhistory
SRC=$(wildcard *.go)
GO=go

.PHONY: all clean gofiles
.NOTPARALLEL: clean

all: $(OUTPUT)


$(OUTPUT): $(SRC)
	$(GO) build -ldflags "-s -w" -o $@

gofiles:
	$(GO) list -f '{{.GoFiles}}'

clean:
	-rm $(OUTPUT)
