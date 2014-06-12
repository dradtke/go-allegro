package game

import (
	"runtime"
)

type GameState interface {
	Enter()
	Render()
	Leave()
}

// NewState() changes the game state.
func NewState(s GameState) {
	if state != nil {
		state.Leave()
	}
	state = s
	state.Enter()
}

// NewStateDelayed() waits for all processes
// to finish without blocking the current goroutine,
// then changes the game state.
func NewStateDelayed(s GameState) {
	go func() {
		for processes.Len() > 0 {
			runtime.Gosched()
		}
		NewState(s)
	}()
}

type BlankState struct{}

func (s *BlankState) Enter()  {}
func (s *BlankState) Render() {}
func (s *BlankState) Leave()  {}
