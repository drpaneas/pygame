package player

import (
	"image"
	"image/color"

	"github.com/drpaneas/pygame/pkg/input"
	"github.com/drpaneas/pygame/pkg/keyboard"
	"github.com/drpaneas/pygame/pkg/tiles"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

// Player is the main character of the game
type Player struct {
	Sprite    *pixel.Sprite
	Position  pixel.Vec
	Direction pixel.Vec
	Speed     float64
	Gravity   float64
	JumpSpeed float64
	Status    string
	OnGround  bool
	OnCeiling bool
}

// NewPlayer creates a new player with the given position
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
		Status:    "idle",
		OnGround:  false,
		OnCeiling: false,
	}
}

// getInputSetDirection gets the input from the keyboard and sets the player's direction
func (p *Player) getInputSetDirection() {
	keys := input.GetKeys()

	if keys[keyboard.Right] {
		p.Direction.X = 1
	} else if keys[keyboard.Left] {
		p.Direction.X = -1
	} else {
		p.Direction.X = 0
	}

	// Player is a able to jump only if he is on the ground (sorry, no double jumps)
	if keys[keyboard.Space] && p.OnGround {
		p.jump()
	}
}

// GetStatus figures out what the player is doind
// a. jumping (up) - falling (down) -         walking (left/right)         -         idle (none)
// direction.y > 0   direction.y <0   direction.y ==0 && direction.x != 0   direction.y ==0 && direction.x == 0
func (p *Player) GetStatus() {
	if p.Direction.Y > 0 || p.OnCeiling {
		p.Status = "jump"
	} else if p.Direction.Y < 0 {
		p.Status = "fall"
	} else {
		if p.Direction.X != 0 {
			p.Status = "walk"
		} else {
			p.Status = "idle"
		}
	}
}

// ApplyGravity applies gravity to the player
func (p *Player) ApplyGravity() {
	// Falling down, means that the Y-axis is decreasing
	// Gravity is a negative number, so we add it to the direction
	p.Direction.Y += p.Gravity
	p.Position.Y += p.Direction.Y
}

// jump makes the player jump
func (p *Player) jump() {
	p.Direction.Y = p.JumpSpeed
}

// boundBox returns the player's bounding box
func (p *Player) boundBox() pixel.Rect {
	// Calculate the player's bounding box based on its position and size
	return pixel.R(
		p.Position.X-p.Sprite.Frame().W()/2,
		p.Position.Y-p.Sprite.Frame().H()/2,
		p.Position.X+p.Sprite.Frame().W()/2,
		p.Position.Y+p.Sprite.Frame().H()/2,
	)
}

// CollidesWithTile checks if the player collides with a tile
func (p *Player) CollidesWithTile(tile *tiles.Tile) bool {
	// Check if the player's bounding box intersects with the tile's bounding box
	intersection := p.boundBox().Intersect(tile.BoundBox())
	return intersection.Area() > 0
}

// Update updates the player logic
func (p *Player) Update() {
	p.getInputSetDirection()
}

// Draw draws the player
func (p *Player) Draw(surface *opengl.Window) {
	mat := pixel.IM.Moved(p.Position)
	p.Sprite.Draw(surface, mat)
}
