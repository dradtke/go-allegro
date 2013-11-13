package dialog

/*
#cgo pkg-config: allegro_dialog-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_native_dialog.h>

void _al_free_string(char *data) {
	al_free(data);
}

void _al_append_native_text_log(ALLEGRO_TEXTLOG *log, char *str) {
	al_append_native_text_log(log, str);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"runtime"
	"strings"
	"unsafe"
)

type FileChooser C.ALLEGRO_FILECHOOSER
type TextLog C.ALLEGRO_TEXTLOG

type FileChooserFlags int

const (
	FILECHOOSER_FILE_MUST_EXIST FileChooserFlags = C.ALLEGRO_FILECHOOSER_FILE_MUST_EXIST
	FILECHOOSER_SAVE            FileChooserFlags = C.ALLEGRO_FILECHOOSER_SAVE
	FILECHOOSER_FOLDER          FileChooserFlags = C.ALLEGRO_FILECHOOSER_FOLDER
	FILECHOOSER_PICTURES        FileChooserFlags = C.ALLEGRO_FILECHOOSER_PICTURES
	FILECHOOSER_SHOW_HIDDEN     FileChooserFlags = C.ALLEGRO_FILECHOOSER_SHOW_HIDDEN
	FILECHOOSER_MULTIPLE        FileChooserFlags = C.ALLEGRO_FILECHOOSER_MULTIPLE
)

type MessageBoxFlags int

const (
	MESSAGEBOX_WARN      MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_WARN
	MESSAGEBOX_ERROR     MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_ERROR
	MESSAGEBOX_QUESTION  MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_QUESTION
	MESSAGEBOX_OK_CANCEL MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_OK_CANCEL
	MESSAGEBOX_YES_NO    MessageBoxFlags = C.ALLEGRO_MESSAGEBOX_YES_NO
)

type MessageBoxResult int

const (
	RESPONSE_NONE      MessageBoxResult = 0
	RESPONSE_YES_OK    MessageBoxResult = 1
	RESPONSE_NO_CANCEL MessageBoxResult = 2
)

type TextLogFlags int

const (
	TEXTLOG_NO_CLOSE  TextLogFlags = C.ALLEGRO_TEXTLOG_NO_CLOSE
	TEXTLOG_MONOSPACE TextLogFlags = C.ALLEGRO_TEXTLOG_MONOSPACE
)

func Version() uint32 {
	return uint32(C.al_get_allegro_native_dialog_version())
}

func CreateNativeFileDialog(initial_path, title, patterns string, flags FileChooserFlags) (*FileChooser, error) {
	initial_path_ := C.CString(initial_path)
	title_ := C.CString(title)
	patterns_ := C.CString(patterns)
	defer C._al_free_string(initial_path_)
	defer C._al_free_string(title_)
	defer C._al_free_string(patterns_)
	d := C.al_create_native_file_dialog(initial_path_, title_, patterns_, C.int(flags))
	if d == nil {
		return nil, errors.New("failed to create native file chooser dialog")
	}
	dialog := (*FileChooser)(d)
	runtime.SetFinalizer(dialog, dialog.Destroy)
	return dialog, nil
}

func ShowNativeFileDialog(display *allegro.Display, dialog *FileChooser) error {
	ok := bool(C.al_show_native_file_dialog((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		(*C.ALLEGRO_FILECHOOSER)(dialog)))
	if !ok {
		return errors.New("failed to show native file dialog")
	}
	return nil
}

func ShowNativeMessageBox(display *allegro.Display, title, heading, text string, flags MessageBoxFlags) MessageBoxResult {
	title_ := C.CString(title)
	heading_ := C.CString(heading)
	text_ := C.CString(text)
	defer C._al_free_string(title_)
	defer C._al_free_string(heading_)
	defer C._al_free_string(text_)
	res := C.al_show_native_message_box((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		title_, heading_, text_, nil, C.int(flags))
	return MessageBoxResult(res)
}

func ShowNativeMessageBoxWithButtons(display *allegro.Display, title, heading, text string, buttons []string, flags MessageBoxFlags) string {
	title_ := C.CString(title)
	heading_ := C.CString(heading)
	text_ := C.CString(text)
	buttons_ := C.CString(strings.Join(buttons, "|"))
	defer C._al_free_string(title_)
	defer C._al_free_string(heading_)
	defer C._al_free_string(text_)
	defer C._al_free_string(buttons_)
	res := int(C.al_show_native_message_box((*C.ALLEGRO_DISPLAY)(unsafe.Pointer(display)),
		title_, heading_, text_, buttons_, C.int(flags))) - 1

	if res > -1 && res < len(buttons) {
		return buttons[res]
	} else {
		return ""
	}
}

func OpenNativeTextLog(title string, flags TextLogFlags) (*TextLog, error) {
	title_ := C.CString(title)
	defer C._al_free_string(title_)
	l := C.al_open_native_text_log(title_, C.int(flags))
	if l == nil {
		return nil, errors.New("failed to open native text log")
	}
	log := (*TextLog)(l)
	return log, nil
}

func (dialog *FileChooser) Count() int {
	return int(C.al_get_native_file_dialog_count((*C.ALLEGRO_FILECHOOSER)(dialog)))
}

func (dialog *FileChooser) Path(i int) (string, error) {
	path := C.al_get_native_file_dialog_path((*C.ALLEGRO_FILECHOOSER)(dialog), C.size_t(i))
	if path == nil {
		return "", fmt.Errorf("failed to get path %d from dialog", i)
	}
	return C.GoString(path), nil
}

func (dialog *FileChooser) Destroy() {
	C.al_destroy_native_file_dialog((*C.ALLEGRO_FILECHOOSER)(dialog))
}

func (log *TextLog) Close() {
	C.al_close_native_text_log((*C.ALLEGRO_TEXTLOG)(log))
}

func (log *TextLog) Append(format string, a ...interface{}) {
	var text string
	if len(a) == 0 {
		text = format + "\n"
	} else {
		text = fmt.Sprintf(format + "\n", a)
	}
	text_ := C.CString(text)
	defer C._al_free_string(text_)
	// C.al_append_native_text_log()
	C._al_append_native_text_log((*C.ALLEGRO_TEXTLOG)(log), text_)
}

func (log *TextLog) EventSource() (*allegro.EventSource) {
	return (*allegro.EventSource)(unsafe.Pointer(
		C.al_get_native_text_log_event_source((*C.ALLEGRO_TEXTLOG)(log))))
}
