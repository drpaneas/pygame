package tiles

import (
	"image"
	"image/color"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

// Size of the tile
const Size = 64

// Tile is a major component of the level design
type Tile struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
}

// NewTile creates a new tile with the given position and size
func NewTile(pos *pixel.Vec, tilesize int) *Tile {
	// create a pixel.Picture with pos Bounds and size
	surface := image.Rect(int(pos.X), int(pos.Y), int(pos.X)+tilesize, int(pos.Y)+tilesize)
	// create a tmpPic with the surface
	tmpPic := image.NewRGBA(surface)
	// set color for the tmpPic, e.g. white
	// loop
	for x := surface.Min.X; x < surface.Max.X; x++ {
		for y := surface.Min.Y; y < surface.Max.Y; y++ {
			tmpPic.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	// create a pixel.PictureData from the tmpPic
	pic := pixel.PictureDataFromImage(tmpPic)
	// create a pixel.Sprite from the pic
	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &Tile{
		Sprite: sprite,
		Position: pixel.Vec{
			X: sprite.Picture().Bounds().Center().X,
			Y: sprite.Picture().Bounds().Center().Y,
		},
	}
}

// BoundBox returns the bounding box of the tile
func (t *Tile) BoundBox() pixel.Rect {
	// Calculate the player's bounding box based on its position and size
	return pixel.R(
		t.Position.X-t.Sprite.Frame().W()/2,
		t.Position.Y-t.Sprite.Frame().H()/2,
		t.Position.X+t.Sprite.Frame().W()/2,
		t.Position.Y+t.Sprite.Frame().H()/2,
	)
}

// Update updates the tile's position
func (t *Tile) Update(xShift float64) {
	velocity := pixel.Vec{
		X: xShift,
		Y: 0,
	}

	t.Position = t.Position.Add(velocity)

}

// Draw draws the tile
func (t *Tile) Draw(surface *opengl.Window) {
	mat := pixel.IM.Moved(t.Position)
	t.Sprite.Draw(surface, mat)
}
