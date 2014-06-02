// +build console

package main

import (
    al "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/example/ongoing/subsystems/console"
    "unicode"
)

func initConsole(eventQueue *al.EventQueue) {
    console.Init(eventQueue)
}

func consoleHandled(ev interface{}) bool {
    switch e := ev.(type) {

    case al.KeyDownEvent:
        switch e.KeyCode() {

        case al.KEY_F12:
            console.Toggle()
            return true
        }

    case al.KeyCharEvent:
        if console.Visible() {
            switch e.KeyCode() {

            case al.KEY_BACKSPACE:
                console.BackspaceCmd()
                return true

            case al.KEY_ENTER:
                console.SubmitCmd()
                return true

            default:
                unichar := rune(e.Unichar())
                if unicode.IsPrint(unichar) {
                    console.WriteCmd(string(unichar))
                    return true
                }
            }
        }

    case al.TimerEvent:
        if e.Source() == console.Blinker() {
            console.Blink()
            return true
        }
    }

    return false
}

func renderConsole() {
    console.Render()
}

func saveConsole() {
    console.Save("console.txt")
}
