package input

import (
	"github.com/drpaneas/pygame/pkg/display"
	"github.com/gopxl/pixel/v2/pixelgl"
)

// GetKeys returns a map of the pressed keys
// so if you press a key, the value of the key is true
// it will keep being true until you release the key.
func GetKeys() map[string]bool {
	keys := make(map[string]bool)

	if display.Screen.Pressed(pixelgl.KeyEscape) {
		keys["Esc"] = true
	}

	if display.Screen.Pressed(pixelgl.KeyUp) {
		keys["Up"] = true
	}

	if display.Screen.Pressed(pixelgl.KeyDown) {
		keys["Down"] = true
	}

	if display.Screen.Pressed(pixelgl.KeyLeft) {
		keys["Left"] = true
	}

	if display.Screen.Pressed(pixelgl.KeyRight) {
		keys["Right"] = true
	}

	if display.Screen.Pressed(pixelgl.KeySpace) {
		keys["Space"] = true
	}

	return keys
}
