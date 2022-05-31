package ui

import (
	"image/color"
	"github.com/lusingander/colorpicker"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
)

func SetupColorPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(c color.Color) {
		app.State.BrushColor = c
		app.Swatches[app.State.SwatchSelected].SetColor(c)
	})
	return container.NewVBox(picker)
}