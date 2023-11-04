package level

import (
	"github.com/drpaneas/pygame/pkg/display"
	"github.com/drpaneas/pygame/pkg/player"
	"github.com/drpaneas/pygame/pkg/tiles"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/pixelgl"
)

// Create a Map list of strings
// Each string is a row of the map
// Each character of the string is representing a tile
// 11 rows of 28 tiles each
var Map = [][]string{
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " "},
	{" ", "X", "X", "X", "X", " ", " ", "P", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", "X", " ", " ", " "},
	{" ", " ", " ", " ", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
	{"X", "X", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
}

type Level struct {
	Layout     [][]string
	Surface    *pixelgl.Window
	Tiles      tiles.TilesGroup
	Player     *player.Player
	WorldShift float64
}

func NewLevel(layout [][]string, surface *pixelgl.Window) *Level {
	l := &Level{
		Layout:     layout,
		Surface:    surface,
		WorldShift: 0,
	}

	l.SetupLevel(l.Layout)

	return l
}

func (l *Level) SetupLevel(layout [][]string) {
	for row_index, row := range layout {
		for col_index, cell := range row {
			x := col_index * tiles.Size
			y := (len(layout) - row_index - 1) * tiles.Size // invert Y-axis

			pos := &pixel.Vec{
				X: float64(x),
				Y: float64(y),
			}

			if cell == "X" {
				tile := tiles.NewTile(pos, tiles.Size)
				l.Tiles = append(l.Tiles, tile)
			}

			if cell == "P" {
				player := player.NewPlayer(pos)
				l.Player = player
			}
		}
	}
}

// ScrollX is the horizontal camera movement
func (l *Level) ScrollX() {
	playerX := l.Player.Position.X
	screenwidth := display.Screen.Bounds().W()
	directionX := l.Player.Direction.X

	if playerX < screenwidth/4 && directionX < 0 {
		l.WorldShift = 8
		l.Player.Speed = 0
	} else if playerX > screenwidth-(screenwidth/4) && directionX > 0 {
		l.WorldShift = -8
		l.Player.Speed = 0
	} else {
		l.WorldShift = 0
		l.Player.Speed = 8
	}
}

func (l *Level) HorizontalMovementCollision() {
	// l.Player.Position.X += l.Player.Direction.X * l.Player.Speed

	// Create a new velocity vector
	// Directions are either -1, 0 or 1 (left, none, right) and we multiply them by the speed
	velocity := pixel.Vec{
		X: l.Player.Direction.X * l.Player.Speed,
	}

	l.Player.Position = l.Player.Position.Add(velocity)

	// loop through the titles I could possibly collide with
	for _, tile := range l.Tiles {
		// if the player collides with a tile
		if l.Player.CollidesWith(tile) {
			// if the player is moving right, stop the player at the left side of the tile
			if l.Player.Direction.X > 0 {
				l.Player.Position.X = tile.Position.X - (tile.Sprite.Frame().W() / 2) - (l.Player.Sprite.Frame().W() / 2)
			} else if l.Player.Direction.X < 0 {
				// if the player is moving left, stop the player at the right side of the tile
				l.Player.Position.X = tile.Position.X + (tile.Sprite.Frame().W() / 2) + (l.Player.Sprite.Frame().W() / 2)
			}
		}
	}
}

func (l *Level) VerticalMovementCollision() {
	l.Player.ApplyGravity()

	// loop through the titles I could possibly collide with
	for _, tile := range l.Tiles {
		// if the player collides with a tile
		if l.Player.CollidesWith(tile) {
			// if the player is moving up, stop the player at the bottom side of the tile
			if l.Player.Direction.Y > 0 {
				l.Player.Position.Y = tile.Position.Y - (tile.Sprite.Frame().H() / 2) - (l.Player.Sprite.Frame().H() / 2)
				l.Player.Direction.Y = 0
			} else if l.Player.Direction.Y < 0 {
				// if the player is moving down, stop the player at the top side of the tile
				l.Player.Position.Y = tile.Position.Y + (tile.Sprite.Frame().H() / 2) + (l.Player.Sprite.Frame().H() / 2)
				l.Player.Direction.Y = 0
			}
		}
	}
}

func (l *Level) Run() {
	// Level Tiles
	l.Tiles.Update(l.WorldShift)
	l.Tiles.Draw(l.Surface)
	l.ScrollX()

	// Player
	l.Player.Update()
	l.HorizontalMovementCollision()
	l.VerticalMovementCollision()
	l.Player.Draw(l.Surface)
}
