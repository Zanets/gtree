#/bin/sh

set -ex

version=0.27.7
curl -OL https://github.com/libgit2/libgit2/archive/v$version.tar.gz
tar zxf v$version.tar.gz
rm v$version.tar.gz

cd libgit2-$version
mkdir build && cd build

cmake -DBUILD_SHARED_LIBS=OFF \
	  -DBUILD_CLAR=OFF \
	  -DUSE_SSH=OFF \
	  -DUSE_HTTPS=OFF \
	  -DCMAKE_INSTALL_PREFIX=$GOPATH \
	  ..

cmake --build . --target install
