package main

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"github.com/dradtke/go-allegro/allegro/font"
	prim "github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/dradtke/go-allegro/example/ongoing/src/config"
	"github.com/dradtke/go-allegro/example/ongoing/src/game"
	"github.com/dradtke/go-allegro/example/ongoing/src/subsystems/console"
	"github.com/dradtke/go-allegro/example/ongoing/src/game/states/loading"
)

func main() {
    {
        if err := al.Install(); err != nil {
            panic(err)
        }
        defer al.Uninstall()

        if err := dialog.Install(); err != nil {
            panic(err)
        }

        if err := prim.Install(); err != nil {
            game.Fatal(err)
        }
        defer prim.Uninstall()

        font.Install()
        defer font.Uninstall()
    }

    game.Init()
    defer game.Cleanup()

	// Initialize subsystems.
	console.Init(game.EventQueue())

	// Set the screen to black.
	al.ClearToColor(config.BlankColor())
	al.FlipDisplay()

    game.NewState(&loading.LoadingState{})
    game.Loop()

	game.Display().SetWindowTitle("Shutting down...")
	console.Save()
}
