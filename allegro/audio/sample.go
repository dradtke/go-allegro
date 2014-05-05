package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>

void free_string(char *str) {
	al_free(str);
}

void *_al_malloc(unsigned int size) {
    return al_malloc(size);
}
*/
import "C"
import (
	"errors"
	"fmt"
)

type Sample C.ALLEGRO_SAMPLE
type SampleInstance C.ALLEGRO_SAMPLE_INSTANCE

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

func CreateSampleInstance(sample_data *Sample) *SampleInstance {
    return (*SampleInstance)(C.al_create_sample_instance((*C.ALLEGRO_SAMPLE)(sample_data)))
}

func (s *SampleInstance) Frequency() uint {
    return uint(C.al_get_sample_instance_frequency((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) Length() uint {
    return uint(C.al_get_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetLength(val uint) error {
    if !bool(C.al_set_sample_instance_length((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.unsigned(val))) {
        return fmt.Errorf("failed to set sample instance length to %d", val)
    }
    return nil
}

func (s *SampleInstance) Position() uint {
    return uint(C.al_get_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetPosition(val uint) error {
    if !bool(C.al_set_sample_instance_position((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.unsigned(val))) {
        return fmt.Errorf("failed to set sample instance position to %d", val)
    }
    return nil
}

func (s *SampleInstance) Speed() float32 {
    return float32(C.al_get_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetSpeed(val float32) error {
    if !bool(C.al_set_sample_instance_speed((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
        return fmt.Errorf("failed to set sample instance speed to %d", val)
    }
    return nil
}

func (s *SampleInstance) Gain() float32 {
    return float32(C.al_get_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetGain(val float32) error {
    if !bool(C.al_set_sample_instance_gain((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
        return fmt.Errorf("failed to set sample instance gain to %d", val)
    }
    return nil
}

func (s *SampleInstance) Pan() float32 {
    return float32(C.al_get_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetPan(val float32) error {
    if !bool(C.al_set_sample_instance_pan((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.float(val))) {
        return fmt.Errorf("failed to set sample instance pan to %d", val)
    }
    return nil
}

func (s *SampleInstance) Depth() Depth {
    return Depth(C.al_get_sample_instance_depth((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) Channels() ChannelConf {
    return ChannelConf(C.al_get_sample_instance_channels((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) PlayMode() PlayMode {
    return PlayMode(C.al_get_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetPlayMode(val PlayMode) error {
    if !bool(C.al_set_sample_instance_playmode((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.ALLEGRO_PLAYMODE(val))) {
        return errors.New("failed to set sample instance playmode")
    }
    return nil
}

func (s *SampleInstance) Playing() bool {
    return bool(C.al_get_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) SetPlaying(val bool) error {
    if !bool(C.al_set_sample_instance_playing((*C.ALLEGRO_SAMPLE_INSTANCE)(s), C.bool(val))) {
        return fmt.Errorf("failed to set sample instance playing to %v", val)
    }
    return nil
}

func (s *SampleInstance) Attached() bool {
    return bool(C.al_get_sample_instance_attached((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) Time() float32 {
    return float32(C.al_get_sample_instance_time((*C.ALLEGRO_SAMPLE_INSTANCE)(s)))
}

func (s *SampleInstance) Detach() error {
    if !bool(C.al_detach_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))) {
        return errors.New("failed to detach sample instance")
    }
    return nil
}

func (s *SampleInstance) Destroy() {
    C.al_destroy_sample_instance((*C.ALLEGRO_SAMPLE_INSTANCE)(s))
}

func LoadSample(filename string) (*Sample, error) {
	filename_ := C.CString(filename)
	defer C.free_string(filename_)
	s := C.al_load_sample(filename_)
	if s == nil {
		return nil, fmt.Errorf("failed to load sample '%s'", filename)
	}
	return (*Sample)(s), nil
}

func (s *Sample) Destroy() {
	C.al_destroy_sample((*C.ALLEGRO_SAMPLE)(s))
}

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

func (s *Sample) Channels() ChannelConf {
	return ChannelConf(C.al_get_sample_channels((*C.ALLEGRO_SAMPLE)(s)))
}

func (s *Sample) Depth() Depth {
	return Depth(C.al_get_sample_depth((*C.ALLEGRO_SAMPLE)(s)))
}

func (s *Sample) Frequency() uint {
	return uint(C.al_get_sample_frequency((*C.ALLEGRO_SAMPLE)(s)))
}

func (s *Sample) Length() uint {
	return uint(C.al_get_sample_length((*C.ALLEGRO_SAMPLE)(s)))
}

func (s *Sample) Data() uintptr {
	return uintptr(C.al_get_sample_data((*C.ALLEGRO_SAMPLE)(s)))
}

func (s *SampleID) Stop() {
	C.al_stop_sample((*C.ALLEGRO_SAMPLE_ID)(s))
}

func StopSamples() {
	C.al_stop_samples()
}
