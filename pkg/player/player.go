package player

import (
	"image"
	"image/color"

	"github.com/drpaneas/pygame/pkg/input"
	"github.com/drpaneas/pygame/pkg/keyboard"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

type Player struct {
	Sprite    *pixel.Sprite
	Position  pixel.Vec
	Direction pixel.Vec
	Speed     float64
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
		Direction: pixel.Vec{
			X: 0,
			Y: 0,
		},
		Speed: 8,
	}
}

func (p *Player) GetInput() {
	keys := input.GetKeys()

	if keys[keyboard.Right] {
		p.Direction.X = 1
	} else if keys[keyboard.Left] {
		p.Direction.X = -1
	} else {
		p.Direction.X = 0
	}
}

func (p *Player) Update() {
	p.GetInput()

	// Create a new velocity vector
	// Directions are either -1, 0 or 1 (left, none, right) and we multiply them by the speed
	velocity := pixel.Vec{
		X: p.Direction.X * p.Speed,
		Y: p.Direction.Y * p.Speed,
	}

	p.Position = p.Position.Add(velocity)
}

func (p *Player) Draw(surface *pixelgl.Window) {
	mat := pixel.IM.Moved(p.Position)
	p.Sprite.Draw(surface, mat)
}
