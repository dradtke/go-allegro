package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
#include "../util.c"
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
)

type Sample C.ALLEGRO_SAMPLE
type SampleInstance C.ALLEGRO_SAMPLE_INSTANCE

// Create a sample data structure from the supplied buffer. If free_buf is true
// then the buffer will be freed with al_free when the sample data structure is
// destroyed. For portability (especially Windows), the buffer should have been
// allocated with al_malloc. Otherwise you should free the sample data yourself.
func CreateSample(samples, freq uint, depth Depth, chan_conf ChannelConf) *Sample {
	buf := C._al_malloc(C.uint(samples * chan_conf.ChannelCount() * depth.Size()))
	return (*Sample)(C.al_create_sample(
		buf,
		C.uint(samples),
		C.uint(freq),
		C.ALLEGRO_AUDIO_DEPTH(depth),
		C.ALLEGRO_CHANNEL_CONF(chan_conf),
		C.bool(true)))
}

// Loads a few different audio file formats based on their extension.
func LoadSample(filename string) (*Sample, error) {
	filename_ := C.CString(filename)
	defer C.free_string(filename_)
	s := C.al_load_sample(filename_)
	if s == nil {
		return nil, fmt.Errorf("failed to load sample '%s'", filename)
	}
	return (*Sample)(s), nil
}

// Loads an audio file from an ALLEGRO_FILE stream into an ALLEGRO_SAMPLE. The
// file type is determined by the passed 'ident' parameter, which is a file
// name extension including the leading dot.
func LoadSampleF(f *allegro.File, ident string) (*Sample, error) {
	ident_ := C.CString(ident)
	defer C.free_string(ident_)
	if sample := C.al_load_sample_f((*C.ALLEGRO_FILE)(f), ident_); sample != nil {
		return (*Sample)(sample), nil
	}
	return nil, errors.New("failed to load sample from file")
}

// Writes a sample into a file. Currently, wav is the only supported format,
// and the extension must be ".wav".
func (s *Sample) Save(filename string) error {
	filename_ := C.CString(filename)
	defer C.free_string(filename_)
	if !bool(C.al_save_sample(filename_, (*C.ALLEGRO_SAMPLE)(s))) {
		return fmt.Errorf("failed to save sample to '%s'", filename)
	}
	return nil
}

// Writes a sample into a ALLEGRO_FILE filestream. Currently, wav is the only
// supported format, and the extension must be ".wav".
func (s *Sample) SaveF(f *allegro.File, ident string) error {
	ident_ := C.CString(ident)
	defer C.free_string(ident_)
	if !bool(C.al_save_sample_f((*C.ALLEGRO_FILE)(f), ident_, (*C.ALLEGRO_SAMPLE)(s))) {
		return errors.New("failed to save sample to file")
	}
	return nil
}

// Creates a sample stream, using the supplied data. This must be attached to a
// voice or mixer before it can be played. The argument may be NULL. You can
// then set the data later with al_set_sample.
func CreateSampleInstance(sample_data *Sample) *SampleInstance {
	return (*SampleInstance)(C.al_create_sample_instance((*C.ALLEGRO_SAMPLE)(sample_data)))
}

// Play an instance of a sample data. Returns true on success, false on failure.
func (s *SampleInstance) Play() error {
	if !bool(C.al_play_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))) {
		return errors.New("failed to play sample instance")
	}
	return nil
}

// Stop an sample instance playing.
func (s *SampleInstance) Stop() error {
	if !bool(C.al_stop_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))) {
		return errors.New("failed to stop sample instance")
	}
	return nil
}

