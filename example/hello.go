
package main

import (
    "github.com/dradtke/go-allegro/allegro"
)

const FPS int = 60

func main() {
    if err := allegro.Install(); err != nil {
        panic(err)
    }
    defer allegro.Uninstall()

    var (
        display    *allegro.Display
        eventQueue *allegro.EventQueue
        running    bool = true
        err        error
    )

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
        found := eventQueue.WaitForEventUntil(allegro.NewTimeout(timeout), &event)
        if found {
            switch event.Type {
            case allegro.EVENT_DISPLAY_CLOSE:
                running = false
                break

                // Handle other events here.
            }
        }

        if !running {
            return
        }
    }
}
