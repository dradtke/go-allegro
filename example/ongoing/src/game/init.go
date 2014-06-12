package game

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/example/ongoing/src/config"
)

// Init() initializes the game by creating the event queue, installing
// input systems, creating the display, and starting the FPS timer.
func Init() {
	var err error

	// Create an event queue. All of the event sources we care about should
	// register themselves to this queue.
	if eventQueue, err = al.CreateEventQueue(); err != nil {
		Fatal(err)
	}

	// Install the Keyboard driver.
	if err = al.InstallKeyboard(); err != nil {
		Fatal(err)
	}
	if keyboard, err := al.KeyboardEventSource(); err != nil {
		Fatal(err)
	} else {
		eventQueue.RegisterEventSource(keyboard)
	}

	// Create a 640x480 window and give it a title.
	al.SetNewDisplayFlags(al.WINDOWED)
	if display, err = al.CreateDisplay(config.DisplayWidth(), config.DisplayHeight()); err != nil {
		Fatal(err)
	}
	display.SetWindowTitle(config.GameName())
	eventQueue.Register(display)

	// Create the FPS timer.
	if fpsTimer, err = al.CreateTimer(1.0 / float64(FPS)); err != nil {
		Fatal(err)
	}
	eventQueue.Register(fpsTimer)
	fpsTimer.Start()
}

// Cleanup() destroys some common resources.
func Cleanup() {
	if fpsTimer != nil {
		fpsTimer.Destroy()
	}

	if display != nil {
        // TODO: investigate why it sometimes takes forever to close the display
		display.Destroy()
	}

	if eventQueue != nil {
		eventQueue.Destroy()
	}
}
