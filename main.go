package main

import (
	"fmt"
	"os"
	"time"

	_ "image/png"

	"github.com/drpaneas/pygame/pkg/display"
	"github.com/drpaneas/pygame/pkg/level"
	"github.com/gopxl/pixel/v2/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	display.Screen = display.SetMode(screenWidth, screenHeight)
	// fmt.Printf("resolution: %v x %v\n", screenWidth, screenHeight)
	// Current level map, resolution: 1280 x 704

	// sprite.Draw(screen, pixel.IM.Moved(screen.Bounds().Center()))

	// Configure the game loop to run in 60 FPS
	timeframe := time.Second / framesPerSecond
	clock := time.NewTicker(timeframe) // ticks every frame
	start := time.Now()
	frame := 0

	world := level.NewLevel(level.Map, display.Screen)

	for !display.Screen.Closed() {
		for range clock.C {
			// Called every tick of the clock (every frame)
			frame++

			if display.Screen.JustPressed(pixelgl.KeyK) {
				fmt.Println("K pressed - Frame: ", frame)
			}

			// if you get an event, and this event is pressing the quit button, then close the window
			if display.Screen.JustPressed(pixelgl.KeyEscape) {
				display.Screen.SetClosed(true)
				os.Exit(1)
			}

			display.Screen.Clear(colornames.Black)

			world.Run()

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
	pixelgl.Run(run)
}
