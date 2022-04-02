package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (r *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	r.canvasCursor = objects
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
	objects = append(objects, r.canvasCursor...)

	return objects
}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Destroy() {}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Layout(size fyne.Size) {
	r.LayoutCanvas(size)
	r.LayoutBorder(size)
}

// WidgetRenderer interface implementation
func (r *PxCanvasRenderer) Refresh() {
	if r.pxCanvas.reloadImage {
		r.canvasImage = canvas.NewImageFromImage(r.pxCanvas.PixelData)
		r.canvasImage.ScaleMode = canvas.ImageScalePixels
		r.canvasImage.FillMode = canvas.ImageFillContain
		r.pxCanvas.reloadImage = false
	}

	r.Layout(r.pxCanvas.Size())
	canvas.Refresh(r.canvasImage)
}

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
