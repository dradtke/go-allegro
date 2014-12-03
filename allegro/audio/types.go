package audio

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
import "C"

type Depth C.ALLEGRO_AUDIO_DEPTH

const (
	AUDIO_DEPTH_INT8     Depth = C.ALLEGRO_AUDIO_DEPTH_INT8
	AUDIO_DEPTH_INT16          = C.ALLEGRO_AUDIO_DEPTH_INT16
	AUDIO_DEPTH_INT24          = C.ALLEGRO_AUDIO_DEPTH_INT24
	AUDIO_DEPTH_FLOAT32        = C.ALLEGRO_AUDIO_DEPTH_FLOAT32
	AUDIO_DEPTH_UNSIGNED       = C.ALLEGRO_AUDIO_DEPTH_UNSIGNED
	AUDIO_DEPTH_UINT8          = C.ALLEGRO_AUDIO_DEPTH_UINT8
	AUDIO_DEPTH_UINT16         = C.ALLEGRO_AUDIO_DEPTH_UINT16
	AUDIO_DEPTH_UINT24         = C.ALLEGRO_AUDIO_DEPTH_UINT24
)

// Return the size of a sample, in bytes, for the given format. The format is
// one of the values listed under ALLEGRO_AUDIO_DEPTH.
func (depth Depth) Size() uint {
	return uint(C.al_get_audio_depth_size(C.ALLEGRO_AUDIO_DEPTH(depth)))
}

type ChannelConf C.ALLEGRO_CHANNEL_CONF

const (
	CHANNEL_CONF_1   ChannelConf = C.ALLEGRO_CHANNEL_CONF_1
	CHANNEL_CONF_2               = C.ALLEGRO_CHANNEL_CONF_2
	CHANNEL_CONF_3               = C.ALLEGRO_CHANNEL_CONF_3
	CHANNEL_CONF_4               = C.ALLEGRO_CHANNEL_CONF_4
	CHANNEL_CONF_5_1             = C.ALLEGRO_CHANNEL_CONF_5_1
	CHANNEL_CONF_6_1             = C.ALLEGRO_CHANNEL_CONF_6_1
	CHANNEL_CONF_7_1             = C.ALLEGRO_CHANNEL_CONF_7_1
)

// Return the number of channels for the given channel configuration, which is
// one of the values listed under ALLEGRO_CHANNEL_CONF.
func (conf ChannelConf) ChannelCount() uint {
	return uint(C.al_get_channel_count(C.ALLEGRO_CHANNEL_CONF(conf)))
}

type PlayMode C.ALLEGRO_PLAYMODE

const (
	PLAYMODE_ONCE  PlayMode = C.ALLEGRO_PLAYMODE_ONCE
	PLAYMODE_LOOP           = C.ALLEGRO_PLAYMODE_LOOP
	PLAYMODE_BIDIR          = C.ALLEGRO_PLAYMODE_BIDIR
)

type SampleID C.ALLEGRO_SAMPLE_ID

const AUDIO_PAN_NONE float32 = -1000.0
