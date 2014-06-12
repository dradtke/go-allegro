package game

import (
	"container/list"
	al "github.com/dradtke/go-allegro/allegro"
)

// Frames per second.
const FPS int = 60

var (
	display    *al.Display
	eventQueue *al.EventQueue
	fpsTimer   *al.Timer

	processes list.List
	mailboxes list.List

	state GameState
	event al.Event

	messengers = make(map[Process]chan interface{})
)

// Display() returns a reference to the game's display.
func Display() *al.Display {
	return display
}

// EventQueue() returns a reference to the game's event queue.
func EventQueue() *al.EventQueue {
	return eventQueue
}

// State() returns a reference to the game's current state.
func State() GameState {
    return state
}
