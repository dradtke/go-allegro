// Package dialog provides support for Allegro's native dialog addon.
package dialog

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_native_dialog.h>
// #include "../util.c"
/*
static void append_to_log(ALLEGRO_TEXTLOG *log, char *str) {
	al_append_native_text_log(log, str);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	//"runtime"
	"strings"
	"unsafe"
)

type FileChooser C.ALLEGRO_FILECHOOSER
type TextLog C.ALLEGRO_TEXTLOG

type FileChooserFlags int

const (
	FILECHOOSER_FILE_MUST_EXIST FileChooserFlags = C.ALLEGRO_FILECHOOSER_FILE_MUST_EXIST
	FILECHOOSER_SAVE                             = C.ALLEGRO_FILECHOOSER_SAVE
	FILECHOOSER_FOLDER                           = C.ALLEGRO_FILECHOOSER_FOLDER
	FILECHOOSER_PICTURES                         = C.ALLEGRO_FILECHOOSER_PICTURES
	FILECHOOSER_SHOW_HIDDEN                      = C.ALLEGRO_FILECHOOSER_SHOW_HIDDEN
	FILECHOOSER_MULTIPLE                         = C.ALLEGRO_FILECHOOSER_MULTIPLE
)

type MessageBoxFlags int

const (
	MESSAGEBOX_WARN      MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_WARN
	MESSAGEBOX_ERROR                     = C.ALLEGRO_MESSAGEBOX_ERROR
	MESSAGEBOX_QUESTION                  = C.ALLEGRO_MESSAGEBOX_QUESTION
	MESSAGEBOX_OK_CANCEL                 = C.ALLEGRO_MESSAGEBOX_OK_CANCEL
	MESSAGEBOX_YES_NO                    = C.ALLEGRO_MESSAGEBOX_YES_NO
)

type MessageBoxResult int

const (
	RESPONSE_NONE      MessageBoxResult = 0
	RESPONSE_YES_OK                     = 1
	RESPONSE_NO_CANCEL                  = 2
)

type TextLogFlags int

const (
	TEXTLOG_NO_CLOSE  TextLogFlags = C.ALLEGRO_TEXTLOG_NO_CLOSE
	TEXTLOG_MONOSPACE              = C.ALLEGRO_TEXTLOG_MONOSPACE
)

// Initialise the native dialog addon.
func Install() error {
	if !bool(C.al_init_native_dialog_addon()) {
		return errors.New("failed to initialize native dialog addon!")
	}
	return nil
}

// Shut down the native dialog addon.
func Shutdown() {
	C.al_shutdown_native_dialog_addon()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_native_dialog_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}

// Creates a new native file dialog. You should only have one such dialog
// opened at a time.
func CreateNativeFileDialog(initial_path, title, patterns string, flags FileChooserFlags) (*FileChooser, error) {
	initial_path_ := C.CString(initial_path)
	title_ := C.CString(title)
	patterns_ := C.CString(patterns)
	defer C.free_string(initial_path_)
	defer C.free_string(title_)
	defer C.free_string(patterns_)
	d := C.al_create_native_file_dialog(initial_path_, title_, patterns_, C.int(flags))
	if d == nil {
		return nil, errors.New("failed to create native file chooser dialog")
	}
	dialog := (*FileChooser)(d)
	//runtime.SetFinalizer(dialog, dialog.Destroy)
	return dialog, nil
}

// Show the dialog window. The display may be NULL, otherwise the given display
// is treated as the parent if possible.
func ShowNativeFileDialog(display *allegro.Display, dialog *FileChooser) error {
	ok := bool(C.al_show_native_file_dialog((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		(*C.ALLEGRO_FILECHOOSER)(dialog)))
	if !ok {
		return errors.New("failed to show native file dialog")
	}
	return nil
}

// Show a native GUI message box. This can be used for example to display an
// error message if creation of an initial display fails. The display may be
// NULL, otherwise the given display is treated as the parent if possible.
func ShowNativeMessageBox(display *allegro.Display, title, heading, text string, flags MessageBoxFlags) MessageBoxResult {
	title_ := C.CString(title)
	heading_ := C.CString(heading)
	text_ := C.CString(text)
	defer C.free_string(title_)
	defer C.free_string(heading_)
	defer C.free_string(text_)
	res := C.al_show_native_message_box((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		title_, heading_, text_, nil, C.int(flags))
	return MessageBoxResult(res)
}

func ShowNativeMessageBoxWithButtons(display *allegro.Display, title, heading, text string, buttons []string, flags MessageBoxFlags) string {
	title_ := C.CString(title)
	heading_ := C.CString(heading)
	text_ := C.CString(text)
	buttons_ := C.CString(strings.Join(buttons, "|"))
	defer C.free_string(title_)
	defer C.free_string(heading_)
	defer C.free_string(text_)
	defer C.free_string(buttons_)
	res := int(C.al_show_native_message_box((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		title_, heading_, text_, buttons_, C.int(flags))) - 1

	if res > -1 && res < len(buttons) {
		return buttons[res]
	} else {
		return ""
	}
}

// Opens a window to which you can append log messages with
// al_append_native_text_log. This can be useful for debugging if you don't
// want to depend on a console being available.
func OpenNativeTextLog(title string, flags TextLogFlags) (*TextLog, error) {
	title_ := C.CString(title)
	defer C.free_string(title_)
	l := C.al_open_native_text_log(title_, C.int(flags))
	if l == nil {
		return nil, errors.New("failed to open native text log")
	}
	log := (*TextLog)(l)
	return log, nil
}

// Returns the number of files selected, or 0 if the dialog was cancelled.
func (dialog *FileChooser) Count() int {
	return int(C.al_get_native_file_dialog_count((*C.ALLEGRO_FILECHOOSER)(dialog)))
}

// Returns one of the selected paths.
func (dialog *FileChooser) Path(i int) (string, error) {
	path := C.al_get_native_file_dialog_path((*C.ALLEGRO_FILECHOOSER)(dialog), C.size_t(i))
	if path == nil {
		return "", fmt.Errorf("failed to get path %d from dialog", i)
	}
	return C.GoString(path), nil
}

// Frees up all resources used by the file dialog.
func (dialog *FileChooser) Destroy() {
	C.al_destroy_native_file_dialog((*C.ALLEGRO_FILECHOOSER)(dialog))
}

// Closes a message log window opened with al_open_native_text_log earlier.
func (log *TextLog) Close() {
	C.al_close_native_text_log((*C.ALLEGRO_TEXTLOG)(log))
}

// Appends a line of text to the message log window and scrolls to the bottom
// (if the line would not be visible otherwise). This works like printf. A line
// is continued until you add a newline character.
func (log *TextLog) Append(format string, a ...interface{}) {
	text_ := C.CString(fmt.Sprintf(format, a...))
	defer C.free_string(text_)
	// C.al_append_native_text_log()
	C.append_to_log((*C.ALLEGRO_TEXTLOG)(log), text_)
}

func (log *TextLog) Appendln(format string, a ...interface{}) {
	log.Append(format+"\n", a...)
}

// Get an event source for a text log window. The possible events are:
func (log *TextLog) EventSource() *allegro.EventSource {
	return (*allegro.EventSource)(unsafe.Pointer(
		C.al_get_native_text_log_event_source((*C.ALLEGRO_TEXTLOG)(log))))
}
