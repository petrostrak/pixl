package ui

import (
	"fyne.io/fyne/v2"
	"github.com/petrostrak/pixl/apptype"
	"github.com/petrostrak/pixl/pxcanvas"
	"github.com/petrostrak/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
