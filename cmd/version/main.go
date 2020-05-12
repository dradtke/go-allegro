package main

import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
)

func main() {
	major, minor, revision, _ := allegro.Version()
	fmt.Printf("Allegro %d.%d.%d is installed.\n", major, minor, revision)
}
