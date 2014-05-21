package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

void set_user_data(ALLEGRO_EVENT *event, intptr_t data1, intptr_t data2, intptr_t data3, intptr_t data4) {
	event->user.data1 = data1;
	event->user.data2 = data2;
	event->user.data3 = data3;
	event->user.data4 = data4;
}

//{{{ event getters
// Common

unsigned int get_event_type(ALLEGRO_EVENT *event) {
	return event->type;
}

ALLEGRO_EVENT_SOURCE *get_event_source(ALLEGRO_EVENT *event) {
	return event->any.source;
}

double get_event_timestamp(ALLEGRO_EVENT *event) {
	return event->any.timestamp;
}

// Joystick

ALLEGRO_JOYSTICK *get_event_joystick_id(ALLEGRO_EVENT *event) {
	return event->joystick.id;
}

int get_event_joystick_stick(ALLEGRO_EVENT *event) {
	return event->joystick.stick;
}

int get_event_joystick_axis(ALLEGRO_EVENT *event) {
	return event->joystick.axis;
}

float get_event_joystick_pos(ALLEGRO_EVENT *event) {
	return event->joystick.pos;
}

int get_event_joystick_button(ALLEGRO_EVENT *event) {
	return event->joystick.button;
}

// Keys

int get_event_keyboard_keycode(ALLEGRO_EVENT *event) {
	return event->keyboard.keycode;
}

int get_event_keyboard_unichar(ALLEGRO_EVENT *event) {
	return event->keyboard.unichar;
}

unsigned int get_event_keyboard_modifiers(ALLEGRO_EVENT *event) {
	return event->keyboard.modifiers;
}

bool get_event_keyboard_repeat(ALLEGRO_EVENT *event) {
	return event->keyboard.repeat;
}

ALLEGRO_DISPLAY *get_event_keyboard_display(ALLEGRO_EVENT *event) {
	return event->keyboard.display;
}

// Mouse

int get_event_mouse_x(ALLEGRO_EVENT *event) {
	return event->mouse.x;
}

int get_event_mouse_y(ALLEGRO_EVENT *event) {
	return event->mouse.y;
}

int get_event_mouse_z(ALLEGRO_EVENT *event) {
	return event->mouse.z;
}

int get_event_mouse_w(ALLEGRO_EVENT *event) {
	return event->mouse.w;
}

int get_event_mouse_dx(ALLEGRO_EVENT *event) {
	return event->mouse.dx;
}

int get_event_mouse_dy(ALLEGRO_EVENT *event) {
	return event->mouse.dy;
}

int get_event_mouse_dz(ALLEGRO_EVENT *event) {
	return event->mouse.dz;
}

int get_event_mouse_dw(ALLEGRO_EVENT *event) {
	return event->mouse.dw;
}

unsigned int get_event_mouse_button(ALLEGRO_EVENT *event) {
	return event->mouse.button;
}

ALLEGRO_DISPLAY *get_event_mouse_display(ALLEGRO_EVENT *event) {
	return event->mouse.display;
}

// Timer

ALLEGRO_TIMER *get_event_timer_source(ALLEGRO_EVENT *event) {
	return event->timer.source;
}

int64_t get_event_timer_count(ALLEGRO_EVENT *event) {
	return event->timer.count;
}

// Display

ALLEGRO_DISPLAY *get_event_display_source(ALLEGRO_EVENT *event) {
	return event->display.source;
}

int get_event_display_x(ALLEGRO_EVENT *event) {
	return event->display.x;
}

int get_event_display_y(ALLEGRO_EVENT *event) {
	return event->display.y;
}

int get_event_display_width(ALLEGRO_EVENT *event) {
	return event->display.width;
}

int get_event_display_height(ALLEGRO_EVENT *event) {
	return event->display.height;
}

int get_event_display_orientation(ALLEGRO_EVENT *event) {
	return event->display.orientation;
}

// User

ALLEGRO_EVENT_SOURCE *get_event_user_source(ALLEGRO_EVENT *event) {
	return event->user.source;
}

intptr_t get_event_user_data1(ALLEGRO_EVENT *event) {
	return event->user.data1;
}

intptr_t get_event_user_data2(ALLEGRO_EVENT *event) {
	return event->user.data2;
}

intptr_t get_event_user_data3(ALLEGRO_EVENT *event) {
	return event->user.data3;
}

intptr_t get_event_user_data4(ALLEGRO_EVENT *event) {
	return event->user.data4;
}

