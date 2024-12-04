package player

import (
	"learning-raylib/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var DebugMode bool = false

type Metadata struct {
	CurrentFrame int8
	FrameSpeed   int8
	SrcRect      rl.Rectangle
	DestRect     rl.Rectangle
	Origin       rl.Vector2
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
	Camera   rl.Camera2D
	Metadata Metadata
	State    State
	Textures Textures
}

func Initialize(config config.Config, debugMode bool) Player {
	// set the debug mode
	DebugMode = debugMode

	// load the textures before setting the player's metadata
	textures := LoadTextures()

	// set the player's metadata
	frameWidth := float32(textures.Idle.Width) / 6
	frameHeight := float32(textures.Idle.Height)
	destRect := rl.Rectangle{
		X:      float32(config.WindowWidth / 2),
		Y:      float32(config.WindowHeight / 2),
		Width:  frameWidth * 2,
		Height: frameHeight * 2,
	}
	origin := rl.Vector2{
		X: frameWidth,
		Y: frameHeight,
	}
	position := rl.Vector2{
		X: float32(config.WindowWidth / 2),
		Y: float32(config.WindowHeight / 2),
	}
	srcRect := rl.Rectangle{
		X:      0.0,
		Y:      0.0,
		Width:  frameWidth,
		Height: frameHeight,
	}

	// set the camera
	camera := rl.Camera2D{
		Offset:   position,
		Rotation: 0.0,
		Target:   position,
		Zoom:     1.0,
	}

	metadata := Metadata{
		CurrentFrame: 0,
		DestRect:     destRect,
		FrameSpeed:   15,
		Origin:       origin,
		Position:     position,
		SrcRect:      srcRect,
	}

	return Player{
		Camera:   camera,
		Metadata: metadata,
		State:    IDLE,
		Textures: textures,
	}
}

func (player *Player) Idle(pos rl.Vector2) {
	if DebugMode {
		rl.DrawRectangleLines(int32(player.Metadata.DestRect.X-player.Metadata.SrcRect.Width), int32(player.Metadata.DestRect.Y-player.Metadata.SrcRect.Height), int32(player.Metadata.DestRect.Width), int32(player.Metadata.DestRect.Height), rl.Red)
	}
	rl.DrawTexturePro(player.Textures.Idle, player.Metadata.SrcRect, player.Metadata.DestRect, player.Metadata.Origin, 0.0, rl.White)
}

func (player *Player) Run(pos rl.Vector2) {
	if DebugMode {
		rl.DrawRectangleLines(int32(player.Metadata.DestRect.X-player.Metadata.SrcRect.Width), int32(player.Metadata.DestRect.Y-player.Metadata.SrcRect.Height), int32(player.Metadata.DestRect.Width), int32(player.Metadata.DestRect.Height), rl.Red)
	}
	rl.DrawTexturePro(player.Textures.Run, player.Metadata.SrcRect, player.Metadata.DestRect, player.Metadata.Origin, 0.0, rl.White)
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
