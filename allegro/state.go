package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"fmt"
)

type State C.ALLEGRO_STATE

type StateFlags int

const (
	STATE_NEW_DISPLAY_PARAMETERS StateFlags = C.ALLEGRO_STATE_NEW_DISPLAY_PARAMETERS
	STATE_NEW_BITMAP_PARAMETERS             = C.ALLEGRO_STATE_NEW_BITMAP_PARAMETERS
	STATE_DISPLAY                           = C.ALLEGRO_STATE_DISPLAY
	STATE_TARGET_BITMAP                     = C.ALLEGRO_STATE_TARGET_BITMAP
	STATE_BLENDER                           = C.ALLEGRO_STATE_BLENDER
	STATE_TRANSFORM                         = C.ALLEGRO_STATE_TRANSFORM
	STATE_NEW_FILE_INTERFACE                = C.ALLEGRO_STATE_NEW_FILE_INTERFACE
	STATE_BITMAP                            = C.ALLEGRO_STATE_BITMAP
	STATE_ALL                               = C.ALLEGRO_STATE_ALL
)

type Error struct {
	Errno int
}

func (e *Error) Error() string {
	return fmt.Sprintf("errno = %d", e.Errno)
}

// Stores part of the state of the current thread in the given ALLEGRO_STATE
// objects. The flags parameter can take any bit-combination of these flags:
func StoreState(flags StateFlags) *State {
	var state C.ALLEGRO_STATE
	C.al_store_state(&state, C.int(flags))
	return (*State)(&state)
}

// Restores part of the state of the current thread from the given
// ALLEGRO_STATE object.
func RestoreState(state *State) {
	C.al_restore_state((*C.ALLEGRO_STATE)(state))
}

// Some Allegro functions will set an error number as well as returning an
// error code. Call this function to retrieve the last error number set for the
// calling thread.
func LastError() error {
	return &Error{int(C.al_get_errno())}
}

// Set the error number for for the calling thread.
func SetLastError(e *Error) {
	C.al_set_errno(C.int(e.Errno))
}
