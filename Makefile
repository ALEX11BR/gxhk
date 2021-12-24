BIN := gxhk
PREFIX := /usr/local

.PHONY: build install

build:
	go build

install: $(BIN)
	install -Dm755 $(BIN) $(PREFIX)/bin/$(BIN)