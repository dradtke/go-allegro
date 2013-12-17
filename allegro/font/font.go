package font

/*
#cgo pkg-config: allegro_font-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_font.h>

void _al_free_string(char *data) {
	al_free(data);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"runtime"
	"unsafe"
)

type Font C.ALLEGRO_FONT

type DrawFlags int

const (
	ALIGN_LEFT    DrawFlags = C.ALLEGRO_ALIGN_LEFT
	ALIGN_CENTRE  DrawFlags = C.ALLEGRO_ALIGN_CENTRE
	ALIGN_RIGHT   DrawFlags = C.ALLEGRO_ALIGN_RIGHT
	ALIGN_INTEGER DrawFlags = C.ALLEGRO_ALIGN_INTEGER
)

// Initialise the font addon.
func Init() {
	C.al_init_font_addon()
}

// Shut down the font addon. This is done automatically at program exit, but
// can be called any time the user wishes as well.
func Shutdown() {
	C.al_shutdown_font_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_font_version())
}

// Creates a monochrome bitmap font (8x8 pixels per character).
func Builtin() (*Font, error) {
	f := C.al_create_builtin_font()
	if f == nil {
		return nil, errors.New("failed to create builtin font")
	}
	return (*Font)(f), nil
}

// Loads a font from disk. This will use al_load_bitmap_font if you pass the
// name of a known bitmap format, or else al_load_ttf_font.
func LoadFont(filename string, size, flags int) (*Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_font(filename_, C.int(size), C.int(flags))
	if f == nil {
		return nil, fmt.Errorf("failed to load font '%s'", filename)
	}
	font := (*Font)(f)
	runtime.SetFinalizer(font, font.Destroy)
	return font, nil
}

// Load a bitmap font from. It does this by first calling al_load_bitmap and
// then al_grab_font_from_bitmap. If you want to for example load an old A4
// font, you could load the bitmap yourself, then call al_convert_mask_to_alpha
// on it and only then pass it to al_grab_font_from_bitmap.
func LoadBitmapFont(filename string) (*Font, error) {
	filename_ := C.CString(filename)
	defer C._al_free_string(filename_)
	f := C.al_load_bitmap_font(filename_)
	if f == nil {
		return nil, fmt.Errorf("failed to load bitmap font '%s'", filename)
	}
	font := (*Font)(f)
	runtime.SetFinalizer(font, font.Destroy)
	return font, nil
}

// Creates a new font from an Allegro bitmap. You can delete the bitmap after
// the function returns as the font will contain a copy for itself.
func GrabFontFromBitmap(bmp *allegro.Bitmap, ranges [][2]int) (*Font, error) {
	n_ranges := len(ranges) * 2
	if n_ranges == 0 {
		return nil, errors.New("no ranges specified")
	}
	c_ranges := make([]C.int, n_ranges)
	for i := 0; i<len(ranges); i++ {
		for j := 0; j<len(ranges[i]); j++ {
			c_ranges[2*i + j] = C.int(ranges[i][j])
		}
	}
	f := C.al_grab_font_from_bitmap((*C.ALLEGRO_BITMAP)(unsafe.Pointer(bmp)),
		C.int(n_ranges), (*C.int)(unsafe.Pointer(&c_ranges[0])))
	if f == nil {
		return nil, errors.New("failed to grab font from bitmap")
	}
	return (*Font)(f), nil
}

// Writes the NUL-terminated string text onto bmp at position x, y, using the
// specified font.
func DrawText(font *Font, color allegro.Color, x, y float32, flags DrawFlags, text string) {
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	C.al_draw_text((*C.ALLEGRO_FONT)(font),
		*((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color))), // is there an easier way to get this converted?
		C.float(x),
		C.float(y),
		C.int(flags),
		text_)
}

// Like al_draw_text, but justifies the string to the region x1-x2.
func DrawJustifiedText(font *Font, color allegro.Color, x1, x2, y, diff float32, flags DrawFlags, text string) {
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	C.al_draw_justified_text((*C.ALLEGRO_FONT)(font),
		*((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color))), // is there an easier way to get this converted?
		C.float(x1),
		C.float(x2),
		C.float(y),
		C.float(diff),
		C.int(flags),
		text_)
}

func DrawTextf(font *Font, color allegro.Color, x, y float32, flags DrawFlags, format string, a ...interface{}) {
	// C.al_draw_textf
	text_ := C.CString(fmt.Sprintf(format, a))
	defer C._al_free_string(text_)
	C.al_draw_text((*C.ALLEGRO_FONT)(font),
		*((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color))), // is there an easier way to get this converted?
		C.float(x),
		C.float(y),
		C.int(flags),
		text_)
}

func DrawJustifiedTextf(font *Font, color allegro.Color, x1, x2, y, diff float32, flags DrawFlags, format string, a ...interface{}) {
	// C.al_draw_justified_textf
	text_ := C.CString(fmt.Sprintf(format, a))
	defer C._al_free_string(text_)
	C.al_draw_justified_text((*C.ALLEGRO_FONT)(font),
		*((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color))), // is there an easier way to get this converted?
		C.float(x1),
		C.float(x2),
		C.float(y),
		C.float(diff),
		C.int(flags),
		text_)
}

// Frees the memory being used by a font structure. Does nothing if passed NULL.
func (f *Font) Destroy() {
	C.al_destroy_font((*C.ALLEGRO_FONT)(f))
}

// Returns the usual height of a line of text in the specified font. For bitmap
// fonts this is simply the height of all glyph bitmaps. For truetype fonts it
// is whatever the font file specifies. In particular, some special glyphs may
// be higher than the height returned here.
func (f *Font) LineHeight() int {
	return int(C.al_get_font_line_height((*C.ALLEGRO_FONT)(f)))
}

// Returns the ascent of the specified font.
func (f *Font) Ascent() int {
	return int(C.al_get_font_ascent((*C.ALLEGRO_FONT)(f)))
}

// Returns the descent of the specified font.
func (f *Font) Descent() int {
	return int(C.al_get_font_descent((*C.ALLEGRO_FONT)(f)))
}

// Calculates the length of a string in a particular font, in pixels.
func (f *Font) TextWidth(text string) int {
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	return int(C.al_get_text_width((*C.ALLEGRO_FONT)(f), text_))
}

// Sometimes, the al_get_text_width and al_get_font_line_height functions are
// not enough for exact text placement, so this function returns some
// additional information.
func (f *Font) TextDimensions(text string) (bbx, bby, bbw, bbh int) {
	var cbbx, cbby, cbbw, cbbh C.int
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	C.al_get_text_dimensions((*C.ALLEGRO_FONT)(f), text_,
		&cbbx, &cbby, &cbbw, &cbbh)
	return int(cbbx), int(cbby), int(cbbw), int(cbbh)
}

