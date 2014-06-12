package main

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/example/ongoing/src/subsystems/console"
)

func render() {
	al.HoldBitmapDrawing(true)
	defer al.HoldBitmapDrawing(false)

	console.Render()
}
