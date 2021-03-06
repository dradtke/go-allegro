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

func frees(xs ...unsafe.Pointer) {
	for _, x := range xs {
		free(x)
	}
}

func freeString(data *C.char) {
	C._al_free(unsafe.Pointer(data))
}

func freeStrings(xs ...*C.char) {
	for _, x := range xs {
		freeString(x)
	}
}

// Override the memory management functions with implementations of
// al_malloc_with_context, al_free_with_context, al_realloc_with_context and
// al_calloc_with_context. The context arguments may be used for debugging. The
// new functions should be thread safe.
//
// See https://liballeg.org/a5docs/5.2.6/memory.html#al_set_memory_interface
func SetMemoryInterface(memory_interface *C.ALLEGRO_MEMORY_INTERFACE) {
	C.al_set_memory_interface(memory_interface)
}
