package ui

import (
	"fyne.io/fyne/v2"
	"github.com/tclohm/paint/apptype"
	"github.com/tclohm/paint/swatch"
	"github.com/tclohm/paint/paintcanvas"
)

type AppInit struct {
	PaintCanvas *paintcanvas.PaintCanvas
	PaintWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}