package paintcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/driver/desktop"
	"image"
	"github.com/tclohm/paint/apptype"
	"image/color"
)

type CanvasMouseState struct {
	prevCoord *fyne.PointEvent
}

type PaintCanvas struct {
	widget.BaseWidget
	apptype.CanvasConfig
	renderer 	*CanvasRenderer
	PixelData 	image.Image
	mouseState 	CanvasMouseState
	appState	*apptype.State
	reloadImage bool
}

func (paintcanvas *PaintCanvas) Bounds() image.Rectangle {
	x0 := int(paintcanvas.CanvasOffset.X)
	y0 := int(paintcanvas.CanvasOffset.Y)
	x1 := int(paintcanvas.PxCols * paintcanvas.PxSize + int(paintcanvas.CanvasOffset.X))
	y1 := int(paintcanvas.PxRows * paintcanvas.PxSize + int(paintcanvas.CanvasOffset.Y))
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

func NewCanvas(state *apptype.State, config apptype.CanvasConfig) *PaintCanvas {
	paintcanvas := &PaintCanvas{
		CanvasConfig: config,
		appState: state,
	}
	paintcanvas.PixelData = NewBlankImage(paintcanvas.PxCols, paintcanvas.PxRows, color.NRGBA{128, 128, 128, 255})
	paintcanvas.ExtendBaseWidget(paintcanvas)
	return paintcanvas
}

func (paintcanvas *PaintCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(paintcanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0 ; i < len(canvasBorder) ; i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &CanvasRenderer {
		paintCanvas: paintcanvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,
	}

	paintcanvas.renderer = renderer
	return renderer
}

func (paintcanvas *PaintCanvas) TryPan(prevCoord *fyne.PointEvent, event *desktop.MouseEvent) {
	if prevCoord != nil && event.Button == desktop.MouseButtonSecondary {
		paintcanvas.Pan(*prevCoord, event.PointEvent)
	}
}

// Brushable interface
func (paintcanvas *PaintCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := paintcanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}

	if rgba, ok := paintcanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}

	paintcanvas.Refresh()
}

func (paintcanvas *PaintCanvas) MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int) {
	bounds := paintcanvas.Bounds()

	if !InBounds(event.Position, bounds) {
		return nil, nil
	}

	canvasSize := float32(paintcanvas.PxSize)
	xOffset := paintcanvas.CanvasOffset.X
	yOffset := paintcanvas.CanvasOffset.Y

	x := int((event.Position.X - xOffset) / canvasSize)
	y := int((event.Position.Y - yOffset) / canvasSize)

	return &x, &y
}


