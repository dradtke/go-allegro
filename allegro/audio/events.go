package audio

/*
#cgo pkg-config: allegro_audio-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
#include "../util.c"
*/
import "C"
import (
    "github.com/dradtke/go-allegro/allegro"
)

const (
    EVENT_AUDIO_STREAM_FRAGMENT allegro.EventType = C.ALLEGRO_EVENT_AUDIO_STREAM_FRAGMENT
    EVENT_AUDIO_STREAM_FINISHED allegro.EventType = C.ALLEGRO_EVENT_AUDIO_STREAM_FINISHED
)

