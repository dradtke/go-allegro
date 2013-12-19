package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/color"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/allegro/image"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"log"
	"math"
)

var (
	display    *allegro.Display
	eventQueue *allegro.EventQueue
	timer      *allegro.Timer

	gordon      *allegro.Bitmap
	builtinFont *font.Font
	floralWhite allegro.Color
	black       allegro.Color
	green       allegro.Color
	yellow      allegro.Color
	red         allegro.Color

	running  bool = true
	playing  bool = true
	gameWon  bool
	clicks   int
	timeLeft int = 60
	err      error

	centerX float32
	centerY float32
)

const (
	SCORE_TIME = 0.5
)

// Handy method that loads a known color, quitting the program if it fails.
func loadColor(name color.Name) allegro.Color {
	var r, g, b float32
	if r, g, b, err = color.NameToRgb(name); err != nil {
		log.Fatal(err)
	}
	return allegro.MapRGBf(r, g, b)
}

func loadAssets() {
	// colors
	floralWhite = loadColor(color.FLORAL_WHITE)
	green = loadColor(color.GREEN)
	yellow = loadColor(color.YELLOW)
	red = loadColor(color.RED)
	black = allegro.MapRGB(0, 0, 0)

	// fonts
	font.Init()
	if builtinFont, err = font.Builtin(); err != nil {
		log.Fatal(err)
	}

	// images
	image.Init()
	if gordon, err = allegro.LoadBitmap("img/gopher.png"); err != nil {
		log.Fatal(err)
	}
}

func render() {
	allegro.ClearToColor(black)

	if playing {
		gordon.Draw(centerX-float32(gordon.Width()/2), centerY-float32(gordon.Height()/2), 0)

		font.DrawTextf(builtinFont, floralWhite, centerX, centerY-100, font.ALIGN_CENTRE, "Clicks: %3d", clicks)
		font.DrawText(builtinFont, floralWhite, centerX, centerY+100, font.ALIGN_CENTRE,
			"Can you reach 100 clicks before the time runs out?!")

		start_theta := float32(-math.Pi / 2)
		delta_theta := -2 * math.Pi * (float32(timeLeft) / 60)

		var pieColor allegro.Color
		switch {
		case timeLeft < 15:
			pieColor = red
		case timeLeft < 30:
			pieColor = yellow
		default:
			pieColor = green
		}

		primitives.DrawFilledPieslice(primitives.Point{centerX, 60}, 30, start_theta, delta_theta, pieColor)
	} else {
		if gameWon {
			font.DrawText(builtinFont, floralWhite, centerX, centerY, font.ALIGN_CENTRE,
				"Congratulations!")
		} else {
			font.DrawText(builtinFont, floralWhite, centerX, centerY, font.ALIGN_CENTRE,
				"Better luck next time. :-(")
		}
	}

	allegro.FlipDisplay()
}

func endGame(won bool) {
	playing = false
	gameWon = won
}

func main() {
	allegro.SetNewDisplayFlags(allegro.WINDOWED)
	if display, err = allegro.CreateDisplay(640, 480); err != nil {
		log.Fatal(err)
	} else {
		defer display.Destroy()
		display.SetWindowTitle("Hello, Go!")
		centerX = float32(display.Width() / 2)
		centerY = float32(display.Height() / 2)
	}

	if err = allegro.InstallMouse(); err != nil {
		log.Fatal(err)
	}

	loadAssets()

	if eventQueue, err = allegro.CreateEventQueue(); err != nil {
		log.Fatal(err)
	} else {
		defer eventQueue.Destroy()
	}

	if timer, err = allegro.CreateTimer(SCORE_TIME); err != nil {
		log.Fatal(err)
	}

	var mouseEventSource *allegro.EventSource
	if mouseEventSource, err = allegro.MouseEventSource(); err != nil {
		log.Fatal(err)
	}
	eventQueue.RegisterEventSource(display.EventSource())
	eventQueue.RegisterEventSource(timer.EventSource())
	eventQueue.RegisterEventSource(mouseEventSource)

	render()
	redraw := false
	timer.Start()

	for {
		event, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(0.06))
		if found {
			switch event.Type {
			case allegro.DisplayCloseEvent:
				running = false
				break
			case allegro.TimerEvent:
				if playing {
					timeLeft--
					redraw = true
					if timeLeft == 0 {
						endGame(false)
					}
				}
			case allegro.MouseButtonDownEvent:
				if playing {
					clicks++
					redraw = true
					if clicks == 100 {
						endGame(true)
					}
				}
			}
		}

		if !running {
			break
		}

		if redraw && eventQueue.IsEmpty() {
			render()
			redraw = false
		}
	}
}
