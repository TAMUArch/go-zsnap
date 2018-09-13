VERSION = '0.1.2'

all: build package

test:
		@echo "Running tests"
		go test zsnap/*

build: clean
		@echo "Building new binary"
		bash --norc ./scripts/build.sh

package:
		@echo "Packaging up binary"
		@tar -czf zfsnap-$(VERSION).omnios.tar.gz bin/

deps:
		@echo "Getting dependencies"
		@go get github.com/jarosser06/go-zsnap/zsnap

clean:
		@echo "Cleaning up"
		@rm -fr bin
		@rm -f *tar.gz
