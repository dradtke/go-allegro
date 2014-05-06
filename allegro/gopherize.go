package allegro

// This file contains tools for making the library more idiomatic by
// doing things like grouping common functionality into interfaces.

// EventGenerator represents anything that can be registered with an
// event queue.
type EventGenerator interface {
    EventSource() *EventSource
}
