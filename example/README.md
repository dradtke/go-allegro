This folder contains examples for using `go-allegro`. The import paths for these examples assume that the `go-allegro` source is contained inside some `src/github.com/dradtke/go-allegro` in your `GOPATH`. In addition to that, they must be run from the `example` folder in order for the image to load correctly.

hello.go
--------

Just displays a 640x480 window with the Go gopher displayed in the center.

resize.go
---------

Same as `hello.go`, but you can resize the window and the gopher should stay in the center.

flip.go
-------

Press the Left or Down keys to flip the gopher horizontally or vertically respectively, and press Escape to close the window. In addition, any keypress will print the name of the pressed key to standard output.
