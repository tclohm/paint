package ui

import (
	"fyne.io/fyne/v2"
	"github.com/tclohm/paint/apptype"
	"github.com/tclohm/paint/swatch"
)

type AppInit struct {
	PaintWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}