package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

void free_string(char *str) {
	al_free(str);
}
*/
import "C"

func FreeString(str *C.char) {
	C.free_string(str)
}
