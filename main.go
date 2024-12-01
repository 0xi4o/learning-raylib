package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameScreen int

const (
	LOGO GameScreen = iota
	TITLE
	GAMEPLAY
	PAUSED
	ENDING
)

type Metadata struct {
	DefaultFont        rl.Font
	WindowTitle        string
	WindowWidth        int32
	WindowHeight       int32
	DefaultFontSize    int32
	DefaultFontSpacing int32
	FrameCounter       int8
	CurrentGameScreen  GameScreen
}

type Hero struct {
	CurrentFrame int8
	FrameSpeed   int8
	FrameRect    rl.Rectangle
	Position     rl.Vector2
	Texture      rl.Texture2D
}

type Game struct {
	Metadata Metadata
	Hero     Hero
}

func main() {
	metadata := Metadata{
		WindowTitle:  "learning raylib",
		WindowWidth:  1280,
		WindowHeight: 720,
	}
	game := Game{
		Metadata: metadata,
	}

	rl.InitWindow(int32(game.Metadata.WindowWidth), int32(game.Metadata.WindowHeight), game.Metadata.WindowTitle)
	game.LoadTextures()
	game.BasicWindow()
	game.UnloadTextures()
	rl.CloseWindow()
}
