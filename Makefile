default: compile

prepare:
	mkdir -p build_lib

prepare_lib: prepare
	cd build_lib && \
	../scripts/install_libgit2.sh

compile:
	PKG_CONFIG_PATH="${GOPATH}/lib/pkgconfig:${PKG_CONFIG_PATH}" go build
