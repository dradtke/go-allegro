package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
*/
import "C"

func Version() uint32 {
	return uint32(C.al_get_allegro_audio_version())
}

func (depth Depth) Size() uint {
	return uint(C.al_get_audio_depth_size(C.ALLEGRO_AUDIO_DEPTH(depth)))
}

func (conf ChannelConf) ChannelCount() uint {
	return uint(C.al_get_channel_count(C.ALLEGRO_CHANNEL_CONF(conf)))
}
