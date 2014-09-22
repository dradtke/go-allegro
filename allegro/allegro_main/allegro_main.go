/*
Package allegro_main provides support for Allegro's main addon.

On OS X, you may run into an NSInternalInconsistencyException if
you attempt to initialize and use Allegro from the main() function
directly. To get around that, declare your program like this:

    package main

    import (
        "github.com/dradtke/go-allegro/allegro/allegro_main"
    )

    func main() {
        allegro_main.Run(func() {
            // initialize and use Allegro here
        })
    }

Note: this has not been fully tested as I do not have access
to OS X, so please let me know if it doesn't work for you.

For more information, go to https://github.com/dradtke/go-allegro/issues/3.
*/
package allegro_main

// #include "main.c"
import "C"

var _main func()

//export go_main
func go_main() {
	if _main != nil {
		_main()
	}
}

// Run() runs the provided function via al_run_main().
func Run(f func()) {
	_main = f
	C.run_main()
}
