package allegro

// #include <allegro5/allegro.h>
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

// Convert a path to its string representation, i.e. optional drive, followed
// by directory components separated by 'delim', followed by an optional
// filename.
func pathStr(path *C.ALLEGRO_PATH) string {
	return C.GoString(C.al_path_cstr(path, C.ALLEGRO_NATIVE_PATH_SEP))
}

// Creates and opens a file (real or virtual) given the path and mode. The
// current file interface is used to open the file.
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

// Make a temporary randomly named file given a filename 'template'.
func MakeTempFile(template string) string {
	template_ := C.CString(template)
	defer freeString(template_)
	var path *C.ALLEGRO_PATH
	C.al_make_temp_file(template_, &path)
	return pathStr(path)
}

// Close the given file, writing any buffered output data (if any).
func (f *File) Close() error {
	C.al_fclose((*C.ALLEGRO_FILE)(f))
	if f.HasError() {
		return LastError()
	}
	return nil
}

// Returns true if the end-of-file indicator has been set on the file, i.e. we
// have attempted to read past the end of the file.
func (f *File) Eof() bool {
	return bool(C.al_feof((*C.ALLEGRO_FILE)(f)))
}

// Returns true if the error indicator is set on the given file, i.e. there was
// some sort of previous error.
func (f *File) HasError() bool {
	return bool(C.al_ferror((*C.ALLEGRO_FILE)(f)))
}

// Clear the error indicator for the given file.
func (f *File) ClearError() {
	C.al_fclearerr((*C.ALLEGRO_FILE)(f))
}

// Read 'size' bytes into the buffer pointed to by 'ptr', from the given file.
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

// Write 'size' bytes from the buffer pointed to by 'ptr' into the given file.
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

// Read and return next byte in the given file. Returns EOF on end of file or
// if an error occurred.
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

// Ungets a single byte from a file. Pushed-back bytes are not written to the
// file, only made available for subsequent reads, in reverse order.
func (f *File) Ungetc(b byte) byte {
	return byte(C.al_fungetc((*C.ALLEGRO_FILE)(f), C.int(b)))
}

// Write a single byte to the given file. The byte written is the value of c
// cast to an unsigned char.
func (f *File) Putc(b byte) byte {
	return byte(C.al_fputc((*C.ALLEGRO_FILE)(f), C.int(b)))
}

// Flush any pending writes to the given file.
func (f *File) Flush() error {
	ok := bool(C.al_fflush((*C.ALLEGRO_FILE)(f)))
	if !ok {
		return LastError()
	}
	return nil
}

// Returns the current position in the given file, or -1 on error. errno is set
// to indicate the error.
func (f *File) Tell() (int64, error) {
	pos := int64(C.al_ftell((*C.ALLEGRO_FILE)(f)))
	if pos == -1 {
		return 0, LastError()
	}
	return pos, nil
}

// Set the current position of the given file to a position relative to that
// specified by 'whence', plus 'offset' number of bytes.
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

// Return the size of the file, if it can be determined, or -1 otherwise.
func (f *File) Size() (int64, error) {
	size := int64(C.al_fsize((*C.ALLEGRO_FILE)(f)))
	if size == -1 {
		return 0, errors.New("File size could not be determined")
	}
	return size, nil
}

// Opens a slice (subset) of an already open random access file as if it were a
// stand alone file. While the slice is open, the parent file handle must not
// be used in any way.
func (f *File) Slice(initial_size int, mode FileMode) (*File, error) {
	mode_ := C.CString(mode.String())
	defer freeString(mode_)
	s := C.al_fopen_slice((*C.ALLEGRO_FILE)(f), C.size_t(initial_size), mode_)
	if s == nil {
		return nil, errors.New("failed to slice File")
	}
	return (*File)(s), nil
}
