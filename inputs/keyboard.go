package inputs

import (
	"learning-raylib/characters/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Keyboard struct{}

func (inputs *Inputs) OnKeyPress(p *player.Player) {
	switch {
	case rl.IsKeyDown(rl.KeyD):
		p.State = player.RUNNING
	default:
		p.State = player.IDLE
	}
}
