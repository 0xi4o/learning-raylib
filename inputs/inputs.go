package inputs

type Inputs struct {
	Keyboard Keyboard
	Gamepad  Gamepad
}

func (inputs *Inputs) Initialize() {
	inputs.Gamepad = Gamepad{}
	inputs.Keyboard = Keyboard{}
}
