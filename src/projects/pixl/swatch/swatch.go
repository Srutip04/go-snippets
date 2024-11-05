package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"srutip04.io/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected bool
	Color color.Color
	SwatchIndex int
	clickHandler func(s* Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func  NewSwatch(color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{Color: color, SwatchIndex: swatchIndex, clickHandler: clickHandler}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	object := []fyne.CanvasObject{square}
	return &SwatchRenderer{square: *square, objects: object, parent: s}
}