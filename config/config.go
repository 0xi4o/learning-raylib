package config

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	DefaultFont        rl.Font
	WindowTitle        string
	WindowWidth        int32
	WindowHeight       int32
	DefaultFontSize    int32
	DefaultFontSpacing int32
}

func GetGameConfig() Config {
	return Config{
		WindowTitle:        "learning raylib",
		WindowWidth:        1920,
		WindowHeight:       1080,
		DefaultFont:        rl.GetFontDefault(),
		DefaultFontSize:    20,
		DefaultFontSpacing: 1,
	}
}
