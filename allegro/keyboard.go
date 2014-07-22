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
	KEY_B            KeyCode = C.ALLEGRO_KEY_B
	KEY_C            KeyCode = C.ALLEGRO_KEY_C
	KEY_D            KeyCode = C.ALLEGRO_KEY_D
	KEY_E            KeyCode = C.ALLEGRO_KEY_E
	KEY_F            KeyCode = C.ALLEGRO_KEY_F
	KEY_G            KeyCode = C.ALLEGRO_KEY_G
	KEY_H            KeyCode = C.ALLEGRO_KEY_H
	KEY_I            KeyCode = C.ALLEGRO_KEY_I
	KEY_J            KeyCode = C.ALLEGRO_KEY_J
	KEY_K            KeyCode = C.ALLEGRO_KEY_K
	KEY_L            KeyCode = C.ALLEGRO_KEY_L
	KEY_M            KeyCode = C.ALLEGRO_KEY_M
	KEY_N            KeyCode = C.ALLEGRO_KEY_N
	KEY_O            KeyCode = C.ALLEGRO_KEY_O
	KEY_P            KeyCode = C.ALLEGRO_KEY_P
	KEY_Q            KeyCode = C.ALLEGRO_KEY_Q
	KEY_R            KeyCode = C.ALLEGRO_KEY_R
	KEY_S            KeyCode = C.ALLEGRO_KEY_S
	KEY_T            KeyCode = C.ALLEGRO_KEY_T
	KEY_U            KeyCode = C.ALLEGRO_KEY_U
	KEY_V            KeyCode = C.ALLEGRO_KEY_V
	KEY_W            KeyCode = C.ALLEGRO_KEY_W
	KEY_X            KeyCode = C.ALLEGRO_KEY_X
	KEY_Y            KeyCode = C.ALLEGRO_KEY_Y
	KEY_Z            KeyCode = C.ALLEGRO_KEY_Z
	KEY_0            KeyCode = C.ALLEGRO_KEY_0
	KEY_1            KeyCode = C.ALLEGRO_KEY_1
	KEY_2            KeyCode = C.ALLEGRO_KEY_2
	KEY_3            KeyCode = C.ALLEGRO_KEY_3
	KEY_4            KeyCode = C.ALLEGRO_KEY_4
	KEY_5            KeyCode = C.ALLEGRO_KEY_5
	KEY_6            KeyCode = C.ALLEGRO_KEY_6
	KEY_7            KeyCode = C.ALLEGRO_KEY_7
	KEY_8            KeyCode = C.ALLEGRO_KEY_8
	KEY_9            KeyCode = C.ALLEGRO_KEY_9
	KEY_PAD_0        KeyCode = C.ALLEGRO_KEY_PAD_0
	KEY_PAD_1        KeyCode = C.ALLEGRO_KEY_PAD_1
	KEY_PAD_2        KeyCode = C.ALLEGRO_KEY_PAD_2
	KEY_PAD_3        KeyCode = C.ALLEGRO_KEY_PAD_3
	KEY_PAD_4        KeyCode = C.ALLEGRO_KEY_PAD_4
	KEY_PAD_5        KeyCode = C.ALLEGRO_KEY_PAD_5
	KEY_PAD_6        KeyCode = C.ALLEGRO_KEY_PAD_6
	KEY_PAD_7        KeyCode = C.ALLEGRO_KEY_PAD_7
	KEY_PAD_8        KeyCode = C.ALLEGRO_KEY_PAD_8
	KEY_PAD_9        KeyCode = C.ALLEGRO_KEY_PAD_9
	KEY_F1           KeyCode = C.ALLEGRO_KEY_F1
	KEY_F2           KeyCode = C.ALLEGRO_KEY_F2
	KEY_F3           KeyCode = C.ALLEGRO_KEY_F3
	KEY_F4           KeyCode = C.ALLEGRO_KEY_F4
	KEY_F5           KeyCode = C.ALLEGRO_KEY_F5
	KEY_F6           KeyCode = C.ALLEGRO_KEY_F6
	KEY_F7           KeyCode = C.ALLEGRO_KEY_F7
	KEY_F8           KeyCode = C.ALLEGRO_KEY_F8
	KEY_F9           KeyCode = C.ALLEGRO_KEY_F9
	KEY_F10          KeyCode = C.ALLEGRO_KEY_F10
	KEY_F11          KeyCode = C.ALLEGRO_KEY_F11
	KEY_F12          KeyCode = C.ALLEGRO_KEY_F12
	KEY_ESCAPE       KeyCode = C.ALLEGRO_KEY_ESCAPE
	KEY_TILDE        KeyCode = C.ALLEGRO_KEY_TILDE
	KEY_MINUS        KeyCode = C.ALLEGRO_KEY_MINUS
	KEY_EQUALS       KeyCode = C.ALLEGRO_KEY_EQUALS
	KEY_BACKSPACE    KeyCode = C.ALLEGRO_KEY_BACKSPACE
	KEY_TAB          KeyCode = C.ALLEGRO_KEY_TAB
	KEY_OPENBRACE    KeyCode = C.ALLEGRO_KEY_OPENBRACE
	KEY_CLOSEBRACE   KeyCode = C.ALLEGRO_KEY_CLOSEBRACE
	KEY_ENTER        KeyCode = C.ALLEGRO_KEY_ENTER
	KEY_SEMICOLON    KeyCode = C.ALLEGRO_KEY_SEMICOLON
	KEY_QUOTE        KeyCode = C.ALLEGRO_KEY_QUOTE
	KEY_BACKSLASH    KeyCode = C.ALLEGRO_KEY_BACKSLASH
	KEY_BACKSLASH2   KeyCode = C.ALLEGRO_KEY_BACKSLASH2
	KEY_COMMA        KeyCode = C.ALLEGRO_KEY_COMMA
	KEY_FULLSTOP     KeyCode = C.ALLEGRO_KEY_FULLSTOP
	KEY_SLASH        KeyCode = C.ALLEGRO_KEY_SLASH
	KEY_SPACE        KeyCode = C.ALLEGRO_KEY_SPACE
	KEY_INSERT       KeyCode = C.ALLEGRO_KEY_INSERT
	KEY_DELETE       KeyCode = C.ALLEGRO_KEY_DELETE
	KEY_HOME         KeyCode = C.ALLEGRO_KEY_HOME
	KEY_END          KeyCode = C.ALLEGRO_KEY_END
	KEY_PGUP         KeyCode = C.ALLEGRO_KEY_PGUP
	KEY_PGDN         KeyCode = C.ALLEGRO_KEY_PGDN
	KEY_LEFT         KeyCode = C.ALLEGRO_KEY_LEFT
	KEY_RIGHT        KeyCode = C.ALLEGRO_KEY_RIGHT
	KEY_UP           KeyCode = C.ALLEGRO_KEY_UP
	KEY_DOWN         KeyCode = C.ALLEGRO_KEY_DOWN
	KEY_PAD_SLASH    KeyCode = C.ALLEGRO_KEY_PAD_SLASH
	KEY_PAD_ASTERISK KeyCode = C.ALLEGRO_KEY_PAD_ASTERISK
	KEY_PAD_MINUS    KeyCode = C.ALLEGRO_KEY_PAD_MINUS
	KEY_PAD_PLUS     KeyCode = C.ALLEGRO_KEY_PAD_PLUS
	KEY_PAD_DELETE   KeyCode = C.ALLEGRO_KEY_PAD_DELETE
	KEY_PAD_ENTER    KeyCode = C.ALLEGRO_KEY_PAD_ENTER
	KEY_PRINTSCREEN  KeyCode = C.ALLEGRO_KEY_PRINTSCREEN
	KEY_PAUSE        KeyCode = C.ALLEGRO_KEY_PAUSE
	KEY_ABNT_C1      KeyCode = C.ALLEGRO_KEY_ABNT_C1
	KEY_YEN          KeyCode = C.ALLEGRO_KEY_YEN
	KEY_KANA         KeyCode = C.ALLEGRO_KEY_KANA
	KEY_CONVERT      KeyCode = C.ALLEGRO_KEY_CONVERT
	KEY_NOCONVERT    KeyCode = C.ALLEGRO_KEY_NOCONVERT
	KEY_AT           KeyCode = C.ALLEGRO_KEY_AT
	KEY_CIRCUMFLEX   KeyCode = C.ALLEGRO_KEY_CIRCUMFLEX
	KEY_COLON2       KeyCode = C.ALLEGRO_KEY_COLON2
	KEY_KANJI        KeyCode = C.ALLEGRO_KEY_KANJI
	KEY_LSHIFT       KeyCode = C.ALLEGRO_KEY_LSHIFT
	KEY_RSHIFT       KeyCode = C.ALLEGRO_KEY_RSHIFT
	KEY_LCTRL        KeyCode = C.ALLEGRO_KEY_LCTRL
	KEY_RCTRL        KeyCode = C.ALLEGRO_KEY_RCTRL
	KEY_ALT          KeyCode = C.ALLEGRO_KEY_ALT
	KEY_ALTGR        KeyCode = C.ALLEGRO_KEY_ALTGR
	KEY_LWIN         KeyCode = C.ALLEGRO_KEY_LWIN
	KEY_RWIN         KeyCode = C.ALLEGRO_KEY_RWIN
	KEY_MENU         KeyCode = C.ALLEGRO_KEY_MENU
	KEY_SCROLLLOCK   KeyCode = C.ALLEGRO_KEY_SCROLLLOCK
	KEY_NUMLOCK      KeyCode = C.ALLEGRO_KEY_NUMLOCK
	KEY_CAPSLOCK     KeyCode = C.ALLEGRO_KEY_CAPSLOCK
	KEY_PAD_EQUALS   KeyCode = C.ALLEGRO_KEY_PAD_EQUALS
	KEY_BACKQUOTE    KeyCode = C.ALLEGRO_KEY_BACKQUOTE
	KEY_SEMICOLON2   KeyCode = C.ALLEGRO_KEY_SEMICOLON2
	KEY_COMMAND      KeyCode = C.ALLEGRO_KEY_COMMAND
)

