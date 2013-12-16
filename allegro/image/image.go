package image

/*
#cgo pkg-config: allegro_image-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_image.h>
*/
import "C"
import (
	"errors"
)

func Init() error {
	ok := bool(C.al_init_image_addon())
	if !ok {
		return errors.New("failed to initialize image addon")
	}
	return nil
}

func Shutdown() {
	C.al_shutdown_image_addon()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_image_version())
}
