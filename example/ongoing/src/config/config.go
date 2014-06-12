package config

import (
	al "github.com/dradtke/go-allegro/allegro"
)

var blank_color al.Color

const CONSOLE_FILE = "build/console.txt"

func init() {
    blank_color = al.MapRGB(0, 0, 0)
}

func BlankColor() al.Color {
    return blank_color
}

func DisplayWidth() int {
    return 640
}

func DisplayHeight() int {
    return 480
}

func GameName() string {
    return "Hello World"
}