//}}}

//{{{ key modifiers
type KeyModifier uint

const (
	KEYMOD_SHIFT      KeyModifier = C.ALLEGRO_KEYMOD_SHIFT
	KEYMOD_CTRL       KeyModifier = C.ALLEGRO_KEYMOD_CTRL
	KEYMOD_ALT        KeyModifier = C.ALLEGRO_KEYMOD_ALT
	KEYMOD_LWIN       KeyModifier = C.ALLEGRO_KEYMOD_LWIN
	KEYMOD_RWIN       KeyModifier = C.ALLEGRO_KEYMOD_RWIN
	KEYMOD_MENU       KeyModifier = C.ALLEGRO_KEYMOD_MENU
	KEYMOD_ALTGR      KeyModifier = C.ALLEGRO_KEYMOD_ALTGR
	KEYMOD_COMMAND    KeyModifier = C.ALLEGRO_KEYMOD_COMMAND
	KEYMOD_SCROLLLOCK KeyModifier = C.ALLEGRO_KEYMOD_SCROLLLOCK
	KEYMOD_NUMLOCK    KeyModifier = C.ALLEGRO_KEYMOD_NUMLOCK
	KEYMOD_CAPSLOCK   KeyModifier = C.ALLEGRO_KEYMOD_CAPSLOCK
	KEYMOD_INALTSEQ   KeyModifier = C.ALLEGRO_KEYMOD_INALTSEQ
	KEYMOD_ACCENT1    KeyModifier = C.ALLEGRO_KEYMOD_ACCENT1
	KEYMOD_ACCENT2    KeyModifier = C.ALLEGRO_KEYMOD_ACCENT2
	KEYMOD_ACCENT3    KeyModifier = C.ALLEGRO_KEYMOD_ACCENT3
	KEYMOD_ACCENT4    KeyModifier = C.ALLEGRO_KEYMOD_ACCENT4
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
