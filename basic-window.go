package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) BasicWindow() {
	rl.InitWindow(int32(game.Metadata.WindowWidth), int32(game.Metadata.WindowHeight), game.Metadata.WindowTitle)

	currentScreen := LOGO
	framesCounter := 0

	game.Metadata.DefaultFont = rl.GetFontDefault()
	game.Metadata.DefaultFontSize = 20
	game.Metadata.DefaultFontSpacing = 1

	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	for !rl.WindowShouldClose() {
		switch currentScreen {
		case LOGO:
			framesCounter += 1
			if framesCounter > 120 {
				currentScreen = TITLE
			}
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		switch currentScreen {
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
				currentScreen = GAMEPLAY
			}
		case GAMEPLAY:
			pressedKey := rl.GetKeyPressed()
			if pressedKey == 256 {
				currentScreen = PAUSED
			}
			game.Gameplay()
		case PAUSED:
			text := "Game Paused"
			pos := game.getCenteredTextPos(text, game.Metadata.DefaultFontSize, float32(game.Metadata.DefaultFontSpacing))
			rl.DrawText(text, pos.X, pos.Y-int32(0.25*float64(game.Metadata.WindowHeight)), game.Metadata.DefaultFontSize, rl.LightGray)
			pressedKey := rl.GetKeyPressed()
			if pressedKey == 256 {
				currentScreen = GAMEPLAY
			}
		default:
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
