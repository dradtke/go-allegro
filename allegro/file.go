package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

type File C.ALLEGRO_FILE

type FileMode int

const (
	FILE_READ FileMode = 1 << iota
	FILE_WRITE
	FILE_BINARY
	FILE_EXPANDABLE // for slicing files
)

func (m FileMode) String() string {
	var buf bytes.Buffer
	if (m & FILE_READ) != 0 {
		buf.WriteString("r")
	}
	if (m & FILE_WRITE) != 0 {
		buf.WriteString("w")
	}
	if (m & FILE_BINARY) != 0 {
		buf.WriteString("b")
	}
	if (m & FILE_EXPANDABLE) != 0 {
		buf.WriteString("e")
	}
	return buf.String()
}

func OpenFile(path string, mode FileMode) (*File, error) {
	path_ := C.CString(path)
	mode_ := C.CString(mode.String())
	defer freeString(path_)
	defer freeString(mode_)
	f := C.al_fopen(path_, mode_)
	if f == nil {
		return nil, fmt.Errorf("failed to open File '%s'", path)
	}
	return (*File)(f), nil
}

func (f *File) Close() error {
	C.al_fclose((*C.ALLEGRO_FILE)(f))
	if f.HasError() {
		return LastError()
	}
	return nil
}

func (f *File) Eof() bool {
	return bool(C.al_feof((*C.ALLEGRO_FILE)(f)))
}

func (f *File) HasError() bool {
	return bool(C.al_ferror((*C.ALLEGRO_FILE)(f)))
}

func (f *File) ClearError() {
	C.al_fclearerr((*C.ALLEGRO_FILE)(f))
}

func (f *File) Read(b []byte) (n int, err error) {
	size := len(b)
	if size == 0 {
		return 0, errors.New("cannot read into empty buffer")
	}
	r := int(C.al_fread((*C.ALLEGRO_FILE)(f),
		unsafe.Pointer(&b[0]),
		C.size_t(size)))
	if r == 0 && f.Eof() {
		return r, io.EOF
	} else if f.HasError() {
		return r, errors.New("read() encountered an error")
	} else {
		return r, nil
	}
}

func (f *File) Write(b []byte) (n int, err error) {
	size := len(b)
	written := int(C.al_fwrite((*C.ALLEGRO_FILE)(f),
		unsafe.Pointer(&b[0]),
		C.size_t(size)))
	if written < size {
		return written, errors.New("write() encountered an error")
	} else {
		return written, nil
	}
}

func (f *File) Getc() (byte, error) {
	b := byte(C.al_fgetc((*C.ALLEGRO_FILE)(f)))
	if f.Eof() {
		return 0, io.EOF
	} else if f.HasError() {
		return 0, errors.New("error occurred when reading byte")
	} else {
		return b, nil
	}
}

func (f *File) Ungetc(b byte) byte {
	return byte(C.al_fungetc((*C.ALLEGRO_FILE)(f), C.int(b)))
}

func (f *File) Putc(b byte) byte {
	return byte(C.al_fputc((*C.ALLEGRO_FILE)(f), C.int(b)))
}

func (f *File) Flush() error {
	ok := bool(C.al_fflush((*C.ALLEGRO_FILE)(f)))
	if !ok {
		return LastError()
	}
	return nil
}

func (f *File) Tell() (int64, error) {
	pos := int64(C.al_ftell((*C.ALLEGRO_FILE)(f)))
	if pos == -1 {
		return 0, LastError()
	}
	return pos, nil
}

func (f *File) Seek(offset int64, whence int) (ret int64, err error) {
	var whence_ C.int
	switch whence {
	case 0:
		whence_ = C.ALLEGRO_SEEK_SET
	case 1:
		whence_ = C.ALLEGRO_SEEK_CUR
	case 2:
		whence_ = C.ALLEGRO_SEEK_END
	default:
		return 0, fmt.Errorf("unrecognized whence value: %d", whence)
	}
	ok := bool(C.al_fseek((*C.ALLEGRO_FILE)(f), C.int64_t(offset), whence_))
	if !ok {
		return 0, LastError()
	}
	pos, err := f.Tell()
	if err != nil {
		return 0, err
	}
	return pos, nil
}

func (f *File) Size() (int64, error) {
	size := int64(C.al_fsize((*C.ALLEGRO_FILE)(f)))
	if size == -1 {
		return 0, errors.New("File size could not be determined")
	}
	return size, nil
}

func (f *File) Slice(initial_size int, mode FileMode) (*File, error) {
	mode_ := C.CString(mode.String())
	defer freeString(mode_)
	s := C.al_fopen_slice((*C.ALLEGRO_FILE)(f), C.size_t(initial_size), mode_)
	if s == nil {
		return nil, errors.New("failed to slice File")
	}
	return (*File)(s), nil
}
