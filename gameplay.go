package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) Gameplay() {
	// pos := rl.NewVector2(0, 0)

	staticPos := rl.NewVector2(float32(game.Metadata.WindowWidth/2)-float32(game.Hero.Texture.Width/2), 20)
	rl.DrawTexture(game.Hero.Texture, int32(staticPos.X), int32(staticPos.Y), rl.White)

	runningPos := rl.NewVector2(float32(game.Metadata.WindowWidth/2)-32.0, float32(game.Metadata.WindowHeight/2)-float32(game.Hero.Texture.Height/2))
	rl.DrawTextureRec(game.Hero.Texture, game.Hero.FrameRect, runningPos, rl.White)
}
