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

type Voice C.ALLEGRO_VOICE

func CreateVoice(freq uint, depth Depth, chan_conf ChannelConf) *Voice {
	return (*Voice)(C.al_create_voice(
		C.uint(freq),
		C.ALLEGRO_AUDIO_DEPTH(depth),
		C.ALLEGRO_CHANNEL_CONF(chan_conf)))
}

func (v *Voice) Destroy() {
	C.al_destroy_voice((*C.ALLEGRO_VOICE)(v))
}

func (v *Voice) Detach() {
	C.al_detach_voice((*C.ALLEGRO_VOICE)(v))
}

func (v *Voice) Frequency() uint {
	return uint(C.al_get_voice_frequency((*C.ALLEGRO_VOICE)(v)))
}

func (v *Voice) Channels() ChannelConf {
	return ChannelConf(C.al_get_voice_channels((*C.ALLEGRO_VOICE)(v)))
}

func (v *Voice) Depth() Depth {
	return Depth(C.al_get_voice_depth((*C.ALLEGRO_VOICE)(v)))
}

func (v *Voice) IsPlaying() bool {
	return bool(C.al_get_voice_playing((*C.ALLEGRO_VOICE)(v)))
}

func (v *Voice) SetPlaying(val bool) error {
	ok := bool(C.al_set_voice_playing((*C.ALLEGRO_VOICE)(v), C.bool(val)))
	if !ok {
		return errors.New("failed to set voice playing status")
	}
	return nil
}

func (v *Voice) Position() uint {
	return uint(C.al_get_voice_position((*C.ALLEGRO_VOICE)(v)))
}

func (v *Voice) SetPosition(val uint) error {
	ok := bool(C.al_set_voice_position((*C.ALLEGRO_VOICE)(v), C.uint(val)))
	if !ok {
		return errors.New("failed to set voice position")
	}
	return nil
}
