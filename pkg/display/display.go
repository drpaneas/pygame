package display

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

// Screen is the window of the game that is being displayed to the user on the screen of the computer or device they are using
var Screen *opengl.Window

// SetMode sets the mode of the screen
// and returns the window
func SetMode(screenWidth, screenHeight float64) *opengl.Window {
	cfg := opengl.WindowConfig{
		Title:       "Platformer",
		Bounds:      pixel.R(0, 0, screenWidth, screenHeight),
		VSync:       true,
		Resizable:   false,
		Undecorated: false,
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}

	win.SetSmooth(false) // For scaling pixel art, it is recommended to use nearest neighbor interpolation.

	return win
}
