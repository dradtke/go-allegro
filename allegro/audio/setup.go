package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
*/
import "C"
import (
	"errors"
)

func Install() error {
	ok := bool(C.al_install_audio())
	if !ok {
		return errors.New("failed to install audio subsystem")
	}
	return nil
}

func Uninstall() {
	C.al_uninstall_audio()
}

func IsAudioInstalled() bool {
	return bool(C.al_is_audio_installed())
}

func ReserveSamples(reserve_samples int) error {
	ok := bool(C.al_reserve_samples(C.int(reserve_samples)))
	if !ok {
		return errors.New("failed to reserve audio samples")
	}
	return nil
}
