package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

bool init() {
	return al_init();
}
*/
import "C"

func Init() bool {
	return gobool(C.init())
}
