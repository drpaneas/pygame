package level

import (
	"github.com/drpaneas/pygame/pkg/player"
	"github.com/drpaneas/pygame/pkg/tiles"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

// Level is a major component of the game
type Level struct {
	Layout     [][]string       // the map
	Surface    *opengl.Window   // where to draw the level
	Tiles      tiles.TilesGroup // the tiles for the whole level
	Player     *player.Player   // the player
	WorldShift float64          // the horizontal camera movement
}

// NewLevel creates a new level with the given layout and surface
func NewLevel(layout [][]string, surface *opengl.Window) *Level {
	l := &Level{
		Layout:     layout,
		Surface:    surface,
		WorldShift: 0,
	}

	l.putTilesAndPlayer(l.Layout)

	return l
}

// putTilesAndPlayer puts the tiles and the player into the level based on the layout
func (l *Level) putTilesAndPlayer(layout [][]string) {
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

func (l *Level) UpdateAndDraw() {
	// Level Tiles
	l.Tiles.Update(l.WorldShift)
	l.Tiles.Draw(l.Surface)
	l.scrollX()

	// Player
	l.Player.Update()
	l.horizontalMovementCollision()
	l.verticalMovementCollision()
	l.Player.GetStatus()
	l.Player.Draw(l.Surface)
}
