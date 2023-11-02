package player

import (
	"image"
	"image/color"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

type Player struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
}

func NewPlayer(pos *pixel.Vec) *Player {
	// 1. Create an image with the the size of 32x64 for now
	width := 32
	heigh := 64
	surface := image.Rect(int(pos.X), int(pos.Y), int(pos.X)+width, int(pos.Y)+heigh)

	// 2. Create a tmpPic with the surface
	tmpPic := image.NewRGBA(surface)

	// 3. Set color for the tmpPic, e.g. Red
	// loop
	for x := surface.Min.X; x < surface.Max.X; x++ {
		for y := surface.Min.Y; y < surface.Max.Y; y++ {
			tmpPic.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}

	// 4. Create a pixel.PictureData from the tmpPic
	pic := pixel.PictureDataFromImage(tmpPic)

	// 5. Create a pixel.Sprite from the pic
	sprite := pixel.NewSprite(pic, pic.Bounds())

	return &Player{
		Sprite: sprite,
		Position: pixel.Vec{
			X: sprite.Picture().Bounds().Center().X,
			Y: sprite.Picture().Bounds().Center().Y,
		},
	}
}

func (p *Player) Update(xShift float64) {
	velocity := pixel.Vec{
		X: xShift,
		Y: 0,
	}

	p.Position = p.Position.Add(pixel.V(velocity.X, velocity.Y))

}

func (p *Player) Draw(surface *pixelgl.Window) {
	mat := pixel.IM.Moved(p.Position)
	p.Sprite.Draw(surface, mat)
}
