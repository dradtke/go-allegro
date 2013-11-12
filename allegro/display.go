package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
)

type Display C.ALLEGRO_DISPLAY

type DisplayFlags int

const (
	WINDOWED                  DisplayFlags = C.ALLEGRO_WINDOWED
	FULLSCREEN                DisplayFlags = C.ALLEGRO_FULLSCREEN
	FULLSCREEN_WINDOW         DisplayFlags = C.ALLEGRO_FULLSCREEN_WINDOW
	RESIZABLE                 DisplayFlags = C.ALLEGRO_RESIZABLE
	OPENGL                    DisplayFlags = C.ALLEGRO_OPENGL
	OPENGL_3_0                DisplayFlags = C.ALLEGRO_OPENGL_3_0
	OPENGL_FORWARD_COMPATIBLE DisplayFlags = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	FRAMELESS                 DisplayFlags = C.ALLEGRO_FRAMELESS
	NOFRAME                   DisplayFlags = C.ALLEGRO_NOFRAME
	GENERATE_EXPOSE_EVENTS    DisplayFlags = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
)

type DisplayMode struct {
	Width, Height, Format, RefreshRate int
	ptr                                C.ALLEGRO_DISPLAY_MODE
}

type MonitorInfo struct {
	X1, Y1, X2, Y2 int
	ptr            C.ALLEGRO_MONITOR_INFO
}

type DisplayOption C.int

const (
	COLOR_SIZE             DisplayOption = C.ALLEGRO_COLOR_SIZE
	RED_SIZE               DisplayOption = C.ALLEGRO_RED_SIZE
	GREEN_SIZE             DisplayOption = C.ALLEGRO_GREEN_SIZE
	BLUE_SIZE              DisplayOption = C.ALLEGRO_BLUE_SIZE
	ALPHA_SIZE             DisplayOption = C.ALLEGRO_ALPHA_SIZE
	RED_SHIFT              DisplayOption = C.ALLEGRO_RED_SHIFT
	GREEN_SHIFT            DisplayOption = C.ALLEGRO_GREEN_SHIFT
	BLUE_SHIFT             DisplayOption = C.ALLEGRO_BLUE_SHIFT
	ALPHA_SHIFT            DisplayOption = C.ALLEGRO_ALPHA_SHIFT
	ACC_RED_SIZE           DisplayOption = C.ALLEGRO_ACC_RED_SIZE
	ACC_GREEN_SIZE         DisplayOption = C.ALLEGRO_ACC_GREEN_SIZE
	ACC_BLUE_SIZE          DisplayOption = C.ALLEGRO_ACC_BLUE_SIZE
	ACC_ALPHA_SIZE         DisplayOption = C.ALLEGRO_ACC_ALPHA_SIZE
	STEREO                 DisplayOption = C.ALLEGRO_STEREO
	AUX_BUFFERS            DisplayOption = C.ALLEGRO_AUX_BUFFERS
	DEPTH_SIZE             DisplayOption = C.ALLEGRO_DEPTH_SIZE
	STENCIL_SIZE           DisplayOption = C.ALLEGRO_STENCIL_SIZE
	SAMPLE_BUFFERS         DisplayOption = C.ALLEGRO_SAMPLE_BUFFERS
	SAMPLES                DisplayOption = C.ALLEGRO_SAMPLES
	RENDER_METHOD          DisplayOption = C.ALLEGRO_RENDER_METHOD
	FLOAT_COLOR            DisplayOption = C.ALLEGRO_FLOAT_COLOR
	FLOAT_DEPTH            DisplayOption = C.ALLEGRO_FLOAT_DEPTH
	SINGLE_BUFFER          DisplayOption = C.ALLEGRO_SINGLE_BUFFER
	SWAP_METHOD            DisplayOption = C.ALLEGRO_SWAP_METHOD
	COMPATIBLE_DISPLAY     DisplayOption = C.ALLEGRO_COMPATIBLE_DISPLAY
	UPDATE_DISPLAY_REGION  DisplayOption = C.ALLEGRO_UPDATE_DISPLAY_REGION
	VSYNC                  DisplayOption = C.ALLEGRO_VSYNC
	MAX_BITMAP_SIZE        DisplayOption = C.ALLEGRO_MAX_BITMAP_SIZE
	SUPPORT_NPOT_BITMAP    DisplayOption = C.ALLEGRO_SUPPORT_NPOT_BITMAP
	CAN_DRAW_INTO_BITMAP   DisplayOption = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
	SUPPORT_SEPARATE_ALPHA DisplayOption = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
)

type Importance C.int

const (
	REQUIRE  Importance = C.ALLEGRO_REQUIRE
	SUGGEST  Importance = C.ALLEGRO_SUGGEST
	DONTCARE Importance = C.ALLEGRO_DONTCARE
)

type DisplayOrientation C.int

