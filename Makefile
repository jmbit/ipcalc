ipcalc: deps
	go build -o ipcalc main.go

clean:
	rm -f ipcalc

install: ipcalc
	install -m 0755 ipcalc /usr/local/bin/ipcalc

dev: clean ipcalc

deps:
	go get .
	go mod tidy
