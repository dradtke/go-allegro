package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Display C.ALLEGRO_DISPLAY

type DisplayFlags int

const (
	WINDOWED                  DisplayFlags = C.ALLEGRO_WINDOWED
	FULLSCREEN                             = C.ALLEGRO_FULLSCREEN
	FULLSCREEN_WINDOW                      = C.ALLEGRO_FULLSCREEN_WINDOW
	RESIZABLE                              = C.ALLEGRO_RESIZABLE
	OPENGL                                 = C.ALLEGRO_OPENGL
	OPENGL_3_0                             = C.ALLEGRO_OPENGL_3_0
	OPENGL_FORWARD_COMPATIBLE              = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
	FRAMELESS                              = C.ALLEGRO_FRAMELESS
	NOFRAME                                = C.ALLEGRO_NOFRAME
	GENERATE_EXPOSE_EVENTS                 = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
)

type DisplayMode C.struct_ALLEGRO_DISPLAY_MODE

type DisplayOption C.int

const (
	COLOR_SIZE             DisplayOption = C.ALLEGRO_COLOR_SIZE
	RED_SIZE                             = C.ALLEGRO_RED_SIZE
	GREEN_SIZE                           = C.ALLEGRO_GREEN_SIZE
	BLUE_SIZE                            = C.ALLEGRO_BLUE_SIZE
	ALPHA_SIZE                           = C.ALLEGRO_ALPHA_SIZE
	RED_SHIFT                            = C.ALLEGRO_RED_SHIFT
	GREEN_SHIFT                          = C.ALLEGRO_GREEN_SHIFT
	BLUE_SHIFT                           = C.ALLEGRO_BLUE_SHIFT
	ALPHA_SHIFT                          = C.ALLEGRO_ALPHA_SHIFT
	ACC_RED_SIZE                         = C.ALLEGRO_ACC_RED_SIZE
	ACC_GREEN_SIZE                       = C.ALLEGRO_ACC_GREEN_SIZE
	ACC_BLUE_SIZE                        = C.ALLEGRO_ACC_BLUE_SIZE
	ACC_ALPHA_SIZE                       = C.ALLEGRO_ACC_ALPHA_SIZE
	STEREO                               = C.ALLEGRO_STEREO
	AUX_BUFFERS                          = C.ALLEGRO_AUX_BUFFERS
	DEPTH_SIZE                           = C.ALLEGRO_DEPTH_SIZE
	STENCIL_SIZE                         = C.ALLEGRO_STENCIL_SIZE
	SAMPLE_BUFFERS                       = C.ALLEGRO_SAMPLE_BUFFERS
	SAMPLES                              = C.ALLEGRO_SAMPLES
	RENDER_METHOD                        = C.ALLEGRO_RENDER_METHOD
	FLOAT_COLOR                          = C.ALLEGRO_FLOAT_COLOR
	FLOAT_DEPTH                          = C.ALLEGRO_FLOAT_DEPTH
	SINGLE_BUFFER                        = C.ALLEGRO_SINGLE_BUFFER
	SWAP_METHOD                          = C.ALLEGRO_SWAP_METHOD
	COMPATIBLE_DISPLAY                   = C.ALLEGRO_COMPATIBLE_DISPLAY
	UPDATE_DISPLAY_REGION                = C.ALLEGRO_UPDATE_DISPLAY_REGION
	VSYNC                                = C.ALLEGRO_VSYNC
	MAX_BITMAP_SIZE                      = C.ALLEGRO_MAX_BITMAP_SIZE
	SUPPORT_NPOT_BITMAP                  = C.ALLEGRO_SUPPORT_NPOT_BITMAP
	CAN_DRAW_INTO_BITMAP                 = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
	SUPPORT_SEPARATE_ALPHA               = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
)

type Importance C.int

const (
	REQUIRE  Importance = C.ALLEGRO_REQUIRE
	SUGGEST             = C.ALLEGRO_SUGGEST
	DONTCARE            = C.ALLEGRO_DONTCARE
)

type DisplayOrientation C.int