const (
	DISPLAY_ORIENTATION_0_DEGREES   DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
	DISPLAY_ORIENTATION_90_DEGREES  DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
	DISPLAY_ORIENTATION_180_DEGREES DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
	DISPLAY_ORIENTATION_270_DEGREES DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
	DISPLAY_ORIENTATION_FACE_UP     DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
	DISPLAY_ORIENTATION_FACE_DOWN   DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

func CreateDisplay(w, h int) (*Display, error) {
	d := C.al_create_display(C.int(w), C.int(h))
	if d == nil {
		return nil, errors.New("failed to create display!")
	}
	display := (*Display)(d)
	runtime.SetFinalizer(display, display.Destroy)
	return display, nil
}

func FlipDisplay() {
	C.al_flip_display()
}

func UpdateDisplayRegion(x, y, width, height int) {
	C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

func NewDisplayFlags() DisplayFlags {
	return DisplayFlags(C.al_get_new_display_flags())
}

func NewDisplayOption(option DisplayOption) (int, Importance) {
	var im C.int
	result := C.al_get_new_display_option(C.int(option), &im)
	return int(result), (Importance)(im)
}

func SetNewDisplayFlags(flags DisplayFlags) {
	C.al_set_new_display_flags(C.int(flags))
}

func SetNewDisplayOption(option DisplayOption, value int, im Importance) {
	C.al_set_new_display_option(C.int(option), C.int(value), C.int(im))
}

func ResetNewDisplayOptions() {
	C.al_reset_new_display_options()
}

func NewDisplayRefreshRate() int {
	return int(C.al_get_new_display_refresh_rate())
}

func SetNewDisplayRefreshRate(rate int) {
	C.al_set_new_display_refresh_rate(C.int(rate))
}

func NewWindowPosition() (int, int) {
	var x, y C.int
	C.al_get_new_window_position(&x, &y)
	return int(x), int(y)
}

func SetNewWindowPosition(x, y int) {
	C.al_set_new_window_position(C.int(x), C.int(y))
}

func ResetNewWindowPosition() {
	C.al_set_new_window_position(C.INT_MAX, C.INT_MAX)
}

func ResetDisplayFlags() {
	C.al_set_new_display_flags(C.int(0))
}

func InhibitScreensaver(inhibit bool) error {
	success := bool(C.al_inhibit_screensaver(C.bool(inhibit)))
	if !success {
		return errors.New("failed to inhibit screensaver!")
	}
	return nil
}

func WaitForVSync() error {
	success := bool(C.al_wait_for_vsync())
	if !success {
		return errors.New("cannot wait for vsync!")
	}
	return nil
}

// updates the display mode
func (mode *DisplayMode) Get(index int) error {
	result := C.al_get_display_mode(C.int(index), &mode.ptr)
	if result == nil {
		return fmt.Errorf("error getting display mode '%d'", index)
	}
	mode.Width = int(result.width)
	mode.Height = int(result.height)
	mode.Format = int(result.format)
	mode.RefreshRate = int(result.refresh_rate)
	return nil
}

func NumDisplayModes() int {
	return int(C.al_get_num_display_modes())
}

func NewDisplayAdapter() int {
	return int(C.al_get_new_display_adapter())
}

func SetNewDisplayAdapter(adapter int) {
	C.al_set_new_display_adapter(C.int(adapter))
}

func (info *MonitorInfo) Get(adapter int) error {
	success := bool(C.al_get_monitor_info(C.int(adapter), &info.ptr))
	if !success {
		return fmt.Errorf("error getting monitor info for adapter '%d'", adapter)
	}
	info.X1 = int(info.ptr.x1)
	info.X2 = int(info.ptr.x2)
	info.Y1 = int(info.ptr.y1)
	info.Y2 = int(info.ptr.y2)
	return nil
}

func NumVideoAdapters() int {
	return int(C.al_get_num_video_adapters())
}

// Display Instance Methods {{{

func (d *Display) Destroy() {
	C.al_destroy_display((*C.ALLEGRO_DISPLAY)(d))
}

func (d *Display) SetDisplayFlag(flags DisplayFlags, onoff bool) error {
	success := bool(C.al_set_display_flag((*C.ALLEGRO_DISPLAY)(d), C.int(flags), C.bool(onoff)))
	if !success {
		return errors.New("failed to set display flag!")
	}
	return nil
}

func (d *Display) DisplayOption(option DisplayOption) int {
	return int(C.al_get_display_option((*C.ALLEGRO_DISPLAY)(d), C.int(option)))
}

func (d *Display) RefreshRate() int {
	return int(C.al_get_display_refresh_rate((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) WindowPosition() (int, int) {
	var x, y C.int
	C.al_get_window_position((*C.ALLEGRO_DISPLAY)(d), &x, &y)
	return int(x), int(y)
}

func (d *Display) SetWindowPosition(x, y int) {
	C.al_set_window_position((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y))
}

func (d *Display) EventSource() *EventSource {
	return (*EventSource)(C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) Width() int {
	return int(C.al_get_display_width((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) Height() int {
	return int(C.al_get_display_height((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) AcknowledgeResize() bool {
	return bool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) SetWindowTitle(title string) {
	title_ := C.CString(title)
	defer freeString(title_)
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(d), title_)
}

func (d *Display) Backbuffer() *Bitmap {
	return (*Bitmap)(C.al_get_backbuffer((*C.ALLEGRO_DISPLAY)(d)))
}

func (d *Display) Resize(width, height int) error {
	success := bool(C.al_resize_display((*C.ALLEGRO_DISPLAY)(d), C.int(width), C.int(height)))
	if !success {
		return errors.New("failed to resize display!")
	}
	return nil
}

func (d *Display) SetDisplayIcon(icon *Bitmap) {
	C.al_set_display_icon((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_BITMAP)(icon))
}

func (d *Display) DisplayFormat() PixelFormat {
	return PixelFormat(C.al_get_display_format((*C.ALLEGRO_DISPLAY)(d)))
}

//}}}
