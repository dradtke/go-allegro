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
		gordon *allegro.Bitmap
		running bool = true
		err error
	)

	if err = allegro.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	allegro.SetNewDisplayFlags(allegro.Windowed)
	if display, err = allegro.CreateDisplay(640, 480); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	} else {
		defer display.Destroy()
		display.SetWindowTitle("Hello, Go!")
	}

	if eventQueue, err = allegro.CreateEventQueue(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	} else {
		defer eventQueue.Destroy()
	}

	image.Init()
	if gordon, err = allegro.LoadBitmap("img/gordon-the-gopher.png"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	eventQueue.RegisterEventSource(display.EventSource())
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
	allegro.FlipDisplay()

	for {
		event, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(0.06))
		if found {
			switch event.Type {
				case allegro.DisplayCloseEvent:
					running = false
					break
			}
		}

		if !running {
			break
		}

		allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

		gordon.Draw(float32((display.Width()-gordon.Width())/2), float32((display.Height()-gordon.Height())/2), 0)

		allegro.FlipDisplay()
	}
}
