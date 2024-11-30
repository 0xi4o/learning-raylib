package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) Gameplay() {
	texture := rl.LoadTexture("assets/hero/run.png")
	fmt.Println(texture)
}
