package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

type Joystick C.ALLEGRO_JOYSTICK

type JoyFlags int

const (
	JOYFLAG_DIGITAL JoyFlags = C.ALLEGRO_JOYFLAG_DIGITAL
	JOYFLAG_ANALOGUE  JoyFlags = C.ALLEGRO_JOYFLAG_ANALOGUE
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

func InstallJoystick() error {
	success := bool(C.al_install_joystick())
	if !success {
		return errors.New("failed to install joystick!")
	}
	return nil
}

func UninstallJoystick() {
	C.al_uninstall_joystick()
}

func IsJoystickInstalled() bool {
	return bool(C.al_is_joystick_installed())
}

func ReconfigureJoysticks() bool {
	return bool(C.al_reconfigure_joysticks())
}

func NumJoysticks() int {
	return int(C.al_get_num_joysticks())
}

func GetJoystick(stick int) (*Joystick, error) {
	joystick := (*Joystick)(C.al_get_joystick(C.int(stick)))
	if joystick == nil {
		return nil, fmt.Errorf("joystick '%d' not found", stick)
	}
	return joystick, nil
}

func (j *Joystick) Release() {
	C.al_release_joystick((*C.ALLEGRO_JOYSTICK)(j))
}

func (j *Joystick) Active() bool {
	return bool(C.al_get_joystick_active((*C.ALLEGRO_JOYSTICK)(j)))
}

func (j *Joystick) Name() string {
	return C.GoString(C.al_get_joystick_name((*C.ALLEGRO_JOYSTICK)(j)))
}

func (j *Joystick) StickName(stick int) (string, error) {
	name_ := C.al_get_joystick_stick_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick))
	if name_ == nil {
		return "", fmt.Errorf("stick '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

func (j *Joystick) AxisName(stick, axis int) (string, error) {
	name_ := C.al_get_joystick_axis_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick), C.int(axis))
	if name_ == nil {
		return "", fmt.Errorf("axis '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

func (j *Joystick) ButtonName(stick int) (string, error) {
	name_ := C.al_get_joystick_button_name((*C.ALLEGRO_JOYSTICK)(j), C.int(stick))
	if name_ == nil {
		return "", fmt.Errorf("button '%d' not found on joystick '%s'", stick, j.Name())
	}
	return C.GoString(name_), nil
}

func (j *Joystick) NumSticks() int {
	return int(C.al_get_joystick_num_sticks((*C.ALLEGRO_JOYSTICK)(j)))
}

func (j *Joystick) NumAxes(stick int) int {
	return int(C.al_get_joystick_num_axes((*C.ALLEGRO_JOYSTICK)(j), C.int(stick)))
}

func (j *Joystick) NumButtons() int {
	return int(C.al_get_joystick_num_buttons((*C.ALLEGRO_JOYSTICK)(j)))
}

func (j *Joystick) StickFlags(stick int) JoyFlags {
	return JoyFlags(C.al_get_joystick_stick_flags((*C.ALLEGRO_JOYSTICK)(j), C.int(stick)))
}

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
