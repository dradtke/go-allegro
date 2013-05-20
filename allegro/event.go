package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
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

//}}}
*/
import "C"
import (
	"errors"
)

type Event struct {
	Type EventType
	Source *EventSource
	Timestamp float64
	// TODO: add the other event types
	Joystick JoystickEventInfo
	Keyboard KeyboardEventInfo
	Mouse MouseEventInfo
	Timer TimerEventInfo
	Display DisplayEventInfo
}

type JoystickEventInfo struct {
	Id *Joystick
	Stick, Axis, Button int
	Pos float32
}

type KeyboardEventInfo struct {
	KeyCode KeyCode
	Display *Display
	Unichar int
	Modifiers KeyModifier
	Repeat bool
}

type MouseEventInfo struct {
	X, Y, Z, W, Dx, Dy, Dz, Dw int
	Button uint
	Display *Display
}

type TimerEventInfo struct {
	Source *Timer
	Count int64
}

type DisplayEventInfo struct {
	Source *Display
	X, Y, Width, Height int
	Orientation DisplayOrientation
}

type EventType int
const (
	JoystickAxisEvent           EventType = C.ALLEGRO_EVENT_JOYSTICK_AXIS
	JoystickButtonDownEvent     EventType = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
	JoystickButtonUpEvent       EventType = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
	JoystickConfigurationEvent  EventType = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION
	KeyDownEvent                EventType = C.ALLEGRO_EVENT_KEY_DOWN
	KeyUpEvent                  EventType = C.ALLEGRO_EVENT_KEY_UP
	KeyCharEvent                EventType = C.ALLEGRO_EVENT_KEY_CHAR
	MouseAxesEvent              EventType = C.ALLEGRO_EVENT_MOUSE_AXES
	MouseButtonDownEvent        EventType = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
	MouseButtonUpEvent          EventType = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
	MouseWarpedEvent            EventType = C.ALLEGRO_EVENT_MOUSE_WARPED
	MouseEnterDisplayEvent      EventType = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
	MouseLeaveDisplayEvent      EventType = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
	TimerEvent                  EventType = C.ALLEGRO_EVENT_TIMER
	DisplayExposeEvent          EventType = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
	DisplayResizeEvent          EventType = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	DisplayCloseEvent           EventType = C.ALLEGRO_EVENT_DISPLAY_CLOSE
	DisplayLostEvent            EventType = C.ALLEGRO_EVENT_DISPLAY_LOST
	DisplayFoundEvent           EventType = C.ALLEGRO_EVENT_DISPLAY_FOUND
	DisplaySwitchOutEvent       EventType = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
	DisplaySwitchInEvent        EventType = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
	DisplayOrientationEvent     EventType = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
)

type EventSource C.ALLEGRO_EVENT_SOURCE

type EventQueue struct {
	ptr *C.ALLEGRO_EVENT_QUEUE
	event C.ALLEGRO_EVENT
}

func CreateEventQueue() (*EventQueue, error) {
	queue := C.al_create_event_queue()
	if queue == nil {
		return nil, errors.New("failed to create event queue!")
	}
	return &EventQueue{ptr:queue}, nil
}

func (queue *EventQueue) Destroy() {
	C.al_destroy_event_queue(queue.ptr)
}

func (queue *EventQueue) RegisterEventSource(source *EventSource) {
	C.al_register_event_source(queue.ptr, (*C.ALLEGRO_EVENT_SOURCE)(source))
}

func (queue *EventQueue) GetNextEvent() (*Event, bool) {
	success := gobool(C.al_get_next_event(queue.ptr, &queue.event))
	if !success {
		return nil, false
	}
	return queue.newEvent(), true
}

func (queue *EventQueue) WaitForEvent() *Event {
	C.al_wait_for_event(queue.ptr, &queue.event)
	return queue.newEvent()
}

// wait for an event, but don't take it off the queue
// better name for this?
func (queue *EventQueue) JustWaitForEvent() {
	C.al_wait_for_event(queue.ptr, nil)
}

func (queue *EventQueue) WaitForEventUntil(timeout *Timeout) (*Event, bool) {
	success := C.al_wait_for_event_until(queue.ptr, &queue.event, (*C.ALLEGRO_TIMEOUT)(timeout))
	if !success {
		return nil, false
	}
	return queue.newEvent(), true
}

// like WaitForEventUntil, but don't return an event and leave everything on the queue
func (queue *EventQueue) JustWaitForEventUntil(timeout *Timeout) bool {
	return gobool(C.al_wait_for_event_until(queue.ptr, nil, (*C.ALLEGRO_TIMEOUT)(timeout)))
}

