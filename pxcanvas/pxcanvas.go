package pxcanvas

import (
	"image"
	"image/color"

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

func (pxc *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxc.CanvasOffset.X)
	y0 := int(pxc.CanvasOffset.Y)
	x1 := int(pxc.PxCols*pxc.PxSize + int(pxc.CanvasOffset.X))
	y1 := int(pxc.PxRows*pxc.PxSize + int(pxc.CanvasOffset.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Min.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.Y < float32(bounds.Min.Y) {
		return true
	}

	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}

	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}

	pxCanvas.PixelData = NewBlankImage(config.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)

	return pxCanvas
}
