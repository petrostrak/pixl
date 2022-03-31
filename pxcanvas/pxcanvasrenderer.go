package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) MinSize() fyne.Size {
	return r.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)

	for i := range r.canvasBorder {
		objects = append(objects, &r.canvasBorder[i])
	}

	objects = append(objects, r.canvasImage)

	return objects
}
