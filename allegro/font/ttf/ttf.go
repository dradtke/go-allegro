package ttf

/*
#cgo pkg-config: allegro_font-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_ttf.h>
*/
import "C"

func Init() {
	C.al_init_ttf_addon()
}

func Shutdown() {
	C.al_shutdown_ttf_addon()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_ttf_version())
}
