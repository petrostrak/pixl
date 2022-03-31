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

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Destroy() {}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Layout(size fyne.Size) {}

func (r *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := r.pxCanvas.PxCols
	imgPxHeight := r.pxCanvas.PxRows
	pxSize := r.pxCanvas.PxSize

	r.canvasImage.Move(fyne.NewPos(r.pxCanvas.CanvasOffset.X, r.pxCanvas.CanvasOffset.Y))
	r.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}
