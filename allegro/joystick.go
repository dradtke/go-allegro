package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
	"fmt"
)

type Joystick C.ALLEGRO_JOYSTICK

type JoyFlags int

const (
	JOYFLAG_DIGITAL  JoyFlags = C.ALLEGRO_JOYFLAG_DIGITAL
	JOYFLAG_ANALOGUE          = C.ALLEGRO_JOYFLAG_ANALOGUE
)

type JoystickState struct {
	Stick    []stickState
	Button   []int
	joystick *Joystick
	ptr      C.ALLEGRO_JOYSTICK_STATE
}

type stickState struct {
	Axis []float32
}

// Install a joystick driver, returning true if successful. If a joystick
// driver was already installed, returns true immediately.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_install_joystick
func InstallJoystick() error {
	success := bool(C.al_install_joystick())
	if !success {
		return errors.New("failed to install joystick!")
	}
	return nil
}

// Uninstalls the active joystick driver. All outstanding ALLEGRO_JOYSTICK
// structures are invalidated. If no joystick driver was active, this function
// does nothing.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_uninstall_joystick
func UninstallJoystick() {
	C.al_uninstall_joystick()
}

// Returns true if al_install_joystick was called successfully.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_is_joystick_installed
func IsJoystickInstalled() bool {
	return bool(C.al_is_joystick_installed())
}

// Allegro is able to cope with users connecting and disconnected joystick
// devices on-the-fly. On existing platforms, the joystick event source will
// generate an event of type ALLEGRO_EVENT_JOYSTICK_CONFIGURATION when a device
// is plugged in or unplugged. In response, you should call
// al_reconfigure_joysticks.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_reconfigure_joysticks
func ReconfigureJoysticks() bool {
	return bool(C.al_reconfigure_joysticks())
}

// Return the number of joysticks currently on the system (or potentially on
// the system). This number can change after al_reconfigure_joysticks is
// called, in order to support hotplugging.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_num_joysticks
func NumJoysticks() int {
	return int(C.al_get_num_joysticks())
}

// Get a handle for a joystick on the system. The number may be from 0 to
// al_get_num_joysticks-1. If successful a pointer to a joystick object is
// returned, which represents a physical device. Otherwise NULL is returned.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick
func GetJoystick(stick int) (*Joystick, error) {
	joystick := (*Joystick)(C.al_get_joystick(C.int(stick)))
	if joystick == nil {
		return nil, fmt.Errorf("joystick '%d' not found", stick)
	}
	return joystick, nil
}

// This function currently does nothing.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_release_joystick
func (j *Joystick) Release() {
	C.al_release_joystick((*C.ALLEGRO_JOYSTICK)(j))
}

// Return if the joystick handle is "active", i.e. in the current
// configuration, the handle represents some physical device plugged into the
// system. al_get_joystick returns active handles. After reconfiguration,
// active handles may become inactive, and vice versa.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_active
func (j *Joystick) Active() bool {
	return bool(C.al_get_joystick_active((*C.ALLEGRO_JOYSTICK)(j)))
}

// Return the name of the given joystick.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_name
func (j *Joystick) Name() string {
	return C.GoString(C.al_get_joystick_name((*C.ALLEGRO_JOYSTICK)(j)))
}

// Return the name of the given "stick". If the stick doesn't exist, NULL is
// returned.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_stick_name
func (j *Joystick) StickName(stick int) (string, error) {
	name_ := C.al_get_joystick_stick_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick))
	if name_ == nil {
		return "", fmt.Errorf("stick '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

// Return the name of the given axis. If the axis doesn't exist, NULL is
// returned. Indices begin from 0.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_axis_name
func (j *Joystick) AxisName(stick, axis int) (string, error) {
	name_ := C.al_get_joystick_axis_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick), C.int(axis))
	if name_ == nil {
		return "", fmt.Errorf("axis '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

// Return the name of the given button. If the button doesn't exist, NULL is
// returned. Indices begin from 0.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_button_name
func (j *Joystick) ButtonName(stick int) (string, error) {
	name_ := C.al_get_joystick_button_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick))
	if name_ == nil {
		return "", fmt.Errorf("button '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

// Return the number of "sticks" on the given joystick. A stick has one or more
// axes.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_num_sticks
func (j *Joystick) NumSticks() int {
	return int(C.al_get_joystick_num_sticks((*C.ALLEGRO_JOYSTICK)(j)))
}

// Return the number of axes on the given "stick". If the stick doesn't exist,
// 0 is returned.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_num_axes
func (j *Joystick) NumAxes(stick int) int {
	return int(C.al_get_joystick_num_axes((*C.ALLEGRO_JOYSTICK)(j), C.int(stick)))
}

// Return the number of buttons on the joystick.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_num_buttons
func (j *Joystick) NumButtons() int {
	return int(C.al_get_joystick_num_buttons((*C.ALLEGRO_JOYSTICK)(j)))
}

// Return the flags of the given "stick". If the stick doesn't exist, NULL is
// returned. Indices begin from 0.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_stick_flags
func (j *Joystick) StickFlags(stick int) JoyFlags {
	return JoyFlags(C.al_get_joystick_stick_flags((*C.ALLEGRO_JOYSTICK)(j), C.int(stick)))
}

// Returns the global joystick event source. All joystick events are generated
// by this event source.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_event_source
func JoystickEventSource() *EventSource {
	return (*EventSource)(C.al_get_joystick_event_source())
}

func (j *Joystick) State() *JoystickState {
	numSticks := j.NumSticks()
	numButtons := j.NumButtons()
	state := JoystickState{joystick: j, Stick: make([]stickState, numSticks), Button: make([]int, numButtons)}
	for i := 0; i < numSticks; i++ {
		state.Stick[i].Axis = make([]float32, j.NumAxes(i))
	}
	return &state
}

// Get the current joystick state.
//
// See https://liballeg.org/a5docs/5.2.6/joystick.html#al_get_joystick_state
func (state *JoystickState) Get() {
	C.al_get_joystick_state((*C.ALLEGRO_JOYSTICK)(state.joystick), &state.ptr)
	for i := 0; i < len(state.Button); i++ {
		state.Button[i] = int(state.ptr.button[C.int(i)])
	}
	for i := 0; i < len(state.Stick); i++ {
		for j := 0; j < len(state.Stick[i].Axis); j++ {
			state.Stick[i].Axis[j] = float32(state.ptr.stick[C.int(i)].axis[C.int(j)])
		}
	}
}
