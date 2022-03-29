package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/petrostrak/pixl/swatch"
)

func BuildSwatches(a *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(a.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(a.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(a.Swatches); j++ {
				a.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			a.State.SwatchSelected = s.SwatchIndex
			a.State.BrushColor = s.Color
		})

		if i == 0 {
			s.Selected = true
			a.State.SwatchSelected = 0
			s.Refresh()
		}

		a.Swatches = append(a.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
