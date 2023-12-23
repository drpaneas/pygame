package main

import (
	"os"
	"time"

	_ "image/png"

	"github.com/drpaneas/pygame/pkg/display"
	"github.com/drpaneas/pygame/pkg/level"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)

func run() {
	display.Screen = display.SetMode(screenWidth, screenHeight)

	// Configure the game loop to run in 60 FPS
	timeframe := time.Second / framesPerSecond
	clock := time.NewTicker(timeframe) // ticks every frame
	start := time.Now()
	frame := 0

	gameLevel := level.NewLevel(level.Map, display.Screen)

	for !display.Screen.Closed() {
		for range clock.C {
			// Called every tick of the clock (every frame)
			frame++

			// Quit the game if the user presses the Esc key
			if display.Screen.JustPressed(pixel.KeyEscape) {
				display.Screen.SetClosed(true)
				os.Exit(1)
			}

			// Clear the screen with black color
			display.Screen.Clear(colornames.Black)

			// All the logic of the game (update, draw, etc is here)
			gameLevel.UpdateAndDraw()

			// Game engine stuff, do not touch
			display.Screen.Update()

			// If a second has passed, reset the frame (counter)
			since := time.Since(start)
			if since > time.Second {
				start = time.Now()
				frame = 0
			}
		}
	}
}

func main() {
	opengl.Run(run)
}
