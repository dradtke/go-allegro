go-allegro
==========

[![GoDoc](https://godoc.org/github.com/dradtke/go-allegro?status.png)](https://godoc.org/github.com/dradtke/go-allegro)

This repository contains bindings for writing [Allegro 5](http://alleg.sourceforge.net) games in Go. Obviously, in order for them to work, you'll need to already have a working Allegro 5 development environment set up.

Function documentation is included in the source, but it's pulled directly from Allegro's C API documentation, so not everything will line up as far as parameters and return values. However, the C API maps pretty well to the Go API, so if you're familiar with the patterns (e.g. `error`'s instead of boolean success values, multiple return values instead of output parameters, object functions as instance methods on structs), then it shouldn't be hard to figure out what's going on.

A number of Allegro functions are blacklisted (via the `blacklist` file) because they either a) overlap with existing Go functionality, like UTF-8 support, or b) are too low-level and probably shouldn't be implemented in Go anyway, like registering new bitmap loaders. These exceptions aside, the bindings officially have 100% method coverage as of Allegro 5.0.10. You can verify this by running `go test coverage_test.go`; any methods found in a header file that aren't covered somewhere in the bindings will show up as unit test failures.