// Attach a sample instance to a mixer. The instance must not already be
// attached to anything.
func (s *SampleInstance) AttachToMixer(mixer *Mixer) error {
	if mixer == nil {
		return errors.New("cannot attach sample instance to null mixer")
	}
	if !bool(C.al_attach_sample_instance_to_mixer((*C.ALLEGRO_SAMPLE_INSTANCE)(s), (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach sample instance to mixer")
	}
	return nil
}

// Attaches a sample to a voice, and allows it to play. The sample's volume and
// loop mode will be ignored, and it must have the same frequency and depth
// (including signed-ness) as the voice. This function may fail if the selected
// driver doesn't support preloading sample data.
func (s *SampleInstance) AttachToVoice(voice *Voice) error {
	if voice == nil {
		return errors.New("cannot attach sample instance to null voice")
	}
	if !bool(C.al_attach_sample_instance_to_voice((*C.ALLEGRO_SAMPLE_INSTANCE)(s), (*C.ALLEGRO_VOICE)(voice))) {
		return errors.New("failed to attach sample instance to voice")
	}
	return nil
}

// Return the frequency of the sample instance.
func (s *SampleInstance) Frequency() uint {
	return uint(C.al_get_sample_instance_frequency((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Return the length of the sample instance in sample values.
func (s *SampleInstance) Length() uint {
	return uint(C.al_get_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the length of the sample instance in sample values.
func (s *SampleInstance) SetLength(val uint) error {
	if !bool(C.al_set_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.unsigned(val))) {
		return fmt.Errorf("failed to set sample instance length to %d", val)
	}
	return nil
}

// Get the playback position of a sample instance.
func (s *SampleInstance) Position() uint {
	return uint(C.al_get_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the playback position of a sample instance.
func (s *SampleInstance) SetPosition(val uint) error {
	if !bool(C.al_set_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.unsigned(val))) {
		return fmt.Errorf("failed to set sample instance position to %d", val)
	}
	return nil
}

// Return the relative playback speed.
func (s *SampleInstance) Speed() float32 {
	return float32(C.al_get_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the relative playback speed. 1.0 is normal speed.
func (s *SampleInstance) SetSpeed(val float32) error {
	if !bool(C.al_set_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
		return fmt.Errorf("failed to set sample instance speed to %d", val)
	}
	return nil
}

// Return the playback gain.
func (s *SampleInstance) Gain() float32 {
	return float32(C.al_get_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the playback gain.
func (s *SampleInstance) SetGain(val float32) error {
	if !bool(C.al_set_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
		return fmt.Errorf("failed to set sample instance gain to %d", val)
	}
	return nil
}

// Get the pan value.
func (s *SampleInstance) Pan() float32 {
	return float32(C.al_get_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the pan value on a sample instance. A value of -1.0 means to play the
// sample only through the left speaker; +1.0 means only through the right
// speaker; 0.0 means the sample is centre balanced. A special value
// ALLEGRO_AUDIO_PAN_NONE disables panning and plays the sample at its original
// level. This will be louder than a pan value of 0.0.
func (s *SampleInstance) SetPan(val float32) error {
	if !bool(C.al_set_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
		return fmt.Errorf("failed to set sample instance pan to %d", val)
	}
	return nil
}

// Return the audio depth.
func (s *SampleInstance) Depth() Depth {
	return Depth(C.al_get_sample_instance_depth((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Return the channel configuration.
func (s *SampleInstance) Channels() ChannelConf {
	return ChannelConf(C.al_get_sample_instance_channels((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Return the playback mode.
func (s *SampleInstance) PlayMode() PlayMode {
	return PlayMode(C.al_get_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Set the playback mode.
func (s *SampleInstance) SetPlayMode(val PlayMode) error {
	if !bool(C.al_set_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.ALLEGRO_PLAYMODE(val))) {
		return errors.New("failed to set sample instance playmode")
	}
	return nil
}

// Return true if the sample instance is playing.
func (s *SampleInstance) Playing() bool {
	return bool(C.al_get_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Change whether the sample instance is playing.
func (s *SampleInstance) SetPlaying(val bool) error {
	if !bool(C.al_set_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.bool(val))) {
		return fmt.Errorf("failed to set sample instance playing to %v", val)
	}
	return nil
}

// Return whether the sample instance is attached to something.
func (s *SampleInstance) Attached() bool {
	return bool(C.al_get_sample_instance_attached((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Return the length of the sample instance in seconds, assuming a playback
// speed of 1.0.
func (s *SampleInstance) Time() float32 {
	return float32(C.al_get_sample_instance_time((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

// Detach the sample instance from whatever it's attached to, if anything.
func (s *SampleInstance) Detach() error {
	if !bool(C.al_detach_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))) {
		return errors.New("failed to detach sample instance")
	}
	return nil
}

// Detaches the sample stream from anything it may be attached to and frees it
// (the sample data is not freed!).
func (s *SampleInstance) Destroy() {
	C.al_destroy_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))
}

// Free the sample data structure. If it was created with the free_buf
// parameter set to true, then the buffer will be freed with al_free.
func (s *Sample) Destroy() {
	C.al_destroy_sample((*C.ALLEGRO_SAMPLE)(s))
}

// Plays a sample on one of the sample instances created by al_reserve_samples.
// Returns true on success, false on failure. Playback may fail because all the
// reserved sample instances are currently used.
func (s *Sample) Play(gain, pan, speed float32, loop PlayMode) (*SampleID, error) {
	var id SampleID
	ok := bool(C.al_play_sample(
		(*C.ALLEGRO_SAMPLE)(s),
		C.float(gain),
		C.float(pan),
		C.float(speed),
		C.ALLEGRO_PLAYMODE(loop),
		(*C.ALLEGRO_SAMPLE_ID)(&id)))
	if !ok {
		return nil, errors.New("failed to play sample")
	}
	return &id, nil
}

// Return the channel configuration.
func (s *Sample) Channels() ChannelConf {
	return ChannelConf(C.al_get_sample_channels((*C.ALLEGRO_SAMPLE)(s)))
}

// Return the audio depth.
func (s *Sample) Depth() Depth {
	return Depth(C.al_get_sample_depth((*C.ALLEGRO_SAMPLE)(s)))
}

// Return the frequency of the sample.
func (s *Sample) Frequency() uint {
	return uint(C.al_get_sample_frequency((*C.ALLEGRO_SAMPLE)(s)))
}

// Return the length of the sample in sample values.
func (s *Sample) Length() uint {
	return uint(C.al_get_sample_length((*C.ALLEGRO_SAMPLE)(s)))
}

// Return a pointer to the raw sample data.
func (s *Sample) Data() uintptr {
	return uintptr(C.al_get_sample_data((*C.ALLEGRO_SAMPLE)(s)))
}

// Stop the sample started by al_play_sample.
func (s *SampleID) Stop() {
	C.al_stop_sample((*C.ALLEGRO_SAMPLE_ID)(s))
}

// Stop all samples started by al_play_sample.
func StopSamples() {
	C.al_stop_samples()
}
