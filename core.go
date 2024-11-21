package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameScreen int

const (
	LOGO gameScreen = iota
	TITLE
	GAMEPLAY
	ENDING
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
)

func main() {
	windowTitle := "raylib [texture] example - sprite animation"

	rl.InitWindow(int32(WindowWidth), int32(WindowHeight), windowTitle)

	currentScreen := LOGO
	framesCounter := 0

	rl.SetTargetFPS(60)

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
			pos := getCenteredTextPos(text, 20.0, 1.0)
			rl.DrawText(text, pos.X, pos.Y, 20, rl.Gray)
		case TITLE:
			text := "Congrats! You created your first window!"
			pos := getCenteredTextPos(text, 20.0, 1.0)
			rl.DrawText("Congrats! You created your first window!", pos.X, pos.Y, 20.0, rl.LightGray)
		default:
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

type Position struct {
	X int32
	Y int32
}

func getCenteredTextPos(text string, fontSize float32, spacing float32) Position {
	defaultFont := rl.GetFontDefault()
	textMeasurements := rl.MeasureTextEx(defaultFont, text, fontSize, spacing)
	return Position{
		X: int32(WindowWidth/2) - int32(textMeasurements.X/2),
		Y: int32(WindowHeight/2) - int32(textMeasurements.Y/2),
	}
}
