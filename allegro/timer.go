package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

// TODO: maintain an internal list of timers, like displays
type Timer struct {
	ptr *C.ALLEGRO_TIMER
}
