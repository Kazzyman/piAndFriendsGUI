package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// @formatter:off

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{Theme: theme.DefaultTheme()})

	myWindow := myApp.NewWindow("Archimedes Pi")
	myWindow.Resize(fyne.NewSize(1900, 1600))

	// Label for output
	outputLabel := widget.NewLabel("Press a button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord

	// Scrollable container for the label
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300))

	// Container for input fields - initially hidden
	inputContainer := container.NewVBox()
	inputContainer.Hide()

	// Fyne printing callback function
	var outputText string
	callBkPrn2canvas := func(oneLineSansCR string) {
		outputText += oneLineSansCR + "\n"
		outputLabel.SetText(outputText)
		fmt.Println(oneLineSansCR)
		scrollContainer.ScrollToBottom()
	}

	// Input handling function - returns values entered by user
	getInputValues := func(prompts []string) []string {
		// Clear any previous input fields
		inputContainer.Objects = nil

		values := make([]string, len(prompts))
		entryFields := make([]*widget.Entry, len(prompts))

		// Create a channel to signal when inputs are complete
		done := make(chan bool)

		// Add input fields and labels for each prompt
		for i, prompt := range prompts {
			label := widget.NewLabel(prompt)
			entryFields[i] = widget.NewEntry()
			inputContainer.Add(container.NewVBox(label, entryFields[i]))
		}

		// Add a submit button
		submitBtn := widget.NewButton("Submit", func() {
			for i, entry := range entryFields {
				values[i] = entry.Text
			}
			inputContainer.Hide()
			done <- true
		})

		inputContainer.Add(submitBtn)
		inputContainer.Show()

		// Wait for user to submit values
		<-done
		return values
	}

	// Create six buttons in a horizontal container
	buttonArchimedes := widget.NewButton("Archimedes", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		go func() {
			ArchimedesBig(callBkPrn2canvas)
		}()
	})

	buttonLeibniz := widget.NewButton("Gottfried Wilhelm Leibniz", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		go func() {
			GottfriedWilhelmLeibniz(callBkPrn2canvas)
		}()
	})

	buttonNilakantha := widget.NewButton("Nilakantha", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		go func() {
			callBkPrn2canvas("\nYou have selected the Nilakantha Somayaji method using big.Float types, and with some ")
			callBkPrn2canvas("patience one can generate 31 correct digits of pi using it.\n\n")

			inputs := getInputValues([]string{
				"Enter the number of iterations (suggest between 100,000 and 100,000,000)",
				"Enter the precision: (suggest between 128 and 512)",
			})

			iters, _ := strconv.Atoi(inputs[0])
			precision, _ := strconv.Atoi(inputs[1])

			// Replace this with your actual implementation
			NilakanthaBig(callBkPrn2canvas, iters, precision)
		}()
	})

	// Add the other buttons similarly...
	buttonGregory := widget.NewButton("Gregory-Leibniz", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		go GregoryLeibniz(callBkPrn2canvas)
	})

	buttonChudnovsky := widget.NewButton("Chudnovsky", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		// go ChudnovskyBig(callBkPrn2canvas)
	})

	buttonMonteCarlo := widget.NewButton("Monte Carlo", func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		// go MonteCarloBig(callBkPrn2canvas)
	})

	// Create horizontal container for buttons
	buttonContainer := container.New(layout.NewHBoxLayout(),
		buttonArchimedes,
		buttonLeibniz,
		buttonGregory,
		buttonNilakantha,
		buttonChudnovsky,
		buttonMonteCarlo)

	// Vertical layout: buttons on top, inputs in middle (when shown), scrollable output below 
	content := container.New(layout.NewVBoxLayout(), buttonContainer, inputContainer, scrollContainer)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
