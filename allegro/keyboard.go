package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
)

type KeyboardState C.ALLEGRO_KEYBOARD_STATE

//{{{ key codes
type KeyCode int
const (
	KeyA            KeyCode = C.ALLEGRO_KEY_A
	KeyB            KeyCode = C.ALLEGRO_KEY_B
	KeyC            KeyCode = C.ALLEGRO_KEY_C
	KeyD            KeyCode = C.ALLEGRO_KEY_D
	KeyE            KeyCode = C.ALLEGRO_KEY_E
	KeyF            KeyCode = C.ALLEGRO_KEY_F
	KeyG            KeyCode = C.ALLEGRO_KEY_G
	KeyH            KeyCode = C.ALLEGRO_KEY_H
	KeyI            KeyCode = C.ALLEGRO_KEY_I
	KeyJ            KeyCode = C.ALLEGRO_KEY_J
	KeyK            KeyCode = C.ALLEGRO_KEY_K
	KeyL            KeyCode = C.ALLEGRO_KEY_L
	KeyM            KeyCode = C.ALLEGRO_KEY_M
	KeyN            KeyCode = C.ALLEGRO_KEY_N
	KeyO            KeyCode = C.ALLEGRO_KEY_O
	KeyP            KeyCode = C.ALLEGRO_KEY_P
	KeyQ            KeyCode = C.ALLEGRO_KEY_Q
	KeyR            KeyCode = C.ALLEGRO_KEY_R
	KeyS            KeyCode = C.ALLEGRO_KEY_S
	KeyT            KeyCode = C.ALLEGRO_KEY_T
	KeyU            KeyCode = C.ALLEGRO_KEY_U
	KeyV            KeyCode = C.ALLEGRO_KEY_V
	KeyW            KeyCode = C.ALLEGRO_KEY_W
	KeyX            KeyCode = C.ALLEGRO_KEY_X
	KeyY            KeyCode = C.ALLEGRO_KEY_Y
	KeyZ            KeyCode = C.ALLEGRO_KEY_Z
	Key0            KeyCode = C.ALLEGRO_KEY_0
	Key1            KeyCode = C.ALLEGRO_KEY_1
	Key2            KeyCode = C.ALLEGRO_KEY_2
	Key3            KeyCode = C.ALLEGRO_KEY_3
	Key4            KeyCode = C.ALLEGRO_KEY_4
	Key5            KeyCode = C.ALLEGRO_KEY_5
	Key6            KeyCode = C.ALLEGRO_KEY_6
	Key7            KeyCode = C.ALLEGRO_KEY_7
	Key8            KeyCode = C.ALLEGRO_KEY_8
	Key9            KeyCode = C.ALLEGRO_KEY_9
	KeyPad0         KeyCode = C.ALLEGRO_KEY_PAD_0
	KeyPad1         KeyCode = C.ALLEGRO_KEY_PAD_1
	KeyPad2         KeyCode = C.ALLEGRO_KEY_PAD_2
	KeyPad3         KeyCode = C.ALLEGRO_KEY_PAD_3
	KeyPad4         KeyCode = C.ALLEGRO_KEY_PAD_4
	KeyPad5         KeyCode = C.ALLEGRO_KEY_PAD_5
	KeyPad6         KeyCode = C.ALLEGRO_KEY_PAD_6
	KeyPad7         KeyCode = C.ALLEGRO_KEY_PAD_7
	KeyPad8         KeyCode = C.ALLEGRO_KEY_PAD_8
	KeyPad9         KeyCode = C.ALLEGRO_KEY_PAD_9
	KeyF1           KeyCode = C.ALLEGRO_KEY_F1
	KeyF2           KeyCode = C.ALLEGRO_KEY_F2
	KeyF3           KeyCode = C.ALLEGRO_KEY_F3
	KeyF4           KeyCode = C.ALLEGRO_KEY_F4
	KeyF5           KeyCode = C.ALLEGRO_KEY_F5
	KeyF6           KeyCode = C.ALLEGRO_KEY_F6
	KeyF7           KeyCode = C.ALLEGRO_KEY_F7
	KeyF8           KeyCode = C.ALLEGRO_KEY_F8
	KeyF9           KeyCode = C.ALLEGRO_KEY_F9
	KeyF10          KeyCode = C.ALLEGRO_KEY_F10
	KeyF11          KeyCode = C.ALLEGRO_KEY_F11
	KeyF12          KeyCode = C.ALLEGRO_KEY_F12
	KeyEscape       KeyCode = C.ALLEGRO_KEY_ESCAPE
	KeyTilde        KeyCode = C.ALLEGRO_KEY_TILDE
	KeyMinus        KeyCode = C.ALLEGRO_KEY_MINUS
	KeyEquals       KeyCode = C.ALLEGRO_KEY_EQUALS
	KeyBackspace    KeyCode = C.ALLEGRO_KEY_BACKSPACE
	KeyTab          KeyCode = C.ALLEGRO_KEY_TAB
	KeyOpenBrace    KeyCode = C.ALLEGRO_KEY_OPENBRACE
	KeyCloseBrace   KeyCode = C.ALLEGRO_KEY_CLOSEBRACE
	KeyEnter        KeyCode = C.ALLEGRO_KEY_ENTER
	KeySemicolon    KeyCode = C.ALLEGRO_KEY_SEMICOLON
	KeyQuote        KeyCode = C.ALLEGRO_KEY_QUOTE
	KeyBackslash    KeyCode = C.ALLEGRO_KEY_BACKSLASH
	KeyBackslash2   KeyCode = C.ALLEGRO_KEY_BACKSLASH2
	KeyComma        KeyCode = C.ALLEGRO_KEY_COMMA
	KeyFullstop     KeyCode = C.ALLEGRO_KEY_FULLSTOP
	KeySlash        KeyCode = C.ALLEGRO_KEY_SLASH
	KeySpace        KeyCode = C.ALLEGRO_KEY_SPACE
	KeyInsert       KeyCode = C.ALLEGRO_KEY_INSERT
	KeyDelete       KeyCode = C.ALLEGRO_KEY_DELETE
	KeyHome         KeyCode = C.ALLEGRO_KEY_HOME
	KeyEnd          KeyCode = C.ALLEGRO_KEY_END
	KeyPgup         KeyCode = C.ALLEGRO_KEY_PGUP
	KeyPgdn         KeyCode = C.ALLEGRO_KEY_PGDN
	KeyLeft         KeyCode = C.ALLEGRO_KEY_LEFT
	KeyRight        KeyCode = C.ALLEGRO_KEY_RIGHT
	KeyUp           KeyCode = C.ALLEGRO_KEY_UP
	KeyDown         KeyCode = C.ALLEGRO_KEY_DOWN
	KeyPadSlash     KeyCode = C.ALLEGRO_KEY_PAD_SLASH
	KeyPadAsterisk  KeyCode = C.ALLEGRO_KEY_PAD_ASTERISK
	KeyPadMinus     KeyCode = C.ALLEGRO_KEY_PAD_MINUS
	KeyPadPlus      KeyCode = C.ALLEGRO_KEY_PAD_PLUS
	KeyPadDelete    KeyCode = C.ALLEGRO_KEY_PAD_DELETE
	KeyPadEnter     KeyCode = C.ALLEGRO_KEY_PAD_ENTER
	KeyPrintScreen  KeyCode = C.ALLEGRO_KEY_PRINTSCREEN
	KeyPause        KeyCode = C.ALLEGRO_KEY_PAUSE
	KeyAbntC1       KeyCode = C.ALLEGRO_KEY_ABNT_C1
	KeyYen          KeyCode = C.ALLEGRO_KEY_YEN
	KeyKana         KeyCode = C.ALLEGRO_KEY_KANA
	KeyConvert      KeyCode = C.ALLEGRO_KEY_CONVERT
	KeyNoConvert    KeyCode = C.ALLEGRO_KEY_NOCONVERT
	KeyAt           KeyCode = C.ALLEGRO_KEY_AT
	KeyCircumflex   KeyCode = C.ALLEGRO_KEY_CIRCUMFLEX
	KeyColon2       KeyCode = C.ALLEGRO_KEY_COLON2
	KeyKanji        KeyCode = C.ALLEGRO_KEY_KANJI
	KeyLShift       KeyCode = C.ALLEGRO_KEY_LSHIFT
	KeyRShift       KeyCode = C.ALLEGRO_KEY_RSHIFT
	KeyLCtrl        KeyCode = C.ALLEGRO_KEY_LCTRL
	KeyRCtrl        KeyCode = C.ALLEGRO_KEY_RCTRL
	KeyAlt          KeyCode = C.ALLEGRO_KEY_ALT
	KeyAltgr        KeyCode = C.ALLEGRO_KEY_ALTGR
	KeyLWin         KeyCode = C.ALLEGRO_KEY_LWIN
	KeyRWin         KeyCode = C.ALLEGRO_KEY_RWIN
	KeyMenu         KeyCode = C.ALLEGRO_KEY_MENU
	KeyScrolllock   KeyCode = C.ALLEGRO_KEY_SCROLLLOCK
	KeyNumlock      KeyCode = C.ALLEGRO_KEY_NUMLOCK
	KeyCapslock     KeyCode = C.ALLEGRO_KEY_CAPSLOCK
	KeyPadEquals    KeyCode = C.ALLEGRO_KEY_PAD_EQUALS
	KeyBackquote    KeyCode = C.ALLEGRO_KEY_BACKQUOTE
	KeySemicolon2   KeyCode = C.ALLEGRO_KEY_SEMICOLON2
	KeyCommand      KeyCode = C.ALLEGRO_KEY_COMMAND
)
//}}}

