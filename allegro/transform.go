package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Transform C.ALLEGRO_TRANSFORM

func UseTransform(trans *Transform) {
	C.al_use_transform((*C.ALLEGRO_TRANSFORM)(trans))
}

// Convenience function for getting a new identity transformation.
func IdentityTransform() *Transform {
	var t Transform
	t.Identity()
	return &t
}

func BuildTransform(x, y, sx, sy, theta float32) *Transform {
	var t Transform
	C.al_build_transform((*C.ALLEGRO_TRANSFORM)(&t),
		C.float(x),
		C.float(y),
		C.float(sx),
		C.float(sy),
		C.float(theta),
	)
	return &t
}

func CurrentTransform() *Transform {
	return (*Transform)(C.al_get_current_transform())
}

func (t *Transform) Copy() *Transform {
	var dest C.ALLEGRO_TRANSFORM
	C.al_copy_transform((*C.ALLEGRO_TRANSFORM)(t), &dest)
	return (*Transform)(&dest)
}

// Turns a transformation into the identity transformation.
func (t *Transform) Identity() {
	C.al_identity_transform((*C.ALLEGRO_TRANSFORM)(t))
}

func (t *Transform) Translate(x, y float32) {
	C.al_translate_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(x), C.float(y))
}

func (t *Transform) Rotate(theta float32) {
	C.al_rotate_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(theta))
}

func (t *Transform) Scale(sx, sy float32) {
	C.al_scale_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(sx), C.float(sy))
}

// Takes a pair of coordinates and returns that same pair transformed by
// this transformation.
func (t *Transform) Coordinates(x, y float32) (float32, float32) {
	var cx, cy = C.float(x), C.float(y)
	C.al_transform_coordinates((*C.ALLEGRO_TRANSFORM)(t), &cx, &cy)
	return float32(x), float32(y)
}

func (t *Transform) Compose(other *Transform) {
	C.al_compose_transform((*C.ALLEGRO_TRANSFORM)(t), (*C.ALLEGRO_TRANSFORM)(other))
}

func (t *Transform) Invert() {
	C.al_invert_transform((*C.ALLEGRO_TRANSFORM)(t))
}

func (t *Transform) CheckInverse(tol float32) bool {
	return int(C.al_check_inverse((*C.ALLEGRO_TRANSFORM)(t), C.float(tol))) != 0
}
