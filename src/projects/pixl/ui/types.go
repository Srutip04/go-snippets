package ui

import (
	"fyne.io/fyne/v2"
	"srutip04.io/pixl/apptype"
	"srutip04.io/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State 	*apptype.State
	Swatches []*swatch.Swatch
}