//{{{ key modifiers
type KeyModifier uint
const (
	KeymodShift         KeyModifier = C.ALLEGRO_KEYMOD_SHIFT
	KeymodCtrl          KeyModifier = C.ALLEGRO_KEYMOD_CTRL
	KeymodAlt           KeyModifier = C.ALLEGRO_KEYMOD_ALT
	KeymodLwin          KeyModifier = C.ALLEGRO_KEYMOD_LWIN
	KeymodRwin          KeyModifier = C.ALLEGRO_KEYMOD_RWIN
	KeymodMenu          KeyModifier = C.ALLEGRO_KEYMOD_MENU
	KeymodAltgr         KeyModifier = C.ALLEGRO_KEYMOD_ALTGR
	KeymodCommand       KeyModifier = C.ALLEGRO_KEYMOD_COMMAND
	KeymodScrolllock    KeyModifier = C.ALLEGRO_KEYMOD_SCROLLLOCK
	KeymodNumlock       KeyModifier = C.ALLEGRO_KEYMOD_NUMLOCK
	KeymodCapslock      KeyModifier = C.ALLEGRO_KEYMOD_CAPSLOCK
	KeymodInaltseq      KeyModifier = C.ALLEGRO_KEYMOD_INALTSEQ
	KeymodAccent1       KeyModifier = C.ALLEGRO_KEYMOD_ACCENT1
	KeymodAccent2       KeyModifier = C.ALLEGRO_KEYMOD_ACCENT2
	KeymodAccent3       KeyModifier = C.ALLEGRO_KEYMOD_ACCENT3
	KeymodAccent4       KeyModifier = C.ALLEGRO_KEYMOD_ACCENT4
)
//}}}

func InstallKeyboard() error {
	success := bool(C.al_install_keyboard())
	if !success {
		return errors.New("failed to install keyboard!")
	}
	return nil
}

func IsKeyboardInstalled() bool {
	return bool(C.al_is_keyboard_installed())
}

func UninstallKeyboard() {
	C.al_uninstall_keyboard()
}

func (state *KeyboardState) Get() {
	C.al_get_keyboard_state((*C.ALLEGRO_KEYBOARD_STATE)(state))
}

func (state *KeyboardState) IsDown(key KeyCode) bool {
	return bool(C.al_key_down((*C.ALLEGRO_KEYBOARD_STATE)(state), C.int(key)))
}

func (key KeyCode) Name() string {
	name := C.al_keycode_to_name(C.int(key))
	return C.GoString(name)
}

func SetKeyboardLeds(leds int) bool {
	return bool(C.al_set_keyboard_leds(C.int(leds)))
}

func KeyboardEventSource() (*EventSource, error) {
	source := C.al_get_keyboard_event_source()
	if source == nil {
		return nil, errors.New("cannot get keyboard event source; did you call InstallKeyboard() first?")
	}
	return (*EventSource)(source), nil
}
