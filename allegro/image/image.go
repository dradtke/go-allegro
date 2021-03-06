// Package image provides support for Allegro's image addon.
package image

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_image.h>
import "C"
import (
	"errors"
)

// Initializes the image addon. This registers bitmap format handlers for
// al_load_bitmap, al_load_bitmap_f, al_save_bitmap, al_save_bitmap_f.
//
// See https://liballeg.org/a5docs/5.2.6/image.html#al_init_image_addon
func Install() error {
	ok := bool(C.al_init_image_addon())
	if !ok {
		return errors.New("failed to initialize image addon")
	}
	return nil
}

// Returns true if the image addon is initialized, otherwise returns false.
//
// See https://liballeg.org/a5docs/5.2.6/image.html#al_is_image_addon_initialized
func Installed() bool {
	return bool(C.al_is_image_addon_initialized())
}

// Shut down the image addon. This is done automatically at program exit, but
// can be called any time the user wishes as well.
//
// See https://liballeg.org/a5docs/5.2.6/image.html#al_shutdown_image_addon
func Uninstall() {
	C.al_shutdown_image_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
//
// See https://liballeg.org/a5docs/5.2.6/image.html#al_get_allegro_image_version
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_image_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}
