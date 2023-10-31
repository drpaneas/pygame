Trying to convert the Pygame tutorial https://www.youtube.com/watch?v=YWN8GcmJ-jA into Golang and GoPxl


One of the the most important topics for pixel-freaks is the scaling 1:1 the resolution.

```python
def set_mode(
    size: Coordinate = (0, 0),
    flags: int = 0,
    depth: int = 0,
    display: int = 0,
    vsync: int = 0,
) -> Surface: ...
```

This looks quite familiar with:

```go
type Window struct {
	window *glfw.Window
	bounds             pixel.Rect
	canvas             *Canvas
	vsync              bool
	cursorVisible      bool
	cursorInsideWindow bool
}
```

So from the pygame's:

```python
screen = pygame.display.set_mode((screen_width, screen_height))
```

I can do the Go equiv:

```go
screen := display.SetMode(screenWidth, screenHeight)
```