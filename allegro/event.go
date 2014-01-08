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

type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue struct {
	raw   *C.ALLEGRO_EVENT_QUEUE
	event C.ALLEGRO_EVENT
}

// Decrease the reference count of a user-defined event. This must be called on
// any user event that you get from al_get_next_event, al_peek_next_event,
// al_wait_for_event, etc. which is reference counted. This function does
// nothing if the event is not reference counted.
func (ev *Event) UnrefUserEvent() {
	C.al_unref_user_event(ev.User.Raw)
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
	queue := &EventQueue{raw: q}
	//runtime.SetFinalizer(queue, queue.Destroy)
	return queue, nil
}

// Destroy the event queue specified. All event sources currently registered
// with the queue will be automatically unregistered before the queue is
// destroyed.
func (queue *EventQueue) Destroy() {
	C.al_destroy_event_queue(queue.raw)
}

// Register the event source with the event queue specified. An event source
// may be registered with any number of event queues simultaneously, or none.
// Trying to register an event source with the same event queue more than once
// does nothing.
func (queue *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source(queue.raw, (*C.ALLEGRO_EVENT_SOURCE)(source))
}

// Unregister an event source with an event queue. If the event source is not
// actually registered with the event queue, nothing happens.
func (queue *EventQueue) UnregisterEventSource(source *EventSource) {
	C.al_unregister_event_source(queue.raw, (*C.ALLEGRO_EVENT_SOURCE)(source))
}

// Return true if the event queue specified is currently empty.
func (queue *EventQueue) IsEmpty() bool {
	return bool(C.al_is_event_queue_empty(queue.raw))
}

// Copy the contents of the next event in the event queue specified into
// ret_event and return true. The original event packet will remain at the head
// of the queue. If the event queue is actually empty, this function returns
// false and the contents of ret_event are unspecified.
func (queue *EventQueue) PeekNextEvent() (*Event, error) {
	ok := bool(C.al_peek_next_event(queue.raw, &queue.event))
	if !ok {
		return nil, EmptyQueue
	}
	return queue.newEvent(), nil
}

// Drop (remove) the next event from the queue. If the queue is empty, nothing
// happens. Returns true if an event was dropped.
func (queue *EventQueue) DropNextEvent() bool {
	return bool(C.al_drop_next_event(queue.raw))
}

// Drops all events, if any, from the queue.
func (queue *EventQueue) Flush() {
	C.al_flush_event_queue(queue.raw)
}

