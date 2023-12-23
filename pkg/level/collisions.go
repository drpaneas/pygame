package level

import "github.com/gopxl/pixel/v2"

// horizontalMovementCollision checks for horizontal collisions
func (l *Level) horizontalMovementCollision() {
	// Create a new velocity vector
	// Directions are either -1, 0 or 1 (left, none, right) and we multiply them by the speed
	velocity := pixel.Vec{
		X: l.Player.Direction.X * l.Player.Speed,
	}

	l.Player.Position = l.Player.Position.Add(velocity)

	// loop through the titles I could possibly collide with
	for _, tile := range l.Tiles {
		// if the player collides with a tile
		if l.Player.CollidesWithTile(tile) {
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

// verticalMovementCollision checks for vertical collisions
func (l *Level) verticalMovementCollision() {
	l.Player.ApplyGravity()

	// loop through the titles I could possibly collide with
	for _, tile := range l.Tiles {
		// if the player collides with a tile
		if l.Player.CollidesWithTile(tile) {
			// if the player is moving up, he has hit the ceiling, stop the player at the bottom side of the tile
			if l.Player.Direction.Y > 0 {
				l.Player.Position.Y = tile.Position.Y - (tile.Sprite.Frame().H() / 2) - (l.Player.Sprite.Frame().H() / 2)
				l.Player.Direction.Y = 0
				l.Player.OnCeiling = true
			} else if l.Player.Direction.Y < 0 {
				// if the player is moving down, he has hit the floor, stop the player at the top side of the tile
				l.Player.Position.Y = tile.Position.Y + (tile.Sprite.Frame().H() / 2) + (l.Player.Sprite.Frame().H() / 2)
				l.Player.Direction.Y = 0
				l.Player.OnGround = true
			}
		}
	}

	// If the player has collided with the floor, he is on the ground
	// but then if he either jumps or falls, he is no longer on the ceiling
	if l.Player.OnGround {
		if l.Player.Direction.Y > 0 {
			l.Player.OnGround = false
		} else if l.Player.Direction.Y < 0 {
			l.Player.OnGround = false
		}
	}

	// If the player has collided with the ceiling, he is on the ceiling
	// but if he falls, he is no longer on the ceiling
	if l.Player.OnCeiling {
		if l.Player.Direction.Y < 0 {
			l.Player.OnCeiling = false
		}
	}
}
