// Package memfile provides support for Allegro's memfile addon.
package memfile

/*
#cgo pkg-config: allegro_memfile-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_memfile.h>

void _al_free_string(char *data) {
	al_free(data);
}
*/
import "C"
import (
	"bytes"
	"errors"
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

type FileMode int

const (
	FILE_READ FileMode = 1 << iota
	FILE_WRITE
)

func (m FileMode) String() string {
	var buf bytes.Buffer
	if (m & FILE_READ) != 0 {
		buf.WriteString("r")
	}
	if (m & FILE_WRITE) != 0 {
		buf.WriteString("w")
	}
	return buf.String()
}

// Returns a file handle to the block of memory. All read and write operations
// act upon the memory directly, so it must not be freed while the file remains
// open.
func Open(mem unsafe.Pointer, size int64, mode FileMode) (*allegro.File, error) {
	mode_ := C.CString(mode.String())
	defer C._al_free_string(mode_)
	f := C.al_open_memfile(mem, C.int64_t(size), mode_)
	if f == nil {
		return nil, errors.New("failed to open memfile")
	}
	return (*allegro.File)(unsafe.Pointer(f)), nil
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_memfile_version())
}

