package allegro

// #include <allegro5/allegro.h>
import "C"

type Transform C.ALLEGRO_TRANSFORM

// Sets the transformation to be used for the the drawing operations on the
// target bitmap (each bitmap maintains its own transformation). Every drawing
// operation after this call will be transformed using this transformation.
// Call this function with an identity transformation to return to the default
// behaviour.
func UseTransform(trans *Transform) {
	C.al_use_transform((*C.ALLEGRO_TRANSFORM)(trans))
}

// Convenience function for getting a new identity transformation.
func IdentityTransform() *Transform {
	var t Transform
	t.Identity()
	return &t
}

// Builds a transformation given some parameters. This call is equivalent to
// calling the transformations in this order: make identity, scale, rotate,
// translate. This method is faster, however, than actually calling those
// functions.
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

// Returns the transformation of the current target bitmap, as set by
// al_use_transform. If there is no target bitmap, this function returns NULL.
func CurrentTransform() *Transform {
	return (*Transform)(C.al_get_current_transform())
}

// Makes a copy of a transformation.
func (t *Transform) Copy() *Transform {
	var dest C.ALLEGRO_TRANSFORM
	C.al_copy_transform((*C.ALLEGRO_TRANSFORM)(t), &dest)
	return (*Transform)(&dest)
}

// Sets the transformation to be the identity transformation. This is the
// default transformation. Use al_use_transform on an identity transformation
// to return to the default.
func (t *Transform) Identity() {
	C.al_identity_transform((*C.ALLEGRO_TRANSFORM)(t))
}

// Apply a translation to a transformation.
func (t *Transform) Translate(x, y float32) {
	C.al_translate_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(x), C.float(y))
}

// Apply a rotation to a transformation.
func (t *Transform) Rotate(theta float32) {
	C.al_rotate_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(theta))
}

// Apply a scale to a transformation.
func (t *Transform) Scale(sx, sy float32) {
	C.al_scale_transform((*C.ALLEGRO_TRANSFORM)(t), C.float(sx), C.float(sy))
}

// Transform a pair of coordinates.
func (t *Transform) Coordinates(x, y float32) (float32, float32) {
	var cx, cy = C.float(x), C.float(y)
	C.al_transform_coordinates((*C.ALLEGRO_TRANSFORM)(t), &cx, &cy)
	return float32(x), float32(y)
}

// Compose (combine) two transformations by a matrix multiplication.
func (t *Transform) Compose(other *Transform) {
	C.al_compose_transform((*C.ALLEGRO_TRANSFORM)(t), (*C.ALLEGRO_TRANSFORM)(other))
}

// Inverts the passed transformation. If the transformation is nearly singular
// (close to not having an inverse) then the returned transformation may be
// invalid. Use al_check_inverse to ascertain if the transformation has an
// inverse before inverting it if you are in doubt.
func (t *Transform) Invert() {
	C.al_invert_transform((*C.ALLEGRO_TRANSFORM)(t))
}

// Checks if the transformation has an inverse using the supplied tolerance.
// Tolerance should be a small value between 0 and 1, with 1e-7 being
// sufficient for most applications.
func (t *Transform) CheckInverse(tol float32) bool {
	return int(C.al_check_inverse((*C.ALLEGRO_TRANSFORM)(t), C.float(tol))) != 0
}
