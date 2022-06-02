package apptype

import (
	"fyne.io/fyne/v2"
	"image/color"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

type CanvasConfig struct {
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSize int
}

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int)
}