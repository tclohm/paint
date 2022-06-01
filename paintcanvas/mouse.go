package paintcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/tclohm/paint/paintcanvas/brush"
)

func (paintcanvas *PaintCanvas) Scrolled(event *fyne.ScrollEvent) {
	paintcanvas.scale(int(event.Scrolled.DY))
	paintcanvas.Refresh()
}

func (paintcanvas *PaintCanvas) MouseMoved(event *desktop.MouseEvent) {
	// hold and paint
	if x, y := paintcanvas.MouseToCanvasXY(event); x != nil && y != nil {
		brush.TryBrush(paintcanvas.appState, paintcanvas, event)
		cursor := brush.Cursor(paintcanvas.CanvasConfig, paintcanvas.appState.BrushType, event, *x, *y)
		paintcanvas.renderer.SetCursor(cursor)
	} else {
		paintcanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	// pan
	paintcanvas.TryPan(paintcanvas.mouseState.prevCoord, event)
	paintcanvas.Refresh()
	paintcanvas.mouseState.prevCoord = &event.PointEvent
}

func (paintcanvas *PaintCanvas) MouseIn(event *desktop.MouseEvent) {}
func (paintcanvas *PaintCanvas) MouseOut() {}

func (paintcanvas *PaintCanvas) MouseUp(event *desktop.MouseEvent) {}
func (paintcanvas *PaintCanvas) MouseDown(event *desktop.MouseEvent) {
	brush.TryBrush(paintcanvas.appState, paintcanvas, event)
}
