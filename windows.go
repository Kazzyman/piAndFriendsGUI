package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

// Three Additional Windows: 
// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow2(myApp fyne.App) fyne.Window {
	window2 := myApp.NewWindow("Classic Pi calculators")
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel2 := widget.NewLabel("Classic Pi calculators, make a selection")
	outputLabel2.Wrapping = fyne.TextWrapWord
	scrollContainer2 := container.NewScroll(outputLabel)
	scrollContainer2.SetMinSize(fyne.NewSize(1900, 1100))

	// ::: Fixed single input dialog
		getSingleInput := func(title, prompt, defaultValue string, callback func(string, bool)) {
			fmt.Println("Showing single input dialog")
			confirmed := false // Track if OK was clicked
			d := dialog.NewEntryDialog(title, prompt, func(value string) {
				fmt.Println("Single dialog OK with value:", value)
				confirmed = true
				callback(value, true)
			}, window2)
			d.SetText(defaultValue)
			d.SetOnClosed(func() {
				if !confirmed { // Only trigger cancel if OK wasn’t clicked
					fmt.Println("Single dialog canceled")
					callback("", false)
				}
			})
			d.Show()
		}

	
	// Dual input dialog
	getDualInput := func(title, prompt1, prompt2, default1, default2 string, callback func(string, string, bool)) {
		fmt.Println("Showing dual input dialog")
		entry1 := widget.NewEntry()
		entry1.SetText(default1)
		entry2 := widget.NewEntry()
		entry2.SetText(default2)
		submitButton := widget.NewButton("Submit", func() {
			fmt.Println("Dual dialog submitted with values:", entry1.Text, entry2.Text)
			callback(entry1.Text, entry2.Text, true)
			dialog.NewInformation("Submitted", "Values submitted", window2).Hide() // Hack to close dialog
		})
		cancelButton := widget.NewButton("Cancel", func() {
			fmt.Println("Dual dialog canceled")
			callback("", "", false)
			dialog.NewInformation("Canceled", "Dialog canceled", window2).Hide() // Hack to close dialog
		})
		form := container.NewVBox(
			widget.NewLabel(prompt1), entry1,
			widget.NewLabel(prompt2), entry2,
			container.NewHBox(submitButton, cancelButton),
		)
		d := dialog.NewCustom(title, "Close", form, window2)
		d.Resize(fyne.NewSize(400, 300))
		d.Show()
	}

	// Custom colored ::: Buttons
				// Bailey chan -- will go here
				archimedesBtn2 := NewColoredButton(
					"modified Archimedes \n-- by Rick Woolley\n three\n four", color.RGBA{255, 100, 100, 215},
					func() {
						if calculating {
							return
						}
						calculating = true
						for _, btn := range buttons {
							btn.Disable()
						}
						updateOutput("\nRunning ArchimedesBig...\n\n")
						go func() {
							ArchimedesBig(updateOutput)
							calculating = false
							for _, btn := range buttons {
								btn.Enable()
							}
						}()
					},
				)

	// nila 3 goes here
	NilakanthaBtn2 := NewColoredButton("Nilakantha -- takes input", color.RGBA{255, 255, 100, 235}, func() {
		fmt.Println("Archimedes button tapped")
		getDualInput("Input Required", "Number of iterations (suggest 100,000 -> 100,000,000):", "Precision (suggest 128 -> 512):", "100000", "256", func(itersStr, precStr string, ok bool) {
			if !ok {
				updateOutput("Nilakantha calculation canceled")
				return
			}
			iters := 100000
			precision := 256
			val1, err1 := strconv.Atoi(itersStr)
			if err1 != nil {
				fmt.Println("Error converting input1:", err1)
				fmt.Println("setting iters to 40,000,555")
				iters = 40000555
			} else {
				fmt.Println("Value of input1:", val1)
				iters = val1
			}
			val2, err2 := strconv.Atoi(precStr)
			if err2 != nil {
				fmt.Println("Error converting input2:", err2)
				fmt.Println("setting precision to 512")
				precision = 512
			} else {
				fmt.Println("Value of input2:", val2)
				precision = val2
			}
			go NilakanthaBig(updateOutput, iters, precision)
		})
	})
	GregLeibnizBtn2 := NewColoredButton("Gregory Leibniz\n-- circa 1676\n quick\n pi", color.RGBA{100, 255, 100, 215}, // Greenish for variety
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons {
				btn.Disable()
			}
			updateOutput("\nRunning Gregory Leibniz...\n\n")
			go func() {
				GregoryLeibniz(updateOutput)
				calculating = false
				for _, btn := range buttons {
					btn.Enable()
				}
			}()
		},
	)
	GottfriedWilhelmLeibnizBtn2 := NewColoredButton("Gottfried Wilhelm Leibniz -- runs long", color.RGBA{100, 255, 100, 225}, 
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons {
				btn.Disable()
			}
			updateOutput("\nRunning Gregory Leibniz...\n\n")
			go func() {
				GottfriedWilhelmLeibniz(updateOutput)
				calculating = false
				for _, btn := range buttons {
					btn.Enable()
				}
			}()
		},
	)
	ChudnovskyBtn2 := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235}, func() {
		getSingleInput("Input Required", "Enter the number of digits for the chudnovsky calculation (e.g., 46):", "46",
			func(digitsStr string, ok bool) {
				if !ok {
					fmt.Println("Houston, we have a problem")
					updateOutput("chudnovsky calculation canceled")
					return
				}
				digits = 46
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err)
					updateOutput("Invalid input, using default 46 digits")
				} else if val <= 0 {
					updateOutput("Input must be positive, using default 46 digits")
				} else if val > 10000 {
					updateOutput("Input must be less than 10,001, using default 46 digits")
				} else {
					digits = val
				}
				go func() {
					chudnovskyBig(updateOutput, digits)
					calculating = false
					for _, btn := range buttons {
						btn.Enable()
					}
				}()
			})
	})
	buttons = []*ColoredButton{archimedesBtn2, NilakanthaBtn2, GregLeibnizBtn2, GottfriedWilhelmLeibnizBtn2}

	// ::: Layout
	content := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),

		container.NewGridWithColumns(4, archimedesBtn2, NilakanthaBtn2, GregLeibnizBtn2, GottfriedWilhelmLeibnizBtn2,
			ChudnovskyBtn2),

		scrollContainer2, // ::: was and probably should be scrollContainer1 ?????
	)
	window2.SetContent(content)
	// Main-thread update loop using Fyne's lifecycle
	window2.Canvas().SetOnTypedRune(func(r rune) {
		// Dummy handler to keep canvas active
	})

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
			window2.Show() // ::: ???
		}),
		fyne.NewMenuItem("Classic Pi calculators", func() {
			createWindow2(myApp).Show() // ::: ???
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
