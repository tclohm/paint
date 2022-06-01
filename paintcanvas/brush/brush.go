package brush

import (
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/tclohm/paint/apptype"
)

const (
	Pixel = iota
)

func TryBrush(appState *apptype.State, canvas apptype.Brushable, event *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, event)
	default:
		return false
	}
}

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, event *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(event)
	if x != nil && y != nil && event.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}