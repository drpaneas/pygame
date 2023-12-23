package tiles

import "github.com/gopxl/pixel/v2/backends/opengl"

type TilesGroup []*Tile

// Draw draws the tilesgroup
func (t TilesGroup) Draw(surface *opengl.Window) {
	for _, tile := range t {
		tile.Draw(surface)
	}
}

// Update updates the tilesgroup
func (t TilesGroup) Update(xShift float64) {
	for _, tile := range t {
		tile.Update(xShift)
	}
}
