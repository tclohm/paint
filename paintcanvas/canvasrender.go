package paintcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CanvasRenderer struct {
	paintCanvas *PaintCanvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line
}

// widgetrenderer interface implementation
func (renderer *CanvasRenderer) MinSize() fyne.Size {
	return renderer.paintCanvas.DrawingArea
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
	if renderer.paintCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.paintCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.paintCanvas.reloadImage = false
	}
	renderer.Layout(renderer.paintCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

func (renderer *CanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgWidth := renderer.paintCanvas.PxCols
	imgHeight := renderer.paintCanvas.PxRows
	paintSize := renderer.paintCanvas.PxSize
	renderer.canvasImage.Move(fyne.NewPos(renderer.paintCanvas.CanvasOffset.X, renderer.paintCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgWidth*paintSize), float32(imgHeight*paintSize)))
}

func (renderer *CanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.paintCanvas.CanvasOffset
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