package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
)

type Color C.ALLEGRO_COLOR
type Bitmap C.ALLEGRO_BITMAP

type DrawFlags int
const (
	FlipNone       DrawFlags = 0
	FlipHorizontal DrawFlags = C.ALLEGRO_FLIP_HORIZONTAL
	FlipVertical   DrawFlags = C.ALLEGRO_FLIP_VERTICAL
)

func (bmp *Bitmap) Destroy() {
	C.al_destroy_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

func (bmp *Bitmap) Width() int {
	return (int)(C.al_get_bitmap_width((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Height() int {
	return (int)(C.al_get_bitmap_height((*C.ALLEGRO_BITMAP)(bmp)))
}

func MapRGB(r, g, b byte) *Color {
	color := (Color)(C.al_map_rgb(cbyte(r), cbyte(g), cbyte(b)))
	return &color
}

func ClearToColor(c *Color) {
	C.al_clear_to_color(*(*C.ALLEGRO_COLOR)(c))
}

func (bmp *Bitmap) Draw(dx, dy float32, flags DrawFlags) {
	C.al_draw_bitmap((*C.ALLEGRO_BITMAP)(bmp), cfloat(dx), cfloat(dy), cint(int(flags)))
}

func LoadBitmap(filename string) (*Bitmap, error) {
	filename_ := C.CString(filename) ; defer FreeString(filename_)
	bmp := C.al_load_bitmap(filename_)
	if bmp == nil {
		return nil, errors.New("failed to load bitmap at '" + filename + "'")
	}
	return (*Bitmap)(bmp), nil
}
