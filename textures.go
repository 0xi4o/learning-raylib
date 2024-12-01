package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (game *Game) LoadTextures() {
	game.Hero.Texture = rl.LoadTexture("assets/hero/run.png")
}

func (game *Game) UnloadTextures() {
	rl.UnloadTexture(game.Hero.Texture)
}
