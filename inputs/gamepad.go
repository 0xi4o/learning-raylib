package inputs

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Gamepad struct {
	Status string
}

func (inputs *Inputs) InitializeGamepad() {
	isGamePadAvailable := rl.IsGamepadAvailable(0)

	if isGamePadAvailable {
		fmt.Println("Gamepad Connected!")
		inputs.Gamepad.Status = "connected"
	}
}
