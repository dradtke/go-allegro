package ttf

/*
#cgo pkg-config: allegro_ttf-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_ttf.h>

void _al_free_string(char *data) {
	al_free(data);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

type TtfFlags int

const (
	TTF_NO_KERNING  TtfFlags = C.ALLEGRO_TTF_NO_KERNING
	TTF_MONOCHROME  TtfFlags = C.ALLEGRO_TTF_MONOCHROME
	TTF_NO_AUTOHINT TtfFlags = C.ALLEGRO_TTF_NO_AUTOHINT
)

func Init() {
	C.al_init_ttf_addon()
}

func Shutdown() {
	C.al_shutdown_ttf_addon()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_ttf_version())
}

func LoadFont(filename string, size int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font(filename_, C.int(size), C.int(flags))
	if f == nil {
		return nil, fmt.Errorf("failed to load ttf font at '%s'", filename)
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}

func LoadFontF(file *allegro.File, filename string, size int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font_f((*C.ALLEGRO_FILE)(unsafe.Pointer(file)), filename_,
		C.int(size), C.int(flags))
	if f == nil {
		return nil, errors.New("failed to load font from file")
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}

func LoadFontStretch(filename string, w, h int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font_stretch(filename_, C.int(w), C.int(h), C.int(flags))
	if f == nil {
		return nil, fmt.Errorf("failed to load ttf font at '%s'", filename)
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}

func LoadFontStretchF(file *allegro.File, filename string, w, h int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font_stretch_f((*C.ALLEGRO_FILE)(unsafe.Pointer(file)),
		filename_, C.int(w), C.int(h), C.int(flags))
	if f == nil {
		return nil, errors.New("failed to load font from file")
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}
