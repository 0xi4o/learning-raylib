package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	Idle    rl.Texture2D
	Run     rl.Texture2D
	Jump    rl.Texture2D
	Attack1 rl.Texture2D
	Attack2 rl.Texture2D
}

func LoadTextures() Textures {
	idle := rl.LoadTexture("assets/player/idle.png")
	run := rl.LoadTexture("assets/player/run.png")
	jump := rl.LoadTexture("assets/player/jump.png")
	attack1 := rl.LoadTexture("assets/player/attack1.png")
	attack2 := rl.LoadTexture("assets/player/attack2.png")

	return Textures{
		Idle:    idle,
		Run:     run,
		Jump:    jump,
		Attack1: attack1,
		Attack2: attack2,
	}
}

func (player *Player) UnloadTextures() {
	rl.UnloadTexture(player.Textures.Idle)
	rl.UnloadTexture(player.Textures.Run)
	rl.UnloadTexture(player.Textures.Jump)
	rl.UnloadTexture(player.Textures.Attack1)
	rl.UnloadTexture(player.Textures.Attack2)
}
