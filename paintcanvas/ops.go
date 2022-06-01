package paintcanvas

import "fyne.io/fyne/v2"

func (paintcanvas *PaintCanvas) Pan(prevCoord, currCoord fyne.PointEvent) {
	xDiff := currCoord.Position.X - prevCoord.Position.X
	yDiff := currCoord.Position.Y - prevCoord.Position.Y

	paintcanvas.CanvasOffset.X += xDiff
	paintcanvas.CanvasOffset.Y += yDiff

	paintcanvas.Refresh()
}

func (paintcanvas *PaintCanvas) scale(direction int) {
	switch {
	case direction > 0:
		paintcanvas.PxSize += 1
	case direction < 0:
		if paintcanvas.PxSize > 2 {
			paintcanvas.PxSize -= 1
		}
	default:
		paintcanvas.PxSize = 10
	}
}