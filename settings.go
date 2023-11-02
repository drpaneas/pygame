package main

import (
	"github.com/drpaneas/pygame/pkg/level"
	"github.com/drpaneas/pygame/pkg/tiles"
)

var (
	screenWidth  float64 = 1200
	screenHeight float64 = float64(len(level.Map)) * tiles.Size // to have integer scaling
)

const (
	framesPerSecond = 60
)
