package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
)

type MouseCursor C.ALLEGRO_MOUSE_CURSOR

type MouseState C.struct_ALLEGRO_MOUSE_STATE

type SystemMouseCursor C.ALLEGRO_SYSTEM_MOUSE_CURSOR

const (
	MouseCursorDefault     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT
	MouseCursorArrow                         = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW
	MouseCursorBusy                          = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY
	MouseCursorQuestion                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION
	MouseCursorEdit                          = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT
	MouseCursorMove                          = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE
	MouseCursorResizeN                       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N
	MouseCursorResizeW                       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W
	MouseCursorResizeS                       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S
	MouseCursorResizeE                       = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E
	MouseCursorResizeNw                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW
	MouseCursorResizeSw                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW
	MouseCursorResizeSe                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE
	MouseCursorResizeNe                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE
	MouseCursorProgress                      = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS
	MouseCursorPrecision                     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION
	MouseCursorLink                          = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK
	MouseCursorAltSelect                     = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT
	MouseCursorUnavailable                   = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE
)

// Install a mouse driver.
func InstallMouse() error {
	success := bool(C.al_install_mouse())
	if !success {
		return errors.New("failed to install mouse!")
	}
	return nil
}

// Returns true if al_install_mouse was called successfully.
func IsMouseInstalled() bool {
	return bool(C.al_is_mouse_installed())
}

// Uninstalls the active mouse driver, if any. This will automatically
// unregister the mouse event source with any event queues.
func UninstallMouse() {
	C.al_uninstall_mouse()
}

// Return the number of buttons on the mouse. The first axis is 0.
func MouseNumAxes() uint {
	return uint(C.al_get_mouse_num_axes())
}

// Return the number of buttons on the mouse. The first button is 1.
func MouseNumButtons() uint {
	return uint(C.al_get_mouse_num_buttons())
}

// Save the state of the mouse specified at the time the function is called
// into the given structure.
func (state *MouseState) Get() {
	C.al_get_mouse_state((*C.struct_ALLEGRO_MOUSE_STATE)(state))
}

func (state *MouseState) X() int {
	return int(state.x)
}

func (state *MouseState) Y() int {
	return int(state.y)
}

func (state *MouseState) W() int {
	return int(state.w)
}

func (state *MouseState) Z() int {
	return int(state.z)
}

func (state *MouseState) Buttons() int {
	return int(state.buttons)
}

func (state *MouseState) Pressure() float32 {
	return float32(state.pressure)
}

// Extract the mouse axis value from the saved state. The axes are numbered
// from 0, in this order: x-axis, y-axis, z-axis, w-axis.
func (state *MouseState) Axis(axis int) int {
	return int(C.al_get_mouse_state_axis((*C.struct_ALLEGRO_MOUSE_STATE)(state), C.int(axis)))
}

// Return true if the mouse button specified was held down in the state
// specified. Unlike most things, the first mouse button is numbered 1.
func (state *MouseState) ButtonDown(button int) bool {
	return bool(C.al_mouse_button_down((*C.struct_ALLEGRO_MOUSE_STATE)(state), C.int(button)))
}

// Try to position the mouse at the given coordinates on the given display. The
// mouse movement resulting from a successful move will generate an
// ALLEGRO_EVENT_MOUSE_WARPED event.
func (d *Display) SetMouseXY(x, y int) error {
	success := C.al_set_mouse_xy((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y))
	if !success {
		return errors.New("failed to set new mouse position!")
	}
	return nil
}

// Set the given mouse axis to the given value.
func SetMouseAxis(which, value int) error {
	success := C.al_set_mouse_axis(C.int(which), C.int(value))
	if !success {
		return errors.New("failed to set mouse axis!")
	}
	return nil
}

// Retrieve the mouse event source.
func MouseEventSource() (*EventSource, error) {
	source := C.al_get_mouse_event_source()
	if source == nil {
		return nil, errors.New("failed to get mouse event source; did you call InstallMouse() first?")
	}
	return (*EventSource)(source), nil
}

// Create a mouse cursor from the bitmap provided.
func CreateMouseCursor(bmp *Bitmap, x_focus, y_focus int) (*MouseCursor, error) {
	c := C.al_create_mouse_cursor((*C.ALLEGRO_BITMAP)(bmp), C.int(x_focus), C.int(y_focus))
	if c == nil {
		return nil, errors.New("failed to create mouse cursor!")
	}
	cursor := (*MouseCursor)(c)
	//runtime.SetFinalizer(cursor, cursor.Destroy)
	return cursor, nil
}

// Free the memory used by the given cursor.
func (cursor *MouseCursor) Destroy() {
	C.al_destroy_mouse_cursor((*C.ALLEGRO_MOUSE_CURSOR)(cursor))
}

// Set the given mouse cursor to be the current mouse cursor for the given
// display.
func (d *Display) SetMouseCursor(cursor *MouseCursor) error {
	success := C.al_set_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_MOUSE_CURSOR)(cursor))
	if !success {
		return errors.New("failed to set display mouse cursor!")
	}
	return nil
}

// Set the given system mouse cursor to be the current mouse cursor for the
// given display. If the cursor is currently 'shown' (as opposed to 'hidden')
// the change is immediately visible.
func (d *Display) SetSystemMouseCursor(cursor SystemMouseCursor) error {
	success := C.al_set_system_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), (C.ALLEGRO_SYSTEM_MOUSE_CURSOR)(cursor))
	if !success {
		return errors.New("failed to set display system mouse cursor!")
	}
	return nil
}

// On platforms where this information is available, this function returns the
// global location of the mouse cursor, relative to the desktop. You should not
// normally use this function, as the information is not useful except for
// special scenarios as moving a window.
func MouseCursorPosition() (int, int, error) {
	var x, y C.int
	success := bool(C.al_get_mouse_cursor_position(&x, &y))
	if !success {
		return 0, 0, errors.New("failed to get mouse cursor position!")
	}
	return int(x), int(y), nil
}

// Hide the mouse cursor in the given display. This has no effect on what the
// current mouse cursor looks like; it just makes it disappear.
func (d *Display) HideMouseCursor() error {
	success := bool(C.al_hide_mouse_cursor((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to hide mouse cursor!")
	}
	return nil
}

// Make a mouse cursor visible in the given display.
func (d *Display) ShowMouseCursor() error {
	success := bool(C.al_show_mouse_cursor((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to show mouse cursor!")
	}
	return nil
}

// Confine the mouse cursor to the given display. The mouse cursor can only be
// confined to one display at a time.
func (d *Display) GrabMouse() error {
	success := bool(C.al_grab_mouse((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to grab mouse!")
	}
	return nil
}

// Stop confining the mouse cursor to any display belonging to the program.
func UngrabMouse() error {
	success := bool(C.al_ungrab_mouse())
	if !success {
		return errors.New("failed to ungrab mouse!")
	}
	return nil
}
