package main

/*
#cgo pkg-config: libgit2 zlib
#include "git2.h"
*/
import "C"
import (
	"unsafe"
)

type Repository struct {
	ptr *C.git_repository
}

func newRepositoryFromC(ptr *C.git_repository) *Repository {
	repo := &Repository{ptr: ptr}

	return repo
}

func OpenRepository(path string) (*Repository, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var ptr *C.git_repository
	ret := C.git_repository_open(&ptr, cpath)
	if ret < 0 {
		//return nil, MakeGitError(ret)

	}

	return newRepositoryFromC(ptr), nil
}

func (v *Repository) IsPathIgnored(path string) (bool, error) {
	var ignored C.int

	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	ret := C.git_ignore_path_is_ignored(&ignored, v.ptr, cpath)
	if ret < 0 {
		//return false, MakeGitError(ret)
	}
	return ignored == 1, nil
}

