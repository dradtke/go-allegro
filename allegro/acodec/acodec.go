package acodec

/*
#cgo pkg-config: allegro_acodec-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_acodec.h>
*/
import "C"

func Init() error {
	ok := bool(C.al_init_acodec_addon())
	if !ok {
		return errors.New("failed to initialize acodec addon")
	}
	return nil
}

func Shutdown() {
	C.al_shutdown_acodec_addon()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_acodec_version())
}
