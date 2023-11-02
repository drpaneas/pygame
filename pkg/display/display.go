package display

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

var Screen *pixelgl.Window

func SetMode(screenWidth, screenHeight float64) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:       "Platformer",
		Bounds:      pixel.R(0, 0, screenWidth, screenHeight),
		VSync:       true,
		Resizable:   false,
		Undecorated: false,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}

	win.SetSmooth(false)

	return win
}
