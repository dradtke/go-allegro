package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
)

type Keyboard C.ALLEGRO_KEYBOARD

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

//{{{ key codes
type KeyCode int

const (
	KEY_A            KeyCode = C.ALLEGRO_KEY_A
	KEY_B                    = C.ALLEGRO_KEY_B
	KEY_C                    = C.ALLEGRO_KEY_C
	KEY_D                    = C.ALLEGRO_KEY_D
	KEY_E                    = C.ALLEGRO_KEY_E
	KEY_F                    = C.ALLEGRO_KEY_F
	KEY_G                    = C.ALLEGRO_KEY_G
	KEY_H                    = C.ALLEGRO_KEY_H
	KEY_I                    = C.ALLEGRO_KEY_I
	KEY_J                    = C.ALLEGRO_KEY_J
	KEY_K                    = C.ALLEGRO_KEY_K
	KEY_L                    = C.ALLEGRO_KEY_L
	KEY_M                    = C.ALLEGRO_KEY_M
	KEY_N                    = C.ALLEGRO_KEY_N
	KEY_O                    = C.ALLEGRO_KEY_O
	KEY_P                    = C.ALLEGRO_KEY_P
	KEY_Q                    = C.ALLEGRO_KEY_Q
	KEY_R                    = C.ALLEGRO_KEY_R
	KEY_S                    = C.ALLEGRO_KEY_S
	KEY_T                    = C.ALLEGRO_KEY_T
	KEY_U                    = C.ALLEGRO_KEY_U
	KEY_V                    = C.ALLEGRO_KEY_V
	KEY_W                    = C.ALLEGRO_KEY_W
	KEY_X                    = C.ALLEGRO_KEY_X
	KEY_Y                    = C.ALLEGRO_KEY_Y
	KEY_Z                    = C.ALLEGRO_KEY_Z
	KEY_0                    = C.ALLEGRO_KEY_0
	KEY_1                    = C.ALLEGRO_KEY_1
	KEY_2                    = C.ALLEGRO_KEY_2
	KEY_3                    = C.ALLEGRO_KEY_3
	KEY_4                    = C.ALLEGRO_KEY_4
	KEY_5                    = C.ALLEGRO_KEY_5
	KEY_6                    = C.ALLEGRO_KEY_6
	KEY_7                    = C.ALLEGRO_KEY_7
	KEY_8                    = C.ALLEGRO_KEY_8
	KEY_9                    = C.ALLEGRO_KEY_9
	KEY_PAD_0                = C.ALLEGRO_KEY_PAD_0
	KEY_PAD_1                = C.ALLEGRO_KEY_PAD_1
	KEY_PAD_2                = C.ALLEGRO_KEY_PAD_2
	KEY_PAD_3                = C.ALLEGRO_KEY_PAD_3
	KEY_PAD_4                = C.ALLEGRO_KEY_PAD_4
	KEY_PAD_5                = C.ALLEGRO_KEY_PAD_5
	KEY_PAD_6                = C.ALLEGRO_KEY_PAD_6
	KEY_PAD_7                = C.ALLEGRO_KEY_PAD_7
	KEY_PAD_8                = C.ALLEGRO_KEY_PAD_8
	KEY_PAD_9                = C.ALLEGRO_KEY_PAD_9
	KEY_F1                   = C.ALLEGRO_KEY_F1
	KEY_F2                   = C.ALLEGRO_KEY_F2
	KEY_F3                   = C.ALLEGRO_KEY_F3
	KEY_F4                   = C.ALLEGRO_KEY_F4
	KEY_F5                   = C.ALLEGRO_KEY_F5
	KEY_F6                   = C.ALLEGRO_KEY_F6
	KEY_F7                   = C.ALLEGRO_KEY_F7
	KEY_F8                   = C.ALLEGRO_KEY_F8
	KEY_F9                   = C.ALLEGRO_KEY_F9
	KEY_F10                  = C.ALLEGRO_KEY_F10
	KEY_F11                  = C.ALLEGRO_KEY_F11
	KEY_F12                  = C.ALLEGRO_KEY_F12
	KEY_ESCAPE               = C.ALLEGRO_KEY_ESCAPE
	KEY_TILDE                = C.ALLEGRO_KEY_TILDE
	KEY_MINUS                = C.ALLEGRO_KEY_MINUS
	KEY_EQUALS               = C.ALLEGRO_KEY_EQUALS
	KEY_BACKSPACE            = C.ALLEGRO_KEY_BACKSPACE
	KEY_TAB                  = C.ALLEGRO_KEY_TAB
	KEY_OPENBRACE            = C.ALLEGRO_KEY_OPENBRACE
	KEY_CLOSEBRACE           = C.ALLEGRO_KEY_CLOSEBRACE
	KEY_ENTER                = C.ALLEGRO_KEY_ENTER
	KEY_SEMICOLON            = C.ALLEGRO_KEY_SEMICOLON
	KEY_QUOTE                = C.ALLEGRO_KEY_QUOTE
	KEY_BACKSLASH            = C.ALLEGRO_KEY_BACKSLASH
	KEY_BACKSLASH2           = C.ALLEGRO_KEY_BACKSLASH2
	KEY_COMMA                = C.ALLEGRO_KEY_COMMA
	KEY_FULLSTOP             = C.ALLEGRO_KEY_FULLSTOP
	KEY_SLASH                = C.ALLEGRO_KEY_SLASH
	KEY_SPACE                = C.ALLEGRO_KEY_SPACE
	KEY_INSERT               = C.ALLEGRO_KEY_INSERT
	KEY_DELETE               = C.ALLEGRO_KEY_DELETE
	KEY_HOME                 = C.ALLEGRO_KEY_HOME
	KEY_END                  = C.ALLEGRO_KEY_END
	KEY_PGUP                 = C.ALLEGRO_KEY_PGUP
	KEY_PGDN                 = C.ALLEGRO_KEY_PGDN
	KEY_LEFT                 = C.ALLEGRO_KEY_LEFT
	KEY_RIGHT                = C.ALLEGRO_KEY_RIGHT
	KEY_UP                   = C.ALLEGRO_KEY_UP
	KEY_DOWN                 = C.ALLEGRO_KEY_DOWN
	KEY_PAD_SLASH            = C.ALLEGRO_KEY_PAD_SLASH
	KEY_PAD_ASTERISK         = C.ALLEGRO_KEY_PAD_ASTERISK
	KEY_PAD_MINUS            = C.ALLEGRO_KEY_PAD_MINUS
	KEY_PAD_PLUS             = C.ALLEGRO_KEY_PAD_PLUS
	KEY_PAD_DELETE           = C.ALLEGRO_KEY_PAD_DELETE
	KEY_PAD_ENTER            = C.ALLEGRO_KEY_PAD_ENTER
	KEY_PRINTSCREEN          = C.ALLEGRO_KEY_PRINTSCREEN
	KEY_PAUSE                = C.ALLEGRO_KEY_PAUSE
	KEY_ABNT_C1              = C.ALLEGRO_KEY_ABNT_C1
	KEY_YEN                  = C.ALLEGRO_KEY_YEN
	KEY_KANA                 = C.ALLEGRO_KEY_KANA
	KEY_CONVERT              = C.ALLEGRO_KEY_CONVERT
	KEY_NOCONVERT            = C.ALLEGRO_KEY_NOCONVERT
	KEY_AT                   = C.ALLEGRO_KEY_AT
	KEY_CIRCUMFLEX           = C.ALLEGRO_KEY_CIRCUMFLEX
	KEY_COLON2               = C.ALLEGRO_KEY_COLON2
	KEY_KANJI                = C.ALLEGRO_KEY_KANJI
	KEY_LSHIFT               = C.ALLEGRO_KEY_LSHIFT
	KEY_RSHIFT               = C.ALLEGRO_KEY_RSHIFT
	KEY_LCTRL                = C.ALLEGRO_KEY_LCTRL
	KEY_RCTRL                = C.ALLEGRO_KEY_RCTRL
	KEY_ALT                  = C.ALLEGRO_KEY_ALT
	KEY_ALTGR                = C.ALLEGRO_KEY_ALTGR
	KEY_LWIN                 = C.ALLEGRO_KEY_LWIN
	KEY_RWIN                 = C.ALLEGRO_KEY_RWIN
	KEY_MENU                 = C.ALLEGRO_KEY_MENU
	KEY_SCROLLLOCK           = C.ALLEGRO_KEY_SCROLLLOCK
	KEY_NUMLOCK              = C.ALLEGRO_KEY_NUMLOCK
	KEY_CAPSLOCK             = C.ALLEGRO_KEY_CAPSLOCK
	KEY_PAD_EQUALS           = C.ALLEGRO_KEY_PAD_EQUALS
	KEY_BACKQUOTE            = C.ALLEGRO_KEY_BACKQUOTE
	KEY_SEMICOLON2           = C.ALLEGRO_KEY_SEMICOLON2
	KEY_COMMAND              = C.ALLEGRO_KEY_COMMAND
)

