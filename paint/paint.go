package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"image/color"
	"github.com/tclohm/paint/ui"
	"github.com/tclohm/paint/apptype"
	"github.com/tclohm/paint/swatch"
	"github.com/tclohm/paint/paintcanvas"
)


func main() {
	paintApp := app.New()
	paintWindow := paintApp.NewWindow("Paint")

	state := apptype.State{
		BrushColor: color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	paintCanvasConfig := apptype.CanvasConfig{
		DrawingArea: fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0,0),
		PxRows: 10,
		PxCols: 10,
		PxSize: 30,
	}

	paintCanvas := paintcanvas.NewCanvas(&state, paintCanvasConfig)

	appInit := ui.AppInit{
		PaintCanvas: paintCanvas,
		PaintWindow: paintWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}