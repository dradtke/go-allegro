// Package acodec provides support for Allegro's acodec addon.
package acodec

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_acodec.h>
import "C"
import (
	"errors"
)

// TODO: get Allegro to recognize the .oga extension.

// This function registers all the known audio file type handlers for
// al_load_sample, al_save_sample, al_load_audio_stream, etc.
func Install() error {
	ok := bool(C.al_init_acodec_addon())
	if !ok {
		return errors.New("failed to initialize acodec addon")
	}
	return nil
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_acodec_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}
