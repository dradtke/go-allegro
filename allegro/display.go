package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
)

type Display C.ALLEGRO_DISPLAY

type DisplayFlags int
const (
	Windowed DisplayFlags = C.ALLEGRO_WINDOWED
	Fullscreen DisplayFlags = C.ALLEGRO_FULLSCREEN
	FullscreenWindow DisplayFlags = C.ALLEGRO_FULLSCREEN_WINDOW
	Resizable DisplayFlags = C.ALLEGRO_RESIZABLE
	// TODO: add the rest of these flags
)

type DisplayOrientation int
const (
	ZeroDegrees                 DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
	NinetyDegrees               DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
	OneHundredEightyDegrees     DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
	TwoHundredSeventyDegrees    DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
	FaceUp                      DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
	FaceDown                    DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

func CreateDisplay(w, h int) (*Display, error) {
	d := C.al_create_display(cint(w), cint(h))
	if d == nil {
		return nil, errors.New("failed to create display!")
	}
	return (*Display)(d), nil
}

func (d *Display) Destroy() {
	C.al_destroy_display((*C.ALLEGRO_DISPLAY)(d))
}

func FlipDisplay() {
	C.al_flip_display()
}

func GetNewDisplayFlags() DisplayFlags {
	return (DisplayFlags)(C.al_get_new_display_flags())
}

func SetNewDisplayFlags(flags DisplayFlags) {
	C.al_set_new_display_flags(cint(int(flags)))
}

func ResetDisplayFlags() {
	C.al_set_new_display_flags(cint(0))
}

func (d *Display) GetEventSource() *EventSource {
	return (*EventSource)(C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) Width() int {
	return (int)(C.al_get_display_width((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) Height() int {
	return (int)(C.al_get_display_height((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) AcknowledgeResize() bool {
	return gobool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) SetWindowTitle(title string) {
	title_ := C.CString(title) ; defer FreeString(title_)
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(d), title_)
}
