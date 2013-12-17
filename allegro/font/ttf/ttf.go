// Package ttf provides support for Allegro's TTF font addon.
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

// Call this after al_init_font_addon to make al_load_font recognize ".ttf" and
// other formats supported by al_load_ttf_font.
func Init() {
	C.al_init_ttf_addon()
}

// Unloads the ttf addon again. You normally don't need to call this.
func Shutdown() {
	C.al_shutdown_ttf_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_ttf_version())
}

// Loads a TrueType font from a file using the FreeType library. Quoting from
// the FreeType FAQ this means support for many different font formats:
func LoadFont(filename string, size int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font(filename_, C.int(size), C.int(flags))
	if f == nil {
		return nil, fmt.Errorf("failed to load ttf font at '%s'", filename)
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}

// Like al_load_ttf_font, but the font is read from the file handle. The
// filename is only used to find possible additional files next to a font file.
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

// Like al_load_ttf_font, except it takes separate width and height parameters
// instead of a single size parameter.
func LoadFontStretch(filename string, w, h int, flags TtfFlags) (*font.Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_ttf_font_stretch(filename_, C.int(w), C.int(h), C.int(flags))
	if f == nil {
		return nil, fmt.Errorf("failed to load ttf font at '%s'", filename)
	}
	return (*font.Font)(unsafe.Pointer(f)), nil
}

// Like al_load_ttf_font_stretch, but the font is read from the file handle.
// The filename is only used to find possible additional files next to a font
// file.
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


