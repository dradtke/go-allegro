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
	"io"
	"unsafe"
)

var ErrNoAvailableFragments = errors.New("no available fragments for audio stream")

var ErrCantWriteFragment = errors.New("failed to write to audio stream fragment")

type Stream struct {
	ptr         *C.ALLEGRO_AUDIO_STREAM
	buffer_size uint
}

// Creates an ALLEGRO_AUDIO_STREAM. The stream will be set to play by default.
// It will feed audio data from a buffer, which is split into a number of
// fragments.
func CreateStream(fragment_count, frag_samples, freq uint, depth Depth, chan_conf ChannelConf) *Stream {
	sample_size := chan_conf.ChannelCount() * depth.Size()
	buffer_size := frag_samples * sample_size
	return &Stream{
		buffer_size: buffer_size,
		ptr: C.al_create_audio_stream(
			C.size_t(fragment_count),
			C.unsigned(frag_samples),
			C.unsigned(freq),
			C.ALLEGRO_AUDIO_DEPTH(depth),
			C.ALLEGRO_CHANNEL_CONF(chan_conf)),
	}
}

// Loads an audio file from disk as it is needed.
func LoadStream(filename string, buffer_count, samples uint) (*Stream, error) {
	filename_ := C.CString(filename)
	defer C.free_string(filename_)
	ptr := C.al_load_audio_stream(filename_, C.size_t(buffer_count), C.unsigned(samples))
	if ptr == nil {
		return nil, fmt.Errorf("failed to load audio stream at '%s'", filename)
	}
	return &Stream{ptr: ptr, buffer_size: 0}, nil
}

// Loads an audio file from ALLEGRO_FILE stream as it is needed.
func LoadStreamF(f *allegro.File, ident string, buffer_count, samples uint) (*Stream, error) {
    ident_ := C.CString(ident)
    defer C.free_string(ident_)
	ptr := C.al_load_audio_stream_f((*C.ALLEGRO_FILE)(f), ident_, C.size_t(buffer_count), C.unsigned(samples))
	if ptr == nil {
		return nil, errors.New("failed to load audio stream from file")
	}
	return &Stream{ptr: ptr, buffer_size: 0}, nil
}

// TODO: generalize a "Sound" interface that supports audio stream, sample instance, etc.

// Destroy an audio stream which was created with al_create_audio_stream or
// al_load_audio_stream.
func (s *Stream) Destroy() {
	C.al_destroy_audio_stream(s.ptr)
}

// Retrieve the associated event source.
func (s *Stream) EventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(C.al_get_audio_stream_event_source(s.ptr)))
}

// You should call this to finalise an audio stream that you will no longer be
// feeding, to wait for all pending buffers to finish playing. The stream's
// playing state will change to false.
func (s *Stream) Drain() {
	C.al_drain_audio_stream(s.ptr)
}

// Set the streaming file playing position to the beginning. Returns true on
// success. Currently this can only be called on streams created with
// al_load_audio_stream, al_load_audio_stream_f and the format-specific
// functions underlying those functions.
func (s *Stream) Rewind() error {
	if !bool(C.al_rewind_audio_stream(s.ptr)) {
		return errors.New("failed to rewind audio stream")
	}
	return nil
}

// Return the stream frequency.
func (s *Stream) Frequency() uint {
	return uint(C.al_get_audio_stream_frequency(s.ptr))
}

// Return the stream channel configuration.
func (s *Stream) Channels() ChannelConf {
	return ChannelConf(C.al_get_audio_stream_channels(s.ptr))
}

// Return the stream audio depth.
func (s *Stream) Depth() Depth {
	return Depth(C.al_get_audio_stream_depth(s.ptr))
}

// Return the stream length in samples.
func (s *Stream) Length() uint {
	return uint(C.al_get_audio_stream_length(s.ptr))
}

// Return the relative playback speed.
func (s *Stream) Speed() float32 {
	return float32(C.al_get_audio_stream_speed(s.ptr))
}

