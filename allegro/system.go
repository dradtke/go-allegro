// Package allegro provides bindings to the core functionality of the
// Allegro 5 API. A bare-bones program might look something like this:
//
//     package main
//
//     import (
//     	"github.com/dradtke/go-allegro/allegro"
//     )
//
//     const FPS int = 60
//
//     func main() {
//     	var (
//     		display    *allegro.Display
//     		eventQueue *allegro.EventQueue
//     		running    bool = true
//     		err        error
//     	)
//
//     	// Create a 640x480 window and give it a title.
//     	allegro.SetNewDisplayFlags(allegro.WINDOWED)
//     	if display, err = allegro.CreateDisplay(640, 480); err == nil {
//     		defer display.Destroy()
//     		display.SetWindowTitle("Hello World")
//     	} else {
//     		panic(err)
//     	}
//
//     	// Create an event queue. All of the event sources we care about should
//     	// register themselves to this queue.
//     	if eventQueue, err = allegro.CreateEventQueue(); err == nil {
//     		defer eventQueue.Destroy()
//     	} else {
//     		panic(err)
//     	}
//
//     	// Calculate the timeout value based on the desired FPS.
//     	timeout := float64(1) / float64(FPS)
//
//     	// Register event sources.
//     	eventQueue.RegisterEventSource(display.EventSource())
//
//     	// Set the screen to black.
//     	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
//     	allegro.FlipDisplay()
//
//     	// Main loop.
//     	for {
//     		event, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout))
//     		if found {
//     			switch event.Type {
//     				case allegro.EVENT_DISPLAY_CLOSE:
//     					running = false
//     					break
//
//     				// Handle other events here.
//     			}
//     		}
//
//     		if !running {
//     			return
//     		}
//     	  }
//     }
//
package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

bool _al_init() {
	return al_init();
}
*/
import "C"
import (
	"errors"
)

func init() {
	if !bool(C._al_init()) {
		panic("failed to initialize allegro!")
	}
}

// Returns the (compiled) version of the Allegro library, packed into a single
// integer as groups of 8 bits in the form (major << 24) | (minor << 16) |
// (revision << 8) | release.
func Version() uint32 {
	return uint32(C.al_get_allegro_version())
}

// Returns the current system configuration structure, or NULL if there is no
// active system driver. This is mainly used for configuring Allegro and its
// addons.
func SystemConfig() (*Config, error) {
	cfg := C.al_get_system_config()
	if cfg == nil {
		return nil, errors.New("no system config found")
	}
	return (*Config)(cfg), nil
}

// This override the executable name used by al_get_standard_path for
// ALLEGRO_EXENAME_PATH and ALLEGRO_RESOURCES_PATH.
func SetExeName(path string) {
	path_ := C.CString(path)
	defer freeString(path_)
	C.al_set_exe_name(path_)
}

// Sets the global organization name.
func SetOrgName(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_set_org_name(name_)
}

// Sets the global application name.
func SetAppName(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_set_app_name(name_)
}

// Returns the global organization name string.
func OrgName() string {
	return C.GoString(C.al_get_org_name())
}

// Returns the global application name string.
func AppName() string {
	return C.GoString(C.al_get_app_name())
}

