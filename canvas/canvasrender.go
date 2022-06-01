package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CanvasRenderer struct {
	canvas *Canvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line
}

// widgetrenderer interface implementation
func (renderer *CanvasRenderer) MinSize() fyne.Size {
	return renderer.canvas.DrawingArea
}

// widgetrenderer interface implementation
func (renderer *CanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0 ; i < len(renderer.canvasBorder) ; i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

// widgetrenderer interface implementation
func (renderer *CanvasRenderer) Destroy() {}

// widgetrenderer interface implementation
func (renderer *CanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

func (renderer *CanvasRenderer) Refresh() {
	if renderer.canvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.canvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.canvas.reloadImage = false
	}
	renderer.Layout(renderer.canvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

func (renderer *CanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgWidth := renderer.canvas.PxCols
	imgHeight := renderer.canvas.PxRows
	size := renderer.canvas.PxSize
	renderer.canvasImage.Move(fyne.NewPos(renderer.canvas.CanvasOffset.X, renderer.canvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgWidth*size), float32(imgHeight*PxSize)))
}

func (renderer *CanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.canvas.CanvasOffset
	imgWidth := renderer.canvasImage.Size().Width
	imgHeight := renderer.canvasImage.Size().Height
	
	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y + imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X + imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y + imgHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y + imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y + imgHeight)

}