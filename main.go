package main

import rl "github.com/gen2brain/raylib-go/raylib"

type gameScreen int

const (
	LOGO gameScreen = iota
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
}

type Game struct {
	Metadata Metadata
}

func main() {
	metadata := Metadata{
		WindowTitle:  "learning raylib",
		WindowWidth:  1920,
		WindowHeight: 1080,
	}
	game := Game{
		Metadata: metadata,
	}

	game.BasicWindow()
}
