package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

func cint(x int) C.int {
	return (C.int)(x)
}

func cbyte(x byte) C.uchar {
	return (C.uchar)(x)
}

func cfloat(x float32) C.float {
	return (C.float)(x)
}

func cdouble(x float64) C.double {
	return (C.double)(x)
}

func gobool(x C.bool) bool {
	return (bool)(x)
}

func godouble(x C.double) float64 {
	return (float64)(x)
}
