// Package image provides support for Allegro's image addon.
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

// Initializes the image addon. This registers bitmap format handlers for
// al_load_bitmap, al_load_bitmap_f, al_save_bitmap, al_save_bitmap_f.
func Init() error {
	ok := bool(C.al_init_image_addon())
	if !ok {
		return errors.New("failed to initialize image addon")
	}
	return nil
}

// Shut down the image addon. This is done automatically at program exit, but
// can be called any time the user wishes as well.
func Shutdown() {
	C.al_shutdown_image_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_image_version())
}

