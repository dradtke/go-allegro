package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Joystick struct {
	ptr *C.ALLEGRO_JOYSTICK
}
