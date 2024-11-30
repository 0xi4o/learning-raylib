package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) Gamepad() {
	isGamePadAvailable := rl.IsGamepadAvailable(0)

	if isGamePadAvailable {
		text := "Gamepad connected!"
		textMeasurements := rl.MeasureTextEx(game.Metadata.DefaultFont, text, float32(game.Metadata.DefaultFontSize), float32(game.Metadata.DefaultFontSpacing))
		rl.DrawText(text, game.Metadata.WindowWidth-int32(textMeasurements.X)-game.Metadata.DefaultFontSize-20, 20, game.Metadata.DefaultFontSize, rl.LightGray)
	}
}
