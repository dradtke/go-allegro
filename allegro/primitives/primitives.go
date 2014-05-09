// Package primitives provides support for Allegro's primitives addon.
package primitives

/*
#cgo pkg-config: allegro_primitives-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>

int get_stride() {
	return 2 * sizeof(float);
}
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

type Vertex struct {
	X, Y, Z float32
	Color   allegro.Color
	U, V    float32

	raw    C.ALLEGRO_VERTEX
	inited bool
}

func (v Vertex) init() {
	if v.inited {
		return
	}
	v.raw.x = C.float(v.X)
	v.raw.y = C.float(v.Y)
	v.raw.z = C.float(v.Z)
	v.raw.color = *((*C.ALLEGRO_COLOR)(unsafe.Pointer(&v.Color)))
	v.raw.u = C.float(v.U)
	v.raw.v = C.float(v.V)
	v.inited = true
}

type VertexElement struct {
    Attribute PrimAttr
    Storage PrimStorage
    Offset int

	raw    C.ALLEGRO_VERTEX_ELEMENT
	inited bool
}

func (v VertexElement) init() {
    if v.inited {
        return
    }
    v.raw.attribute = C.int(v.Attribute)
    v.raw.storage = C.int(v.Storage)
    v.raw.offset = C.int(v.Offset)
    v.inited = true
}

type VertexDecl C.ALLEGRO_VERTEX_DECL

type PrimType int

const (
	PRIM_POINT_LIST     PrimType = C.ALLEGRO_PRIM_POINT_LIST
	PRIM_LINE_LIST      PrimType = C.ALLEGRO_PRIM_LINE_LIST
	PRIM_LINE_STRIP     PrimType = C.ALLEGRO_PRIM_LINE_STRIP
	PRIM_LINE_LOOP      PrimType = C.ALLEGRO_PRIM_LINE_LOOP
	PRIM_TRIANGLE_LIST  PrimType = C.ALLEGRO_PRIM_TRIANGLE_LIST
	PRIM_TRIANGLE_STRIP PrimType = C.ALLEGRO_PRIM_TRIANGLE_STRIP
	PRIM_TRIANGLE_FAN   PrimType = C.ALLEGRO_PRIM_TRIANGLE_FAN
)

type PrimAttr int

const (
	PRIM_POSITION        PrimAttr = C.ALLEGRO_PRIM_POSITION
	PRIM_COLOR_ATTR      PrimAttr = C.ALLEGRO_PRIM_COLOR_ATTR
	PRIM_TEX_COORD       PrimAttr = C.ALLEGRO_PRIM_TEX_COORD
	PRIM_TEX_COORD_PIXEL PrimAttr = C.ALLEGRO_PRIM_TEX_COORD_PIXEL
)

type PrimStorage int

const (
	PRIM_FLOAT_2 PrimStorage = C.ALLEGRO_PRIM_FLOAT_2
	PRIM_FLOAT_3 PrimStorage = C.ALLEGRO_PRIM_FLOAT_3
	PRIM_SHORT_2 PrimStorage = C.ALLEGRO_PRIM_SHORT_2
)

// Initializes the primitives addon.
func Install() error {
	ok := bool(C.al_init_primitives_addon())
	if !ok {
		return errors.New("failed to initialize primitives addon")
	}
	return nil
}

// Shut down the primitives addon. This is done automatically at program exit,
// but can be called any time the user wishes as well.
func Uninstall() {
	C.al_shutdown_primitives_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_primitives_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
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
func CalculateArc(center Point, rx, ry, start_theta, delta_theta, thickness float32, num_points int) []Point {
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
		C.get_stride(),
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		C.float(start_theta),
		C.float(delta_theta),
		C.float(thickness),
		C.int(num_points))
	buf := make([]Point, buffer_length)
	for i := 0; i < (buffer_length * 2); i += 2 {
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

// Draws a filled pieslice (filled circular sector).
func DrawFilledPieslice(center Point, r, start_theta, delta_theta float32, color allegro.Color) {
	C.al_draw_filled_pieslice(
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

// Draws a filled ellipse.
func DrawFilledEllipse(center Point, rx, ry float32, color allegro.Color) {
	C.al_draw_filled_ellipse(
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		col(color))
}

// Draws an outlined circle.
func DrawCircle(center Point, r float32, color allegro.Color, thickness float32) {
	C.al_draw_circle(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		col(color),
		C.float(thickness))
}

// Draws a filled circle.
func DrawFilledCircle(center Point, r float32, color allegro.Color) {
	C.al_draw_filled_circle(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		col(color))
}

// Draws an arc.
func DrawArc(center Point, r, start_theta, delta_theta float32, color allegro.Color, thickness float32) {
	C.al_draw_arc(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		C.float(start_theta),
		C.float(delta_theta),
		col(color),
		C.float(thickness))
}

// Draws an elliptical arc.
func DrawEllipticalArc(center Point, rx, ry, start_theta, delta_theta float32, color allegro.Color, thickness float32) {
	C.al_draw_elliptical_arc(
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		C.float(start_theta),
		C.float(delta_theta),
		col(color),
		C.float(thickness))
}

// Calculates a Bézier spline given 4 control points. If thickness <= 0, then
// num_segments of points are required in the destination, otherwise twice as
// many are needed. The destination buffer should consist of regularly spaced
// (by distance of stride bytes) doublets of floats, corresponding to x and y
// coordinates of the vertices.
func CalculateSpline(points [4]Point, thickness float32, num_segments int) []Point {
	if num_segments == 0 {
		return make([]Point, 0)
	}
	buffer_length := num_segments
	if thickness <= 0 {
		buffer_length *= 2
	}
	cpoints := []C.float{
		C.float(points[0].X), C.float(points[0].Y),
		C.float(points[1].X), C.float(points[1].Y),
		C.float(points[2].X), C.float(points[2].Y),
		C.float(points[3].X), C.float(points[3].Y),
	}
	cbuf := make([]C.float, buffer_length*2)
	C.al_calculate_spline(
		(*C.float)(unsafe.Pointer(&cbuf[0])),
		C.get_stride(),
		(*C.float)(unsafe.Pointer(&cpoints[0])),
		C.float(thickness),
		C.int(num_segments))
	buf := make([]Point, buffer_length)
	for i := 0; i < (buffer_length * 2); i += 2 {
		buf[i/2] = Point{float32(cbuf[i]), float32(cbuf[i+1])}
	}
	return buf
}

// Draws a Bézier spline given 4 control points.
func DrawSpline(points [4]Point, color allegro.Color, thickness float32) {
	cpoints := []C.float{
		C.float(points[0].X), C.float(points[0].Y),
		C.float(points[1].X), C.float(points[1].Y),
		C.float(points[2].X), C.float(points[2].Y),
		C.float(points[3].X), C.float(points[3].Y),
	}
	C.al_draw_spline(
		(*C.float)(unsafe.Pointer(&cpoints[0])),
		col(color),
		C.float(thickness))
}

// Calculates a ribbon given an array of points. The ribbon will go through all
// of the passed points. If thickness <= 0, then num_segments of points are
// required in the destination buffer, otherwise twice as many are needed. The
// destination and the points buffer should consist of regularly spaced
// doublets of floats, corresponding to x and y coordinates of the vertices.
func CalculateRibbon(points []Point, color allegro.Color, thickness float32, num_segments int) []Point {
	if num_segments == 0 {
		return make([]Point, 0)
	}
	buffer_length := num_segments
	if thickness <= 0 {
		buffer_length *= 2
	}
	cpoints := []C.float{
		C.float(points[0].X), C.float(points[0].Y),
		C.float(points[1].X), C.float(points[1].Y),
		C.float(points[2].X), C.float(points[2].Y),
		C.float(points[3].X), C.float(points[3].Y),
	}
	cbuf := make([]C.float, buffer_length*2)
	C.al_calculate_ribbon(
		(*C.float)(unsafe.Pointer(&cbuf[0])),
		C.get_stride(),
		(*C.float)(unsafe.Pointer(&cpoints[0])),
		C.get_stride(),
		C.float(thickness),
		C.int(num_segments))
	buf := make([]Point, buffer_length)
	for i := 0; i < (buffer_length * 2); i += 2 {
		buf[i/2] = Point{float32(cbuf[i]), float32(cbuf[i+1])}
	}
	return buf
}

// Draws a series of straight lines given an array of points. The ribbon will
// go through all of the passed points.
func DrawRibbon(points []Point, color allegro.Color, thickness float32, num_segments int) {
	cpoints := make([]C.float, len(points)*2)
	for i := 0; i < len(points)*2; i += 2 {
		cpoints[i] = C.float(points[i/2].X)
		cpoints[i+1] = C.float(points[i/2].Y)
	}
	C.al_draw_ribbon(
		(*C.float)(unsafe.Pointer(&cpoints[0])),
		C.get_stride(),
		col(color),
		C.float(thickness),
		C.int(num_segments))
}

// Draws a subset of the passed vertex buffer.
func DrawPrim(vertices []Vertex, decl *VertexDecl, texture *allegro.Bitmap, start, end int, prim_type PrimType) int {
	vertices_ := make([]C.ALLEGRO_VERTEX, len(vertices))
	for i, vertex := range vertices {
		// how does this perform?
		vertex.init()
		vertices_[i] = vertex.raw
	}
	drawn := C.al_draw_prim(unsafe.Pointer(&vertices_[0]),
		(*C.ALLEGRO_VERTEX_DECL)(decl),
		(*C.ALLEGRO_BITMAP)(texture),
		C.int(start),
		C.int(end),
		C.int(prim_type))
	return int(drawn)
}

// Draws a subset of the passed vertex buffer. This function uses an index
// array to specify which vertices to use.
func DrawIndexedPrim(vertices []Vertex, decl *VertexDecl, texture *allegro.Bitmap, indices []int, num_vertices int, prim_type PrimType) int {
	vertices_ := make([]C.ALLEGRO_VERTEX, len(vertices))
	for i, vertex := range vertices {
		// how does this perform?
		vertex.init()
		vertices_[i] = vertex.raw
	}
	indices_ := make([]C.int, len(indices))
	for i, index := range indices {
		indices_[i] = C.int(index)
	}
	drawn := C.al_draw_indexed_prim(unsafe.Pointer(&vertices_[0]),
		(*C.ALLEGRO_VERTEX_DECL)(decl),
		(*C.ALLEGRO_BITMAP)(texture),
		(*C.int)(unsafe.Pointer(&indices_[0])),
		C.int(num_vertices),
		C.int(prim_type))
	return int(drawn)
}

// Creates a vertex declaration, which describes a custom vertex format.
func CreateVertexDecl(elements []VertexElement, stride int) *VertexDecl {
    elements_ := make([]C.ALLEGRO_VERTEX_ELEMENT, len(elements))
    for i, element := range elements {
		// how does this perform?
        element.init()
        elements_[i] = element.raw
    }
    return (*VertexDecl)(C.al_create_vertex_decl((*C.ALLEGRO_VERTEX_ELEMENT)(unsafe.Pointer(&elements_[0])), C.int(stride)))
}

// Destroys a vertex declaration.
func (v *VertexDecl) Destroy() {
    C.al_destroy_vertex_decl((*C.ALLEGRO_VERTEX_DECL)(v))
}

