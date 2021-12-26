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
	install -Dm644 gxhk.1 $(DESTDIR)$(PREFIX)/share/man/man1/gxhk.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BIN)
	rm -f $(DESTDIR)/etc/gxhkrc
	rm -f $(DESTDIR)$(PREFIX)/share/man/man1/gxhk.1