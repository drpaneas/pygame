# Choose a resolution

For 2D retro-style games, the graphic assets (aka `sprites`) are usually in the following sizes:

* 8x8
* 16x16
* 32x32
* 64x64
* 128x128

Back in the 80s and 90s, this was fine, but today we have a problem.

The problem is that people have 4K monitors, so all these graphics will look very tiny on their display.

To remedy this, we have to scale the graphics up to an appropriate size, so they look better in modern TV screens.

Just make sure you respect the original aspect ratio, and don't stretch it.

In other words, apply: Integer Scaling.

Integer scaling is a method of upscaling pixel art without introducing blurring or artifacts.

As the name implies, it is achieved by multiplying each pixel by a whole number (by saying "whole" I mean a number of type `int`), such as 2, 3, or 4.

To do that, you need two things:

1. Correct resolution
2. Nearest Neighbor Interpolation

To figoure out our resolution, we look at our levelmap.
How many pixels high and how many pixels wide should be displayed on screen?

Here's my tilemap:

```go
// Trapeza stin diey8ynsi Bausparr
// Stefan giorti xristougennon
// Create a Map list of strings
//  - Each character of the string is representing a tile
//  - 11 rows of 28 tiles each
var Map = [][]string{
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " "},
	{" ", "X", "X", " ", "P", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", "X", " ", " ", " "},
	{" ", " ", " ", " ", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
	{"X", "X", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
}
```

The tilemap is a 2D grid of tiles, where each tile represents a small part of the game level.

The tilemap is represented as a list of strings:

```go
var Map []][]string{str1, str2, str3 ...}
```

where each string represents a tile.


The characters in each string represent the type of tile at that position, such as an empty space, a wall, or a player character.

```go
" " // empty space where the player can freely move
"X" // wall, the player cannot go there
"P" // the starting position for the player when the game starts
```

In this specific tilemap, the X characters represent walls, the P character represents the player character, and the empty spaces represent areas where the player can move.

The tilemap is used to create the game level by rendering each tile at its appropriate position.

For example, the player is position at Map[4][4]

```go
for row_index, row := range layout {
	for col_index, cell := range row {
		if cell == "P" {
			// print coordinates
			fmt.Printf("Player position: %v, %v\n", row_index, col_index)
		}
    }
}
```

But this kind of translating the position is not very helpful.

We need to know which pixels we have to draw in the screen.

Thankfully, there is a way to answer that.

Each tile, in my case is made of 64x64pixels:

```go
const Size = 64
```

That means the first tile, at position `Map[0][0]` that is the top left corner of the screen, has screen coordinates from `x=y=0` to `x=y=63`.

The second one, at position `Map[0][1]`, so first row and second column, will be from `x=64` to `x=127` and on the same heigh as previously, from `y=0` to `y=63`.

So the `Map[4][4]` would be X=256 to X=319 and Y=256 to Y=319.

To calculate that, the math formula is:

```go
x := col_index * tiles.Size
y := (len(Map) - row_index - 1) * tiles.Size // invert Y-axis
```

I have to invert the Y-axis because the origin for drawing using `gopixel` is at the bottom left corner, instead of the top left corner.

So the resolution, to be 1:1, it should be:

```go
var (
	screenWidth  float64 = 1280
	screenHeight float64 = float64(len(level.Map)) * tiles.Size // to have integer scaling
)
```

That said, the function becomes like this:

```go
func (l *Level) SetupLevel(layout [][]string) {
	for row_index, row := range layout {
		for col_index, cell := range row {
            // Get the 'from', X and Y coordinates:
			x := col_index * tiles.Size
			y := (len(layout) - row_index - 1) * tiles.Size // invert Y-axis
			 
			if cell == "X" {
				// load the tile represented by the 'X' string
			}
		}
	}
}
```

Integer scaling is a popular technique for pixel art games because it preserves the sharp edges and clarity of the pixel art.

This is especially important for retro-style games, where the pixel art is a key part of the aesthetic.

Ok, how can I do it?

Great question.
The most common approach is to use the  **nearest neighbor** interpolation.

For scaling pixel art, it is recommended to use nearest neighbor interpolation. This is because it preserves the sharp edges and clarity of the pixel art, which is especially important for retro-style games where the pixel art is a key part of the aesthetic. Linear interpolation can introduce blurring or artifacts, which is not desirable for pixel art.

The library we are going to use, [uses by default nearest (pixely) neighbor interpolation](https://github.com/gopxl/glhf/blob/90c5b7c2543df34700f5ac1562870dc349b46abc/texture.go#L120-L133).

```go
// SetSmooth sets whether the Texture should be drawn "smoothly" or "pixely".
//
// It affects how the Texture is drawn when zoomed. Smooth interpolates between the neighbour
// pixels, while pixely always chooses the nearest pixel.
func (t *Texture) SetSmooth(smooth bool) {
	t.smooth = smooth
	if smooth {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	} else {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	}
}
```

In Go, the default value for boolean variables is `false`, hence the `else` part is the default codepath to be triggered here.

Personally, I prefer to be explicit, so I am going to hardcode this to be sure of it.

```go
win, err := pixelgl.NewWindow(cfg)
if err != nil {
	log.Fatalf("Failed to create window: %v", err)
}

win.SetSmooth(false) // use Nearest Neighbor Interpolation (aka pixelated)
```