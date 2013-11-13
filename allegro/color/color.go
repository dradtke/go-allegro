package color

/*
#cgo pkg-config: allegro_color-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_color.h>
*/
import "C"
import (
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

func al(col C.ALLEGRO_COLOR) allegro.Color {
	return *(*allegro.Color)(unsafe.Pointer(&col))
}

func Version() uint32 {
	return uint32(C.al_get_allegro_color_version())
}

func Cmyk(cyan, magenta, yellow, key float32) allegro.Color {
	col := C.al_color_cmyk(
		C.float(cyan),
		C.float(magenta),
		C.float(yellow),
		C.float(key))
	return al(col)
}

func CmykToRgb(cyan, magenta, yellow, key float32) (red, green, blue float32) {
	var cred, cgreen, cblue C.float
	C.al_color_cmyk_to_rgb(
		C.float(cyan),
		C.float(magenta),
		C.float(yellow),
		C.float(key),
		&cred,
		&cgreen,
		&cblue)
	return float32(cred), float32(cgreen), float32(cblue)
}

func Hsl(hue, saturation, lightness float32) allegro.Color {
	col := C.al_color_hsl(
		C.float(hue),
		C.float(saturation),
		C.float(lightness))
	return al(col)
}

func HslToRgb(hue, saturation, lightness float32) (red, green, blue float32) {
	var cred, cgreen, cblue C.float
	C.al_color_hsl_to_rgb(
		C.float(hue),
		C.float(saturation),
		C.float(lightness),
		&cred,
		&cgreen,
		&cblue)
	return float32(cred), float32(cgreen), float32(cblue)
}