ALLEGRO_USER_EVENT *get_user_event(ALLEGRO_EVENT *event) {
	return &event->user;
}

//}}}
*/
import "C"
import (
	"errors"
	"fmt"
)

var EmptyQueue = errors.New("event queue is empty")

type Event struct {
	Type      EventType
	Source    *EventSource
	Timestamp float64
	// TODO: add the other event types
	Joystick JoystickEventInfo
	Keyboard KeyboardEventInfo
	Mouse    MouseEventInfo
	Timer    TimerEventInfo
	Display  DisplayEventInfo
	User     UserEventInfo
	raw      C.ALLEGRO_EVENT
}

type JoystickEventInfo struct {
	Id                  *Joystick
	Stick, Axis, Button int
	Pos                 float32
}

type KeyboardEventInfo struct {
	KeyCode   KeyCode
	Display   *Display
	Unichar   int
	Modifiers KeyModifier
	Repeat    bool
}

type MouseEventInfo struct {
	X, Y, Z, W, Dx, Dy, Dz, Dw int
	Button                     uint
	Display                    *Display
}

type TimerEventInfo struct {
	Source *Timer
	Count  int64
}

type DisplayEventInfo struct {
	Source              *Display
	X, Y, Width, Height int
	Orientation         DisplayOrientation
}

type UserEventInfo struct {
	Source                     *EventSource
	Data1, Data2, Data3, Data4 uintptr
	Raw                        *C.ALLEGRO_USER_EVENT // used for memory management
}

type EventType int

const (
	EVENT_JOYSTICK_AXIS          EventType = C.ALLEGRO_EVENT_JOYSTICK_AXIS
	EVENT_JOYSTICK_BUTTON_DOWN   EventType = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
	EVENT_JOYSTICK_BUTTON_UP     EventType = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
	EVENT_JOYSTICK_CONFIGURATION EventType = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION
	EVENT_KEY_DOWN               EventType = C.ALLEGRO_EVENT_KEY_DOWN
	EVENT_KEY_UP                 EventType = C.ALLEGRO_EVENT_KEY_UP
	EVENT_KEY_CHAR               EventType = C.ALLEGRO_EVENT_KEY_CHAR
	EVENT_MOUSE_AXES             EventType = C.ALLEGRO_EVENT_MOUSE_AXES
	EVENT_MOUSE_BUTTON_DOWN      EventType = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
	EVENT_MOUSE_BUTTON_UP        EventType = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
	EVENT_MOUSE_WARPED           EventType = C.ALLEGRO_EVENT_MOUSE_WARPED
	EVENT_MOUSE_ENTER_DISPLAY    EventType = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
	EVENT_MOUSE_LEAVE_DISPLAY    EventType = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
	EVENT_TIMER                  EventType = C.ALLEGRO_EVENT_TIMER
	EVENT_DISPLAY_EXPOSE         EventType = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
	EVENT_DISPLAY_RESIZE         EventType = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	EVENT_DISPLAY_CLOSE          EventType = C.ALLEGRO_EVENT_DISPLAY_CLOSE
	EVENT_DISPLAY_LOST           EventType = C.ALLEGRO_EVENT_DISPLAY_LOST
	EVENT_DISPLAY_FOUND          EventType = C.ALLEGRO_EVENT_DISPLAY_FOUND
	EVENT_DISPLAY_SWITCH_OUT     EventType = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
	EVENT_DISPLAY_SWITCH_IN      EventType = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
	EVENT_DISPLAY_ORIENTATION    EventType = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
)

func (typ EventType) Name() string {
	switch typ {
	case EVENT_JOYSTICK_AXIS:
		return "Joystick Axis"
	case EVENT_JOYSTICK_BUTTON_DOWN:
		return "Joystick Button Down"
	case EVENT_JOYSTICK_BUTTON_UP:
		return "Joystick Button Up"
	case EVENT_JOYSTICK_CONFIGURATION:
		return "Joystick Configuration"
	case EVENT_KEY_DOWN:
		return "Key Down"
	case EVENT_KEY_UP:
		return "Key Up"
	case EVENT_KEY_CHAR:
		return "Key Char"
	case EVENT_MOUSE_AXES:
		return "Mouse Axes"
	case EVENT_MOUSE_BUTTON_DOWN:
		return "Mouse Button Down"
	case EVENT_MOUSE_BUTTON_UP:
		return "Mouse Button Up"
	case EVENT_MOUSE_WARPED:
		return "Mouse Warped"
	case EVENT_MOUSE_ENTER_DISPLAY:
		return "Mouse Enter Display"
	case EVENT_MOUSE_LEAVE_DISPLAY:
		return "Mouse Leave Display"
	case EVENT_TIMER:
		return "Timer"
	case EVENT_DISPLAY_EXPOSE:
		return "Display Expose"
	case EVENT_DISPLAY_RESIZE:
		return "Display Resize"
	case EVENT_DISPLAY_CLOSE:
		return "Display Close"
	case EVENT_DISPLAY_LOST:
		return "Display Lost"
	case EVENT_DISPLAY_FOUND:
		return "Display Found"
	case EVENT_DISPLAY_SWITCH_OUT:
		return "Display Switch Out"
	case EVENT_DISPLAY_SWITCH_IN:
		return "Display Switch In"
	case EVENT_DISPLAY_ORIENTATION:
		return "Display Orientation"
	}
	return "Unknown"
}

