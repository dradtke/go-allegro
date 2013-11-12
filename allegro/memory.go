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

func free(data unsafe.Pointer) {
	C._al_free(data)
}

func freeString(data *C.char) {
	C._al_free(unsafe.Pointer(data))
}

// Allow users to override default C memory management, if they really want to.
func SetMemoryInterface(memory_interface *C.ALLEGRO_MEMORY_INTERFACE) {
	C.al_set_memory_interface(memory_interface)
}
