package audio

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
import "C"
import (
	"errors"
)

type Voice C.ALLEGRO_VOICE

// Creates a voice structure and allocates a voice from the digital sound
// driver. The passed frequency, sample format and channel configuration are
// used as a hint to what kind of data will be sent to the voice. However, the
// underlying sound driver is free to use non-matching values. For example it
// may be the native format of the sound hardware. If a mixer is attached to
// the voice, the mixer will convert from the mixer's format to the voice
// format and care does not have to be taken for this.
func CreateVoice(freq uint, depth Depth, chan_conf ChannelConf) *Voice {
	return (*Voice)(C.al_create_voice(
		C.uint(freq),
		C.ALLEGRO_AUDIO_DEPTH(depth),
		C.ALLEGRO_CHANNEL_CONF(chan_conf)))
}

// Destroys the voice and deallocates it from the digital driver. Does nothing
// if the voice is NULL.
func (v *Voice) Destroy() {
	C.al_destroy_voice((*C.ALLEGRO_VOICE)(v))
}

// Detaches the mixer or sample or stream from the voice.
func (v *Voice) Detach() {
	C.al_detach_voice((*C.ALLEGRO_VOICE)(v))
}

// Return the frequency of the voice, e.g. 44100.
func (v *Voice) Frequency() uint {
	return uint(C.al_get_voice_frequency((*C.ALLEGRO_VOICE)(v)))
}

// Return the channel configuration of the voice.
func (v *Voice) Channels() ChannelConf {
	return ChannelConf(C.al_get_voice_channels((*C.ALLEGRO_VOICE)(v)))
}

// Return the audio depth of the voice.
func (v *Voice) Depth() Depth {
	return Depth(C.al_get_voice_depth((*C.ALLEGRO_VOICE)(v)))
}

// Return true if the voice is currently playing.
func (v *Voice) IsPlaying() bool {
	return bool(C.al_get_voice_playing((*C.ALLEGRO_VOICE)(v)))
}

// Change whether a voice is playing or not. This can only work if the voice
// has a non-streaming object attached to it, e.g. a sample instance. On
// success the voice's current sample position is reset.
func (v *Voice) SetPlaying(val bool) error {
	ok := bool(C.al_set_voice_playing((*C.ALLEGRO_VOICE)(v), C.bool(val)))
	if !ok {
		return errors.New("failed to set voice playing status")
	}
	return nil
}

// When the voice has a non-streaming object attached to it, e.g. a sample,
// returns the voice's current sample position. Otherwise, returns zero.
func (v *Voice) Position() uint {
	return uint(C.al_get_voice_position((*C.ALLEGRO_VOICE)(v)))
}

// Set the voice position. This can only work if the voice has a non-streaming
// object attached to it, e.g. a sample instance.
func (v *Voice) SetPosition(val uint) error {
	ok := bool(C.al_set_voice_position((*C.ALLEGRO_VOICE)(v), C.uint(val)))
	if !ok {
		return errors.New("failed to set voice position")
	}
	return nil
}