// Take the next event out of the event queue specified, and copy the contents
// into ret_event, returning true. The original event will be removed from the
// queue. If the event queue is empty, return false and the contents of
// ret_event are unspecified.
func (queue *EventQueue) GetNextEvent() (*Event, error) {
	ok := bool(C.al_get_next_event(queue.raw, &queue.event))
	if !ok {
		return nil, EmptyQueue
	}
	return queue.newEvent(), nil
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEvent(peek bool) *Event {
	if peek {
		C.al_wait_for_event(queue.raw, nil)
		return nil
	} else {
		C.al_wait_for_event(queue.raw, &queue.event)
		return queue.newEvent()
	}
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventTimed(peek bool, secs float32) (bool, *Event) {
	if peek {
		ok := bool(C.al_wait_for_event_timed(queue.raw, nil, C.float(secs)))
		return ok, nil
	} else {
		ok := bool(C.al_wait_for_event_timed(queue.raw, &queue.event, C.float(secs)))
		return ok, queue.newEvent()
	}
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventUntil(timeout *Timeout) (*Event, bool) {
	ok := C.al_wait_for_event_until(queue.raw, &queue.event, (*C.ALLEGRO_TIMEOUT)(timeout))
	if !ok {
		return nil, false
	}
	return queue.newEvent(), true
}

// like WaitForEventUntil, but don't return an event and leave everything on the queue
func (queue *EventQueue) JustWaitForEventUntil(timeout *Timeout) bool {
	return bool(C.al_wait_for_event_until(queue.raw, nil, (*C.ALLEGRO_TIMEOUT)(timeout)))
}

func (queue *EventQueue) newEvent() *Event {
	ev := Event{
		Type:      (EventType)(C.get_event_type(&queue.event)),
		Source:    (*EventSource)(C.get_event_source(&queue.event)),
		Timestamp: float64(C.get_event_timestamp(&queue.event)),
	}
	switch ev.Type {
	case EVENT_JOYSTICK_AXIS:
		id := (*Joystick)(C.get_event_joystick_id(&queue.event))
		stick := int(C.get_event_joystick_stick(&queue.event))
		axis := int(C.get_event_joystick_axis(&queue.event))
		pos := float32(C.get_event_joystick_pos(&queue.event))
		ev.Joystick = JoystickEventInfo{Id: id, Stick: stick, Axis: axis, Pos: pos}

	case EVENT_JOYSTICK_BUTTON_DOWN, EVENT_JOYSTICK_BUTTON_UP:
		id := (*Joystick)(C.get_event_joystick_id(&queue.event))
		button := int(C.get_event_joystick_button(&queue.event))
		ev.Joystick = JoystickEventInfo{Id: id, Button: button}

	case EVENT_JOYSTICK_CONFIGURATION:
		ev.Joystick = JoystickEventInfo{}

	case EVENT_KEY_DOWN, EVENT_KEY_UP:
		keycode := (KeyCode)(C.get_event_keyboard_keycode(&queue.event))
		display := (*Display)(C.get_event_keyboard_display(&queue.event))
		ev.Keyboard = KeyboardEventInfo{KeyCode: keycode, Display: display}

	case EVENT_KEY_CHAR:
		keycode := (KeyCode)(C.get_event_keyboard_keycode(&queue.event))
		unichar := int(C.get_event_keyboard_unichar(&queue.event))
		modifiers := KeyModifier(C.get_event_keyboard_modifiers(&queue.event))
		repeat := bool(C.get_event_keyboard_repeat(&queue.event))
		display := (*Display)(C.get_event_keyboard_display(&queue.event))
		ev.Keyboard = KeyboardEventInfo{KeyCode: keycode, Unichar: unichar, Modifiers: modifiers, Repeat: repeat, Display: display}

	case EVENT_MOUSE_AXES:
		x := int(C.get_event_mouse_x(&queue.event))
		y := int(C.get_event_mouse_y(&queue.event))
		z := int(C.get_event_mouse_z(&queue.event))
		w := int(C.get_event_mouse_w(&queue.event))
		dx := int(C.get_event_mouse_dx(&queue.event))
		dy := int(C.get_event_mouse_dy(&queue.event))
		dz := int(C.get_event_mouse_dz(&queue.event))
		dw := int(C.get_event_mouse_dw(&queue.event))
		display := (*Display)(C.get_event_mouse_display(&queue.event))
		ev.Mouse = MouseEventInfo{X: x, Y: y, Z: z, W: w, Dx: dx, Dy: dy, Dz: dz, Dw: dw, Display: display}

	case EVENT_MOUSE_BUTTON_DOWN, EVENT_MOUSE_BUTTON_UP:
		x := int(C.get_event_mouse_x(&queue.event))
		y := int(C.get_event_mouse_y(&queue.event))
		z := int(C.get_event_mouse_z(&queue.event))
		w := int(C.get_event_mouse_w(&queue.event))
		button := uint(C.get_event_mouse_button(&queue.event))
		display := (*Display)(C.get_event_mouse_display(&queue.event))
		ev.Mouse = MouseEventInfo{X: x, Y: y, Z: z, W: w, Button: button, Display: display}

	case EVENT_MOUSE_WARPED:
		ev.Mouse = MouseEventInfo{}

	case EVENT_MOUSE_ENTER_DISPLAY, EVENT_MOUSE_LEAVE_DISPLAY:
		x := int(C.get_event_mouse_x(&queue.event))
		y := int(C.get_event_mouse_y(&queue.event))
		z := int(C.get_event_mouse_z(&queue.event))
		w := int(C.get_event_mouse_w(&queue.event))
		display := (*Display)(C.get_event_mouse_display(&queue.event))
		ev.Mouse = MouseEventInfo{X: x, Y: y, Z: z, W: w, Display: display}

	case EVENT_TIMER:
		source := (*Timer)(C.get_event_timer_source(&queue.event))
		count := int64(C.get_event_timer_count(&queue.event))
		ev.Timer = TimerEventInfo{Source: source, Count: count}

	case EVENT_DISPLAY_EXPOSE, EVENT_DISPLAY_RESIZE:
		source := (*Display)(C.get_event_display_source(&queue.event))
		x := int(C.get_event_display_x(&queue.event))
		y := int(C.get_event_display_y(&queue.event))
		width := int(C.get_event_display_width(&queue.event))
		height := int(C.get_event_display_height(&queue.event))
		ev.Display = DisplayEventInfo{Source: source, X: x, Y: y, Width: width, Height: height}

	case EVENT_DISPLAY_CLOSE, EVENT_DISPLAY_LOST, EVENT_DISPLAY_FOUND, EVENT_DISPLAY_SWITCH_OUT, EVENT_DISPLAY_SWITCH_IN:
		source := (*Display)(C.get_event_display_source(&queue.event))
		ev.Display = DisplayEventInfo{Source: source}

	case EVENT_DISPLAY_ORIENTATION:
		source := (*Display)(C.get_event_display_source(&queue.event))
		orientation := (DisplayOrientation)(C.get_event_display_orientation(&queue.event))
		ev.Display = DisplayEventInfo{Source: source, Orientation: orientation}

	default:
		source := (*EventSource)(C.get_event_user_source(&queue.event))
		data1 := uintptr(C.get_event_user_data1(&queue.event))
		data2 := uintptr(C.get_event_user_data2(&queue.event))
		data3 := uintptr(C.get_event_user_data3(&queue.event))
		data4 := uintptr(C.get_event_user_data4(&queue.event))
		raw := C.get_user_event(&queue.event)
		ev.User = UserEventInfo{Source: source, Data1: data1, Data2: data2, Data3: data3, Data4: data4, Raw: raw}
	}
	return &ev
}
