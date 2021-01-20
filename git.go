package main

/*
#cgo pkg-config: libgit2 
#include "git2.h"
*/
import "C"
import (
	"unsafe"
	"errors"
)

type Repository struct {
	instance *C.git_repository
	path string
}

func (repo *Repository) Open(path string) int {
	C.git_libgit2_init()
	
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var ptr *C.git_repository
	ret := C.git_repository_open(&ptr, cpath)
	if ret == 0 {
		repo.instance = ptr
	}

	return int(ret)
}

func (repo *Repository) Close() {
	C.git_libgit2_shutdown()
}

func (repo *Repository) IsIgnored(path string) (bool, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var ignored C.int
	ret := C.git_ignore_path_is_ignored(&ignored, repo.instance, cpath)
	if ret < 0 {
		return false, errors.New("Could not process ignore rules.")
	}
	return ignored == 1, nil
}

