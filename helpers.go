package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Position struct {
	X int32
	Y int32
}

func (game *Game) getCenteredTextPos(text string, fontSize int32, spacing float32) Position {
	defaultFont := rl.GetFontDefault()
	// measure text width and height using default font, font size, and spacing
	textMeasurements := rl.MeasureTextEx(defaultFont, text, float32(fontSize), spacing)
	// return position to center text on screen
	return Position{
		X: int32(game.Metadata.WindowWidth/2) - int32(textMeasurements.X/2) - int32(fontSize),
		Y: int32(game.Metadata.WindowHeight/2) - int32(textMeasurements.Y/2),
	}
}
