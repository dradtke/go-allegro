package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>

void free_string(char *str) {
	al_free(str);
}

void *create_sample_buffer(uint samples, uint channel_count, uint depth_size) {
	return al_malloc(samples * channel_count * depth_size);
}
*/
import "C"
import (
	"errors"
	"fmt"
)

type Sample C.ALLEGRO_SAMPLE

func CreateSample(samples, freq uint, depth Depth, chan_conf ChannelConf) *Sample {
	buf := C.create_sample_buffer(C.uint(samples),
		C.uint(chan_conf.ChannelCount()), C.uint(depth.Size()))
	return (*Sample)(C.al_create_sample(
		buf,
		C.uint(samples),
		C.uint(freq),
		C.ALLEGRO_AUDIO_DEPTH(depth),
		C.ALLEGRO_CHANNEL_CONF(chan_conf),
		C.bool(true)))
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
