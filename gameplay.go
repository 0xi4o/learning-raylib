package main

import (
	"learning-raylib/characters/player"
)

func (game *Game) Gameplay() {
	switch game.Player.State {
	case player.IDLE:
		game.Player.Idle(game.Player.Position)
	case player.RUNNING:
		game.Player.Run(game.Player.Position)
	}
}
