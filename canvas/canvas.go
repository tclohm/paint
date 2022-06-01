package canvas

import (
	"fyne.io/fyne/v2"
	"image"
	"github.com/tclohm/paint/apptype"
	"image/color"
)

type CanvasMouseState struct {
	prevCoord *fyne.PointEvent
}

type Canvas struct {
	widget.BaseWidget
	apptype.CanvasConfig
	renderer 	*CanvasRenderer
	PixelData 	image.Image
	mouseState 	CanvasMouseState
	appState	*apptype.State
	reloadImage bool
}

func (canvas *Canvas) Bounds() image.Rectangle {
	x0 := int(canvas.CanvasOffset.X)
	y0 := int(canvas.CanvasOffset.Y)
	x1 := int(canvas.PxCols * canvas.PxSize + int(canvas.CanvasOffset.X))
	y1 := int(canvas.PxRows * canvas.PxSize + int(canvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
	   pos.X < float32(bounds.Max.X) &&
	   pos.Y >= float32(bounds.Min.Y) &&
	   pos.X < float32(bounds.Max.Y) {
	   	return true
	   }
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0 ; y < rows ; y++ {
		for x := 0 ; x < cols ; x++ {
			
		}
	}
}