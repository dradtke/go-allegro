// Package audio provides support for Allegro's audio addon.
package audio

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
import "C"
import (
	"errors"
)

// Install the audio subsystem.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_install_audio
func Install() error {
	ok := bool(C.al_install_audio())
	if !ok {
		return errors.New("failed to install audio subsystem")
	}
	return nil
}

// Uninstalls the audio subsystem.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_uninstall_audio
func Uninstall() {
	C.al_uninstall_audio()
}

// Returns true if al_install_audio was called previously and returned
// successfully.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_is_audio_installed
func IsAudioInstalled() bool {
	return bool(C.al_is_audio_installed())
}

// Reserves a number of sample instances, attaching them to the default mixer.
// If no default mixer is set when this function is called, then it will create
// one and attach it to the default voice. If no default voice has been set,
// it, too, will be created.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_reserve_samples
func ReserveSamples(reserve_samples int) error {
	ok := bool(C.al_reserve_samples(C.int(reserve_samples)))
	if !ok {
		return errors.New("failed to reserve audio samples")
	}
	return nil
}
