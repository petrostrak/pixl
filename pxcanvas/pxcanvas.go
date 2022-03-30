package pxcanvas

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}
