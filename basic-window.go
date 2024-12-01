package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) BasicWindow() {
	game.Metadata.CurrentGameScreen = LOGO
	game.Metadata.FrameCounter = 0

	game.Metadata.DefaultFont = rl.GetFontDefault()
	game.Metadata.DefaultFontSize = 20
	game.Metadata.DefaultFontSpacing = 1

	// initialize Hero
	game.InitializeHero()

	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	for !rl.WindowShouldClose() {
		game.Metadata.FrameCounter++
		switch game.Metadata.CurrentGameScreen {
		case LOGO:
			if game.Metadata.FrameCounter > 120 {
				game.Metadata.CurrentGameScreen = TITLE
			}
		case GAMEPLAY:
			// update the frame
			if game.Metadata.FrameCounter >= int8(60/game.Hero.FrameSpeed) {
				game.Metadata.FrameCounter = 0
				game.Hero.CurrentFrame++
				if game.Hero.CurrentFrame > 3 {
					game.Hero.CurrentFrame = 0
				}
				game.Hero.FrameRect.X = float32(game.Hero.CurrentFrame) * game.Hero.FrameRect.Width
			}
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		switch game.Metadata.CurrentGameScreen {
		case LOGO:
			rl.DrawText("LOGO SCREEN", 20, 20, 40, rl.LightGray)
			text := "WAIT FOR 2 SECONDS..."
			pos := game.getCenteredTextPos(text, game.Metadata.DefaultFontSize, float32(game.Metadata.DefaultFontSpacing))
			rl.DrawText(text, pos.X, pos.Y, game.Metadata.DefaultFontSize, rl.Gray)
		case TITLE:
			title := "My Super Awesome Game"
			text := "Press Any Key"
			titlePos := game.getCenteredTextPos(title, game.Metadata.DefaultFontSize+20, float32(game.Metadata.DefaultFontSpacing))
			textPos := game.getCenteredTextPos(text, game.Metadata.DefaultFontSize, float32(game.Metadata.DefaultFontSpacing))
			rl.DrawText(title, titlePos.X, titlePos.Y, game.Metadata.DefaultFontSize+20, rl.LightGray)
			rl.DrawText(text, textPos.X, textPos.Y+40, game.Metadata.DefaultFontSize, rl.LightGray)
			// game.Gamepad()
			pressedKey := rl.GetKeyPressed()
			if pressedKey != 0 {
				game.Metadata.CurrentGameScreen = GAMEPLAY
			}
		case GAMEPLAY:
			pressedKey := rl.GetKeyPressed()
			if pressedKey == 256 {
				game.Metadata.CurrentGameScreen = PAUSED
			}
			game.Gameplay()
		case PAUSED:
			text := "Game Paused"
			pos := game.getCenteredTextPos(text, game.Metadata.DefaultFontSize, float32(game.Metadata.DefaultFontSpacing))
			rl.DrawText(text, pos.X, pos.Y-int32(0.25*float64(game.Metadata.WindowHeight)), game.Metadata.DefaultFontSize, rl.LightGray)
			pressedKey := rl.GetKeyPressed()
			if pressedKey == 256 {
				game.Metadata.CurrentGameScreen = GAMEPLAY
			}
		default:
		}

		rl.EndDrawing()
	}
}
