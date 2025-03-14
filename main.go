package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// @formatter:off

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{Theme: theme.DefaultTheme()})

	myWindow := myApp.NewWindow("Archimedes Pi")
	myWindow.Resize(fyne.NewSize(1900, 1600)) 

	// Button to trigger the function
	button := widget.NewButton("Calculate Pi", nil) 
	
	// Label for output
	outputLabel := widget.NewLabel("Press the button to start...\n") 
	outputLabel.Wrapping = fyne.TextWrapWord

	// Scrollable container for the label
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300)) 

	// Vertical layout: button on top, scrollable output below 
	content := container.New(layout.NewVBoxLayout(), button, scrollContainer)

	// Fyne printing callback function to append output to scrollContainer and optionally print to terminal 
	var outputText string
	callBkPrn2canvas := func(oneLineSansCR string) { // Is this a callback func ???
		outputText += oneLineSansCR + "\n"
		outputLabel.SetText(outputText)
		fmt.Println(oneLineSansCR)        // Keep CLI output (print to terminal)
		scrollContainer.ScrollToBottom() // Auto-scroll to the latest output
	}
	
	// Button click handler 
	button.OnTapped = func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText) 
		go ArchimedesBig(callBkPrn2canvas) // I have several more of these, so I would like to have, say six, buttons, one button for each method of computing Pi
	}

	myWindow.SetContent(content) 
	myWindow.ShowAndRun()
}