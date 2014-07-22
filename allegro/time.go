package allegro

// #include <allegro5/allegro.h>
import "C"

type Timeout C.ALLEGRO_TIMEOUT

// Set timeout value of some number of seconds after the function call.
func NewTimeout(seconds float64) *Timeout {
	var timeout C.ALLEGRO_TIMEOUT
	C.al_init_timeout(&timeout, C.double(seconds))
	return (*Timeout)(&timeout)
}

// Waits for the specified number seconds. This tells the system to pause the
// current thread for the given amount of time. With some operating systems,
// the accuracy can be in the order of 10ms. That is, even
func Rest(seconds float64) {
	C.al_rest(C.double(seconds))
}

// Return the number of seconds since the Allegro library was initialised. The
// return value is undefined if Allegro is uninitialised. The resolution
// depends on the used driver, but typically can be in the order of
// microseconds.
func Time() float64 {
	return float64(C.al_get_time())
}
