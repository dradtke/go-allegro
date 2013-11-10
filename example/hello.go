package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/image"
	"fmt"
	"os"
)

func main() {
	var (
		display *allegro.Display
		eventQueue *allegro.EventQueue
		gordon *allegro.Bitmap
		running bool = true
		err error
	)

	allegro.SetNewDisplayFlags(allegro.WINDOWED)
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