type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue C.ALLEGRO_EVENT_QUEUE

// Decrease the reference count of a user-defined event. This must be called on
// any user event that you get from al_get_next_event, al_peek_next_event,
// al_wait_for_event, etc. which is reference counted. This function does
// nothing if the event is not reference counted.
func (e *Event) UnrefUserEvent() {
	C.al_unref_user_event(e.User.Raw)
}

// Initialise an event source for emitting user events. The space for the event
// source must already have been allocated.
func (source EventSource) InitUserEventSource() {
	C.al_init_user_event_source((*C.ALLEGRO_EVENT_SOURCE)(&source))
}

// Emit a user event. The event source must have been initialised with
// al_init_user_event_source. Returns false if the event source isn't
// registered with any queues, hence the event wouldn't have been delivered
// into any queues.
func (source *EventSource) EmitUserEvent(data []uintptr) error {
	var l = len(data)
	if l > 4 {
		return fmt.Errorf("too many parameters: %d > 4", l)
	}
	var data1, data2, data3, data4 C.intptr_t = 0, 0, 0, 0
	if l > 0 {
		data1 = C.intptr_t(data[0])
	}
	if l > 1 {
		data2 = C.intptr_t(data[1])
	}
	if l > 2 {
		data3 = C.intptr_t(data[2])
	}
	if l > 3 {
		data4 = C.intptr_t(data[3])
	}
	var event C.ALLEGRO_EVENT
	C.set_user_data(&event, data1, data2, data3, data4)
	ok := bool(C.al_emit_user_event((*C.ALLEGRO_EVENT_SOURCE)(source), &event, nil))
	if !ok {
		return errors.New("failed to emit user event")
	}
	return nil
}

// Destroy an event source initialised with al_init_user_event_source.
func (source *EventSource) DestroyUserEventSource() {
	C.al_destroy_user_event_source((*C.ALLEGRO_EVENT_SOURCE)(source))
}

// Assign the abstract user data to the event source. Allegro does not use the
// data internally for anything; it is simply meant as a convenient way to
// associate your own data or objects with events.
func (source *EventSource) SetData(data uintptr) {
	C.al_set_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(source), C.intptr_t(data))
}

// Returns the abstract user data associated with the event source. If no data
// was previously set, returns NULL.
func (source *EventSource) Data() uintptr {
	return uintptr(C.al_get_event_source_data((*C.ALLEGRO_EVENT_SOURCE)(source)))
}

// Create a new, empty event queue, returning a pointer to object if
// successful. Returns NULL on error.
func CreateEventQueue() (*EventQueue, error) {
	q := C.al_create_event_queue()
	if q == nil {
		return nil, errors.New("failed to create event queue!")
	}
	return (*EventQueue)(q), nil
}

// Destroy the event queue specified. All event sources currently registered
// with the queue will be automatically unregistered before the queue is
// destroyed.
func (queue *EventQueue) Destroy() {
	C.al_destroy_event_queue((*C.ALLEGRO_EVENT_QUEUE)(queue))
}

// Shorthand method for registering anything with an EventSource() method.
func (queue *EventQueue) Register(obs ...EventGenerator) {
	for _, ob := range obs {
		queue.RegisterEventSource(ob.EventSource())
	}
}

// Register the event source with the event queue specified. An event source
// may be registered with any number of event queues simultaneously, or none.
// Trying to register an event source with the same event queue more than once
// does nothing.
func (queue *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT_SOURCE)(source))
}

