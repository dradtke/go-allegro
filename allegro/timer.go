package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
)

type Timer C.ALLEGRO_TIMER

// Allocates and initializes a timer. If successful, a pointer to a new timer
// object is returned, otherwise NULL is returned. speed_secs is in seconds per
// "tick", and must be positive. The new timer is initially stopped.
func CreateTimer(speed float64) (*Timer, error) {
	t := C.al_create_timer(C.double(speed))
	if t == nil {
		return nil, errors.New("failed to create timer")
	}
	timer := (*Timer)(t)
	//runtime.SetFinalizer(timer, timer.Destroy)
	return timer, nil
}

// Uninstall the timer specified. If the timer is started, it will
// automatically be stopped before uninstallation. It will also automatically
// unregister the timer with any event queues.
func (t *Timer) Destroy() {
	C.al_destroy_timer((*C.ALLEGRO_TIMER)(t))
}

// Start the timer specified. From then, the timer's counter will increment at
// a constant rate, and it will begin generating events. Starting a timer that
// is already started does nothing.
func (t *Timer) Start() {
	C.al_start_timer((*C.ALLEGRO_TIMER)(t))
}

// Stop the timer specified. The timer's counter will stop incrementing and it
// will stop generating events. Stopping a timer that is already stopped does
// nothing.
func (t *Timer) Stop() {
	C.al_stop_timer((*C.ALLEGRO_TIMER)(t))
}

// Return true if the timer specified is currently started.
func (t *Timer) IsStarted() bool {
	return bool(C.al_get_timer_started((*C.ALLEGRO_TIMER)(t)))
}

// Return the timer's speed, in seconds. (The same value passed to
// al_create_timer or al_set_timer_speed.)
func (t *Timer) Speed() float64 {
	return float64(C.al_get_timer_speed((*C.ALLEGRO_TIMER)(t)))
}

// Set the timer's speed, i.e. the rate at which its counter will be
// incremented when it is started. This can be done when the timer is started
// or stopped. If the timer is currently running, it is made to look as though
// the speed change occurred precisely at the last tick.
func (t *Timer) SetSpeed(speed float64) {
	C.al_set_timer_speed((*C.ALLEGRO_TIMER)(t), C.double(speed))
}

// Return the timer's counter value. The timer can be started or stopped.
func (t *Timer) Count() int64 {
	return int64(C.al_get_timer_count((*C.ALLEGRO_TIMER)(t)))
}

// Set the timer's counter value. The timer can be started or stopped. The
// count value may be positive or negative, but will always be incremented by
// +1 at each tick.
func (t *Timer) SetCount(count int64) {
	C.al_set_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(count))
}

// Add diff to the timer's counter value. This is similar to writing:
func (t *Timer) AddCount(diff int64) {
	C.al_add_timer_count((*C.ALLEGRO_TIMER)(t), C.int64_t(diff))
}

// Retrieve the associated event source.
func (t *Timer) EventSource() *EventSource {
	return (*EventSource)(C.al_get_timer_event_source((*C.ALLEGRO_TIMER)(t)))
}
