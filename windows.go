package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
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
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1100))

	/*
	   // Define these locally and include them in the content
	   promptLabel := widget.NewLabel("")
	   inputContainer := container.NewVBox()
	   inputContainer.Hide() // Initially hidden, like in main()
	 */
	
	promptLabel := widget.NewLabel("")
	inputContainer := container.NewVBox()
	inputContainer.Hide() // for Nilakantha's two input fields
	
	// var outputText string
	updateChan := make(chan updateData, 100) // Changed to struct
	
	var mu sync.Mutex // would go with: go func() { ticker  , below
	
		callBkPrn2canvas := func(oneLineSansCR string) {
			updateChan <- updateData{text: oneLineSansCR}
		}

/*
	myApp := app.New()
	myApp.Settings().SetTheme(&highContrastTheme{Theme: theme.LightTheme()}) // ??? points to a custom theme, defined above as a method called color [non-exported due to lower case name] ...
		fmt.Printf("color.Color is: %s\n", "what could I put here to see what is/was returned by the above color method?")
	myWindow := myApp.NewWindow("Fast Pi calculators")					// '&' is the address-of operator. It is used to get the memory address of a variable (see examples below).
	myWindow.Resize(fyne.NewSize(1900, 1600))
	outputLabel := widget.NewLabel("Press a button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1100)) // was 1300

	promptLabel := widget.NewLabel("")
	inputContainer := container.NewVBox()
	inputContainer.Hide() // for Nilakantha's two input fields
*/
	// ::: --- get input values (two strings) ========== = = = = = = = = = = = = = = = = = = = = = = = = =  
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
			inputContainer.Hide() // ::: no effect ?
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
		entryFields[0].Move(fyne.NewPos(0, 0)) // ::: what are these doing ????
		entryFields[1].Move(fyne.NewPos(230, 0))
		submitBtn.Move(fyne.NewPos(390, 0))
		hbox.Resize(fyne.NewSize(500, 40))

		inputContainer.Add(container.NewBorder(nil, nil, nil, nil, hbox))
		inputContainer.Resize(fyne.NewSize(500, 60))
		inputContainer.Show()

		return inputChan
	}


	ArchimedesButton2pg1pos := NewColoredButton("Bailey chan goes here\nnot\nArchimedesBigCopy", color.RGBA{255, 100, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go ArchimedesBig(callBkPrn2canvas)
	})

	buttonGregory := NewColoredButton("Gregory-Leibniz, is quick", color.RGBA{100, 100, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GregoryLeibniz(callBkPrn2canvas)
	})
	
	// nila 3 goes here
	buttonNilakantha := NewColoredButton("Nilakantha -- takes input", color.RGBA{255, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go func() { // this anonymous func concludes with a normal function call
			inputChan := getInputValues([]string{
				"You have selected the Nilakantha Somayaji method...\nPlease fill-in the fields with the number of iterations (suggest 100,000 -> 100,000,000)",
				"And a value for the precision: (suggest 128 -> 512), then hit 'Submit'",
			})

			inputs := <-inputChan // Receive the slice from the channel

			// Error handling for input1
			iters := 100000
			precision := 256
			val1, err1 := strconv.Atoi(inputs[0])
			if err1 != nil {
				fmt.Println("Error converting input1:", err1)
				fmt.Println("setting iters to 40,000,555")
				iters = 40000555
			} else {
				fmt.Println("Value of input1:", val1)
				iters = val1
			}
			// Error handling for input2
			val2, err2 := strconv.Atoi(inputs[1])
			if err2 != nil {
				fmt.Println("Error converting input2:", err2)
				fmt.Println("setting precision to 512")
				precision = 512
			} else {
				fmt.Println("Value of input2:", val2)
				precision = val2
			}

			NilakanthaBig(updateChan, iters, precision)
		}()
	})

	GottfriedWilhelmLeibniz2pg2pos := NewColoredButton("Gottfried Wilhelm Leibniz -- runs long", color.RGBA{100, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GottfriedWilhelmLeibniz(callBkPrn2canvas)
	})
	
	buttonContainer := container.NewGridWithColumns(4,
		ArchimedesButton2pg1pos, buttonNilakantha, buttonGregory, GottfriedWilhelmLeibniz2pg2pos,
	)
	
	/*

	   // Include promptLabel and inputContainer in the layout
	   content := container.NewVBox(buttonContainer, promptLabel, inputContainer, scrollContainer)
	   window2.SetContent(content)
	 */
	
	content := container.NewVBox(buttonContainer, promptLabel, inputContainer, scrollContainer)
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


	// Drop-Down Menus
	logFilesMenu := fyne.NewMenu("Log Files",
		fyne.NewMenuItem("View Log 1", func() {
			// Implement log file viewing here
			dialog.ShowInformation("Log Files", "Viewing Log 1", window2)
		}),
		fyne.NewMenuItem("View Log 2", func() {
			// Implement log file viewing here
			dialog.ShowInformation("Log Files", "Viewing Log 2", window2)
		}),
	)

	windowsMenu := fyne.NewMenu("Collections",
		fyne.NewMenuItem("Fast Pi calculators", func() {
			window2.Show()
		}),
		fyne.NewMenuItem("Classic Pi calculators", func() {
			createWindow2(myApp).Show()
		}),
		fyne.NewMenuItem("Odd Pi calculators", func() {
			createWindow3(myApp).Show()
		}),
		fyne.NewMenuItem("Misc Maths", func() {
			createWindow4(myApp).Show()
		}),
	)

	informationMenu := fyne.NewMenu("Information",
		fyne.NewMenuItem("Help", func() {
			dialog.ShowInformation("Information", "Help...", window2)
		}),
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("Information", "About...", window2)
		}),
	)

	mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
	window2.SetMainMenu(mainMenu)

	return window2
} // end of createWindow2


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
		widget.NewButton("Button 9", func() {}), 
		widget.NewButton("Button 10", func() {}), 
		widget.NewButton("Button 11", func() {}), 
		widget.NewButton("Button 12", func() {}),
		widget.NewButton("Button 13", func() {}), 
		widget.NewButton("Button 14", func() {}), 
		widget.NewButton("Button 15", func() {}), 
		widget.NewButton("Button 16", func() {}),
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
