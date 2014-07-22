// Package color provides support for Allegro's color addon.
package color

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_color.h>
// #include "../util.c"
import "C"
import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

func al(col C.ALLEGRO_COLOR) allegro.Color {
	return *(*allegro.Color)(unsafe.Pointer(&col))
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_color_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}

// Return an ALLEGRO_COLOR structure from CMYK values (cyan, magenta, yellow,
// black).
func Cmyk(cyan, magenta, yellow, key float32) allegro.Color {
	col := C.al_color_cmyk(
		C.float(cyan),
		C.float(magenta),
		C.float(yellow),
		C.float(key))
	return al(col)
}

// Convert CMYK values to RGB values.
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

// Return an ALLEGRO_COLOR structure from HSL (hue, saturation, lightness)
// values.
func Hsl(hue, saturation, lightness float32) allegro.Color {
	col := C.al_color_hsl(
		C.float(hue),
		C.float(saturation),
		C.float(lightness))
	return al(col)
}

// Convert values in HSL color model to RGB color model.
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

// Convert values in HSV color model to RGB color model.
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

// Interprets an HTML styled hex number (e.g. #00faff) as a color. Components
// that are malformed are set to 0.
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

// Each RGB color can be represented in CMYK with a K component of 0 with the
// following formula:
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

// Given an RGB triplet with components in the range 0..1, return the hue in
// degrees from 0..360 and saturation and lightness in the range 0..1.
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

// Given an RGB triplet with components in the range 0..1, return the hue in
// degrees from 0..360 and saturation and value in the range 0..1.
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

// Create an HTML-style string representation of an ALLEGRO_COLOR, e.g. #00faff.
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

// Convert RGB values to YUV color space.
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

// Parameters:
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

// Given an RGB triplet with components in the range 0..1, find a color name
// describing it approximately.
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

// Return an ALLEGRO_COLOR structure from YUV values.
func Yuv(y, u, v float32) allegro.Color {
	col := C.al_color_yuv(
		C.float(y),
		C.float(u),
		C.float(v))
	return al(col)
}

// Convert YUV color values to RGB color space.
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
