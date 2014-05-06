go-allegro
==========

[![GoDoc](https://godoc.org/github.com/dradtke/go-allegro?status.png)](https://godoc.org/github.com/dradtke/go-allegro)

This repository contains bindings for writing [Allegro 5](http://alleg.sourceforge.net) games in Go. Obviously, in order for them to work, you'll need to already have a working Allegro 5 development environment set up.

Function documentation is included in the source, but it's pulled directly from Allegro's C API documentation, so not everything will line up as far as parameters and return values. However, the C API maps pretty well to the Go API, so if you're familiar with the patterns (e.g. `error`'s instead of boolean success values, multiple return values instead of output parameters, object functions as instance methods on structs), then it shouldn't be hard to figure out what's going on.

