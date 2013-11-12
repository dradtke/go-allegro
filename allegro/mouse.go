package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
	"runtime"
)

type MouseCursor C.ALLEGRO_MOUSE_CURSOR

type MouseState struct {
	X, Y, W, Z, Buttons int
	ptr C.ALLEGRO_MOUSE_STATE
}

type SystemMouseCursor C.ALLEGRO_SYSTEM_MOUSE_CURSOR
const (
	MouseCursorDefault      SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT
	MouseCursorArrow        SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW
	MouseCursorBusy         SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY
	MouseCursorQuestion     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION
	MouseCursorEdit         SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT
	MouseCursorMove         SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE
	MouseCursorResizeN      SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N
	MouseCursorResizeW      SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W
	MouseCursorResizeS      SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S
	MouseCursorResizeE      SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E
	MouseCursorResizeNw     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW
	MouseCursorResizeSw     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW
	MouseCursorResizeSe     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE
	MouseCursorResizeNe     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE
	MouseCursorProgress     SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS
	MouseCursorPrecision    SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION
	MouseCursorLink         SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK
	MouseCursorAltSelect    SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT
	MouseCursorUnavailable  SystemMouseCursor = C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE
)

func InstallMouse() error {
	success := bool(C.al_install_mouse())
	if !success {
		return errors.New("failed to install mouse!")
	}
	return nil
}

func IsMouseInstalled() bool {
	return bool(C.al_is_mouse_installed())
}

func UninstallMouse() {
	C.al_uninstall_mouse()
}

func MouseNumAxes() uint {
	return uint(C.al_get_mouse_num_axes())
}

func MouseNumButtons() uint {
	return uint(C.al_get_mouse_num_buttons())
}

func (state *MouseState) Get() {
	C.al_get_mouse_state(&state.ptr)
	state.X = int(state.ptr.x)
	state.Y = int(state.ptr.y)
	state.Z = int(state.ptr.z)
	state.W = int(state.ptr.w)
	state.Buttons = int(state.ptr.buttons)
}

func (state *MouseState) Axis(axis int) int {
	return int(C.al_get_mouse_state_axis(&state.ptr, C.int(axis)))
}

func (state *MouseState) ButtonDown(button int) bool {
	return bool(C.al_mouse_button_down(&state.ptr, C.int(button)))
}

func (d *Display) SetMouseXY(x, y int) error {
	success := C.al_set_mouse_xy((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y))
	if !success {
		return errors.New("failed to set new mouse position!")
	}
	return nil
}

func SetMouseAxis(which, value int) error {
	success := C.al_set_mouse_axis(C.int(which), C.int(value))
	if !success {
		return errors.New("failed to set mouse axis!")
	}
	return nil
}

func MouseEventSource() (*EventSource, error) {
	source := C.al_get_mouse_event_source()
	if source == nil {
		return nil, errors.New("failed to get mouse event source; did you call InstallMouse() first?")
	}
	return (*EventSource)(source), nil
}

func CreateMouseCursor(bmp *Bitmap, x_focus, y_focus int) (*MouseCursor, error) {
	c := C.al_create_mouse_cursor((*C.ALLEGRO_BITMAP)(bmp), C.int(x_focus), C.int(y_focus))
	if c == nil {
		return nil, errors.New("failed to create mouse cursor!")
	}
	cursor := (*MouseCursor)(c)
	runtime.SetFinalizer(cursor, cursor.Destroy)
	return cursor, nil
}

func (cursor *MouseCursor) Destroy() {
	C.al_destroy_mouse_cursor((*C.ALLEGRO_MOUSE_CURSOR)(cursor))
}

func (d *Display) SetMouseCursor(cursor *MouseCursor) error {
	success := C.al_set_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_MOUSE_CURSOR)(cursor))
	if !success {
		return errors.New("failed to set display mouse cursor!")
	}
	return nil
}

func (d *Display) SetSystemMouseCursor(cursor SystemMouseCursor) error {
	success := C.al_set_system_mouse_cursor((*C.ALLEGRO_DISPLAY)(d), (C.ALLEGRO_SYSTEM_MOUSE_CURSOR)(cursor))
	if !success {
		return errors.New("failed to set display system mouse cursor!")
	}
	return nil
}

func MouseCursorPosition() (int, int, error) {
	var x, y C.int
	success := bool(C.al_get_mouse_cursor_position(&x, &y))
	if !success {
		return 0, 0, errors.New("failed to get mouse cursor position!")
	}
	return int(x), int(y), nil
}

func (d *Display) HideMouseCursor() error {
	success := bool(C.al_hide_mouse_cursor((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to hide mouse cursor!")
	}
	return nil
}

func (d *Display) ShowMouseCursor() error {
	success := bool(C.al_show_mouse_cursor((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to show mouse cursor!")
	}
	return nil
}

func (d *Display) GrabMouse() error {
	success := bool(C.al_grab_mouse((*C.ALLEGRO_DISPLAY)(d)))
	if !success {
		return errors.New("failed to grab mouse!")
	}
	return nil
}

func UngrabMouse() error {
	success := bool(C.al_ungrab_mouse())
	if !success {
		return errors.New("failed to ungrab mouse!")
	}
	return nil
}
