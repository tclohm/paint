package ui

import (
	"errors"
	"strconv"
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"

)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PaintCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PaintWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PaintWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}

			err = png.Encode(fh, app.PaintCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		// func literal
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				paintWidth := 0
				paintHeight := 0

				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PaintWindow)
				} else {
					paintWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PaintWindow)
				} else {
					paintWidth, _ = strconv.Atoi(heightEntry.Text)
				}
				app.PaintCanvas.NewDrawing(paintWidth, paintHeight)
			}
		}, app.PaintWindow)

	})
}


func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
	)
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PaintWindow.SetMainMenu(mainMenu)
}