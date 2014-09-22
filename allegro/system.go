/*
Package allegro provides bindings to the core functionality of the
Allegro 5 API.

In order to improve the readability of this API, many methods are annotated
with information pulled directly from Allegro's online C documentation.
Note that this means that method names and some other information may be
C-specific, but otherwise should still be useful.

A bare-bones program might look something like this:

    package main

    import (
    	"github.com/dradtke/go-allegro/allegro"
    )

    const FPS int = 60

    func main() {
    	var (
    		display    *allegro.Display
    		eventQueue *allegro.EventQueue
    		running    bool = true
    		err        error
    	)

    	allegro.Run(func() {
    		// Create a 640x480 window and give it a title.
    		allegro.SetNewDisplayFlags(allegro.WINDOWED)
    		if display, err = allegro.CreateDisplay(640, 480); err == nil {
    			defer display.Destroy()
    			display.SetWindowTitle("Hello World")
    		} else {
    			panic(err)
    		}

    		// Create an event queue. All of the event sources we care about should
    		// register themselves to this queue.
    		if eventQueue, err = allegro.CreateEventQueue(); err == nil {
    			defer eventQueue.Destroy()
    		} else {
    			panic(err)
    		}

    		// Calculate the timeout value based on the desired FPS.
    		timeout := float64(1) / float64(FPS)

    		// Register event sources.
    		eventQueue.Register(display)

    		// Set the screen to black.
    		allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
    		allegro.FlipDisplay()

    		// Main loop.
    		var event allegro.Event
    		for {
    			if e, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout), &event); found {
    				switch e.(type) {
    				case allegro.DisplayCloseEvent:
    					running = false
    					break

    					// Handle other events here.
    				}
    			}

    			if !running {
    				return
    			}
    		}
    	})
    }
*/
package allegro

// #include <allegro5/allegro.h>
/*
bool _al_init() {
	return al_init();
}
*/
import "C"
import (
	"errors"
)

// Returns the (compiled) version of the Allegro library, packed into a single
// integer as groups of 8 bits in the form (major << 24) | (minor << 16) |
// (revision << 8) | release.
func Version() (major, minor, revision, release uint8) {
	v := uint32(C.al_get_allegro_version())
	major = uint8(v >> 24)
	minor = uint8((v >> 16) & 255)
	revision = uint8((v >> 8) & 255)
	release = uint8(v & 255)
	return
}

// Returns the current system configuration structure, or NULL if there is no
// active system driver. The returned configuration should not be destroyed
// with al_destroy_config. This is mainly used for configuring Allegro and its
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

func install() error {
	if !bool(C._al_init()) {
		return errors.New("failed to initialize allegro!")
	}
	return nil
}

// Closes down the Allegro system.
func uninstall() {
	C.al_uninstall_system()
}
