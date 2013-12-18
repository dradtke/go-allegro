package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/color"
	"github.com/dradtke/go-allegro/allegro/font"
	//"github.com/dradtke/go-allegro/allegro/image"
	"log"
)

func main() {
	var (
		display     *allegro.Display
		eventQueue  *allegro.EventQueue
		builtinFont *font.Font
		//gordon      *allegro.Bitmap
		textColor   allegro.Color
		running     bool = true
		score       int  = 0
		err         error
	)

	allegro.SetNewDisplayFlags(allegro.WINDOWED)
	if display, err = allegro.CreateDisplay(640, 480); err != nil {
		log.Fatal(err)
	} else {
		defer display.Destroy()
		display.SetWindowTitle("Hello, Go!")
	}

	if eventQueue, err = allegro.CreateEventQueue(); err != nil {
		log.Fatal(err)
	} else {
		defer eventQueue.Destroy()
	}

	/*
	image.Init()
	if gordon, err = allegro.LoadBitmap("img/gordon-the-gopher.png"); err != nil {
		log.Fatal(err)
	}
	*/

	font.Init()
	if builtinFont, err = font.Builtin(); err != nil {
		log.Fatal(err)
	}

	textR, textG, textB, err := color.NameToRgb(color.FLORAL_WHITE)
	if err != nil {
		log.Fatal(err)
	}
	textColor = allegro.MapRGB(byte(textR), byte(textG), byte(textB))

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
		//gordon.Draw(float32((display.Width()-gordon.Width())/2), float32((display.Height()-gordon.Height())/2), 0)
		font.DrawTextf(builtinFont, textColor, 50, 50, font.ALIGN_LEFT, "Score: %d", score)
		allegro.FlipDisplay()
	}
}
