default: build

prepare:
	mkdir build_lib

lib: prepare
	cd build_lib; ../scripts/install_libgit2.sh

build: lib
	PKG_CONFIG_PATH="${GOPATH}/lib/pkgconfig:${PKG_CONFIG_PATH}" go build

clean:
	rm -r build_lib
	rm gtree
