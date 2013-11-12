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

func Init() {
	C.al_init_font_addon()
}

func Shutdown() {
	C.al_shutdown_font_addon()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_font_version())
}

func Builtin() (*Font, error) {
	f := C.al_create_builtin_font()
	if f == nil {
		return nil, errors.New("failed to create builtin font")
	}
	return (*Font)(f), nil
}

// TODO: find out what flags this supports and create its own type
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

func (f *Font) Destroy() {
	C.al_destroy_font((*C.ALLEGRO_FONT)(f))
}

func (f *Font) LineHeight() int {
	return int(C.al_get_font_line_height((*C.ALLEGRO_FONT)(f)))
}

func (f *Font) Ascent() int {
	return int(C.al_get_font_ascent((*C.ALLEGRO_FONT)(f)))
}

func (f *Font) Descent() int {
	return int(C.al_get_font_descent((*C.ALLEGRO_FONT)(f)))
}

func (f *Font) TextWidth(text string) int {
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	return int(C.al_get_text_width((*C.ALLEGRO_FONT)(f), text_))
}

func (f *Font) TextDimensions(text string) (bbx, bby, bbw, bbh int) {
	var cbbx, cbby, cbbw, cbbh C.int
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	C.al_get_text_dimensions((*C.ALLEGRO_FONT)(f), text_,
		&cbbx, &cbby, &cbbw, &cbbh)
	return int(cbbx), int(cbby), int(cbbw), int(cbbh)
}
