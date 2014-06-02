package main

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/font"
	prim "github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/dradtke/go-allegro/example/ongoing/config"
)

const (
	FPS int = 60
)

func render() {
	al.HoldBitmapDrawing(true)
	defer al.HoldBitmapDrawing(false)

	renderConsole()
}

func main() {
	if err := al.Install(); err != nil {
		panic(err)
	}
	defer al.Uninstall()

	if err := prim.Install(); err != nil {
		panic(err)
	}
	defer prim.Uninstall()

	font.Install()
	defer font.Uninstall()

	var (
		display    *al.Display
		eventQueue *al.EventQueue
		keyboard   *al.EventSource
		fpsTimer   *al.Timer

		event  al.Event
		redraw bool  = false
		err    error = nil
	)

	// Create an event queue. All of the event sources we care about should
	// register themselves to this queue.
	if eventQueue, err = al.CreateEventQueue(); err != nil {
		panic(err)
	}
	defer eventQueue.Destroy()

	// Install the Keyboard driver.
	if err = al.InstallKeyboard(); err != nil {
		panic(err)
	}
	if keyboard, err = al.KeyboardEventSource(); err != nil {
		panic(err)
	}
	eventQueue.RegisterEventSource(keyboard)

	// Create a 640x480 window and give it a title.
	al.SetNewDisplayFlags(al.WINDOWED)
	if display, err = al.CreateDisplay(config.DisplayWidth(), config.DisplayHeight()); err != nil {
		panic(err)
	}
	defer display.Destroy()
	display.SetWindowTitle("Hello World")
	eventQueue.Register(display)

	// Initialize subsystems.
	initConsole(eventQueue)

	// Set the screen to black.
	al.ClearToColor(config.BlankColor())
	al.FlipDisplay()

	// Create the FPS timer.
	if fpsTimer, err = al.CreateTimer(1.0 / float64(FPS)); err != nil {
		panic(err)
	}
	defer fpsTimer.Destroy()
	eventQueue.Register(fpsTimer)
	fpsTimer.Start()

	// Main loop.
mainLoop:
	for {
		ev := eventQueue.WaitForEvent(&event)
		if consoleHandled(ev) {
            // TODO: after this, check to see if any console events need to be processed,
            // then act accordingly (e.g. updating blank_color)
			goto eventHandled
		}

		switch e := ev.(type) {
		case al.TimerEvent:
			if e.Source() == fpsTimer {
				redraw = true
			}

		case al.KeyDownEvent:
			// TODO: handle

		case al.DisplayCloseEvent:
			break mainLoop
		}

	eventHandled:
		if redraw && eventQueue.IsEmpty() {
			al.ClearToColor(config.BlankColor())
			render()
			al.FlipDisplay()
			redraw = false
		}
	}

	display.SetWindowTitle("Shutting down...")
	saveConsole()
}