// Shorthand method for registering anything with an EventSource() method.
func (queue *EventQueue) Unregister(ob EventGenerator) {
	queue.UnregisterEventSource(ob.EventSource())
}

// Unregister an event source with an event queue. If the event source is not
// actually registered with the event queue, nothing happens.
func (queue *EventQueue) UnregisterEventSource(source *EventSource) {
	C.al_unregister_event_source((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT_SOURCE)(source))
}

// Return true if the event queue specified is currently empty.
func (queue *EventQueue) IsEmpty() bool {
	return bool(C.al_is_event_queue_empty((*C.ALLEGRO_EVENT_QUEUE)(queue)))
}

// Copy the contents of the next event in the event queue specified into
// ret_event and return true. The original event packet will remain at the head
// of the queue. If the event queue is actually empty, this function returns
// false and the contents of ret_event are unspecified.
func (queue *EventQueue) PeekNextEvent(event *Event) error {
	ok := bool(C.al_peek_next_event((*C.ALLEGRO_EVENT_QUEUE)(queue), &event.raw))
	if !ok {
		return EmptyQueue
	}
	return nil
}

// Drop (remove) the next event from the queue. If the queue is empty, nothing
// happens. Returns true if an event was dropped.
func (queue *EventQueue) DropNextEvent() bool {
	return bool(C.al_drop_next_event((*C.ALLEGRO_EVENT_QUEUE)(queue)))
}

// Drops all events, if any, from the queue.
func (queue *EventQueue) Flush() {
	C.al_flush_event_queue((*C.ALLEGRO_EVENT_QUEUE)(queue))
}

