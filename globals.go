package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// @formatter:off




var copyOfLastPosition int

// convenience globals:

var usingBigFloats = false // a variable of type bool which is passed by many funcs to printResultStatsLong()

var iterationsForMonte16i int
var iterationsForMonte16j int
var iterationsForMonteTotal int
var four float64 // is initialized to 4 where needed
var π float64    // a var can be any character, as in this Pi symbol/character
var LinesPerSecond float64
var LinesPerIter float64
var iterInt64 int64     // to be used primarily in selections which require modulus calculations
var iterFloat64 float64 // to be used in selections which do not require modulus calculations
var Table_of_perfect_squares = []int{}
var t2 time.Time

// The following globals, are used in multiple funcs of case 18: calculate either square or cube root of any integer

var Tim_win float64             // Time Window
var sortedResults = []Results{} // sortedResults is an array of type Results as defined at the top of this file
var Table_of_perfect_Products = []int{}
var diffOfLarger int
var diffOfSmaller int
var perfectResult2 float64 // will contain the square root result if the workpiece was itself a perfect square
var perfectResult3 float64 // will contain the cube root result if the workpiece was itself a perfect cube
var precisionOfRoot int    // this being global means we do not need to pass it in to the read func
var workPiece int          // the square or cube of which we are to find a root
var skip_redoing_loop int

type Results struct { // define a new structure called Results with two fields; result, and pdiff
	result float64
	pdiff  float64
}


// ColoredButton type
type ColoredButton struct {
	widget.Button
	BackgroundColor color.Color
}

func NewColoredButton(label string, backgroundColor color.Color, tapped func()) *ColoredButton {
	btn := &ColoredButton{BackgroundColor: backgroundColor}
	btn.Text = label
	btn.OnTapped = tapped
	btn.ExtendBaseWidget(btn)
	return btn
}

func (b *ColoredButton) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(b.Text, color.Black)
	text.Alignment = fyne.TextAlignCenter
	background := canvas.NewRectangle(b.BackgroundColor)
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeColor = color.Gray{0x80}
	border.StrokeWidth = 2
	return &coloredButtonRenderer{
		button:     b,
		text:       text,
		background: background,
		border:     border,
		objects:    []fyne.CanvasObject{background, border, text},
	}
}

// Custom renderer
type coloredButtonRenderer struct {
	button     *ColoredButton
	text       *canvas.Text
	background *canvas.Rectangle
	border     *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (r *coloredButtonRenderer) BackgroundColor() color.Color {
	return r.button.BackgroundColor
}

func (r *coloredButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.border.Resize(size)
	textSize := r.text.MinSize()
	r.text.Resize(fyne.NewSize(size.Width-20, textSize.Height))
	r.text.Move(fyne.NewPos(10, (size.Height-textSize.Height)/2)) // Center vertically
}

func (r *coloredButtonRenderer) MinSize() fyne.Size {
	textSize := r.text.MinSize()
	return fyne.NewSize(fyne.Max(textSize.Width+20, 200), fyne.Max(textSize.Height+20, 50))
}

func (r *coloredButtonRenderer) Refresh() {
	r.background.FillColor = r.button.BackgroundColor
	r.text.Text = r.button.Text
	r.background.Refresh()
	r.border.Refresh()
	r.text.Refresh()
}

func (r *coloredButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *coloredButtonRenderer) Destroy() {
	// No-op
}

// Theme
type myTheme struct {
	Theme fyne.Theme
}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{245, 245, 245, 255}
	case theme.ColorNameButton:
		return color.RGBA{200, 200, 200, 255}
	case theme.ColorNameForeground:
		return color.RGBA{0, 0, 0, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 165, 0, 255}
	}
	return m.Theme.Color(name, variant)
}

func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource    { return m.Theme.Font(style) }
func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource { return m.Theme.Icon(name) }
func (m *myTheme) Size(name fyne.ThemeSizeName) float32       { return m.Theme.Size(name) }
