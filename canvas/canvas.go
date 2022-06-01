package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
			img.Set(x, y, c)
		}
	}

	return img
}

func NewCanvas(state *apptype.State, config apptype.CanvasConfig) *Canvas {
	canvas := &Canvas{
		CanvasConfig: config,
		appState: state,
	}
	canvas.PixelData = NewBlankImage(canvas.PxCols, canvas.PxRows, color.NRGBA{128, 128, 128, 255})
	canvas.ExtendBaseWidget(canvas)
	return canvas
}

func (canvas *Canvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(canvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0 ; i < len(canvas) ; i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &CanvasRenderer {
		canvas: canvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,
	}

	canvas.renderer = renderer
	return renderer
}