// Take the next event out of the event queue specified, and copy the contents
// into ret_event, returning true. The original event will be removed from the
// queue. If the event queue is empty, return false and the contents of
// ret_event are unspecified.
func (queue *EventQueue) GetNextEvent(event *Event) error {
	ok := bool(C.al_get_next_event((*C.ALLEGRO_EVENT_QUEUE)(queue), &event.raw))
	if !ok {
		return EmptyQueue
	}
	event.hydrate()
	return nil
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEvent(event *Event) {
	var raw *C.ALLEGRO_EVENT = nil
	if event != nil {
		raw = &event.raw
	}
	C.al_wait_for_event((*C.ALLEGRO_EVENT_QUEUE)(queue), raw)
	if event != nil {
		event.hydrate()
	}
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventTimed(event *Event, secs float32) bool {
	var raw *C.ALLEGRO_EVENT = nil
	if event != nil {
		raw = &event.raw
	}
	ok := bool(C.al_wait_for_event_timed((*C.ALLEGRO_EVENT_QUEUE)(queue), raw, C.float(secs)))
	if !ok {
		return false
	} else {
		if event != nil {
			event.hydrate()
		}
		return true
	}
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventUntil(timeout *Timeout, event *Event) bool {
	var raw *C.ALLEGRO_EVENT = nil
	if event != nil {
		raw = &event.raw
	}
	ok := C.al_wait_for_event_until((*C.ALLEGRO_EVENT_QUEUE)(queue), raw, (*C.ALLEGRO_TIMEOUT)(timeout))
	if !ok {
		return false
	} else {
		if event != nil {
			event.hydrate()
		}
		return true
	}
}

// hydrate() fills the Event struct's public fields with data from its internal
// raw pointer. This is usually called right after the raw event is filled with
// data from one of the above methods.
func (event *Event) hydrate() {
	event.Type = (EventType)(C.get_event_type(&event.raw))
	event.Source = (*EventSource)(C.get_event_source(&event.raw))
	event.Timestamp = float64(C.get_event_timestamp(&event.raw))
	switch event.Type {
	case EVENT_JOYSTICK_AXIS:
		event.Joystick.Id = (*Joystick)(C.get_event_joystick_id(&event.raw))
		event.Joystick.Stick = int(C.get_event_joystick_stick(&event.raw))
		event.Joystick.Axis = int(C.get_event_joystick_axis(&event.raw))
		event.Joystick.Pos = float32(C.get_event_joystick_pos(&event.raw))

	case EVENT_JOYSTICK_BUTTON_DOWN, EVENT_JOYSTICK_BUTTON_UP:
		event.Joystick.Id = (*Joystick)(C.get_event_joystick_id(&event.raw))
		event.Joystick.Button = int(C.get_event_joystick_button(&event.raw))

	case EVENT_JOYSTICK_CONFIGURATION:
		// no information

	case EVENT_KEY_DOWN, EVENT_KEY_UP:
		event.Keyboard.KeyCode = (KeyCode)(C.get_event_keyboard_keycode(&event.raw))
		event.Keyboard.Display = (*Display)(C.get_event_keyboard_display(&event.raw))

	case EVENT_KEY_CHAR:
		event.Keyboard.KeyCode = (KeyCode)(C.get_event_keyboard_keycode(&event.raw))
		event.Keyboard.Unichar = int(C.get_event_keyboard_unichar(&event.raw))
		event.Keyboard.Modifiers = KeyModifier(C.get_event_keyboard_modifiers(&event.raw))
		event.Keyboard.Repeat = bool(C.get_event_keyboard_repeat(&event.raw))
		event.Keyboard.Display = (*Display)(C.get_event_keyboard_display(&event.raw))

	case EVENT_MOUSE_AXES:
		event.Mouse.X = int(C.get_event_mouse_x(&event.raw))
		event.Mouse.Y = int(C.get_event_mouse_y(&event.raw))
		event.Mouse.Z = int(C.get_event_mouse_z(&event.raw))
		event.Mouse.W = int(C.get_event_mouse_w(&event.raw))
		event.Mouse.Dx = int(C.get_event_mouse_dx(&event.raw))
		event.Mouse.Dy = int(C.get_event_mouse_dy(&event.raw))
		event.Mouse.Dz = int(C.get_event_mouse_dz(&event.raw))
		event.Mouse.Dw = int(C.get_event_mouse_dw(&event.raw))
		event.Mouse.Display = (*Display)(C.get_event_mouse_display(&event.raw))

	case EVENT_MOUSE_BUTTON_DOWN, EVENT_MOUSE_BUTTON_UP:
		event.Mouse.X = int(C.get_event_mouse_x(&event.raw))
		event.Mouse.Y = int(C.get_event_mouse_y(&event.raw))
		event.Mouse.Z = int(C.get_event_mouse_z(&event.raw))
		event.Mouse.W = int(C.get_event_mouse_w(&event.raw))
		event.Mouse.Button = uint(C.get_event_mouse_button(&event.raw))
		event.Mouse.Display = (*Display)(C.get_event_mouse_display(&event.raw))

	case EVENT_MOUSE_WARPED:
		// no information

	case EVENT_MOUSE_ENTER_DISPLAY, EVENT_MOUSE_LEAVE_DISPLAY:
		event.Mouse.X = int(C.get_event_mouse_x(&event.raw))
		event.Mouse.Y = int(C.get_event_mouse_y(&event.raw))
		event.Mouse.Z = int(C.get_event_mouse_z(&event.raw))
		event.Mouse.W = int(C.get_event_mouse_w(&event.raw))
		event.Mouse.Display = (*Display)(C.get_event_mouse_display(&event.raw))

	case EVENT_TIMER:
		event.Timer.Source = (*Timer)(C.get_event_timer_source(&event.raw))
		event.Timer.Count = int64(C.get_event_timer_count(&event.raw))

	case EVENT_DISPLAY_EXPOSE, EVENT_DISPLAY_RESIZE:
		event.Display.Source = (*Display)(C.get_event_display_source(&event.raw))
		event.Display.X = int(C.get_event_display_x(&event.raw))
		event.Display.Y = int(C.get_event_display_y(&event.raw))
		event.Display.Width = int(C.get_event_display_width(&event.raw))
		event.Display.Height = int(C.get_event_display_height(&event.raw))

	case EVENT_DISPLAY_CLOSE, EVENT_DISPLAY_LOST, EVENT_DISPLAY_FOUND, EVENT_DISPLAY_SWITCH_OUT, EVENT_DISPLAY_SWITCH_IN:
		event.Display.Source = (*Display)(C.get_event_display_source(&event.raw))

	case EVENT_DISPLAY_ORIENTATION:
		event.Display.Source = (*Display)(C.get_event_display_source(&event.raw))
		event.Display.Orientation = (DisplayOrientation)(C.get_event_display_orientation(&event.raw))

	default:
		event.User.Source = (*EventSource)(C.get_event_user_source(&event.raw))
		event.User.Data1 = uintptr(C.get_event_user_data1(&event.raw))
		event.User.Data2 = uintptr(C.get_event_user_data2(&event.raw))
		event.User.Data3 = uintptr(C.get_event_user_data3(&event.raw))
		event.User.Data4 = uintptr(C.get_event_user_data4(&event.raw))
		event.User.Raw = C.get_user_event(&event.raw)
	}
}
