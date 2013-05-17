package image

/*
#cgo pkg-config: allegro_image-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_image.h>
*/
import "C"

func Init() {
	C.al_init_image_addon()
}

func Shutdown() {
	C.al_shutdown_image_addon()
}
