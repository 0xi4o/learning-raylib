package main

import (
	"learning-raylib/characters/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) Start() {
	game.Metadata.CurrentGameScreen = LOGO
	game.Metadata.FrameCounter = 0

	game.Player.Position = rl.NewVector2(float32(game.Config.WindowWidth/2)-32.0, float32(game.Config.WindowHeight/2)-64.0)

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
			numberOfFrames := 6
			switch game.Player.State {
			case player.RUNNING:
				numberOfFrames = 4
				game.Player.Metadata.FrameRect.Width = float32(game.Player.Textures.Run.Width) / float32(numberOfFrames)
				game.Player.Metadata.FrameRect.Height = float32(game.Player.Textures.Run.Height)
				game.Player.Metadata.FrameSpeed = 20
			default:
				numberOfFrames = 6
				game.Player.Metadata.FrameRect.Width = float32(game.Player.Textures.Idle.Width) / float32(numberOfFrames)
				game.Player.Metadata.FrameRect.Height = float32(game.Player.Textures.Idle.Height)
				game.Player.Metadata.FrameSpeed = 12
			}
			// update the frame
			if game.Metadata.FrameCounter >= int8(60/game.Player.Metadata.FrameSpeed) {
				game.Metadata.FrameCounter = 0
				game.Player.Metadata.CurrentFrame++
				if game.Player.Metadata.CurrentFrame > int8(numberOfFrames) {
					game.Player.Metadata.CurrentFrame = 0
				}
				game.Player.Metadata.FrameRect.X = float32(game.Player.Metadata.CurrentFrame) * game.Player.Metadata.FrameRect.Width
			}
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		switch game.Metadata.CurrentGameScreen {
		case LOGO:
			rl.DrawText("LOGO SCREEN", 20, 20, 40, rl.LightGray)
			text := "WAIT FOR 2 SECONDS..."
			pos := game.getCenteredTextPos(text, game.Config.DefaultFontSize, float32(game.Config.DefaultFontSpacing))
			rl.DrawText(text, pos.X, pos.Y, game.Config.DefaultFontSize, rl.Gray)
		case TITLE:
			title := "My Super Awesome Game"
			text := "Press Any Key"
			titlePos := game.getCenteredTextPos(title, game.Config.DefaultFontSize+20, float32(game.Config.DefaultFontSpacing))
			textPos := game.getCenteredTextPos(text, game.Config.DefaultFontSize, float32(game.Config.DefaultFontSpacing))
			rl.DrawText(title, titlePos.X, titlePos.Y, game.Config.DefaultFontSize+20, rl.LightGray)
			rl.DrawText(text, textPos.X, textPos.Y+40, game.Config.DefaultFontSize, rl.LightGray)
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
			game.Inputs.OnKeyPress(&game.Player)
			game.Gameplay()
		case PAUSED:
			text := "Game Paused"
			pos := game.getCenteredTextPos(text, game.Config.DefaultFontSize, float32(game.Config.DefaultFontSpacing))
			rl.DrawText(text, pos.X, pos.Y-int32(0.25*float64(game.Config.WindowHeight)), game.Config.DefaultFontSize, rl.LightGray)
			pressedKey := rl.GetKeyPressed()
			if pressedKey == 256 {
				game.Metadata.CurrentGameScreen = GAMEPLAY
			}
		default:
		}

		rl.EndDrawing()
	}
}
