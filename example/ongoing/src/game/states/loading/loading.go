package loading

import (
    al "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/example/ongoing/src/game"
)

type LoadingState struct{
    dots string
}

func (s *LoadingState) Enter() {
    font.Install()
    game.RunProcess(&LoadingDotAnimation{DotDelay: 30})
}

func (s *LoadingState) Render() {
    f, _ := font.Builtin()
    font.DrawText(f, al.MapRGB(255, 255, 255), 10, 10, font.ALIGN_LEFT, "Loading" + s.dots)
}

func (s *LoadingState) Leave() {}
