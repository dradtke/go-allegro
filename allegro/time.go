package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type Timeout C.ALLEGRO_TIMEOUT

func NewTimeout(seconds float64) *Timeout {
	var timeout C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&timeout, cdouble(seconds))
	return (*Timeout)(&timeout)
}

func Rest(seconds float64) {
	C.al_rest(cdouble(seconds))
}
