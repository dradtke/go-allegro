package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

bool init() {
	return al_init();
}
*/
import "C"
import "errors"

func Init() error {
	if !gobool(C.init()) {
		return errors.New("failed to initialize allegro!")
	}
	return nil
}
