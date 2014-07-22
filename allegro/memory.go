package allegro

// #include <allegro5/allegro.h>
/*
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

// Override the memory management functions with implementations of
// al_malloc_with_context, al_free_with_context, al_realloc_with_context and
// al_calloc_with_context. The context arguments may be used for debugging.
func SetMemoryInterface(memory_interface *C.ALLEGRO_MEMORY_INTERFACE) {
	C.al_set_memory_interface(memory_interface)
}
