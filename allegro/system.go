package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

bool init() {
	return al_init();
}
*/
import "C"
import (
	"errors"
)

func Init() error {
	if !bool(C.init()) {
		return errors.New("failed to initialize allegro!")
	}
	return nil
}

func Version() uint32 {
	return uint32(C.al_get_allegro_version())
}
