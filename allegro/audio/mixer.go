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

// Creates a mixer to attach sample instances, audio streams, or other mixers
// to. It will mix into a buffer at the requested frequency (in Hz) and channel
// count.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_create_mixer
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
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_default_mixer
func DefaultMixer() *Mixer {
	return (*Mixer)(C.al_get_default_mixer())
}

// Sets the default mixer. All samples started with al_play_sample will be
// stopped and all sample instances returned by al_lock_sample_id will be
// invalidated. If you are using your own mixer, this should be called before
// al_reserve_samples.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_set_default_mixer
func SetDefaultMixer(mixer *Mixer) error {
	if !bool(C.al_set_default_mixer((*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to set new default mixer")
	}
	return nil
}

// Restores Allegro's default mixer and attaches it to the default voice. If
// the default mixer hasn't been created before, it will be created. If the
// default voice hasn't been set via al_set_default_voice or created before, it
// will also be created. All samples started with al_play_sample will be
// stopped and all sample instances returned by al_lock_sample_id will be
// invalidated.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_restore_default_mixer
func RestoreDefaultMixer() error {
	if !bool(C.al_restore_default_mixer()) {
		return errors.New("failed to restore default mixer")
	}
	return nil
}

// Attaches the mixer passed as the first argument onto the mixer passed as the
// second argument. The first mixer (that is going to be attached) must not
// already be attached to anything. Both mixers must use the same frequency,
// audio depth and channel configuration.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_attach_mixer_to_mixer
func (m *Mixer) AttachToMixer(mixer *Mixer) error {
	if !bool(C.al_attach_mixer_to_mixer((*C.ALLEGRO_MIXER)(m), (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach mixer to mixer")
	}
	return nil
}

// Destroys the mixer.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_destroy_mixer
func (m *Mixer) Destroy() {
	C.al_destroy_mixer((*C.ALLEGRO_MIXER)(m))
}

// Return the mixer frequency (in Hz).
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_frequency
func (m *Mixer) Frequency() uint {
	return uint(C.al_get_mixer_frequency((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer frequency (in Hz). This will only work if the mixer is not
// attached to anything.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_set_mixer_frequency
func (m *Mixer) SetFrequency(val uint) error {
	if !bool(C.al_set_mixer_frequency((*C.ALLEGRO_MIXER)(m), C.unsigned(val))) {
		return fmt.Errorf("failed to set mixer frequency to %d", val)
	}
	return nil
}

// Return the mixer gain (amplification factor). The default is 1.0.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_gain
func (m *Mixer) Gain() float32 {
	return float32(C.al_get_mixer_gain((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer gain (amplification factor).
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_set_mixer_gain
func (m *Mixer) SetGain(val float32) error {
	if !bool(C.al_set_mixer_gain((*C.ALLEGRO_MIXER)(m), C.float(val))) {
		return fmt.Errorf("failed to set mixer gain to %d", val)
	}
	return nil
}

// Return the mixer channel configuration.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_channels
func (m *Mixer) Channels() ChannelConf {
	return ChannelConf(C.al_get_mixer_channels((*C.ALLEGRO_MIXER)(m)))
}

// Return the mixer audio depth.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_depth
func (m *Mixer) Depth() Depth {
	return Depth(C.al_get_mixer_depth((*C.ALLEGRO_MIXER)(m)))
}

// Return the mixer quality.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_quality
func (m *Mixer) Quality() MixerQuality {
	return MixerQuality(C.al_get_mixer_quality((*C.ALLEGRO_MIXER)(m)))
}

// Set the mixer quality. This can only succeed if the mixer does not have
// anything attached to it.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_set_mixer_quality
func (m *Mixer) SetQuality(quality MixerQuality) error {
	if !bool(C.al_set_mixer_quality((*C.ALLEGRO_MIXER)(m), C.ALLEGRO_MIXER_QUALITY(quality))) {
		return errors.New("failed to set new mixer quality")
	}
	return nil
}

// Return true if the mixer is playing.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_playing
func (m *Mixer) Playing() bool {
	return bool(C.al_get_mixer_playing((*C.ALLEGRO_MIXER)(m)))
}

// Change whether the mixer is playing.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_set_mixer_playing
func (m *Mixer) SetPlaying(val bool) error {
	if !bool(C.al_set_mixer_playing((*C.ALLEGRO_MIXER)(m), C.bool(val))) {
		return fmt.Errorf("failed to set mixer playing to %v", val)
	}
	return nil
}

// Attaches a mixer to a voice. It must have the same frequency and channel
// configuration, but the depth may be different.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_attach_mixer_to_voice
func (m *Mixer) AttachToVoice(voice *Voice) error {
	if !bool(C.al_attach_mixer_to_voice((*C.ALLEGRO_MIXER)(m), (*C.ALLEGRO_VOICE)(voice))) {
		return errors.New("failed to attach mixer to voice")
	}
	return nil
}

// Return true if the mixer is attached to something.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_get_mixer_attached
func (m *Mixer) Attached() bool {
	return bool(C.al_get_mixer_attached((*C.ALLEGRO_MIXER)(m)))
}

// Detach the mixer from whatever it is attached to, if anything.
//
// See https://liballeg.org/a5docs/5.2.6/audio.html#al_detach_mixer
func (m *Mixer) Detach() error {
	if !bool(C.al_detach_mixer((*C.ALLEGRO_MIXER)(m))) {
		return errors.New("failed to detach mixer")
	}
	return nil
}
