// Package x11 provides support for Allegro's X11 functions.
package x11

// #include <allegro5/allegro_x.h>
import "C"
import (
	"github.com/dradtke/go-allegro/allegro"
)

// Retrieves the XID associated with the Allegro display.
//
// See https://liballeg.org/a5docs/5.2.6/platform.html#al_get_x_window_id
func WindowID(d *allegro.Display) uint64 {
	return uint64(C.al_get_x_window_id((*C.ALLEGRO_DISPLAY)(d)))
}
