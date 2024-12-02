package main

import (
	"learning-raylib/characters/player"
	"learning-raylib/config"
	"learning-raylib/inputs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScreen int

const (
	LOGO GameScreen = iota
	TITLE
	GAMEPLAY
	PAUSED
	ENDING
)

type Metadata struct {
	FrameCounter      int8
	CurrentGameScreen GameScreen
}

type Game struct {
	Inputs   inputs.Inputs
	Config   config.Config
	Metadata Metadata
	Player   player.Player
}

func main() {
	// initialize game config
	gameConfig := config.GetGameConfig()

	// initialize inputs
	inputs := inputs.Inputs{}
	inputs.Initialize()

	// initialize game metadata
	metadata := Metadata{}

	// initialize window
	rl.InitWindow(int32(gameConfig.WindowWidth), int32(gameConfig.WindowHeight), gameConfig.WindowTitle)

	// initialize player
	player := player.Initialize()

	game := Game{
		Config:   gameConfig,
		Inputs:   inputs,
		Metadata: metadata,
		Player:   player,
	}

	// start game
	game.Start()

	// clean up
	game.UnloadTextures()
	rl.CloseWindow()
}
