// Package primitives provides support for Allegro's primitives addon.
package primitives

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_primitives.h>
/*
static inline int get_stride() {
	return 2 * sizeof(float);
}

extern void trangulage_polygon_callback(int x, int y, int z, void *userdata);

// This is a workaround for a compilation type error.
// See https://github.com/golang/go/issues/19835
typedef void (*emit_triangle_callback_t)(int, int, int, void*);
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"

	"github.com/dradtke/go-allegro/allegro"
)

func col(color allegro.Color) C.ALLEGRO_COLOR {
	return *((*C.ALLEGRO_COLOR)(unsafe.Pointer(&color)))
}

func cVertices(vertices []Vertex) []C.ALLEGRO_VERTEX {
	vertices_ := make([]C.ALLEGRO_VERTEX, len(vertices))
	for i, vertex := range vertices {
		// how does this perform?
		vertex.init()
		vertices_[i] = vertex.raw
	}
	return vertices_
}

func cInts(v []int) []C.int {
	x := make([]C.int, len(v))
	for i := range v {
		x[i] = C.int(v[i])
	}
	return x
}

type Point struct {
	X float32
	Y float32
}

type Polyline []Point

func (p Polyline) vertices() []C.float {
	v := make([]C.float, 0, len(p)*2)
	for _, point := range p {
		v = append(v, C.float(point.X), C.float(point.Y))
	}
	return v
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
	Storage   PrimStorage
	Offset    int

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
	PRIM_LINE_LIST               = C.ALLEGRO_PRIM_LINE_LIST
	PRIM_LINE_STRIP              = C.ALLEGRO_PRIM_LINE_STRIP
	PRIM_LINE_LOOP               = C.ALLEGRO_PRIM_LINE_LOOP
	PRIM_TRIANGLE_LIST           = C.ALLEGRO_PRIM_TRIANGLE_LIST
	PRIM_TRIANGLE_STRIP          = C.ALLEGRO_PRIM_TRIANGLE_STRIP
	PRIM_TRIANGLE_FAN            = C.ALLEGRO_PRIM_TRIANGLE_FAN
)

type PrimAttr int

const (
	PRIM_POSITION        PrimAttr = C.ALLEGRO_PRIM_POSITION
	PRIM_COLOR_ATTR               = C.ALLEGRO_PRIM_COLOR_ATTR
	PRIM_TEX_COORD                = C.ALLEGRO_PRIM_TEX_COORD
	PRIM_TEX_COORD_PIXEL          = C.ALLEGRO_PRIM_TEX_COORD_PIXEL
)

type PrimStorage int

const (
	PRIM_FLOAT_2 PrimStorage = C.ALLEGRO_PRIM_FLOAT_2
	PRIM_FLOAT_3             = C.ALLEGRO_PRIM_FLOAT_3
	PRIM_SHORT_2             = C.ALLEGRO_PRIM_SHORT_2
)

type LineJoin int

const (
	LINE_JOIN_NONE  LineJoin = C.ALLEGRO_LINE_JOIN_NONE
	LINE_JOIN_BEVEL          = C.ALLEGRO_LINE_JOIN_BEVEL
	LINE_JOIN_ROUND          = C.ALLEGRO_LINE_JOIN_ROUND
	LINE_JOIN_MITER          = C.ALLEGRO_LINE_JOIN_MITER
)

type LineCap int

const (
	LINE_CAP_NONE     LineCap = C.ALLEGRO_LINE_CAP_NONE
	LINE_CAP_SQUARE           = C.ALLEGRO_LINE_CAP_SQUARE
	LINE_CAP_ROUND            = C.ALLEGRO_LINE_CAP_ROUND
	LINE_CAP_TRIANGLE         = C.ALLEGRO_LINE_CAP_TRIANGLE
	LINE_CAP_CLOSED           = C.ALLEGRO_LINE_CAP_CLOSED
)

// Initializes the primitives addon.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_init_primitives_addon
func Install() error {
	ok := bool(C.al_init_primitives_addon())
	if !ok {
		return errors.New("failed to initialize primitives addon")
	}
	return nil
}

// Returns true if the primitives addon is initialized, otherwise returns false.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_is_primitives_addon_initialized
func Installed() bool {
	return bool(C.al_is_primitives_addon_initialized())
}

// Shut down the primitives addon. This is done automatically at program exit,
// but can be called any time the user wishes as well.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_shutdown_primitives_addon
func Uninstall() {
	C.al_shutdown_primitives_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_get_allegro_primitives_version
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_primitives_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}

// Draws a line segment between two points.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_line
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_triangle
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_triangle
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_rectangle
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_rectangle
func DrawFilledRectangle(p1, p2 Point, color allegro.Color) {
	C.al_draw_filled_rectangle(
		C.float(p1.X),
		C.float(p1.Y),
		C.float(p2.X),
		C.float(p2.Y),
		col(color))
}

// Draws an outlined rounded rectangle.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_rounded_rectangle
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_rounded_rectangle
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

// When thickness <= 0 this function computes positions of num_points regularly
// spaced points on an elliptical arc. When thickness > 0 this function
// computes two sets of points, obtained as follows: the first set is obtained
// by taking the points computed in the thickness <= 0 case and shifting them
// by thickness / 2 outward, in a direction perpendicular to the arc curve. The
// second set is the same, but shifted thickness / 2 inward relative to the
// arc. The two sets of points are interleaved in the destination buffer (i.e.
// the first pair of points will be collinear with the arc center, the first
// point of the pair will be farther from the center than the second point; the
// next pair will also be collinear, but at a different angle and so on).
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_calculate_arc
func CalculateArc(center Point, rx, ry, start_theta, delta_theta, thickness float32, num_points int) Polyline {
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_pieslice
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_pieslice
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_ellipse
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_ellipse
func DrawFilledEllipse(center Point, rx, ry float32, color allegro.Color) {
	C.al_draw_filled_ellipse(
		C.float(center.X),
		C.float(center.Y),
		C.float(rx),
		C.float(ry),
		col(color))
}

// Draws an outlined circle.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_circle
func DrawCircle(center Point, r float32, color allegro.Color, thickness float32) {
	C.al_draw_circle(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		col(color),
		C.float(thickness))
}

// Draws a filled circle.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_circle
func DrawFilledCircle(center Point, r float32, color allegro.Color) {
	C.al_draw_filled_circle(
		C.float(center.X),
		C.float(center.Y),
		C.float(r),
		col(color))
}

// Draws an arc.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_arc
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_elliptical_arc
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_calculate_spline
func CalculateSpline(points [4]Point, thickness float32, num_segments int) Polyline {
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_spline
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_calculate_ribbon
func CalculateRibbon(p Polyline, color allegro.Color, thickness float32, num_segments int) Polyline {
	if num_segments == 0 {
		return make([]Point, 0)
	}
	buffer_length := num_segments
	if thickness <= 0 {
		buffer_length *= 2
	}
	points := p.vertices()
	cbuf := make([]C.float, buffer_length*2)
	C.al_calculate_ribbon(
		(*C.float)(unsafe.Pointer(&cbuf[0])),
		C.get_stride(),
		(*C.float)(unsafe.Pointer(&points[0])),
		C.get_stride(),
		C.float(thickness),
		C.int(num_segments))
	buf := make([]Point, buffer_length)
	for i := 0; i < (buffer_length * 2); i += 2 {
		buf[i/2] = Point{float32(cbuf[i]), float32(cbuf[i+1])}
	}
	return buf
}

// Draws a ribbon given an array of points. The ribbon will go through all of
// the passed points. The points buffer should consist of regularly spaced
// doublets of floats, corresponding to x and y coordinates of the vertices.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_ribbon
func DrawRibbon(p Polyline, color allegro.Color, thickness float32) {
	points := p.vertices()
	C.al_draw_ribbon(
		(*C.float)(unsafe.Pointer(&points[0])),
		C.get_stride(),
		col(color),
		C.float(thickness),
		C.int(len(p)))
}

// Draws a subset of the passed vertex array.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_prim
func DrawPrim(vertices []Vertex, decl *VertexDecl, texture *allegro.Bitmap, start, end int, prim_type PrimType) int {
	vertices_ := cVertices(vertices)
	drawn := C.al_draw_prim(unsafe.Pointer(&vertices_[0]),
		(*C.ALLEGRO_VERTEX_DECL)(decl),
		(*C.ALLEGRO_BITMAP)(texture),
		C.int(start),
		C.int(end),
		C.int(prim_type))
	return int(drawn)
}

// Draws a subset of the passed vertex array. This function uses an index array
// to specify which vertices to use.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_indexed_prim
func DrawIndexedPrim(vertices []Vertex, decl *VertexDecl, texture *allegro.Bitmap, indices []int, num_vertices int, prim_type PrimType) int {
	vertices_ := cVertices(vertices)
	indices_ := cInts(indices)
	drawn := C.al_draw_indexed_prim(unsafe.Pointer(&vertices_[0]),
		(*C.ALLEGRO_VERTEX_DECL)(decl),
		(*C.ALLEGRO_BITMAP)(texture),
		(*C.int)(unsafe.Pointer(&indices_[0])),
		C.int(num_vertices),
		C.int(prim_type))
	return int(drawn)
}

// Creates a vertex declaration, which describes a custom vertex format.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_create_vertex_decl
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
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_destroy_vertex_decl
func (v *VertexDecl) Destroy() {
	C.al_destroy_vertex_decl((*C.ALLEGRO_VERTEX_DECL)(v))
}

// Draw a series of line segments.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_polyline
func DrawPolyline(p Polyline, joinStyle LineJoin, capStyle LineCap, color allegro.Color, thickness float32, miterLimit float32) {
	vertices := p.vertices()
	C.al_draw_polyline(
		(*C.float)(unsafe.Pointer(&vertices[0])),
		C.get_stride(),
		C.int(len(p)),
		C.int(joinStyle),
		C.int(capStyle),
		col(color),
		C.float(thickness),
		C.float(miterLimit),
	)
}

// Draw an unfilled polygon. This is the same as passing
// ALLEGRO_LINE_CAP_CLOSED to al_draw_polyline.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_polygon
func DrawPolygon(p Polyline, joinStyle LineJoin, color allegro.Color, thickness float32, miterLimit float32) {
	vertices := p.vertices()
	C.al_draw_polygon(
		(*C.float)(unsafe.Pointer(&vertices[0])),
		C.int(len(p)),
		C.int(joinStyle),
		col(color),
		C.float(thickness),
		C.float(miterLimit),
	)
}

// Draw a filled, simple polygon. Simple means it does not have to be convex
// but must not be self-overlapping.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_polygon
func DrawFilledPolygon(p Polyline, color allegro.Color) {
	vertices := p.vertices()
	C.al_draw_filled_polygon(
		(*C.float)(unsafe.Pointer(&vertices[0])),
		C.int(len(p)),
		col(color),
	)
}

// Draws a filled simple polygon with zero or more other simple polygons
// subtracted from it - the holes. The holes cannot touch or intersect with the
// outline of the filled polygon.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_draw_filled_polygon_with_holes
func DrawFilledPolygonWithHoles(p Polyline, holes []Polyline, color allegro.Color) {
	vertexCounts := make([]C.int, 0, len(holes)+2)
	vertexCounts = append(vertexCounts, C.int(len(p)))
	vertices := p.vertices()
	for _, hole := range holes {
		vertices = append(vertices, hole.vertices()...)
		vertexCounts = append(vertexCounts, C.int(len(hole)))
	}
	vertexCounts = append(vertexCounts, 0)
	C.al_draw_filled_polygon_with_holes(
		(*C.float)(unsafe.Pointer(&vertices[0])),
		(*C.int)(unsafe.Pointer(&vertexCounts[0])),
		col(color),
	)
}

// This technique was borrowed from https://github.com/golang/go/wiki/cgo#function-variables

var (
	tpcMu    sync.Mutex
	tpcIndex int
	tpcFns   = make(map[int]func(x, y, z int))
)

func registerTriangulatePolygonCallback(f func(x, y, z int)) int {
	tpcMu.Lock()
	defer tpcMu.Unlock()
	tpcIndex++
	for tpcFns[tpcIndex] != nil {
		tpcIndex++
	}
	tpcFns[tpcIndex] = f
	return tpcIndex
}

//export trangulage_polygon_callback
func trangulage_polygon_callback(x, y, z C.int, data unsafe.Pointer) {
	tpcMu.Lock()
	defer tpcMu.Unlock()
	index := int(uintptr(data))
	f := tpcFns[int(index)]
	f(int(x), int(y), int(z))
}

// Divides a simple polygon into triangles, with zero or more other simple
// polygons subtracted from it - the holes. The holes cannot touch or intersect
// with the outline of the main polygon. Simple means the polygon does not have
// to be convex but must not be self-overlapping.
//
// See https://liballeg.org/a5docs/5.2.6/primitives.html#al_triangulate_polygon
func TriangulatePolygon(p Polyline, holes []Polyline, callback func(x, y, z int)) {
	vertexCounts := make([]C.int, 0, len(holes)+2)
	vertexCounts = append(vertexCounts, C.int(len(p)))
	vertices := p.vertices()
	for _, hole := range holes {
		vertices = append(vertices, hole.vertices()...)
		vertexCounts = append(vertexCounts, C.int(len(hole)))
	}
	vertexCounts = append(vertexCounts, 0)
	callbackIndex := registerTriangulatePolygonCallback(callback)
	C.al_triangulate_polygon(
		(*C.float)(unsafe.Pointer(&vertices[0])),
		C.size_t(C.get_stride()),
		(*C.int)(unsafe.Pointer(&vertexCounts[0])),
		C.emit_triangle_callback_t(C.trangulage_polygon_callback),
		unsafe.Pointer(uintptr(callbackIndex)),
	)
}
