VERSION = '0.1.1'

all: build package

test: build
		@echo "Running tests"
		go test zsnap/*

build: clean
		@echo "Building new binary"
		bash --norc ./scripts/build.sh

package:
		@echo "Packaging up binary"
		@tar -czf zfsnap-$(VERSION).omnios.tar.gz bin/

clean:
		@echo "Cleaning up"
		@rm -fr bin
		@rm -f *tar.gz
