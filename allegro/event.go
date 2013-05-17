package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

unsigned int get_event_type(ALLEGRO_EVENT *event) {
	return event->type;
}

ALLEGRO_EVENT_SOURCE *get_event_source(ALLEGRO_EVENT *event) {
	return event->any.source;
}

double get_event_timestamp(ALLEGRO_EVENT *event) {
	return event->any.timestamp;
}

ALLEGRO_DISPLAY *get_event_display_source(ALLEGRO_EVENT *event) {
	return event->display.source;
}
*/
import "C"

var event C.ALLEGRO_EVENT

type Event struct {
	Type EventType
	Source *EventSource
	Timestamp float64
	// TODO: add the other event types
	Display DisplayEvent
}

type DisplayEvent struct {
	Source *Display
}

type EventType int
const (
	DisplayResize EventType = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	DisplayClose EventType = C.ALLEGRO_EVENT_DISPLAY_CLOSE
)

type EventSource struct {
	ptr *C.ALLEGRO_EVENT_SOURCE
}

type EventQueue struct {
	ptr *C.ALLEGRO_EVENT_QUEUE
}

func CreateEventQueue() *EventQueue {
	queue := C.al_create_event_queue()
	if queue == nil {
		return nil
	}
	return &EventQueue{ptr:queue}
}

func (queue *EventQueue) Destroy() {
	C.al_destroy_event_queue(queue.ptr)
}

func (queue *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source(queue.ptr, source.ptr)
}

func (queue *EventQueue) GetNextEvent() (*Event, bool) {
	success := gobool(C.al_get_next_event(queue.ptr, &event))
	if !success {
		return nil, false
	}
	return newEvent(), true
}

func (queue *EventQueue) WaitForEvent() *Event {
	C.al_wait_for_event(queue.ptr, &event)
	return newEvent()
}

// wait for an event, but don't take it off the queue
// better name for this?
func (queue *EventQueue) JustWaitForEvent() {
	C.al_wait_for_event(queue.ptr, nil)
}

func (queue *EventQueue) WaitForEventUntil(timeout *Timeout) (*Event, bool) {
	success := C.al_wait_for_event_until(queue.ptr, &event, &timeout.ptr)
	if !success {
		return nil, false
	}
	return newEvent(), true
}

// like WaitForEventUntil, but don't return an event and leave everything on the queue
func (queue *EventQueue) JustWaitForEventUntil(timeout *Timeout) bool {
	return gobool(C.al_wait_for_event_until(queue.ptr, nil, &timeout.ptr))
}

func newEvent() *Event {
	ev := Event{
		Type: (EventType)(C.get_event_type(&event)),
		Source: &EventSource{ptr:C.get_event_source(&event)},
		Timestamp: godouble(C.get_event_timestamp(&event)),
	}
	switch ev.Type {
		case DisplayResize, DisplayClose:
			source := C.get_event_display_source(&event)
			for e := displayList.Front(); e != nil; e = e.Next() {
				display := e.Value.(*Display)
				if display.ptr == source {
					ev.Display = DisplayEvent{Source:display}
					break
				}
			}
			//e.Display = DisplayEvent{Source:&Display{ptr:C.get_event_display_source(&event)}}
	}
	return &ev
}
