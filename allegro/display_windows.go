package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

const (
	DIRECT3D DisplayFlags = C.ALLEGRO_DIRECT3D_INTERNAL
)
