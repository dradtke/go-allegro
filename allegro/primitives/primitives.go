// Package primitives provides support for Allegro's primitives addon.
package primitives

/*
#cgo pkg-config: allegro_primitives-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>
*/
import "C"
import (
	"errors"
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

func col(color allegro.Color) C.ALLEGRO_COLOR {
	return *((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color)))
}

type Point struct {
	X float32
	Y float32
}

// Initializes the primitives addon.
func Init() error {
	ok := bool(C.al_init_primitives_addon())
	if !ok {
		return errors.New("failed to initialize primitives addon")
	}
	return nil
}

// Shut down the primitives addon. This is done automatically at program exit,
// but can be called any time the user wishes as well.
func Shutdown() {
	C.al_shutdown_primitives_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_primitives_version())
}

// Draws a line segment between two points.
func DrawLine(p1, p2 Point, color allegro.Color, thickness float32) {
	C.al_draw_line(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		col(color),
		C.float(thickness))
}

// Draws an outlined triangle.
func DrawTriangle(p1, p2, p3 Point, color allegro.Color, thickness float32) {
	C.al_draw_triangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		C.float(p3.X),
		C.float(p3.Y),
		col(color),
		C.float(thickness))
}

// Draws a filled triangle.
func DrawFilledTriangle(p1, p2, p3 Point, color allegro.Color) {
	C.al_draw_filled_triangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		C.float(p3.X),
		C.float(p3.Y),
		col(color))
}

// Draws an outlined rectangle.
func DrawRectangle(p1, p2 Point, color allegro.Color, thickness float32) {
	C.al_draw_rectangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		col(color),
		C.float(thickness))
}

// Draws a filled rectangle.
func DrawFilledRectangle(p1, p2 Point, color allegro.Color) {
	C.al_draw_filled_rectangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		col(color))
}

// Draws an outlined rounded rectangle.
func DrawRoundedRectangle(p1, p2 Point, rx, ry float32, color allegro.Color, thickness float32) {
	C.al_draw_rounded_rectangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		C.float(rx),
		C.float(ry),
		col(color),
		C.float(thickness))
}

// Draws an filled rounded rectangle.
func DrawFilledRoundedRectangle(p1, p2 Point, rx, ry float32, color allegro.Color) {
	C.al_draw_filled_rounded_rectangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		C.float(rx),
		C.float(ry),
		col(color))
}

// Calculates an elliptical arc, and sets the vertices in the destination
// buffer to the calculated positions. If thickness <= 0, then num_points of
// points are required in the destination, otherwise twice as many are needed.
// The destination buffer should consist of regularly spaced (by distance of
// stride bytes) doublets of floats, corresponding to x and y coordinates of
// the vertices.
func CalculateArc(stride int, center Point, rx, ry, start_theta, delta_theta, thickness float32, num_points int) []Point {
	if num_points == 0 {
		return make([]Point, 0)
	}
	buffer_length := num_points
	if thickness <= 0 {
		buffer_length *= 2
	}
	cbuf := make([]C.float, buffer_length*2)
	C.al_calculate_arc(
		(*C.float)(unsafe.Pointer(&cbuf[0])),
		C.int(stride),
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		C.float(start_theta),
		C.float(delta_theta),
		C.float(thickness),
		C.int(num_points))
	buf := make([]Point, buffer_length)
	for i := 0; i < (buffer_length*2); i += 2 {
		buf[i/2] = Point{float32(cbuf[i]), float32(cbuf[i+1])}
	}
	return buf
}

// Draws a pieslice (outlined circular sector).
func DrawPieslice(center Point, r, start_theta, delta_theta float32, color allegro.Color, thickness float32) {
	C.al_draw_pieslice(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		C.float(start_theta),
		C.float(delta_theta),
		col(color),
		C.float(thickness))
}

func DrawFilledPieslice(center Point, r, start_theta, delta_theta float32, color allegro.Color) {
	C.al_draw_pieslice(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		C.float(start_theta),
		C.float(delta_theta),
		col(color))
}

// Draws an outlined ellipse.
func DrawEllipse(center Point, rx, ry float32, color allegro.Color, thickness float32) {
	C.al_draw_ellipse(
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		col(color),
		C.float(thickness))
}

