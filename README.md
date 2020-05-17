go-allegro
==========

[![GoDoc](https://godoc.org/github.com/dradtke/go-allegro?status.png)](https://godoc.org/github.com/dradtke/go-allegro)

This repository contains bindings for writing [Allegro 5](https://github.com/liballeg/allegro5) games in Go. Obviously, in order for them to work, you'll need to already have a working Allegro 5 development environment set up.

Function documentation is included in the source, but it's pulled directly from Allegro's C API documentation, so not everything will line up as far as parameters and return values. However, the C API maps pretty well to the Go API, so if you're familiar with the patterns (e.g. `error`'s instead of boolean success values, multiple return values instead of output parameters, object functions as instance methods on structs), then it shouldn't be hard to figure out what's going on.

A number of Allegro functions are blacklisted (via the `blacklist` file) because they either a) overlap with existing Go functionality, like UTF-8 support, or b) are too low-level and probably shouldn't be implemented in Go anyway, like registering new bitmap loaders. These exceptions aside, the bindings officially have 100% method coverage as of Allegro 5.0.10. You can verify this by running `go test coverage_test.go`; any methods found in a header file that aren't covered somewhere in the bindings will show up as unit test failures.

Branching
=========

`master` is the default branch, but you should usually use one corresponding to your Allegro version. Check out the list of branches to see what's available.

Installation
============

Before installation, be sure to get the source by running `go get -d github.com/dradtke/go-allegro`.

\*Nix
----

Install Allegro 5 through your favorite package manager, ensure that it's registered with `pkg-config`, then run `go install github.com/dradtke/go-allegro/allegro`.

Windows
-------

Download the Allegro 5 binaries [here](https://www.allegro.cc/files/) and extract the root folder somewhere.

Set the `ALLEGRO_HOME` environment variable to this folder's absolute path, and set `ALLEGRO_VERSION` to the version of Allegro downloaded, e.g. 5.0.10. You can also optionally set `ALLEGRO_LIB` to reflect which version you want to link against; the default value is `monolith-static-mt-debug`.

Once that's done, run the included `setenv.bat`, and if no errors were reported, then you can then build and install the library as usual.

Unstable APIs
=============

A number of Allegro's functions are defined as [unstable](https://liballeg.org/a5docs/trunk/getting_started.html#unstable-api), and so in the Go library, they live behind the `unstable` build flag:

```bash
$ go build ./allegro                 # no unstable APIs available
$ go build -tags=unstable ./allegro  # unstable APIs now available
```

Function Callbacks
==================

This library now includes an initial example of passing Go functions to an Allegro function requiring a callback (see `primitives.TriangulatePolygon`), using the technique outlined in the Go wiki [here](https://github.com/golang/go/wiki/cgo#function-variables). Initial testing suggests that it works, but has not been stress-tested, so please open an issue, or even better a pull request, if you encounter any issues.
