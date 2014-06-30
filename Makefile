VERSION = '0.1.0'

all: clean build package

test:
		@echo "Running tests"
		go test zsnap/*

build:
		@echo "Building new binary"
		bash --norc ./scripts/build.sh

package:
		@echo "Packaging up binary"
		tar -czf zfsnap-$(VERSION).omnios.tar.gz bin/

clean:
		@echo "Cleaning up"
		@rm -fr bin
		@rm -f *tar.gz
