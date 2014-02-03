package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
*/
import "C"
import (
)

type Mixer C.ALLEGRO_MIXER

func DefaultMixer() *Mixer {
	return (*Mixer)(C.al_get_default_mixer())
}

func (m *Mixer) IsPlaying() bool {
	return bool(C.al_get_mixer_playing((*C.ALLEGRO_MIXER)(m)))
}
