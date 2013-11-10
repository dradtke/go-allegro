package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"

type State C.ALLEGRO_STATE

type StateFlags int

const (
	STATE_NEW_DISPLAY_PARAMETERS StateFlags = C.ALLEGRO_STATE_NEW_DISPLAY_PARAMETERS
	STATE_NEW_BITMAP_PARAMETERS  StateFlags = C.ALLEGRO_STATE_NEW_BITMAP_PARAMETERS
	STATE_DISPLAY                StateFlags = C.ALLEGRO_STATE_DISPLAY
	STATE_TARGET_BITMAP          StateFlags = C.ALLEGRO_STATE_TARGET_BITMAP
	STATE_BLENDER                StateFlags = C.ALLEGRO_STATE_BLENDER
	STATE_TRANSFORM              StateFlags = C.ALLEGRO_STATE_TRANSFORM
	STATE_NEW_FILE_INTERFACE     StateFlags = C.ALLEGRO_STATE_NEW_FILE_INTERFACE
	STATE_BITMAP                 StateFlags = C.ALLEGRO_STATE_BITMAP
	STATE_ALL                    StateFlags = C.ALLEGRO_STATE_ALL
)

func StoreState(flags StateFlags) *State {
	var state C.ALLEGRO_STATE
	C.al_store_state(&state, C.int(flags))
	return (*State)(&state)
}

func RestoreState(state *State) {
	C.al_restore_state((*C.ALLEGRO_STATE)(state))
}
