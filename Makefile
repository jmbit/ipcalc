
PLATFORMS := linux darwin windows
ARCHITECTURES := amd64 arm64
APP_NAME := ipcalc


.PHONY: all clean install test dev deps list no_targets__

ipcalc: deps
	go build -o ipcalc main.go

clean:
	rm -f ipcalc
	rm -f $(OUTPUT_DIR)/$(APP_NAME)-*

install: ipcalc
	install -m 0755 ipcalc /usr/local/bin/ipcalc

test:
	TMPDIR=${PWD}/tmp go test -v ./...

dev: clean ipcalc
	TMPDIR=${PWD}/tmp go run .

deps:
	go get .
	go mod tidy

all: deps $(PLATFORMS)

$(PLATFORMS): 
	@for GOOS in $(PLATFORMS); do \
		for GOARCH in $(ARCHITECTURES); do \
			export GOOS=$$GOOS; \
			export GOARCH=$$GOARCH; \
			OUTPUT_NAME=$(OUTPUT_DIR)/$(APP_NAME)-$$GOOS-$$GOARCH; \
			go build -o $$OUTPUT_NAME; \
		done \
	done

docker:
	docker build --output type=local,dest=release . 
