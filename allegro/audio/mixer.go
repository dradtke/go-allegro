package audio

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
import "C"
import (
	"errors"
	"fmt"
)

type Mixer C.ALLEGRO_MIXER

type MixerQuality C.ALLEGRO_MIXER_QUALITY

const (
	MIXER_QUALITY_POINT  MixerQuality = C.ALLEGRO_MIXER_QUALITY_POINT
	MIXER_QUALITY_LINEAR              = C.ALLEGRO_MIXER_QUALITY_LINEAR
	MIXER_QUALITY_CUBIC               = C.ALLEGRO_MIXER_QUALITY_CUBIC
)

// Creates a mixer stream, to attach sample streams or other mixers to. It will
// mix into a buffer at the requested frequency and channel count.
func CreateMixer(freq uint, depth Depth, chan_conf ChannelConf) (*Mixer, error) {
	mixer := C.al_create_mixer(C.unsigned(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf))
	if mixer == nil {
		return nil, errors.New("failed to create mixer")
	}
	return (*Mixer)(mixer), nil
}

// Return the default mixer, or NULL if one has not been set. Although
// different configurations of mixers and voices can be used, in most cases a
// single mixer attached to a voice is what you want. The default mixer is used
// by al_play_sample.
func DefaultMixer() *Mixer {
	return (*Mixer)(C.al_get_default_mixer())
}

// Sets the default mixer. All samples started with al_play_sample will be
// stopped. If you are using your own mixer, this should be called before
// al_reserve_samples.
func SetDefaultMixer(mixer *Mixer) error {
	if !bool(C.al_set_default_mixer((*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to set new default mixer")
	}
	return nil
}

// Restores Allegro's default mixer. All samples started with al_play_sample
// will be stopped. Returns true on success, false on error.
func RestoreDefaultMixer() error {
	if !bool(C.al_restore_default_mixer()) {
		return errors.New("failed to restore default mixer")
	}
	return nil
}

// Attaches a mixer onto another mixer. The same rules as with
// al_attach_sample_instance_to_mixer apply, with the added caveat that both
// mixers must be the same frequency. Returns true on success, false on error.
func (m *Mixer) AttachToMixer(mixer *Mixer) error {
	if !bool(C.al_attach_mixer_to_mixer((*C.ALLEGRO_MIXER)(m), (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach mixer to mixer")
	}
	return nil
}

// Destroys the mixer stream.
func (m *Mixer) Destroy() {
	C.al_destroy_mixer((*C.ALLEGRO_MIXER)(m))
}

// Return the mixer frequency.
func (m *Mixer) Frequency() uint {
	return uint(C.al_get_mixer_frequency((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer frequency. This will only work if the mixer is not attached to
// anything.
func (m *Mixer) SetFrequency(val uint) error {
	if !bool(C.al_set_mixer_frequency((*C.ALLEGRO_MIXER)(m), C.unsigned(val))) {
		return fmt.Errorf("failed to set mixer frequency to %d", val)
	}
	return nil
}

// Return the mixer gain (amplification factor). The default is 1.0.
func (m *Mixer) Gain() float32 {
	return float32(C.al_get_mixer_gain((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer gain (amplification factor).
func (m *Mixer) SetGain(val float32) error {
	if !bool(C.al_set_mixer_gain((*C.ALLEGRO_MIXER)(m), C.float(val))) {
		return fmt.Errorf("failed to set mixer gain to %d", val)
	}
	return nil
}

// Return the mixer channel configuration.
func (m *Mixer) Channels() ChannelConf {
	return ChannelConf(C.al_get_mixer_channels((*C.ALLEGRO_MIXER)(m)))
}

// Return the mixer audio depth.
func (m *Mixer) Depth() Depth {
	return Depth(C.al_get_mixer_depth((*C.ALLEGRO_MIXER)(m)))
}

// Return the mixer quality.
func (m *Mixer) Quality() MixerQuality {
	return MixerQuality(C.al_get_mixer_quality((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer quality. This can only succeed if the mixer does not have
// anything attached to it.
func (m *Mixer) SetQuality(quality MixerQuality) error {
	if !bool(C.al_set_mixer_quality((*C.ALLEGRO_MIXER)(m), C.ALLEGRO_MIXER_QUALITY(quality))) {
		return errors.New("failed to set new mixer quality")
	}
	return nil
}

// Return true if the mixer is playing.
func (m *Mixer) Playing() bool {
	return bool(C.al_get_mixer_playing((*C.ALLEGRO_MIXER)(m)))
}

// Change whether the mixer is playing.
func (m *Mixer) SetPlaying(val bool) error {
	if !bool(C.al_set_mixer_playing((*C.ALLEGRO_MIXER)(m), C.bool(val))) {
		return fmt.Errorf("failed to set mixer playing to %v", val)
	}
	return nil
}

// Attaches a mixer to a voice. The same rules as
// al_attach_sample_instance_to_voice apply, with the exception of the depth
// requirement.
func (m *Mixer) AttachToVoice(voice *Voice) error {
	if !bool(C.al_attach_mixer_to_voice((*C.ALLEGRO_MIXER)(m), (*C.ALLEGRO_VOICE)(voice))) {
		return errors.New("failed to attach mixer to voice")
	}
	return nil
}

// Return true if the mixer is attached to something.
func (m *Mixer) Attached() bool {
	return bool(C.al_get_mixer_attached((*C.ALLEGRO_MIXER)(m)))
}

// Detach the mixer from whatever it is attached to, if anything.
func (m *Mixer) Detach() error {
	if !bool(C.al_detach_mixer((*C.ALLEGRO_MIXER)(m))) {
		return errors.New("failed to detach mixer")
	}
	return nil
}
