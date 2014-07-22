package audio

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_audio.h>
// #include "../util.c"
import "C"
import (
	"github.com/dradtke/go-allegro/allegro"
	"unsafe"
)

func init() {
	allegro.RegisterEventType(C.ALLEGRO_EVENT_AUDIO_STREAM_FRAGMENT, func(e *allegro.Event) interface{} {
		return (*audio_stream_fragment_event)(unsafe.Pointer(e))
	})
	allegro.RegisterEventType(C.ALLEGRO_EVENT_AUDIO_STREAM_FINISHED, func(e *allegro.Event) interface{} {
		return (*audio_stream_finished_event)(unsafe.Pointer(e))
	})
}