//}}}

//{{{ key modifiers
type KeyModifier uint

const (
	KEYMOD_SHIFT      KeyModifier = C.ALLEGRO_KEYMOD_SHIFT
	KEYMOD_CTRL                   = C.ALLEGRO_KEYMOD_CTRL
	KEYMOD_ALT                    = C.ALLEGRO_KEYMOD_ALT
	KEYMOD_LWIN                   = C.ALLEGRO_KEYMOD_LWIN
	KEYMOD_RWIN                   = C.ALLEGRO_KEYMOD_RWIN
	KEYMOD_MENU                   = C.ALLEGRO_KEYMOD_MENU
	KEYMOD_ALTGR                  = C.ALLEGRO_KEYMOD_ALTGR
	KEYMOD_COMMAND                = C.ALLEGRO_KEYMOD_COMMAND
	KEYMOD_SCROLLLOCK             = C.ALLEGRO_KEYMOD_SCROLLLOCK
	KEYMOD_NUMLOCK                = C.ALLEGRO_KEYMOD_NUMLOCK
	KEYMOD_CAPSLOCK               = C.ALLEGRO_KEYMOD_CAPSLOCK
	KEYMOD_INALTSEQ               = C.ALLEGRO_KEYMOD_INALTSEQ
	KEYMOD_ACCENT1                = C.ALLEGRO_KEYMOD_ACCENT1
	KEYMOD_ACCENT2                = C.ALLEGRO_KEYMOD_ACCENT2
	KEYMOD_ACCENT3                = C.ALLEGRO_KEYMOD_ACCENT3
	KEYMOD_ACCENT4                = C.ALLEGRO_KEYMOD_ACCENT4
)

