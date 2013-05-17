go-allegro
==========

This repository contains experimental bindings for writing [Allegro 5](http://alleg.sourceforge.net) games in Go. Go is a very easy language to write C bindings for, so it's only experimental in the sense that, due to the sheer number of API calls, I've only implemented the bare minimum necessary to get the examples up and running. If you'd like to contribute by writing bindings for some of the lesser-used calls, I would be happy to accept them.

I assume that you have a working Allegro 5 development environment set up. If not, go do that first. Then make sure the repository root is in your `GOPATH` and run `go install allegro` to build the main bindings, or `go install <module>` to build a particular submodule.

Screenshot of the included example (must be run with `example` as your working directory):

![screenshot](https://github.com/dradtke/go-allegro/raw/master/example/screenshot.png)
