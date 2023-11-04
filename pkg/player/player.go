package player

import (
	"image"
	"image/color"

	"github.com/drpaneas/pygame/pkg/input"
	"github.com/drpaneas/pygame/pkg/keyboard"
	"github.com/drpaneas/pygame/pkg/tiles"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

type Player struct {
	Sprite    *pixel.Sprite
	Position  pixel.Vec
	Direction pixel.Vec
	Speed     float64
	Gravity   float64
	JumpSpeed float64
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
		Speed:     8,
		Gravity:   -0.8,
		JumpSpeed: 16,
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

	if keys[keyboard.Space] {
		p.Jump()
	}
}

func (p *Player) ApplyGravity() {
	// Falling down, means that the Y-axis is decreasing
	// So, the direction has to be negative, that is why we subtract the gravity.
	p.Direction.Y += p.Gravity
	p.Position.Y += p.Direction.Y // Adding the direction to the position (aka subtracting the gravity, aka falling down)
}

func (p *Player) Jump() {
	p.Direction.Y = p.JumpSpeed
}

func (p *Player) Bounds() pixel.Rect {
	// Calculate the player's bounding box based on its position and size
	return pixel.R(
		p.Position.X-p.Sprite.Frame().W()/2,
		p.Position.Y-p.Sprite.Frame().H()/2,
		p.Position.X+p.Sprite.Frame().W()/2,
		p.Position.Y+p.Sprite.Frame().H()/2,
	)
}

func (p *Player) CollidesWith(tile *tiles.Tile) bool {
	// Check if the player's bounding box intersects with the tile's bounding box
	intersection := p.Bounds().Intersect(tile.Bounds())
	return intersection.Area() > 0
}

func (p *Player) Update() {
	p.GetInput()

	// // Create a new velocity vector
	// // Directions are either -1, 0 or 1 (left, none, right) and we multiply them by the speed
	// velocity := pixel.Vec{
	// 	X: p.Direction.X * p.Speed,
	// }

	// p.Position = p.Position.Add(velocity)

	// p.ApplyGravity()
}

func (p *Player) Draw(surface *pixelgl.Window) {
	mat := pixel.IM.Moved(p.Position)
	p.Sprite.Draw(surface, mat)
}
