package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

void _al_free(void *data) {
	al_free(data);
}
*/
import "C"
import (
	"unsafe"
)

func freeString(str *C.char) {
	C._al_free(unsafe.Pointer(str))
}

// Allow users to override default C memory management.
func SetMemoryInterface(memory_interface *C.ALLEGRO_MEMORY_INTERFACE) {
	C.al_set_memory_interface(memory_interface)
}
