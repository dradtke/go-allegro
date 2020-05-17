package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/color"
	"github.com/dradtke/go-allegro/allegro/primitives"
)

const FPS = 60

func render() {
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

	primitives.DrawLine(
		primitives.Point{100, 100},
		primitives.Point{200, 100},
		color.ForName(color.BLANCHED_ALMOND),
		3,
	)

	primitives.DrawRectangle(
		primitives.Point{150, 150},
		primitives.Point{250, 250},
		color.ForName(color.CADET_BLUE),
		3,
	)

	primitives.DrawPolygon([]primitives.Point{
		{X: 300, Y: 300},
		{X: 280, Y: 320},
		{X: 220, Y: 350},
	}, primitives.LINE_JOIN_BEVEL, color.ForName(color.GHOST_WHITE), 3, 0)

	allegro.FlipDisplay()

}

func main() {
	allegro.Run(func() {
		var (
			display    *allegro.Display
			eventQueue *allegro.EventQueue
			running    bool = true
			err        error
		)

		if err = primitives.Install(); err != nil {
			panic(err)
		}

		// Create a 640x480 window and give it a title.
		allegro.SetNewDisplayFlags(allegro.WINDOWED)
		if display, err = allegro.CreateDisplay(640, 480); err == nil {
			defer display.Destroy()
			display.SetWindowTitle("Primitives")
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

		timer, err := allegro.CreateTimer(1.0 / FPS)
		if err != nil {
			panic(err)
		}

		// Register event sources.
		eventQueue.Register(display)
		eventQueue.Register(timer)

		redraw := false
		timer.Start()

		// Main loop.
		var event allegro.Event
		for running {
			switch eventQueue.WaitForEvent(&event).(type) {
			case allegro.TimerEvent:
				redraw = true

			case allegro.DisplayCloseEvent:
				running = false
				break

				// Handle other events here.
			}

			if !running {
				break
			}

			if redraw && eventQueue.IsEmpty() {
				render()
				redraw = false
			}
		}
	})
}
