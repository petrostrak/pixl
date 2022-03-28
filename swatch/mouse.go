package swatch

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/pixl/apptype"
)

type clickHandler func(*Swatch)

type Swatch struct {
	widget.BaseWidget
	Selected    bool
	Color       color.Color
	SwatchIndex int
	clickHandler
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHdlr clickHandler) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHdlr,
		SwatchIndex:  swatchIndex,
	}
	swatch.ExtendBaseWidget(swatch)

	return swatch
}
