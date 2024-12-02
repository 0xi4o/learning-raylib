package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Metadata struct {
	CurrentFrame int8
	FrameSpeed   int8
	FrameRect    rl.Rectangle
	Position     rl.Vector2
	Texture      rl.Texture2D
}

type State int

const (
	IDLE State = iota
	RUNNING
	JUMPING
	ATTACKING1
	ATTACKING2
)

type Player struct {
	Metadata Metadata
	Position rl.Vector2
	State    State
	Textures Textures
}

func Initialize() Player {
	textures := LoadTextures()
	frameRect := rl.Rectangle{
		X:      0.0,
		Y:      0.0,
		Width:  float32(textures.Idle.Width) / 6,
		Height: float32(textures.Idle.Height),
	}
	metadata := Metadata{
		CurrentFrame: 0,
		FrameSpeed:   15,
		FrameRect:    frameRect,
	}

	return Player{
		Metadata: metadata,
		State:    IDLE,
		Textures: textures,
	}
}

func (player *Player) Idle(pos rl.Vector2) {
	rl.DrawTextureRec(player.Textures.Idle, player.Metadata.FrameRect, pos, rl.White)
}

func (player *Player) Run(pos rl.Vector2) {
	rl.DrawTextureRec(player.Textures.Run, player.Metadata.FrameRect, pos, rl.White)
}

type MovementDirection int

const (
	LEFT MovementDirection = iota
	RIGHT
)

type Movement struct {
	Direction MovementDirection
}

func (movement *Movement) Move() {
	switch movement.Direction {
	case LEFT:
		MoveLeft()
	case RIGHT:
		MoveRight()
	}
}

func MoveLeft() {
	panic("not implemented")
}

func MoveRight() {
	panic("not implemented")
}
