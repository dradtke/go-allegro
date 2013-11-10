package main

import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/image"
	"math"
	"os"
)

var gameMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

const TILE_SIZE = 30
const START_X = 6
const START_Y = 6
const GOPHER_SPEED = 6

var (
	gopher    *Gopher
	blackTile *allegro.Bitmap
	whiteTile *allegro.Bitmap
)

type Gopher struct {
	x, y, w, h float32
	graphic    *allegro.Bitmap
}

func (g *Gopher) Render() {
	g.graphic.Draw(g.x, g.y, allegro.FLIP_NONE)
}

func (g *Gopher) Move(x, y float32) {
	var tx float32 = g.x + x
	var ty float32 = g.y + y
	if x > 0 {
		xtile := int(math.Floor(float64((tx + g.w) / TILE_SIZE)))
		ytile := int(math.Floor(float64(ty / TILE_SIZE)))
		if gameMap[ytile][xtile] == 0 {
			g.x = tx
		}
	} else if x < 0 {
		xtile := int(math.Floor(float64(tx / TILE_SIZE)))
		ytile := int(math.Floor(float64(ty / TILE_SIZE)))
		if gameMap[ytile][xtile] == 0 {
			g.x = tx
		}
	} else if y > 0 {
		xtile := int(math.Floor(float64(tx / TILE_SIZE)))
		ytile := int(math.Floor(float64((ty + g.h) / TILE_SIZE)))
		if gameMap[ytile][xtile] == 0 {
			g.y = ty
		}
	} else if y < 0 {
		xtile := int(math.Floor(float64(tx / TILE_SIZE)))
		ytile := int(math.Floor(float64(ty / TILE_SIZE)))
		if gameMap[ytile][xtile] == 0 {
			g.y = ty
		}
	}
}

func Render() {
	allegro.HoldBitmapDrawing(true)
	for y, row := range gameMap {
		for x, tile := range row {
			dx := float32(x * TILE_SIZE)
			dy := float32(y * TILE_SIZE)
			if tile == 1 {
				blackTile.Draw(dx, dy, allegro.FLIP_NONE)
			} else {
				whiteTile.Draw(dx, dy, allegro.FLIP_NONE)
			}
		}
	}
	gopher.Render()
	allegro.HoldBitmapDrawing(false)
	allegro.FlipDisplay()
}

func main() {
	var (
		display    *allegro.Display
		eventQueue *allegro.EventQueue
		running    bool = true
		err        error
	)

	image.Init()

	allegro.SetNewDisplayFlags(allegro.WINDOWED)
	if display, err = allegro.CreateDisplay(600, 480); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	} else {
		defer display.Destroy()
		display.SetWindowTitle("Help Gordon Escape!")
	}

	if eventQueue, err = allegro.CreateEventQueue(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	} else {
		defer eventQueue.Destroy()
	}

	eventQueue.RegisterEventSource(display.EventSource())
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
	allegro.FlipDisplay()

	gopher = new(Gopher)
	bmp, err := allegro.LoadBitmap("img/little-gopher.png")
	if err != nil {
		panic(err)
	}
	gopher.graphic = bmp
	gopher.w = float32(gopher.graphic.Width())
	gopher.h = float32(gopher.graphic.Height())
	gopher.x = float32((START_X * TILE_SIZE) - (gopher.w/2))
	gopher.y = float32((START_Y * TILE_SIZE) - (gopher.h/2))

	background := allegro.TargetBitmap()

	blackTile = allegro.CreateBitmap(TILE_SIZE, TILE_SIZE)
	allegro.SetTargetBitmap(blackTile)
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

	whiteTile = allegro.CreateBitmap(TILE_SIZE, TILE_SIZE)
	allegro.SetTargetBitmap(whiteTile)
	allegro.ClearToColor(allegro.MapRGB(0xFF, 0xFF, 0xFF))

	allegro.SetTargetBitmap(background)

	if err := allegro.InstallKeyboard(); err != nil {
		panic(err)
	}
	var keyboard allegro.KeyboardState

	for running {
		event, found := eventQueue.WaitForEventUntil(allegro.NewTimeout(0.03))
		if found {
			switch event.Type {
			case allegro.DisplayCloseEvent:
				running = false
				break
			default:
				// do nothing
			}
		}

		if !running {
			break
		}

		keyboard.Get()
		if keyboard.IsDown(allegro.KEY_RIGHT) {
			gopher.Move(GOPHER_SPEED, 0)
		} else if keyboard.IsDown(allegro.KEY_LEFT) {
			gopher.Move(-GOPHER_SPEED, 0)
		} else if keyboard.IsDown(allegro.KEY_DOWN) {
			gopher.Move(0, GOPHER_SPEED)
		} else if keyboard.IsDown(allegro.KEY_UP) {
			gopher.Move(0, -GOPHER_SPEED)
		}

		Render()
	}
}
