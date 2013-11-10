package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
)

type Timer C.ALLEGRO_TIMER

func CreateTimer(speed float64) (*Timer, error) {
	timer := C.al_create_timer(C.double(speed))
	if timer == nil {
		return nil, errors.New("failed to create timer")
	}
	return (*Timer)(timer), nil
}

func (t *Timer) Destroy() {
	C.al_destroy_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) Start() {
	C.al_start_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) Stop() {
	C.al_stop_timer((*C.ALLEGRO_TIMER)(t))
}

func (t *Timer) IsStarted() bool {
	return bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) Speed() float64 {
	return float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) SetSpeed(speed float64) {
	C.al_set_timer_speed((*C.ALLEGRO_TIMER)(t), C.double(speed))
}

func (t *Timer) Count() int64 {
	return int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(t)))
}

func (t *Timer) SetCount(count int64) {
	C.al_set_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(count))
}

func (t *Timer) AddCount(diff int64) {
	C.al_add_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(diff))
}

func (t *Timer) EventSource() *EventSource {
	return (*EventSource)(C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(t)))
}