// Set the relative playback speed. 1.0 is normal speed.
func (s *Stream) SetSpeed(val float32) error {
	if !bool(C.al_set_audio_stream_speed(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream speed to %f", val)
	}
	return nil
}

// Return the playback gain.
func (s *Stream) Gain() float32 {
	return float32(C.al_get_audio_stream_gain(s.ptr))
}

// Set the playback gain.
func (s *Stream) SetGain(val float32) error {
	if !bool(C.al_set_audio_stream_gain(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream gain to %f", val)
	}
	return nil
}

// Get the pan value.
func (s *Stream) Pan() float32 {
	return float32(C.al_get_audio_stream_pan(s.ptr))
}

// Set the pan value on an audio stream. A value of -1.0 means to play the
// stream only through the left speaker; +1.0 means only through the right
// speaker; 0.0 means the sample is centre balanced. A special value
// ALLEGRO_AUDIO_PAN_NONE disables panning and plays the stream at its original
// level. This will be louder than a pan value of 0.0.
func (s *Stream) SetPan(val float32) error {
	if !bool(C.al_set_audio_stream_pan(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream pan to %f", val)
	}
	return nil
}

// Return true if the stream is playing.
func (s *Stream) Playing() bool {
	return bool(C.al_get_audio_stream_playing(s.ptr))
}

// Change whether the stream is playing.
func (s *Stream) SetPlaying(val bool) error {
	if !bool(C.al_set_audio_stream_playing(s.ptr, C.bool(val))) {
		return fmt.Errorf("failed to set audio stream playing to %v", val)
	}
	return nil
}

// Return the playback mode.
func (s *Stream) PlayMode() PlayMode {
	return PlayMode(C.al_get_audio_stream_playmode(s.ptr))
}

// Set the playback mode.
func (s *Stream) SetPlayMode(val PlayMode) error {
	if !bool(C.al_set_audio_stream_playmode(s.ptr, C.ALLEGRO_PLAYMODE(val))) {
		return fmt.Errorf("failed to set audio stream play mode to %v", val)
	}
	return nil
}

// Return whether the stream is attached to something.
func (s *Stream) Attached() bool {
	return bool(C.al_get_audio_stream_attached(s.ptr))
}

// Attach a stream to a mixer.
func (s *Stream) AttachToMixer(mixer *Mixer) error {
	if !bool(C.al_attach_audio_stream_to_mixer(s.ptr, (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach audio stream to mixer")
	}
	return nil
}

// Attaches an audio stream to a voice. The same rules as
// al_attach_sample_instance_to_voice apply. This may fail if the driver can't
// create a voice with the buffer count and buffer size the stream uses.
func (s *Stream) AttachToVoice(voice *Voice) error {
	if !bool(C.al_attach_audio_stream_to_voice(s.ptr, (*C.ALLEGRO_VOICE)(voice))) {
		return errors.New("failed to attach audio stream to voice")
	}
	return nil
}

// Returns the number of fragments this stream uses. This is the same value as
// passed to al_create_audio_stream when a new stream is created.
func (s *Stream) Fragments() uint {
	return uint(C.al_get_audio_stream_fragments(s.ptr))
}

// Returns the number of available fragments in the stream, that is, fragments
// which are not currently filled with data for playback.
func (s *Stream) AvailableFragments() uint {
	return uint(C.al_get_available_audio_stream_fragments(s.ptr))
}

// Detach the stream from whatever it's attached to, if anything.
func (s *Stream) Detach() error {
	if !bool(C.al_detach_audio_stream(s.ptr)) {
		return errors.New("failed to detach audio stream")
	}
	return nil
}

// This function needs to be called for every successful call of
// al_get_audio_stream_fragment to indicate that the buffer is filled with new
// data.
func (s *Stream) Write(p []byte) (n int, err error) {
	buffer := C.al_get_audio_stream_fragment(s.ptr)
	if buffer == nil {
		return 0, ErrNoAvailableFragments
	}
	defer func() {
		if !bool(C.al_set_audio_stream_fragment(s.ptr, buffer)) {
			n = 0
			err = ErrCantWriteFragment
		}
	}()
	buffer_addr := uintptr(unsafe.Pointer(buffer))
	for i := range p {
		if uint(i) >= s.buffer_size {
			return int(s.buffer_size), io.ErrShortWrite
		}
		*(*C.float)(unsafe.Pointer(buffer_addr + uintptr(i))) = C.float(p[i])
	}
	return len(p), nil
}

// Set the streaming file playing position to time. Returns true on success.
// Currently this can only be called on streams created with
// al_load_audio_stream, al_load_audio_stream_f and the format-specific
// functions underlying those functions.
func (s *Stream) SeekSecs(time float64) error {
	if !bool(C.al_seek_audio_stream_secs(s.ptr, C.double(time))) {
		return fmt.Errorf("failed to seek audio stream to %f", time)
	}
	return nil
}

// Return the position of the stream in seconds. Currently this can only be
// called on streams created with al_load_audio_stream.
func (s *Stream) PositionSecs() float64 {
	return float64(C.al_get_audio_stream_position_secs(s.ptr))
}

// Return the length of the stream in seconds, if known. Otherwise returns zero.
func (s *Stream) LengthSecs() float64 {
	return float64(C.al_get_audio_stream_length_secs(s.ptr))
}

// Sets the loop points for the stream in seconds. Currently this can only be
// called on streams created with al_load_audio_stream, al_load_audio_stream_f
// and the format-specific functions underlying those functions.
func (s *Stream) SetLoopSecs(start, end float64) error {
	if !bool(C.al_set_audio_stream_loop_secs(s.ptr, C.double(start), C.double(end))) {
		return errors.New("failed to set stream loop")
	}
	return nil
}

