package paintcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (paintcanvas *PaintCanvas) Scrolled(event *fyne.ScrollEvent) {
	paintcanvas.scale(int(event.Scrolled.DY))
	paintcanvas.Refresh()
}

func (paintcanvas *PaintCanvas) MouseMoved(event *desktop.MouseEvent) {
	paintcanvas.TryPan(paintcanvas.mouseState.prevCoord, event)
	paintcanvas.Refresh()
	paintcanvas.mouseState.prevCoord = &event.PointEvent
}

func (paintcanvas *PaintCanvas) MouseIn(event *desktop.MouseEvent) {}
func (paintcanvas *PaintCanvas) MouseOut() {}