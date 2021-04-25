package cells

import (
	"golang.org/x/image/colornames"
	"image/color"
)

const (
	Air = iota
	Wall
	Snake
	FoodBanana
	FoodCherry
	FoodCoconut
)

var Palette = []color.Color{
	colornames.Darkcyan,
	colornames.Cyan,
	colornames.Green,
	colornames.Yellowgreen,
	colornames.Darkred,
	colornames.Antiquewhite,
}
