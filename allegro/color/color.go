package color

/*
#cgo pkg-config: allegro_color-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_color.h>

void free_string(char *data) {
	al_free(data);
}
*/
import "C"
import (
	"fmt"
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

func HsvToRgb(hue, saturation, value float32) (red, green, blue float32) {
	var cred, cgreen, cblue C.float
	C.al_color_hsv_to_rgb(
		C.float(hue),
		C.float(saturation),
		C.float(value),
		&cred,
		&cgreen,
		&cblue)
	return float32(cred), float32(cgreen), float32(cblue)
}

func HtmlToRgb(str string) (red, green, blue float32) {
	str_ := C.CString(str)
	defer C.free_string(str_)
	var cred, cgreen, cblue C.float
	C.al_color_html_to_rgb(
		str_,
		&cred,
		&cgreen,
		&cblue)
	return float32(cred), float32(cgreen), float32(cblue)
}

func RgbToCmyk(red, green, blue float32) (cyan, magenta, yellow, key float32) {
	var ccyan, cmagenta, cyellow, ckey C.float
	C.al_color_rgb_to_cmyk(
		C.float(red),
		C.float(green),
		C.float(blue),
		&ccyan,
		&cmagenta,
		&cyellow,
		&ckey)
	return float32(ccyan), float32(cmagenta), float32(cyellow), float32(ckey)
}

func RgbToHsl(red, green, blue float32) (hue, saturation, lightness float32) {
	var chue, csaturation, clightness C.float
	C.al_color_rgb_to_hsl(
		C.float(red),
		C.float(green),
		C.float(blue),
		&chue,
		&csaturation,
		&clightness)
	return float32(chue), float32(csaturation), float32(clightness)
}

func RgbToHsv(red, green, blue float32) (hue, saturation, value float32) {
	var chue, csaturation, cvalue C.float
	C.al_color_rgb_to_hsv(
		C.float(red),
		C.float(green),
		C.float(blue),
		&chue,
		&csaturation,
		&cvalue)
	return float32(chue), float32(csaturation), float32(cvalue)
}

func RgbToHtml(red, green, blue float32) string {
	var cstr *C.char
	C.al_color_rgb_to_html(
		C.float(red),
		C.float(green),
		C.float(blue),
		cstr)
	if cstr != nil {
		return C.GoString(cstr)
	}
	return ""
}

func RgbToYuv(red, green, blue float32) (y, u, v float32) {
	var cy, cu, cv C.float
	C.al_color_rgb_to_yuv(
		C.float(red),
		C.float(green),
		C.float(blue),
		&cy,
		&cu,
		&cv)
	return float32(cy), float32(cu), float32(cv)
}

func NameToRgb(name Name) (red, green, blue float32, err error) {
	name_ := C.CString(string(name))
	defer C.free_string(name_)
	var cred, cgreen, cblue C.float
	ok := bool(C.al_color_name_to_rgb(name_, &cred, &cgreen, &cblue))
	if !ok {
		return 0, 0, 0, fmt.Errorf("unrecognized color name '%s'", name)
	}
	return float32(cred), float32(cgreen), float32(cblue), nil
}

func RgbToName(red, green, blue float32) string {
	cname := C.al_color_rgb_to_name(
		C.float(red),
		C.float(green),
		C.float(blue))
	if cname != nil {
		return C.GoString(cname)
	}
	return ""
}

func Yuv(y, u, v float32) allegro.Color {
	col := C.al_color_yuv(
		C.float(y),
		C.float(u),
		C.float(v))
	return al(col)
}

func YuvToRgb(y, u, v float32) (red, green, blue float32) {
	var cred, cgreen, cblue C.float
	C.al_color_yuv_to_rgb(
		C.float(y),
		C.float(u),
		C.float(v),
		&cred,
		&cgreen,
		&cblue)
	return float32(cred), float32(cgreen), float32(cblue)
}
