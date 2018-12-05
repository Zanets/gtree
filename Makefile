default: build

all: prepre lib build

prepare:
	mkdir build_lib

lib: 
	cd build_lib; ../scripts/install_libgit2.sh

build:
	PKG_CONFIG_PATH="${GOPATH}/lib/pkgconfig:${PKG_CONFIG_PATH}" go build

clean:
	rm -r build_lib
	rm gtree
