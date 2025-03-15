package main

import (
	"fmt"
	"image/color"
	"strconv"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type highContrastTheme struct {
	fyne.Theme
}

func (h *highContrastTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.White
	case theme.ColorNameInputBackground:
		return color.White
	case theme.ColorNameForeground:
		return color.Black
	}
	return h.Theme.Color(name, variant)
}

type updateData struct {
	text      string
	clearText bool
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&highContrastTheme{Theme: theme.LightTheme()})
	myWindow := myApp.NewWindow("Archimedes Pi")
	myWindow.Resize(fyne.NewSize(1900, 1600))

	outputLabel := widget.NewLabel("Press a button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300))

	promptLabel := widget.NewLabel("")
	inputContainer := container.NewVBox()
	inputContainer.Hide()

	// var outputText string
	updateChan := make(chan updateData, 100) // Changed to struct
	var mu sync.Mutex

	callBkPrn2canvas := func(oneLineSansCR string) {
		updateChan <- updateData{text: oneLineSansCR}
	}

	getInputValues := func(prompts []string) chan []string {
		inputContainer.Objects = nil
		promptLabel.SetText(prompts[0] + "\n" + prompts[1])
		values := make([]string, len(prompts))
		entryFields := make([]*widget.Entry, len(prompts))

		entryFields[0] = widget.NewEntry()
		entryFields[0].SetPlaceHolder("e.g., 50,000,000")
		entryFields[0].Resize(fyne.NewSize(220, 40))

		entryFields[1] = widget.NewEntry()
		entryFields[1].SetPlaceHolder("e.g., 256")
		entryFields[1].Resize(fyne.NewSize(150, 40))

		inputChan := make(chan []string)

		submitBtn := widget.NewButton("Submit", func() {
			for i, entry := range entryFields {
				values[i] = entry.Text
				fmt.Println("Input value:", values[i])
			}
			inputContainer.Hide()
			promptLabel.SetText("")
			fmt.Println("Submit button clicked")
			inputChan <- values
			close(inputChan) // Close the channel
		})

		submitBtn.Resize(fyne.NewSize(95, 40))
		submitBtn.Importance = widget.HighImportance

		hbox := container.NewWithoutLayout(
			entryFields[0],
			entryFields[1],
			submitBtn,
		)
		entryFields[0].Move(fyne.NewPos(0, 0))
		entryFields[1].Move(fyne.NewPos(230, 0))
		submitBtn.Move(fyne.NewPos(390, 0))
		hbox.Resize(fyne.NewSize(500, 40))

		inputContainer.Add(container.NewBorder(nil, nil, nil, nil, hbox))
		inputContainer.Resize(fyne.NewSize(500, 60))
		inputContainer.Show()

		return inputChan
	}

	// Buttons
	buttonArchimedes := NewColoredButton("Archimedes", color.RGBA{255, 100, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go ArchimedesBig(callBkPrn2canvas)
	})
	buttonLeibniz := NewColoredButton("Gottfried Wilhelm Leibniz\n runs long", color.RGBA{100, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GottfriedWilhelmLeibniz(callBkPrn2canvas)
	})

	buttonNilakantha := NewColoredButton("Nilakantha", color.RGBA{255, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go func() {
			inputChan := getInputValues([]string{
				"You have selected the Nilakantha Somayaji method...\nPlease fill-in the fields with the number of iterations (suggest 100,000 -> 100,000,000)",
				"And a value for the precision: (suggest 128 -> 512), then hit 'Submit'",
			})

			inputs := <-inputChan // Receive the slice from the channel

			// Error handling for input1
			iters := 100000
			precision := 256
			val1, err1 := strconv.Atoi(inputs[0]) // ::: apparently, we need to pause and wait for user input before doing this !!!!!!
			if err1 != nil {
				fmt.Println("Error converting input1:", err1)
				fmt.Println("setting iters to 40,000,555")
				iters = 40000555
			} else {
				fmt.Println("Value of input1:", val1)
				iters = val1
			}
			// Error handling for input2
			val2, err2 := strconv.Atoi(inputs[1]) // ::: or this !!!!!!!
			if err2 != nil {
				fmt.Println("Error converting input2:", err2)
				fmt.Println("setting precision to 512")
				precision = 512
			} else {
				fmt.Println("Value of input2:", val2)
				precision = val2
			}
			/*
				iters, _ := strconv.Atoi(inputs[0])
				precision, _ := strconv.Atoi(inputs[1])
			*/

			NilakanthaBig(updateChan, iters, precision)
		}()
	})
	buttonGregory := NewColoredButton("Gregory-Leibniz", color.RGBA{100, 100, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GregoryLeibniz(callBkPrn2canvas)
	})
	buttonChudnovsky := NewColoredButton("Chudnovsky", color.RGBA{255, 100, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go ChudnovskyBig(callBkPrn2canvas)
	})
	buttonMonteCarlo := NewColoredButton("Monte Carlo", color.RGBA{100, 255, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go MonteCarloBig(callBkPrn2canvas)
	})
	buttonExtra1 := NewColoredButton("Extra 1", color.RGBA{200, 200, 200, 255}, func() {
		updateChan <- updateData{text: "Extra 1 clicked"}
	})
	buttonExtra2 := NewColoredButton("Extra 2", color.RGBA{150, 150, 150, 255}, func() {
		updateChan <- updateData{text: "Extra 2 clicked"}
	})

	buttonContainer := container.NewGridWithColumns(4,
		buttonArchimedes, buttonLeibniz, buttonGregory, buttonNilakantha,
		buttonChudnovsky, buttonMonteCarlo, buttonExtra1, buttonExtra2,
	)
	content := container.NewVBox(buttonContainer, promptLabel, inputContainer, scrollContainer)
	myWindow.SetContent(content)

	// Main-thread update loop using Fyne's lifecycle
	myWindow.Canvas().SetOnTypedRune(func(r rune) {
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

	myWindow.ShowAndRun()
}