//}}}

// Install a keyboard driver. Returns true if successful. If a driver was
// already installed, nothing happens and true is returned.
func InstallKeyboard() error {
	success := bool(C.al_install_keyboard())
	if !success {
		return errors.New("failed to install keyboard")
	}
	return nil
}

// Returns true if al_install_keyboard was called successfully.
func IsKeyboardInstalled() bool {
	return bool(C.al_is_keyboard_installed())
}

// Uninstalls the active keyboard driver, if any. This will automatically
// unregister the keyboard event source with any event queues.
func UninstallKeyboard() {
	C.al_uninstall_keyboard()
}

// Save the state of the keyboard specified at the time the function is called
// into the structure pointed to by ret_state.
func (state *KeyboardState) Get() {
	C.al_get_keyboard_state((*C.ALLEGRO_KEYBOARD_STATE)(state))
}

// Return true if the key specified was held down in the state specified.
func (state *KeyboardState) IsDown(key KeyCode) bool {
	return bool(C.al_key_down((*C.ALLEGRO_KEYBOARD_STATE)(state), C.int(key)))
}

// Converts the given keycode to a description of the key.
func (key KeyCode) Name() string {
	name := C.al_keycode_to_name(C.int(key))
	return C.GoString(name)
}

// Implement the Stringer interface.
func (key KeyCode) String() string {
	return key.Name()
}

// Overrides the state of the keyboard LED indicators. Set to -1 to return to
// default behavior. False is returned if the current keyboard driver cannot
// set LED indicators.
func SetKeyboardLeds(leds int) bool {
	return bool(C.al_set_keyboard_leds(C.int(leds)))
}

// Retrieve the keyboard event source.
func KeyboardEventSource() (*EventSource, error) {
	source := C.al_get_keyboard_event_source()
	if source == nil {
		return nil, errors.New("failed to get keyboard event source; did you call InstallKeyboard() first?")
	}
	return (*EventSource)(source), nil
}
