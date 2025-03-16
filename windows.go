package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"sync"
	"time"
)

// @formatter:off

// Three Additional Windows: 
// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow2(myApp fyne.App) fyne.Window {
	window2 := myApp.NewWindow("Classic Pi calculators")
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel := widget.NewLabel("Classic Pi calculators, make a selection")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300))
	// var outputText string
	updateChan := make(chan updateData, 100) // Changed to struct
	var mu sync.Mutex
	callBkPrn2canvas := func(oneLineSansCR string) {
		updateChan <- updateData{text: oneLineSansCR}
	}
	buttonContainer := container.NewGridWithColumns(4,
		widget.NewButton("Button 1", func() {}),
		NewColoredButton("Button 2 has two lines\nlike so\nArchimedesBig ", color.RGBA{255, 100, 100, 255}, func() {
			updateChan <- updateData{clearText: true}
			go ArchimedesBig(callBkPrn2canvas)
		}),
		widget.NewButton("Button 3", func() {}),
		widget.NewButton("Button 4", func() {}),
		widget.NewButton("Button 5", func() {}),
		widget.NewButton("Button 6", func() {}),
		widget.NewButton("Button 7", func() {}),
		widget.NewButton("Button 8", func() {}),
	)

	content := container.NewVBox(buttonContainer, scrollContainer)
	window2.SetContent(content)
	// Main-thread update loop using Fyne's lifecycle
	window2.Canvas().SetOnTypedRune(func(r rune) {
		// Dummy handler to keep canvas active
	})
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case data := <-updateChan:
				mu.Lock()
				if data.clearText {
					outputLabel.SetText("") // Clear the label immediately
				}
				outputLabel.SetText(outputLabel.Text + data.text + "\n") // Append and update immediately
				scrollContainer.ScrollToBottom()
				fmt.Println(data.text) // Print each line as it's added
				mu.Unlock()
			default:
				// No need for the default case anymore
			}
		}
	}()
	return window2
}

// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow3(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window3 := myApp.NewWindow("Odd Pi calculators")
	window3.Resize(fyne.NewSize(1900, 1600))
	outputLabel := widget.NewLabel("Odd Pi calculators, make a selection")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer := container.NewGridWithColumns(4,
		widget.NewButton("Button 9", func() {}), widget.NewButton("Button 10", func() {}), widget.NewButton("Button 11", func() {}), widget.NewButton("Button 12", func() {}),
		widget.NewButton("Button 13", func() {}), widget.NewButton("Button 14", func() {}), widget.NewButton("Button 15", func() {}), widget.NewButton("Button 16", func() {}),
	)
	content := container.NewVBox(buttonContainer, scrollContainer)
	window3.SetContent(content)
	return window3
}

// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow4(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window4 := myApp.NewWindow("Misc Maths")
	window4.Resize(fyne.NewSize(1900, 1600))
	outputLabel := widget.NewLabel("Misc Maths, make a selection")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer := container.NewGridWithColumns(4,
		widget.NewButton("Button 17", func() {}), widget.NewButton("Button 18", func() {}), widget.NewButton("Button 19", func() {}), widget.NewButton("Button 20", func() {}),
		widget.NewButton("Button 21", func() {}), widget.NewButton("Button 22", func() {}), widget.NewButton("Button 23", func() {}), widget.NewButton("Button 24", func() {}),
	)
	content := container.NewVBox(buttonContainer, scrollContainer)
	window4.SetContent(content)
	return window4
}
