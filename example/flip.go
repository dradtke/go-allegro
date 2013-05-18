package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/image"
	"fmt"
	"os"
)

/*
// Without this, cgo complains about atexit not being defined.
// Linking explicitly against /usr/lib64/libc_nonshared.a might fix it
void atexit() {}
*/
import "C"

func main() {
	var (
		display *allegro.Display
		eventQueue *allegro.EventQueue
		keyboardState allegro.KeyboardState
		running bool = true
		windowTitle string = "Flip Me Around!"
	)

	if err := allegro.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if !allegro.InstallKeyboard() {
		fmt.Fprintf(os.Stderr, "failed to initialize keyboard\n")
		return
	}

	allegro.SetNewDisplayFlags(allegro.Windowed)
	if display = allegro.CreateDisplay(640, 480); display != nil {
		defer display.Destroy()
		display.SetWindowTitle(windowTitle)
	} else {
		fmt.Fprintf(os.Stderr, "failed to create display\n")
		return
	}

	if eventQueue = allegro.CreateEventQueue(); eventQueue != nil {
		defer eventQueue.Destroy()
	} else {
		fmt.Fprintf(os.Stderr, "failed to create event queue\n")
		return
	}

	image.Init()
	gordon := allegro.LoadBitmap("img/gordon-the-gopher.png") ; defer gordon.Destroy()

	eventQueue.RegisterEventSource(display.GetEventSource())
	eventQueue.RegisterEventSource(allegro.GetKeyboardEventSource())

	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
	allegro.FlipDisplay()

	flipFlags := allegro.FlipNone

	for {
		event, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(0.06))
		if found {
			switch event.Type {
				case allegro.DisplayResizeEvent:
					event.Display.Source.AcknowledgeResize()
				case allegro.DisplayCloseEvent:
					running = false
					break
				case allegro.KeyDownEvent:
					keyboardState.Update()
					fmt.Println(event.Keyboard.KeyCode.Name())
					switch event.Keyboard.KeyCode {
						case allegro.KeyEscape:
							running = false
						case allegro.KeyDown:
							flipFlags = allegro.FlipVertical
							event.Keyboard.Display.SetWindowTitle(windowTitle + " [Flipped Vertically]")
						case allegro.KeyLeft:
							flipFlags = allegro.FlipHorizontal
							event.Keyboard.Display.SetWindowTitle(windowTitle + " [Flipped Horizontally]")
					}
				case allegro.KeyUpEvent:
					keyboardState.Update()
					switch event.Keyboard.KeyCode {
						case allegro.KeyDown, allegro.KeyLeft:
							if !keyboardState.IsDown(allegro.KeyDown) && !keyboardState.IsDown(allegro.KeyLeft) {
								flipFlags = allegro.FlipNone
								event.Keyboard.Display.SetWindowTitle(windowTitle)
							}
					}
			}
		}

		if !running {
			break
		}

		allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
		gordon.Draw(float32((display.Width-gordon.Width)/2), float32((display.Height-gordon.Height)/2), flipFlags)
		allegro.FlipDisplay()
	}
}
