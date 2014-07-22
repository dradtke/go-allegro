package audio

/* -- Audio Stream Fragment -- */

type AudioStreamFragment interface {
	audio_stream_fragment()
}

type audio_stream_fragment_event struct{} // C.ALLEGRO_EVENT_AUDIO_STREAM_FRAGMENT

func (e *audio_stream_fragment_event) audio_stream_fragment() {}

/* -- Audio Stream Finished -- */

type AudioStreamFinished interface {
	audio_stream_finished()
}

type audio_stream_finished_event struct{} // C.ALLEGRO_EVENT_AUDIO_STREAM_FINISHED

func (e *audio_stream_finished_event) audio_stream_finished() {}
