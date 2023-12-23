package input

import (
	"github.com/drpaneas/pygame/pkg/display"
	"github.com/gopxl/pixel/v2"
)

// GetKeys returns a map of the pressed keys
// so if you press a key, the value of the key is true
// it will keep being true until you release the key.
func GetKeys() map[string]bool {
	keys := make(map[string]bool)

	if display.Screen.Pressed(pixel.KeyEscape) {
		keys["Esc"] = true
	}

	if display.Screen.Pressed(pixel.KeyUp) || display.Screen.Pressed(pixel.KeyW) {
		keys["Up"] = true
	}

	if display.Screen.Pressed(pixel.KeyDown) || display.Screen.Pressed(pixel.KeyS) {
		keys["Down"] = true
	}

	if display.Screen.Pressed(pixel.KeyLeft) || display.Screen.Pressed(pixel.KeyA) {
		keys["Left"] = true
	}

	if display.Screen.Pressed(pixel.KeyRight) || display.Screen.Pressed(pixel.KeyD) {
		keys["Right"] = true
	}

	if display.Screen.JustPressed(pixel.KeySpace) {
		keys["Space"] = true
	}

	return keys
}
