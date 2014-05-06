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

// sample_size = al_get_channel_count(chan_conf) * al_get_audio_depth_size(depth)
// buffer_size = frag_samples * sample_size
func CreateAudioStream(fragment_count, frag_samples, freq uint, depth Depth, chan_conf ChannelConf) *Stream {
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

func LoadAudioStream(filename string, buffer_count, samples uint) (*Stream, error) {
	filename_ := C.CString(filename)
	defer C.free_string(filename_)
	ptr := C.al_load_audio_stream(filename_, C.size_t(buffer_count), C.unsigned(samples))
	if ptr == nil {
		return nil, fmt.Errorf("failed to load audio stream at '%s'", filename)
	}
	return &Stream{ptr: ptr, buffer_size: 0}, nil
}

func LoadAudioStreamF(f *allegro.File, ident string, buffer_count, samples uint) (*Stream, error) {
    ident_ := C.CString(ident)
    defer C.free_string(ident_)
	ptr := C.al_load_audio_stream_f((*C.ALLEGRO_FILE)(f), ident_, C.size_t(buffer_count), C.unsigned(samples))
	if ptr == nil {
		return nil, errors.New("failed to load audio stream from file")
	}
	return &Stream{ptr: ptr, buffer_size: 0}, nil
}

// TODO: generalize a "Sound" interface that supports audio stream, sample instance, etc.

func (s *Stream) Destroy() {
	C.al_destroy_audio_stream(s.ptr)
}

func (s *Stream) EventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(C.al_get_audio_stream_event_source(s.ptr)))
}

func (s *Stream) Drain() {
	C.al_drain_audio_stream(s.ptr)
}

func (s *Stream) Rewind() error {
	if !bool(C.al_rewind_audio_stream(s.ptr)) {
		return errors.New("failed to rewind audio stream")
	}
	return nil
}

func (s *Stream) Frequency() uint {
	return uint(C.al_get_audio_stream_frequency(s.ptr))
}

func (s *Stream) Channels() ChannelConf {
	return ChannelConf(C.al_get_audio_stream_channels(s.ptr))
}

func (s *Stream) Depth() Depth {
	return Depth(C.al_get_audio_stream_depth(s.ptr))
}

func (s *Stream) Length() uint {
	return uint(C.al_get_audio_stream_length(s.ptr))
}

func (s *Stream) Speed() float32 {
	return float32(C.al_get_audio_stream_speed(s.ptr))
}

func (s *Stream) SetSpeed(val float32) error {
	if !bool(C.al_set_audio_stream_speed(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream speed to %f", val)
	}
	return nil
}

func (s *Stream) Gain() float32 {
	return float32(C.al_get_audio_stream_gain(s.ptr))
}

func (s *Stream) SetGain(val float32) error {
	if !bool(C.al_set_audio_stream_gain(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream gain to %f", val)
	}
	return nil
}

func (s *Stream) Pan() float32 {
	return float32(C.al_get_audio_stream_pan(s.ptr))
}

func (s *Stream) SetPan(val float32) error {
	if !bool(C.al_set_audio_stream_pan(s.ptr, C.float(val))) {
		return fmt.Errorf("failed to set audio stream pan to %f", val)
	}
	return nil
}

func (s *Stream) Playing() bool {
	return bool(C.al_get_audio_stream_playing(s.ptr))
}

func (s *Stream) SetPlaying(val bool) error {
	if !bool(C.al_set_audio_stream_playing(s.ptr, C.bool(val))) {
		return fmt.Errorf("failed to set audio stream playing to %v", val)
	}
	return nil
}

func (s *Stream) PlayMode() PlayMode {
	return PlayMode(C.al_get_audio_stream_playmode(s.ptr))
}

func (s *Stream) SetPlayMode(val PlayMode) error {
	if !bool(C.al_set_audio_stream_playmode(s.ptr, C.ALLEGRO_PLAYMODE(val))) {
		return fmt.Errorf("failed to set audio stream play mode to %v", val)
	}
	return nil
}

func (s *Stream) Attached() bool {
	return bool(C.al_get_audio_stream_attached(s.ptr))
}

func (s *Stream) AttachToMixer(mixer *Mixer) error {
	if !bool(C.al_attach_audio_stream_to_mixer(s.ptr, (*C.ALLEGRO_MIXER)(mixer))) {
		return errors.New("failed to attach audio stream to mixer")
	}
	return nil
}

func (s *Stream) AttachToVoice(voice *Voice) error {
	if !bool(C.al_attach_audio_stream_to_voice(s.ptr, (*C.ALLEGRO_VOICE)(voice))) {
		return errors.New("failed to attach audio stream to voice")
	}
	return nil
}

func (s *Stream) Fragments() uint {
	return uint(C.al_get_audio_stream_fragments(s.ptr))
}

func (s *Stream) AvailableFragments() uint {
	return uint(C.al_get_available_audio_stream_fragments(s.ptr))
}

func (s *Stream) Detach() error {
	if !bool(C.al_detach_audio_stream(s.ptr)) {
		return errors.New("failed to detach audio stream")
	}
	return nil
}

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
		*(*byte)(unsafe.Pointer(buffer_addr + uintptr(i))) = p[i]
	}
	return len(p), nil
}

func (s *Stream) SeekSecs(time float64) error {
	if !bool(C.al_seek_audio_stream_secs(s.ptr, C.double(time))) {
		return fmt.Errorf("failed to seek audio stream to %f", time)
	}
	return nil
}

func (s *Stream) PositionSecs() float64 {
	return float64(C.al_get_audio_stream_position_secs(s.ptr))
}

func (s *Stream) LengthSecs() float64 {
	return float64(C.al_get_audio_stream_length_secs(s.ptr))
}

func (s *Stream) SetLoopSecs(start, end float64) error {
	if !bool(C.al_set_audio_stream_loop_secs(s.ptr, C.double(start), C.double(end))) {
		return errors.New("failed to set stream loop")
	}
	return nil
}