const (
	DISPLAY_ORIENTATION_0_DEGREES   DisplayOrientation = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
	DISPLAY_ORIENTATION_90_DEGREES                     = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
	DISPLAY_ORIENTATION_180_DEGREES                    = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
	DISPLAY_ORIENTATION_270_DEGREES                    = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
	DISPLAY_ORIENTATION_FACE_UP                        = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
	DISPLAY_ORIENTATION_FACE_DOWN                      = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

// Create a display, or window, with the specified dimensions. The parameters
// of the display are determined by the last calls to al_set_new_display_*.
// Default parameters are used if none are set explicitly. Creating a new
// display will automatically make it the active one, with the backbuffer
// selected for drawing.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_create_display
func CreateDisplay(w, h int) (*Display, error) {
	d := C.al_create_display(C.int(w), C.int(h))
	if d == nil {
		return nil, errors.New("failed to create display!")
	}
	display := (*Display)(d)
	//runtime.SetFinalizer(display, func(d_ *Display) { d_.Destroy() })
	return display, nil
}

// Copies or updates the front and back buffers so that what has been drawn
// previously on the currently selected display becomes visible on screen.
// Pointers to the special back buffer bitmap remain valid and retain their
// semantics as the back buffer, although the contents may have changed.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_flip_display
func FlipDisplay() {
	C.al_flip_display()
}

// Does the same as al_flip_display, but tries to update only the specified
// region. With many drivers this is not possible, but for some it can improve
// performance. If this is not supported, this function falls back to the
// behavior of al_flip_display. You can query the support for this function
// using al_get_display_option(display, ALLEGRO_UPDATE_DISPLAY_REGION).
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_update_display_region
func UpdateDisplayRegion(x, y, width, height int) {
	C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

// Get the display flags to be used when creating new displays on the calling
// thread.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_new_display_flags
func NewDisplayFlags() DisplayFlags {
	return DisplayFlags(C.al_get_new_display_flags())
}

// Retrieve an extra display setting which was previously set with
// al_set_new_display_option.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_new_display_option
func NewDisplayOption(option DisplayOption) (int, Importance) {
	var im C.int
	result := C.al_get_new_display_option(C.int(option), &im)
	return int(result), (Importance)(im)
}

// Sets various flags to be used when creating new displays on the calling
// thread. flags is a bitfield containing any reasonable combination of the
// following:
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_new_display_flags
func SetNewDisplayFlags(flags DisplayFlags) {
	C.al_set_new_display_flags(C.int(flags))
}

// Set an extra display option, to be used when creating new displays on the
// calling thread. Display options differ from display flags, and specify some
// details of the context to be created within the window itself. These mainly
// have no effect on Allegro itself, but you may want to specify them, for
// example if you want to use multisampling.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_new_display_option
func SetNewDisplayOption(option DisplayOption, value int, im Importance) {
	C.al_set_new_display_option(C.int(option), C.int(value), C.int(im))
}

// This undoes any previous call to al_set_new_display_option on the calling
// thread.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_reset_new_display_options
func ResetNewDisplayOptions() {
	C.al_reset_new_display_options()
}

// Get the requested refresh rate to be used when creating new displays on the
// calling thread.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_new_display_refresh_rate
func NewDisplayRefreshRate() int {
	return int(C.al_get_new_display_refresh_rate())
}

// Sets the refresh rate to use when creating new displays on the calling
// thread. If the refresh rate is not available, al_create_display will fail. A
// list of modes with refresh rates can be found with al_get_num_display_modes
// and al_get_display_mode.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_new_display_refresh_rate
func SetNewDisplayRefreshRate(rate int) {
	C.al_set_new_display_refresh_rate(C.int(rate))
}

// Get the position where new non-fullscreen displays created by the calling
// thread will be placed.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_new_window_position
func NewWindowPosition() (int, int) {
	var x, y C.int
	C.al_get_new_window_position(&x, &y)
	return int(x), int(y)
}

// Sets where the top left pixel of the client area of newly created windows
// (non-fullscreen) will be on screen, for displays created by the calling
// thread. Negative values are allowed on some multihead systems.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_new_window_position
func SetNewWindowPosition(x, y int) {
	C.al_set_new_window_position(C.int(x), C.int(y))
}

// Returns the title that will be used when a new display is created. This
// returns the value that al_set_window_title was called with. If that function
// wasn't called yet, the value of al_get_app_name is returned as a default.
// The current implementation returns a pointer to a static buffer of which you
// should make a copy if you want to modify it.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_new_window_title
func NewWindowTitle() string {
	return C.GoString(C.al_get_new_window_title())
}

// Set the title that will be used when a new display is created. Allegro uses
// a static buffer of ALLEGRO_NEW_WINDOW_TITLE_MAX_SIZE to store this, so the
// length of the titme you set must be less than this.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_new_window_title
func SetNewWindowTitle(title string) {
	if len(title) > C.ALLEGRO_NEW_WINDOW_TITLE_MAX_SIZE {
		panic("window title too long: " + title)
	}
	title_ := C.CString(title)
	defer freeString(title_)
	C.al_set_new_window_title(title_)
}

func ResetNewWindowPosition() {
	C.al_set_new_window_position(C.INT_MAX, C.INT_MAX)
}

func ResetDisplayFlags() {
	C.al_set_new_display_flags(C.int(0))
}

// This function allows the user to stop the system screensaver from starting
// up if true is passed, or resets the system back to the default state (the
// state at program start) if false is passed. It returns true if the state was
// set successfully, otherwise false.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_inhibit_screensaver
func InhibitScreensaver(inhibit bool) error {
	success := bool(C.al_inhibit_screensaver(C.bool(inhibit)))
	if !success {
		return errors.New("failed to inhibit screensaver!")
	}
	return nil
}

// Wait for the beginning of a vertical retrace. Some driver/card/monitor
// combinations may not be capable of this.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_wait_for_vsync
func WaitForVSync() error {
	success := bool(C.al_wait_for_vsync())
	if !success {
		return errors.New("cannot wait for vsync!")
	}
	return nil
}

// Retrieves a fullscreen mode. Display parameters should not be changed
// between a call of al_get_num_display_modes and al_get_display_mode. index
// must be between 0 and the number returned from al_get_num_display_modes-1.
// mode must be an allocated ALLEGRO_DISPLAY_MODE structure. This function will
// return NULL on failure, and the mode parameter that was passed in on success.
//
// See https://liballeg.org/a5docs/5.2.6/fullscreen_mode.html#al_get_display_mode
func GetDisplayMode(index int) (*DisplayMode, error) {
	var mode C.struct_ALLEGRO_DISPLAY_MODE
	result := C.al_get_display_mode(C.int(index), &mode)
	if result == nil {
		return nil, fmt.Errorf("error getting display mode '%d'", index)
	}
	return (*DisplayMode)(&mode), nil
}

// Screen width.
func (m *DisplayMode) Width() int {
	return int(m.width)
}

// Screen height.
func (m *DisplayMode) Height() int {
	return int(m.height)
}

// Pixel format.
func (m *DisplayMode) Format() PixelFormat {
	return PixelFormat(m.format)
}

// Refresh rate. May be 0 if unknown.
func (m *DisplayMode) RefreshRate() int {
	return int(m.refresh_rate)
}

// Get the number of available fullscreen display modes for the current set of
// display parameters. This will use the values set with
// al_set_new_display_refresh_rate, and al_set_new_display_flags to find the
// number of modes that match. Settings the new display parameters to zero will
// give a list of all modes for the default driver.
//
// See https://liballeg.org/a5docs/5.2.6/fullscreen_mode.html#al_get_num_display_modes
func NumDisplayModes() int {
	return int(C.al_get_num_display_modes())
}

// Display Instance Methods {{{

// Destroy a display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_destroy_display
func (d *Display) Destroy() {
	C.al_destroy_display((*C.ALLEGRO_DISPLAY)(d))
}

// Enable or disable one of the display flags. The flags are the same as for
// al_set_new_display_flags. The only flags that can be changed after creation
// are:
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_display_flag
func (d *Display) SetDisplayFlag(flags DisplayFlags, onoff bool) error {
	success := bool(C.al_set_display_flag((*C.ALLEGRO_DISPLAY)(d), C.int(flags), C.bool(onoff)))
	if !success {
		return errors.New("failed to set display flag!")
	}
	return nil
}

// Return an extra display setting of the display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_option
func (d *Display) DisplayOption(option DisplayOption) int {
	return int(C.al_get_display_option((*C.ALLEGRO_DISPLAY)(d), C.int(option)))
}

// Gets the refresh rate of the display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_refresh_rate
func (d *Display) RefreshRate() int {
	return int(C.al_get_display_refresh_rate((*C.ALLEGRO_DISPLAY)(d)))
}

// Gets the position of a non-fullscreen display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_window_position
func (d *Display) WindowPosition() (int, int) {
	var x, y C.int
	C.al_get_window_position((*C.ALLEGRO_DISPLAY)(d), &x, &y)
	return int(x), int(y)
}

// Sets the position on screen of a non-fullscreen display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_window_position
func (d *Display) SetWindowPosition(x, y int) {
	C.al_set_window_position((*C.ALLEGRO_DISPLAY)(d), C.int(x), C.int(y))
}

// Retrieve the associated event source. See the documentation on events for a
// list of the events displays will generate.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_event_source
func (d *Display) EventSource() *EventSource {
	return (*EventSource)(C.al_get_display_event_source((*C.ALLEGRO_DISPLAY)(d)))
}

// Gets the width of the display. This is like SCREEN_W in Allegro 4.x.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_width
func (d *Display) Width() int {
	return int(C.al_get_display_width((*C.ALLEGRO_DISPLAY)(d)))
}

// Gets the height of the display. This is like SCREEN_H in Allegro 4.x.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_height
func (d *Display) Height() int {
	return int(C.al_get_display_height((*C.ALLEGRO_DISPLAY)(d)))
}

// When the user receives a resize event from a resizable display, if they wish
// the display to be resized they must call this function to let the graphics
// driver know that it can now resize the display. Returns true on success.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_acknowledge_resize
func (d *Display) AcknowledgeResize() bool {
	return bool(C.al_acknowledge_resize((*C.ALLEGRO_DISPLAY)(d)))
}

// Set the title on a display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_window_title
func (d *Display) SetWindowTitle(title string) {
	title_ := C.CString(title)
	defer freeString(title_)
	C.al_set_window_title((*C.ALLEGRO_DISPLAY)(d), title_)
}

// Return a special bitmap representing the back-buffer of the display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_backbuffer
func (d *Display) Backbuffer() *Bitmap {
	return (*Bitmap)(C.al_get_backbuffer((*C.ALLEGRO_DISPLAY)(d)))
}

// Resize the display. Returns true on success, or false on error. This works
// on both fullscreen and windowed displays, regardless of the
// ALLEGRO_RESIZABLE flag.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_resize_display
func (d *Display) Resize(width, height int) error {
	success := bool(C.al_resize_display((*C.ALLEGRO_DISPLAY)(d), C.int(width), C.int(height)))
	if !success {
		return errors.New("failed to resize display!")
	}
	return nil
}

// Changes the icon associated with the display (window). Same as
// al_set_display_icons with one icon.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_display_icon
func (d *Display) SetDisplayIcon(icon *Bitmap) {
	C.al_set_display_icon((*C.ALLEGRO_DISPLAY)(d), (*C.ALLEGRO_BITMAP)(icon))
}

// Changes the icons associated with the display (window). Multiple icons can
// be provided for use in different contexts, e.g. window frame, taskbar,
// alt-tab popup. The number of icons must be at least one.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_display_icons
func (d *Display) SetDisplayIcons(icons []*Bitmap) {
	n_icons := len(icons)
	icons_ := make([]*C.ALLEGRO_BITMAP, n_icons)
	for i := 0; i < n_icons; i++ {
		icons_[i] = (*C.ALLEGRO_BITMAP)(icons[i])
	}
	C.al_set_display_icons((*C.ALLEGRO_DISPLAY)(d), C.int(n_icons), (**C.ALLEGRO_BITMAP)(unsafe.Pointer(&icons_[0])))
}

// Gets the pixel format of the display.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_display_format
func (d *Display) DisplayFormat() PixelFormat {
	return PixelFormat(C.al_get_display_format((*C.ALLEGRO_DISPLAY)(d)))
}

// This function returns a pointer to a string, allocated with al_malloc with
// the text contents of the clipboard if available. If no text is available on
// the clipboard then this function returns NULL. You must call al_free on the
// returned pointer when you don't need it anymore.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_get_clipboard_text
func (d *Display) ClipboardText() string {
	text := C.al_get_clipboard_text((*C.ALLEGRO_DISPLAY)(d))
	defer freeString(text)
	return C.GoString(text)
}

// This function pastes the text given as an argument to the clipboard.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_set_clipboard_text
func (d *Display) SetClipboardText(text string) {
	text_ := C.CString(text)
	defer freeString(text_)
	C.al_set_clipboard_text((*C.ALLEGRO_DISPLAY)(d), text_)
}

// This function returns true if and only if the clipboard has text available.
//
// See https://liballeg.org/a5docs/5.2.6/display.html#al_clipboard_has_text
func (d *Display) ClipboardHasText() bool {
	return bool(C.al_clipboard_has_text((*C.ALLEGRO_DISPLAY)(d)))
}

//}}}
