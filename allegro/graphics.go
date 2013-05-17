package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Color struct {
	ptr C.ALLEGRO_COLOR
}

type Bitmap struct {
	// bitmap format?
	Width, Height int
	ptr *C.ALLEGRO_BITMAP
}

type DrawFlags int
const (
	FlipNone       DrawFlags = 0
	FlipHorizontal DrawFlags = C.ALLEGRO_FLIP_HORIZONTAL
	FlipVertical   DrawFlags = C.ALLEGRO_FLIP_VERTICAL
)

func (bmp *Bitmap) Destroy() {
	C.al_destroy_bitmap(bmp.ptr)
}

func MapRGB(r, g, b byte) *Color {
	return &Color{ptr:C.al_map_rgb(cbyte(r), cbyte(g), cbyte(b))};
}

func ClearToColor(c *Color) {
	C.al_clear_to_color(c.ptr)
}

func (bmp *Bitmap) Draw(dx, dy float32, flags DrawFlags) {
	C.al_draw_bitmap(bmp.ptr, cfloat(dx), cfloat(dy), cint(int(flags)))
}

func LoadBitmap(filename string) *Bitmap {
	filename_ := C.CString(filename) ; defer FreeString(filename_)
	bmp := C.al_load_bitmap(filename_)
	if bmp == nil {
		return nil
	}
	return &Bitmap{Width:int(C.al_get_bitmap_width(bmp)), Height:int(C.al_get_bitmap_height(bmp)), ptr:bmp}
}
