package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

type Mixer C.ALLEGRO_MIXER

type MixerQuality C.ALLEGRO_MIXER_QUALITY

const (
	MIXER_QUALITY_POINT  MixerQuality = C.ALLEGRO_MIXER_QUALITY_POINT
	MIXER_QUALITY_LINEAR MixerQuality = C.ALLEGRO_MIXER_QUALITY_LINEAR
	MIXER_QUALITY_CUBIC  MixerQuality = C.ALLEGRO_MIXER_QUALITY_CUBIC
)

func CreateMixer(freq uint, depth Depth, chan_conf ChannelConf) (*Mixer, error) {
	mixer := C.al_create_mixer(C.unsigned(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf))
	if mixer == nil {
		return nil, errors.New("failed to create mixer")
	}
	return (*Mixer)(mixer), nil
}

func DefaultMixer() *Mixer {
	return (*Mixer)(C.al_get_default_mixer())
}

func SetDefaultMixer(mixer *Mixer) error {
	if !bool(C.al_set_default_mixer((*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to set new default mixer")
	}
	return nil
}

func RestoreDefaultMixer() error {
	if !bool(C.al_restore_default_mixer()) {
		return errors.New("failed to restore default mixer")
	}
	return nil
}

func (m *Mixer) AttachToMixer(mixer *Mixer) error {
	if !bool(C.al_attach_mixer_to_mixer((*C.ALLEGRO_MIXER)(m), (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach mixer to mixer")
	}
	return nil
}

func (m *Mixer) Destroy() {
	C.al_destroy_mixer((*C.ALLEGRO_MIXER)(m))
}

func (m *Mixer) Frequency() uint {
	return uint(C.al_get_mixer_frequency((*C.ALLEGRO_MIXER)(m)))
}

func (m *Mixer) SetFrequency(val uint) error {
	if !bool(C.al_set_mixer_frequency((*C.ALLEGRO_MIXER)(m), C.unsigned(val))) {
		return fmt.Errorf("failed to set mixer frequency to %d", val)
	}
	return nil
}

func (m *Mixer) Gain() float32 {
	return float32(C.al_get_mixer_gain((*C.ALLEGRO_MIXER)(m)))
}

func (m *Mixer) SetGain(val float32) error {
	if !bool(C.al_set_mixer_gain((*C.ALLEGRO_MIXER)(m), C.float(val))) {
		return fmt.Errorf("failed to set mixer gain to %d", val)
	}
	return nil
}

func (m *Mixer) Channels() ChannelConf {
	return ChannelConf(C.al_get_mixer_channels((*C.ALLEGRO_MIXER)(m)))
}

func (m *Mixer) Depth() Depth {
	return Depth(C.al_get_mixer_depth((*C.ALLEGRO_MIXER)(m)))
}

func (m *Mixer) Quality() MixerQuality {
    return MixerQuality(C.al_get_mixer_quality((*C.ALLEGRO_MIXER)(m)))
}

func (m *Mixer) Playing() bool {
	return bool(C.al_get_mixer_playing((*C.ALLEGRO_MIXER)(m)))
}