func (queue *EventQueue) newEvent() *Event {
	ev := Event{
		Type: (EventType)(C.get_event_type(&queue.event)),
		Source: (*EventSource)(C.get_event_source(&queue.event)),
		Timestamp: godouble(C.get_event_timestamp(&queue.event)),
	}
	switch ev.Type {
		case JoystickAxisEvent:
			id := (*Joystick)(C.get_event_joystick_id(&queue.event))
			stick := int(C.get_event_joystick_stick(&queue.event))
			axis := int(C.get_event_joystick_axis(&queue.event))
			pos := float32(C.get_event_joystick_pos(&queue.event))
			ev.Joystick = JoystickEventInfo{Id:id, Stick:stick, Axis:axis, Pos:pos}

		case JoystickButtonDownEvent, JoystickButtonUpEvent:
			id := (*Joystick)(C.get_event_joystick_id(&queue.event))
			button := int(C.get_event_joystick_button(&queue.event))
			ev.Joystick = JoystickEventInfo{Id:id, Button:button}

		case JoystickConfigurationEvent:
			ev.Joystick = JoystickEventInfo{}

		case KeyDownEvent, KeyUpEvent:
			keycode := (KeyCode)(C.get_event_keyboard_keycode(&queue.event))
			display := (*Display)(C.get_event_keyboard_display(&queue.event))
			ev.Keyboard = KeyboardEventInfo{KeyCode:keycode, Display:display}

		case KeyCharEvent:
			keycode := (KeyCode)(C.get_event_keyboard_keycode(&queue.event))
			unichar := int(C.get_event_keyboard_unichar(&queue.event))
			modifiers := KeyModifier(C.get_event_keyboard_modifiers(&queue.event))
			repeat := bool(C.get_event_keyboard_repeat(&queue.event))
			display := (*Display)(C.get_event_keyboard_display(&queue.event))
			ev.Keyboard = KeyboardEventInfo{KeyCode:keycode, Unichar:unichar, Modifiers:modifiers, Repeat:repeat, Display:display}

		case MouseAxesEvent:
			x := int(C.get_event_mouse_x(&queue.event))
			y := int(C.get_event_mouse_y(&queue.event))
			z := int(C.get_event_mouse_z(&queue.event))
			w := int(C.get_event_mouse_w(&queue.event))
			dx := int(C.get_event_mouse_dx(&queue.event))
			dy := int(C.get_event_mouse_dy(&queue.event))
			dz := int(C.get_event_mouse_dz(&queue.event))
			dw := int(C.get_event_mouse_dw(&queue.event))
			display := (*Display)(C.get_event_mouse_display(&queue.event))
			ev.Mouse = MouseEventInfo{X:x, Y:y, Z:z, W:w, Dx:dx, Dy:dy, Dz:dz, Dw:dw, Display:display}

		case MouseButtonDownEvent, MouseButtonUpEvent:
			x := int(C.get_event_mouse_x(&queue.event))
			y := int(C.get_event_mouse_y(&queue.event))
			z := int(C.get_event_mouse_z(&queue.event))
			w := int(C.get_event_mouse_w(&queue.event))
			button := uint(C.get_event_mouse_button(&queue.event))
			display := (*Display)(C.get_event_mouse_display(&queue.event))
			ev.Mouse = MouseEventInfo{X:x, Y:y, Z:z, W:w, Button:button, Display:display}

		case MouseWarpedEvent:
			ev.Mouse = MouseEventInfo{}

		case MouseEnterDisplayEvent, MouseLeaveDisplayEvent:
			x := int(C.get_event_mouse_x(&queue.event))
			y := int(C.get_event_mouse_y(&queue.event))
			z := int(C.get_event_mouse_z(&queue.event))
			w := int(C.get_event_mouse_w(&queue.event))
			display := (*Display)(C.get_event_mouse_display(&queue.event))
			ev.Mouse = MouseEventInfo{X:x, Y:y, Z:z, W:w, Display:display}

		case TimerEvent:
			source := &Timer{ptr:C.get_event_timer_source(&queue.event)}
			count := int64(C.get_event_timer_count(&queue.event))
			ev.Timer = TimerEventInfo{Source:source, Count:count}

		case DisplayExposeEvent, DisplayResizeEvent:
			source := (*Display)(C.get_event_display_source(&queue.event))
			x := int(C.get_event_display_x(&queue.event))
			y := int(C.get_event_display_y(&queue.event))
			width := int(C.get_event_display_width(&queue.event))
			height := int(C.get_event_display_height(&queue.event))
			ev.Display = DisplayEventInfo{Source:source, X:x, Y:y, Width:width, Height:height}

		case DisplayCloseEvent, DisplayLostEvent, DisplayFoundEvent, DisplaySwitchOutEvent, DisplaySwitchInEvent:
			source := (*Display)(C.get_event_display_source(&queue.event))
			ev.Display = DisplayEventInfo{Source:source}

		case DisplayOrientationEvent:
			source := (*Display)(C.get_event_display_source(&queue.event))
			orientation := (DisplayOrientation)(C.get_event_display_orientation(&queue.event))
			ev.Display = DisplayEventInfo{Source:source, Orientation:orientation}
	}
	return &ev
}
