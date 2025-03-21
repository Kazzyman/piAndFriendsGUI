package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// @formatter:off

// PiCalculator interface
type PiCalculator interface {
	Calculate(callback func(string))
	Name() string
}

// GregoryLeibniz 
type GregoryLeibnizWrapper struct {}
func (g GregoryLeibnizWrapper) Calculate(callback func(string)){
	go GregoryLeibniz(callback)
}
func (g GregoryLeibnizWrapper) Name() string { return "Gregory Leibniz" }

// ArchimedesBigWrapper
type ArchimedesBigWrapper struct{}
func (a ArchimedesBigWrapper) Calculate(callback func(string)) {
	go ArchimedesBig(callback)
}
func (a ArchimedesBigWrapper) Name() string { return "ArchimedesBig" }

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	window := myApp.NewWindow("Pi Estimation Demo")
	window.Resize(fyne.NewSize(1900, 1600))

	outputLabel := widget.NewLabel("\nPress a button to estimate π...\n\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewVScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1200))

	buttons := make([]*widget.Button, 0, 4)

	toggleButtons := func(enable bool) {
				// fmt.Println("Toggling buttons to enabled:", enable)
		for _, btn := range buttons {
			if enable {
				// fmt.Println("Enabling button:", btn.Text)
				btn.Enable()
			} else {
				// fmt.Println("Disabling button:", btn.Text)
				btn.Disable()
			}
		}
	}

	updateOutput := func(text string) {
				// fmt.Println("Updating output with:", text)
		current := outputLabel.Text
		if len(current) > 99500 { // was 1000
			current = current[len(current)-10:] // was -1000
		}
		outputLabel.SetText(current + text)
		outputLabel.Refresh()
		scrollContainer.ScrollToBottom()
	}

	calculators := []PiCalculator{
		ArchimedesBigWrapper{},
		GregoryLeibnizWrapper{},
	}

	// dynamically load the buttons array 
	for _, calc := range calculators {
		calc := calc // should calc be a shadow ??
		btn := widget.NewButton(calc.Name(), func() {
					// fmt.Println("Button clicked:", calc.Name())
			toggleButtons(false)
			updateOutput(fmt.Sprintf("\nRunning %s...\n\n", calc.Name()))
			calc.Calculate(func(result string) {
					// fmt.Println("Callback received with:", result)
				updateOutput(result)
				toggleButtons(true)
			})
		})
		buttons = append(buttons, btn)
	}

	var buttonObjects []fyne.CanvasObject
	for _, btn := range buttons {
		buttonObjects = append(buttonObjects, btn) // Convert []*widget.Button to []fyne.CanvasObject
	}

	content := container.NewVBox(
		widget.NewLabel("\nSelect a method to estimate π:\n"),
		container.NewGridWithColumns(2, buttonObjects...), // Pass properly converted slice
		scrollContainer,                                   // Added back to show outputLabel
	)

	window.SetContent(content)

	toggleButtons(true)
			// fmt.Println("App started, buttons enabled")

	window.ShowAndRun()
}
