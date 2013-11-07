package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timeout C.ALLEGRO_TIMEOUT

func NewTimeout(seconds float64) *Timeout {
	var timeout C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&timeout, C.double(seconds))
	return (*Timeout)(&timeout)
}

func Rest(seconds float64) {
	C.al_rest(C.double(seconds))
}

func Time() float64 {
	return float64(C.al_get_time())
}
