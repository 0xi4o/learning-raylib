package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (game *Game) InitializeHero() {
	game.Hero.FrameRect = rl.Rectangle{
		X:      0.0,
		Y:      0.0,
		Width:  float32(game.Hero.Texture.Width) / 4,
		Height: float32(game.Hero.Texture.Height),
	}
	game.Hero.CurrentFrame = 0
	game.Hero.FrameSpeed = 15
}
