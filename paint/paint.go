package main

import (
	"fyne.io/fyne/v2/app"
	"image/color"
	"github.com/tclohm/paint/ui"
	"github.com/tclohm/paint/apptype"
	"github.com/tclohm/paint/swatch"
)


func main() {
	paintApp := app.New()
	paintWindow := paintApp.NewWindow("Paint")

	state := apptype.State{
		BrushColor: color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PaintWindow: paintWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}