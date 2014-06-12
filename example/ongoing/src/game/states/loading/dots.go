package loading

import (
	"github.com/dradtke/go-allegro/example/ongoing/src/game"
)

type LoadingDotAnimation struct{
    timer int
    DotDelay int
}

func (p *LoadingDotAnimation) HandleMessage(msg interface{}) {
}

func (p *LoadingDotAnimation) HandleEvent(ev interface{}) bool {
    return false
}

func (p *LoadingDotAnimation) Tick(delta float32) (bool, error) {
    p.timer++
    if p.timer >= p.DotDelay {
        state := game.State().(*LoadingState)
        if state.dots == "..." {
            state.dots = ""
        } else {
            state.dots += "."
        }
        p.timer = 0
    }
    return true, nil
}
