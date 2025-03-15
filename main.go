package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

// Custom Theme Definition
type highContrastTheme struct {  // The highContrastTheme struct embeds fyne.Theme, which is an interface.
	fyne.Theme
}
// The Color method overrides specific theme colors:
func (h *highContrastTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:  // theme.ColorNameBackground → color.White
		return color.White
	case theme.ColorNameInputBackground:
		return color.White // White background for inputs, yet they are not!
	case theme.ColorNameForeground:
		return color.Black // Black text
	}
	return h.Theme.Color(name, variant)
}  // For all other colors, it delegates to the embedded h.Theme (set to theme.LightTheme() later).
/*
Theme Delegation: Since fyne.Theme is an interface with methods like Color, Font, Icon, and Size, embedding it means highContrastTheme must either implement all these methods or 
rely on the embedded Theme field (set to theme.LightTheme()) to provide them. Here, only Color is overridden, but because h.Theme is initialized with 
theme.LightTheme() (a concrete implementation), the other methods (Font, Icon, Size) are promoted and handled by theme.LightTheme(). This is correct and should work.
	In Fyne, widget.Entry uses theme colors such that it should result in black text on a white background. But we are seeing black on black, the theme might not be applying correctly (more on this later).
*/

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&highContrastTheme{Theme: theme.LightTheme()}) // sets the custom theme, wrapping theme.LightTheme() as the base theme.
	// 	Since highContrastTheme delegates to theme.LightTheme() for unhandled methods, and theme.LightTheme() uses the light variant, this should enforce your color choices unless something overrides it
	myWindow := myApp.NewWindow("Archimedes Pi")
	myWindow.Resize(fyne.NewSize(1900, 1600))

	// Output and Scroll Container
	outputLabel := widget.NewLabel("Press a button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300)) // Creates a label with initial text and word wrapping enabled. Places it in a scrollable container with a minimum size of 1900x1300.

	// Initialize an empty vertical box container for input fields and hide it initially (shown later when needed).
	inputContainer := container.NewVBox()
	inputContainer.Hide()

	// Output Callback Function
	var outputText string
	callBkPrn2canvas := func(oneLineSansCR string) {
		outputText += oneLineSansCR + "\n"
		outputLabel.SetText(outputText)
		fmt.Println(oneLineSansCR)
		scrollContainer.ScrollToBottom()
	}
	/*
	outputText is a global variable storing the accumulated output. The callBkPrn2canvas function appends a line to outputText, updates the label, prints to console, and scrolls to the bottom.
	*/

	// Input Collection Function :
	/*
	Clears the inputContainer.
	Creates arrays for return values and entry fields based on the number of prompts.
	For each prompt: Adds a label and an entry field to a vertical box, then adds that to inputContainer.
	Adds a "Submit" button that: Collects text from entry fields into values; and hides the input container and signals completion via a channel.
	Shows the container and waits for submission before returning the values.
	*/
	getInputValues := func(prompts []string) []string {
		inputContainer.Objects = nil
		values := make([]string, len(prompts))
		entryFields := make([]*widget.Entry, len(prompts))
		done := make(chan bool)
		for i, prompt := range prompts {
			label := widget.NewLabel(prompt)
			entryFields[i] = widget.NewEntry()
			inputContainer.Add(container.NewVBox(label, entryFields[i]))
		}
		submitBtn := widget.NewButton("Submit", func() {
			for i, entry := range entryFields {
				values[i] = entry.Text
			}
			inputContainer.Hide()
			done <- true
		})
		inputContainer.Add(submitBtn)
		inputContainer.Show()
		<-done
		return values
	}


	// Archimedes Callback, needed ::: per below !!!
	archimedesCallback := func(btn *ColoredButton) {
		outputText = ""
		outputLabel.SetText(outputText)
		btn.BackgroundColor = color.RGBA{255, 50, 50, 255}
		btn.Refresh() // Resets output, changes button color to a darker red, refreshes it.  Runs ArchimedesBig in a goroutine (good for long computations), updates button back to lighter red.
		go func() {
			ArchimedesBig(callBkPrn2canvas)
			btn.BackgroundColor = color.RGBA{255, 100, 100, 255}
			btn.Refresh()
		}()
	}
	// ::: This first button would just fire off if coded like the rest of the buttons!!!
	buttonArchimedes := NewColoredButton("Archimedes", color.RGBA{255, 100, 100, 255}, nil)
        buttonArchimedes.OnTapped = func() {  // Wait for a mouse click -- this traps us here ::: until a tap occurs -- on any button!!!
		archimedesCallback(buttonArchimedes)
		/* ::: Not this: 
			outputText = ""
			outputLabel.SetText(outputText)
			go ArchimedesBig(callBkPrn2canvas)
		 */
	} 

	buttonLeibniz := NewColoredButton("Gottfried Wilhelm Leibniz", color.RGBA{100, 255, 100, 255}, func() {
		outputText = ""
		outputLabel.SetText(outputText)
		go GottfriedWilhelmLeibniz(callBkPrn2canvas)
	})
	
	
	buttonNilakantha := NewColoredButton("Nilakantha", color.RGBA{255, 255, 100, 255}, func() {
		outputText = ""
		outputLabel.SetText(outputText)
		go func() {
			callBkPrn2canvas("\nYou have selected the Nilakantha Somayaji method...")
			inputs := getInputValues([]string{
				"Enter the number of iterations (suggest between 100,000 and 100,000,000)",
				"Enter the precision: (suggest between 128 and 512)",
			})
			iters, err := strconv.Atoi(inputs[0]) // ::: Also need to specifically handle the case of non numerics etc. 
			if err != nil {
				callBkPrn2canvas("Invalid iterations input")
				return
			}
			precision, err := strconv.Atoi(inputs[1]) // ::: Also need to specifically handle the case of non numerics etc.
			if err != nil {
				callBkPrn2canvas("Invalid precision input")
				return
			}
			NilakanthaBig(callBkPrn2canvas, iters, precision)
		}()
	})

	buttonGregory := NewColoredButton("Gregory-Leibniz", color.RGBA{100, 100, 255, 255}, func() {
		outputText = ""
		outputLabel.SetText(outputText)
		go GregoryLeibniz(callBkPrn2canvas)
	})

	buttonChudnovsky := NewColoredButton("Chudnovsky", color.RGBA{255, 100, 255, 255}, func() {
		outputText = ""
		outputLabel.SetText(outputText)
		go ChudnovskyBig(callBkPrn2canvas)
	})

	buttonMonteCarlo := NewColoredButton("Monte Carlo", color.RGBA{100, 255, 255, 255}, func() {
		outputText = ""
		outputLabel.SetText(outputText)
		go MonteCarloBig(callBkPrn2canvas)
	})

	buttonExtra1 := NewColoredButton("Extra 1", color.RGBA{200, 200, 200, 255}, func() {
		callBkPrn2canvas("Extra 1 clicked")
	})
		buttonExtra2 := NewColoredButton("Extra 2", color.RGBA{150, 150, 150, 255}, func() {
			callBkPrn2canvas("Extra 2 clicked")
		})


	buttonContainer := container.NewGridWithColumns(4,
		buttonArchimedes,
		buttonLeibniz,
		buttonGregory,
		buttonNilakantha,
		buttonChudnovsky,
		buttonMonteCarlo,
		buttonExtra1,
		buttonExtra2,
	)

	content := container.NewVBox(buttonContainer, inputContainer, scrollContainer)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// Rest of your code (ColoredButton, renderer, etc.) unchanged...
