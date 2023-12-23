package level

import "github.com/drpaneas/pygame/pkg/display"

// scrollX is the horizontal camera movement
func (l *Level) scrollX() {
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
