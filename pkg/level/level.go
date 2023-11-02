package level

import (
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
	{" ", "X", "X", " ", "P", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", "X", " ", " ", " "},
	{" ", " ", " ", " ", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
	{"X", "X", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
}

type TilesGroup []*tiles.Tile

func (t TilesGroup) Draw(surface *pixelgl.Window) {
	for _, tile := range t {
		tile.Draw(surface)
	}
}

type Level struct {
	Layout     [][]string
	Surface    *pixelgl.Window
	Tiles      TilesGroup
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
		}
	}
}

func (t TilesGroup) Update(xShift float64) {
	for _, tile := range t {
		tile.Update(xShift)
	}
}

func (l *Level) Run() {
	l.Tiles.Update(l.WorldShift)
	l.Tiles.Draw(l.Surface)
}
