package allegro

// #include "main.c"
import "C"

var _main func()

//export go_main
func go_main() {
	if err := install(); err != nil {
		panic(err)
	}
	if _main != nil {
		_main()
	}
	uninstall()
}

/*
On OS X, you may run into an NSInternalInconsistencyException if
you attempt to initialize and use Allegro from the main() function
directly. To get around that, the Run() function is used in conjunction
with the allegro_main module and al_run_main().

Note: this has not been fully tested as I do not have access
to OS X, so please let me know if it doesn't work for you.

For more information, go to https://github.com/dradtke/go-allegro/issues/3.
*/
func Run(f func()) {
	_main = f
	C.run_main()
}
