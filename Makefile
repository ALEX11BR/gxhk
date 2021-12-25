BIN := gxhk
DESTDIR := 
PREFIX := /usr/local

.PHONY: build clean install uninstall

build:
	go build

clean:
	go clean

install: $(BIN)
	install -Dm755 $(BIN) $(DESTDIR)$(PREFIX)/bin/$(BIN)
	install -Dm755 gxhkrc $(DESTDIR)/etc/gxhkrc

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BIN)
	rm -f $(DESTDIR)/etc/gxhkrc