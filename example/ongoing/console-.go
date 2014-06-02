// +build !console

package main

import (
    al "github.com/dradtke/go-allegro/allegro"
)

func initConsole(eventQueue *al.EventQueue) {}

func consoleHandled(ev interface{}) bool {
    return false
}

func renderConsole() {}

func saveConsole() {}
