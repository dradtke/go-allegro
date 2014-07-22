package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var registeredEvents = make(map[C.ALLEGRO_EVENT_TYPE]func(e *Event) interface{})

var EmptyQueue = errors.New("event queue is empty")

type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue C.ALLEGRO_EVENT_QUEUE

// Initialise an event source for emitting user events. The space for the event
// source must already have been allocated.
func (source EventSource) InitUserEventSource() {
	C.al_init_user_event_source((*C.ALLEGRO_EVENT_SOURCE)(&source))
}

// Emit a user event. The event source must have been initialised with
// al_init_user_event_source. Returns false if the event source isn't
// registered with any queues, hence the event wouldn't have been delivered
// into any queues.
func (source *EventSource) EmitUserEvent(data ...uintptr) error {
	data_len := len(data)
	if data_len > 4 {
		return fmt.Errorf("too many parameters: %d > 4", data_len)
	}
	var data1, data2, data3, data4 C.intptr_t = 0, 0, 0, 0
	switch data_len {
	case 4:
		data4 = C.intptr_t(data[3])
		fallthrough
	case 3:
		data3 = C.intptr_t(data[2])
		fallthrough
	case 2:
		data2 = C.intptr_t(data[1])
		fallthrough
	case 1:
		data1 = C.intptr_t(data[0])
	}
	event := C.struct_ALLEGRO_USER_EVENT{data1: data1, data2: data2, data3: data3, data4: data4}
	if ok := bool(C.al_emit_user_event((*C.ALLEGRO_EVENT_SOURCE)(source), (*C.ALLEGRO_EVENT)(unsafe.Pointer(&event)), nil)); !ok {
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
func (queue *EventQueue) PeekNextEvent(event *Event) (interface{}, error) {
	if ok := bool(C.al_peek_next_event((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT)(event))); !ok {
		return nil, EmptyQueue
	}
	return event.cast(), nil
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
func (queue *EventQueue) GetNextEvent(event *Event) (interface{}, error) {
	if ok := bool(C.al_get_next_event((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT)(event))); !ok {
		return nil, EmptyQueue
	}
	return event.cast(), nil
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEvent(event *Event) interface{} {
	C.al_wait_for_event((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT)(event))
	if event == nil {
		return nil
	}
	return event.cast()
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventTimed(event *Event, secs float32) (interface{}, bool) {
	if ok := bool(C.al_wait_for_event_timed((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT)(event), C.float(secs))); !ok {
		return nil, false
	}
	if event == nil {
		return nil, true
	}
	return event.cast(), true
}

// Wait until the event queue specified is non-empty. If ret_event is not NULL,
// the first event in the queue will be copied into ret_event and removed from
// the queue. If ret_event is NULL the first event is left at the head of the
// queue.
func (queue *EventQueue) WaitForEventUntil(timeout *Timeout, event *Event) (interface{}, bool) {
	if ok := C.al_wait_for_event_until((*C.ALLEGRO_EVENT_QUEUE)(queue), (*C.ALLEGRO_EVENT)(event), (*C.ALLEGRO_TIMEOUT)(timeout)); !ok {
		return nil, false
	}
	if event == nil {
		return nil, true
	}
	return event.cast(), true
}

type Event C.union_ALLEGRO_EVENT

// RegisterEventType() lets modules register their own event types.
func RegisterEventType(t C.ALLEGRO_EVENT_TYPE, f func(*Event) interface{}) {
	registeredEvents[t] = f
}

func (e *Event) cast() interface{} {
	switch t := C.ALLEGRO_EVENT_TYPE(e[0]); t {
	case C.ALLEGRO_EVENT_JOYSTICK_AXIS:
		return (*joystick_axis_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN:
		return (*joystick_button_down_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP:
		return (*joystick_button_up_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION:
		return (*joystick_configuration_event)(unsafe.Pointer(e))

	case C.ALLEGRO_EVENT_KEY_DOWN:
		return (*key_down_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_KEY_UP:
		return (*key_up_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_KEY_CHAR:
		return (*key_char_event)(unsafe.Pointer(e))

	case C.ALLEGRO_EVENT_MOUSE_AXES:
		return (*mouse_axes_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN:
		return (*mouse_button_down_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_MOUSE_BUTTON_UP:
		return (*mouse_button_up_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_MOUSE_WARPED:
		return (*mouse_warped_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY:
		return (*mouse_enter_display_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY:
		return (*mouse_leave_display_event)(unsafe.Pointer(e))

	case C.ALLEGRO_EVENT_TIMER:
		return (*timer_event)(unsafe.Pointer(e))

	case C.ALLEGRO_EVENT_DISPLAY_EXPOSE:
		return (*display_expose_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_RESIZE:
		return (*display_resize_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_CLOSE:
		return (*display_close_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_LOST:
		return (*display_lost_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_FOUND:
		return (*display_found_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT:
		return (*display_switch_out_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN:
		return (*display_switch_in_event)(unsafe.Pointer(e))
	case C.ALLEGRO_EVENT_DISPLAY_ORIENTATION:
		return (*display_orientation_event)(unsafe.Pointer(e))

	default:
		if f, ok := registeredEvents[t]; ok {
			return f(e)
		} else {
			return (*user_event)(unsafe.Pointer(e))
		}
	}
}

/* -- Joystick Axis -- */

type JoystickAxisEvent interface {
	joystick_axis()
	Timestamp() float64
	Id() *Joystick
	Stick() int
	Axis() int
	Pos() float32
}

type joystick_axis_event C.struct_ALLEGRO_JOYSTICK_EVENT

func (e *joystick_axis_event) joystick_axis() {}

func (e *joystick_axis_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *joystick_axis_event) Id() *Joystick {
	return (*Joystick)(e.id)
}

func (e *joystick_axis_event) Stick() int {
	return int(e.stick)
}

func (e *joystick_axis_event) Axis() int {
	return int(e.axis)
}

func (e *joystick_axis_event) Pos() float32 {
	return float32(e.pos)
}

/* -- Joystick Button Down -- */

type JoystickButtonDownEvent interface {
	joystick_button_down()
	Timestamp() float64
	Id() *Joystick
	Button() int
}

type joystick_button_down_event C.struct_ALLEGRO_JOYSTICK_EVENT

func (e *joystick_button_down_event) joystick_button_down() {}

func (e *joystick_button_down_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *joystick_button_down_event) Id() *Joystick {
	return (*Joystick)(e.id)
}

func (e *joystick_button_down_event) Button() int {
	return int(e.button)
}

/* -- Joystick Button Up -- */

type JoystickButtonUpEvent interface {
	joystick_button_up()
	Timestamp() float64
	Id() *Joystick
	Button() int
}

type joystick_button_up_event C.struct_ALLEGRO_JOYSTICK_EVENT

func (e *joystick_button_up_event) joystick_button_down() {}

func (e *joystick_button_up_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *joystick_button_up_event) Id() *Joystick {
	return (*Joystick)(e.id)
}

func (e *joystick_button_up_event) Button() int {
	return int(e.button)
}

/* -- Joystick Configuration -- */

type JoystickConfigurationEvent interface {
	joystick_configuration()
	Timestamp() float64
}

type joystick_configuration_event C.struct_ALLEGRO_JOYSTICK_EVENT

func (e *joystick_configuration_event) joystick_configuration() {}

func (e *joystick_configuration_event) Timestamp() float64 {
	return float64(e.timestamp)
}

/* -- Key Down -- */

type KeyDownEvent interface {
	key_down()
	Timestamp() float64
	Source() *Keyboard
	KeyCode() KeyCode
	Display() *Display
}

type key_down_event C.struct_ALLEGRO_KEYBOARD_EVENT

func (e *key_down_event) key_down() {}

func (e *key_down_event) Source() *Keyboard {
	return (*Keyboard)(e.source)
}

func (e *key_down_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *key_down_event) KeyCode() KeyCode {
	return KeyCode(e.keycode)
}

func (e *key_down_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Key Up -- */

type KeyUpEvent interface {
	key_up()
	Timestamp() float64
	Source() *Keyboard
	KeyCode() KeyCode
	Display() *Display
}

type key_up_event C.struct_ALLEGRO_KEYBOARD_EVENT

func (e *key_up_event) key_up() {}

func (e *key_up_event) Source() *Keyboard {
	return (*Keyboard)(e.source)
}

func (e *key_up_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *key_up_event) KeyCode() KeyCode {
	return KeyCode(e.keycode)
}

func (e *key_up_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Key Char -- */

type KeyCharEvent interface {
	key_char()
	Timestamp() float64
	Source() *Keyboard
	KeyCode() KeyCode
	Unichar() int
	Modifiers() KeyModifier
	Repeat() bool
	Display() *Display
}

type key_char_event C.struct_ALLEGRO_KEYBOARD_EVENT

func (e *key_char_event) key_char() {}

func (e *key_char_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *key_char_event) Source() *Keyboard {
	return (*Keyboard)(e.source)
}

func (e *key_char_event) KeyCode() KeyCode {
	return KeyCode(e.keycode)
}

func (e *key_char_event) Unichar() int {
	return int(e.unichar)
}

func (e *key_char_event) Modifiers() KeyModifier {
	return KeyModifier(e.modifiers)
}

func (e *key_char_event) Repeat() bool {
	return bool(e.repeat)
}

func (e *key_char_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Axes -- */

type MouseAxesEvent interface {
	mouse_axes()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Dx() int
	Dy() int
	Dz() int
	Dw() int
	Display() *Display
}

type mouse_axes_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_axes_event) mouse_axes() {}

func (e *mouse_axes_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_axes_event) X() int {
	return int(e.x)
}

func (e *mouse_axes_event) Y() int {
	return int(e.y)
}

func (e *mouse_axes_event) Z() int {
	return int(e.z)
}

func (e *mouse_axes_event) W() int {
	return int(e.w)
}

func (e *mouse_axes_event) Dx() int {
	return int(e.dx)
}

func (e *mouse_axes_event) Dy() int {
	return int(e.dy)
}

func (e *mouse_axes_event) Dz() int {
	return int(e.dz)
}

func (e *mouse_axes_event) Dw() int {
	return int(e.dw)
}

func (e *mouse_axes_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Button Down -- */

type MouseButtonDownEvent interface {
	mouse_button_down()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Button() uint
	Display() *Display
}

type mouse_button_down_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_button_down_event) mouse_button_down() {}

func (e *mouse_button_down_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_button_down_event) X() int {
	return int(e.x)
}

func (e *mouse_button_down_event) Y() int {
	return int(e.y)
}

func (e *mouse_button_down_event) Z() int {
	return int(e.z)
}

func (e *mouse_button_down_event) W() int {
	return int(e.w)
}

func (e *mouse_button_down_event) Button() uint {
	return uint(e.button)
}

func (e *mouse_button_down_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Button Up -- */

type MouseButtonUpEvent interface {
	mouse_button_up()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Button() uint
	Display() *Display
}

type mouse_button_up_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_button_up_event) mouse_button_up() {}

func (e *mouse_button_up_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_button_up_event) X() int {
	return int(e.x)
}

func (e *mouse_button_up_event) Y() int {
	return int(e.y)
}

func (e *mouse_button_up_event) Z() int {
	return int(e.z)
}

func (e *mouse_button_up_event) W() int {
	return int(e.w)
}

func (e *mouse_button_up_event) Button() uint {
	return uint(e.button)
}

func (e *mouse_button_up_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Warped -- */

type MouseWarpedEvent interface {
	mouse_warped()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Dx() int
	Dy() int
	Dz() int
	Dw() int
	Display() *Display
}

type mouse_warped_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_warped_event) mouse_warped() {}

func (e *mouse_warped_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_warped_event) X() int {
	return int(e.x)
}

func (e *mouse_warped_event) Y() int {
	return int(e.y)
}

func (e *mouse_warped_event) Z() int {
	return int(e.z)
}

func (e *mouse_warped_event) W() int {
	return int(e.w)
}

func (e *mouse_warped_event) Dx() int {
	return int(e.dx)
}

func (e *mouse_warped_event) Dy() int {
	return int(e.dy)
}

func (e *mouse_warped_event) Dz() int {
	return int(e.dz)
}

func (e *mouse_warped_event) Dw() int {
	return int(e.dw)
}

func (e *mouse_warped_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Enter Display -- */

type MouseEnterDisplayEvent interface {
	mouse_enter_display()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Display() *Display
}

type mouse_enter_display_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_enter_display_event) mouse_enter_display() {}

func (e *mouse_enter_display_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_enter_display_event) X() int {
	return int(e.x)
}

func (e *mouse_enter_display_event) Y() int {
	return int(e.y)
}

func (e *mouse_enter_display_event) Z() int {
	return int(e.z)
}

func (e *mouse_enter_display_event) W() int {
	return int(e.w)
}

func (e *mouse_enter_display_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Mouse Leave Display -- */

type MouseLeaveDisplayEvent interface {
	mouse_leave_display()
	Timestamp() float64
	X() int
	Y() int
	Z() int
	W() int
	Display() *Display
}

type mouse_leave_display_event C.struct_ALLEGRO_MOUSE_EVENT

func (e *mouse_leave_display_event) mouse_leave_display() {}

func (e *mouse_leave_display_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *mouse_leave_display_event) X() int {
	return int(e.x)
}

func (e *mouse_leave_display_event) Y() int {
	return int(e.y)
}

func (e *mouse_leave_display_event) Z() int {
	return int(e.z)
}

func (e *mouse_leave_display_event) W() int {
	return int(e.w)
}

func (e *mouse_leave_display_event) Display() *Display {
	return (*Display)(e.display)
}

/* -- Timer -- */

type TimerEvent interface {
	timer()
	Timestamp() float64
	Source() *Timer
	Count() int64
}

type timer_event C.struct_ALLEGRO_TIMER_EVENT

func (e *timer_event) timer() {}

func (e *timer_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *timer_event) Source() *Timer {
	return (*Timer)(e.source)
}

func (e *timer_event) Count() int64 {
	return int64(e.count)
}

/* -- Display Expose -- */

type DisplayExposeEvent interface {
	display_expose()
	Timestamp() float64
	Source() *Display
	X() int
	Y() int
	Width() int
	Height() int
}

type display_expose_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_expose_event) display_expose() {}

func (e *display_expose_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_expose_event) Source() *Display {
	return (*Display)(e.source)
}

func (e *display_expose_event) X() int {
	return int(e.x)
}

func (e *display_expose_event) Y() int {
	return int(e.y)
}

func (e *display_expose_event) Width() int {
	return int(e.width)
}

func (e *display_expose_event) Height() int {
	return int(e.height)
}

/* -- Display Resize -- */

type DisplayResizeEvent interface {
	display_resize()
	Timestamp() float64
	Source() *Display
	X() int
	Y() int
	Width() int
	Height() int
}

type display_resize_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_resize_event) display_resize() {}

func (e *display_resize_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_resize_event) Source() *Display {
	return (*Display)(e.source)
}

func (e *display_resize_event) X() int {
	return int(e.x)
}

func (e *display_resize_event) Y() int {
	return int(e.y)
}

func (e *display_resize_event) Width() int {
	return int(e.width)
}

func (e *display_resize_event) Height() int {
	return int(e.height)
}

/* -- Display Close -- */

type DisplayCloseEvent interface {
	display_close()
	Timestamp() float64
	Source() *Display
}

type display_close_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_close_event) display_close() {}

func (e *display_close_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_close_event) Source() *Display {
	return (*Display)(e.source)
}

/* -- Display Lost -- */

type DisplayLostEvent interface {
	display_lost()
	Timestamp() float64
	Source() *Display
}

type display_lost_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_lost_event) display_lost() {}

func (e *display_lost_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_lost_event) Source() *Display {
	return (*Display)(e.source)
}

/* -- Display Found -- */

type DisplayFoundEvent interface {
	display_found()
	Timestamp() float64
	Source() *Display
}

type display_found_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_found_event) display_found() {}

func (e *display_found_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_found_event) Source() *Display {
	return (*Display)(e.source)
}

/* -- Display Switch Out -- */

type DisplaySwitchOutEvent interface {
	display_switch_out()
	Timestamp() float64
	Source() *Display
}

type display_switch_out_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_switch_out_event) display_switch_out() {}

func (e *display_switch_out_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_switch_out_event) Source() *Display {
	return (*Display)(e.source)
}

/* -- Display Switch In -- */

type DisplaySwitchInEvent interface {
	display_switch_in()
	Timestamp() float64
	Source() *Display
}

type display_switch_in_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_switch_in_event) display_switch_in() {}

func (e *display_switch_in_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_switch_in_event) Source() *Display {
	return (*Display)(e.source)
}

/* -- Display Orientation -- */

type DisplayOrientationEvent interface {
	display_orientation()
	Timestamp() float64
	Source() *Display
	Orientation() DisplayOrientation
}

type display_orientation_event C.struct_ALLEGRO_DISPLAY_EVENT

func (e *display_orientation_event) display_orientation() {}

func (e *display_orientation_event) Timestamp() float64 {
	return float64(e.timestamp)
}

func (e *display_orientation_event) Source() *Display {
	return (*Display)(e.source)
}

func (e *display_orientation_event) Orientation() DisplayOrientation {
	return DisplayOrientation(e.orientation)
}

/* -- Audio Stream Fragment -- */

type AudioStreamFragment interface {
	audio_stream_fragment()
}

type audio_stream_fragment_event struct{}

func (e *audio_stream_fragment_event) audio_stream_fragment() {}

/* -- Audio Stream Finished -- */

type AudioStreamFinished interface {
	audio_stream_finished()
}

type audio_stream_finished_event struct{}

func (e *audio_stream_finished_event) audio_stream_finished() {}

/* -- User -- */

type UserEvent interface {
	user()
	Source() *EventSource
	Data1() uintptr
	Data2() uintptr
	Data3() uintptr
	Data4() uintptr
	Unref()
}

type user_event C.struct_ALLEGRO_USER_EVENT

func (e *user_event) user() {}

func (e *user_event) Source() *EventSource {
	return (*EventSource)(e.source)
}

func (e *user_event) Data1() uintptr {
	return uintptr(e.data1)
}

func (e *user_event) Data2() uintptr {
	return uintptr(e.data2)
}

func (e *user_event) Data3() uintptr {
	return uintptr(e.data3)
}

func (e *user_event) Data4() uintptr {
	return uintptr(e.data4)
}

// Decrease the reference count of a user-defined event. This must be called on
// any user event that you get from al_get_next_event, al_peek_next_event,
// al_wait_for_event, etc. which is reference counted. This function does
// nothing if the event is not reference counted.
func (e *user_event) Unref() {
	C.al_unref_user_event((*C.ALLEGRO_USER_EVENT)(e))
}
