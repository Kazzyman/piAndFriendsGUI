package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// @formatter:off

// Define a custom theme with larger text
type myTheme struct {
	fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 20 // Larger default text size
	}
	return theme.DefaultTheme().Size(name)
}

var copyOfLastPosition int

var ricksFirstWords = canvas.NewText("this is ricksFirstWords?", color.Black) // write something on the canvas
