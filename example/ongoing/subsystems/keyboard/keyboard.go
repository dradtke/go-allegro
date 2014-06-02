package keyboard

import (
	"fmt"
	al "github.com/dradtke/go-allegro/allegro"
)

var (
	ctrl_mod  int8
	alt_mod   int8
	shift_mod int8
)

type Mod int

const (
    Ctrl Mod = iota
    Alt
    Shift
)

func update_mods(key al.KeyCode, delta int8) (is_mod bool) {
    is_mod = true
	switch {
	case key == al.KEY_LCTRL || key == al.KEY_RCTRL:
		ctrl_mod += delta
	case key == al.KEY_LALT || key == al.KEY_RALT:
		alt_mod += delta
	case key == al.KEY_LSHIFT || key == al.KEY_RSHIFT:
		shift_mod += delta
    default:
        is_mod = false
	}
    return
}

var pressed = make(map[al.KeyCode]bool)

func Down(key al.KeyCode) {
	if !update_mods(key, 1) {
        pressed[key] = true
    }
}

func Up(key al.KeyCode) {
    if !update_mods(key, -1) {
        delete(pressed, key)
    }
}

func matches(mods []Mod, keys []al.KeyCode) bool {
    var (
        needCtrl bool
        needAlt bool
        needShift bool
    )

    for m := range mods {
        switch m {
        case Ctrl:
            needCtrl = true
        case Alt:
            needAlt = true
        case Shift:
            needShift = true
        }
    }

    if needCtrl != ctrl_mod || needAlt != alt_mod || needShift != shift_mod {
        return
    }

    // we have the modifiers!
}
