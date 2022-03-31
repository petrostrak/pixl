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

func (r *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := r.pxCanvas.CanvasOffset
	imgHeight := r.canvasImage.Size().Height
	imgWidth := r.canvasImage.Size().Width

	left := &r.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &r.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &r.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &r.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
