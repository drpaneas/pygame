package tiles

import (
	"image"
	"image/color"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

const Size = 64

type Tile struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
}

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

func (t *Tile) Update(xShift float64) {
	velocity := pixel.Vec{
		X: xShift,
		Y: 0,
	}

	t.Position = t.Position.Add(pixel.V(velocity.X, velocity.Y))

}

func (t *Tile) Draw(surface *pixelgl.Window) {
	mat := pixel.IM.Moved(t.Position)
	t.Sprite.Draw(surface, mat)
}