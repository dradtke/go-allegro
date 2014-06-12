package game

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/example/ongoing/src/config"
	"github.com/dradtke/go-allegro/example/ongoing/src/util"
	"runtime"
)

// Loop() is the main game loop.
func Loop() {
	var (
		running    = true
		ticking    = false
	)

	for running {
		ev := eventQueue.WaitForEvent(&event)

		switch e := ev.(type) {
		case al.TimerEvent:
			if e.Source() == fpsTimer {
				ticking = true
				goto eventHandled
			}

		case al.DisplayCloseEvent:
			running = false
			goto eventHandled

		case al.UserEvent:
			addr := e.Data1()
			switch util.Retrieve(addr).(type) {
			    // TODO: check on user events
			}

        default:
            for e := processes.Front(); e != nil; e = e.Next() {
                if handled := e.Value.(Process).HandleEvent(ev); handled {
                    break
                }
            }
		}


	eventHandled:
		if running && ticking && eventQueue.IsEmpty() {
			al.ClearToColor(config.BlankColor())
			al.HoldBitmapDrawing(true)

			// TODO: add a delta value somewhere
			Broadcast(&tick{0})
			state.Render()

			al.HoldBitmapDrawing(false)
			al.FlipDisplay()
			ticking = false
		}
	}

	// Tell all processes to quit immediately, then wait
	// for them to finish before exiting.
	Broadcast(&quit{})
	for processes.Len() > 0 {
        runtime.Gosched()
	}
}
