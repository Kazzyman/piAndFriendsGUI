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

var outputLabel2 = widget.NewLabel("Classic Pi calculators, make a selection")
var scrollContainer2 = container.NewScroll(outputLabel2)

// Three Additional Windows: 
// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow2(myApp fyne.App) fyne.Window {
	window2 := myApp.NewWindow("Classic Pi calculators")
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel2.Wrapping = fyne.TextWrapWord
	scrollContainer2.SetMinSize(fyne.NewSize(1900, 1050))

		getSingleInput2 := func(title, prompt, defaultValue string, callback func(string, bool)) {
			confirmed := false // Track if OK was clicked
			d := dialog.NewEntryDialog(title, prompt, func(value string) {
				confirmed = true
				callback(value, true)
			}, window2)
			d.SetText(defaultValue)
			d.SetOnClosed(func() {
				if !confirmed { // Only trigger cancel if OK wasn’t clicked
					callback("", false)
				}
			})
			d.Show()
		}

	
	// Dual input dialog
	getDualInput2 := func(title, prompt1, prompt2, default1, default2 string, callback func(string, string, bool)) {
		calculating = true
		for _, btn := range buttons2 {
			btn.Disable()
		}
		entry1 := widget.NewEntry()
		entry1.SetText(default1)
		entry2 := widget.NewEntry()
		entry2.SetText(default2)
		submitButton := widget.NewButton("Run with those values",
			func() {
				callback(entry1.Text, entry2.Text, true)
				dialog.NewInformation("Submitted", "Values submitted", window2).Hide() // Hack to close dialog
				calculating = true
				for _, btn := range buttons2 {
					btn.Disable()
				}
			})
		form := container.NewVBox(
			widget.NewLabel(prompt1), entry1,
			widget.NewLabel(prompt2), entry2,
			// container.NewHBox(submitButton, cancelButton), // ?? still get Close button ? need no cancelButton ??
			container.NewHBox(submitButton),
		)
		d := dialog.NewCustom(title, "Dismiss dialogBox", form, window2)

		d.Resize(fyne.NewSize(400, 300))
		d.Show()
	}

	// ::: Buttons2
	done := make(chan bool) // kill channel for all goroutines 
	
	// Bailey chan -- will go here
				archimedesBtn2 := NewColoredButton(
					"modified Archimedes \n-- by Rick Woolley\n three\n four", color.RGBA{255, 100, 100, 215},
					func() {
						if calculating {
							return
						}
						calculating = true
						for _, btn := range buttons2 {
							btn.Disable()
						}
						updateOutput2("\nRunning ArchimedesBig...\n\n")
						go func() {
							ArchimedesBig(updateOutput2, done)
							calculating = false
							for _, btn := range buttons2 {
								btn.Enable()
							}
						}()
					},
				)

	// ::: nila 3 goes here ??
	NilakanthaBtn2 := NewColoredButton("Nilakantha -- input iterations\noutput up to 26 digits of pi", color.RGBA{255, 255, 100, 235}, func() {
		if calculating {
			return
		}
		calculating = true
		for _, btn := range buttons2 {
			btn.Disable()
		}
		getDualInput2("Input Required", "Number of iterations (suggest 300,000 -> 30,000,000  -> 300,000,000):", "Precision (suggest 128):", 
			"30000000", "128", // 30,000,000
			func(itersStr, precStr string, ok bool) {
				calculating = true
				for _, btn := range buttons2 {
					btn.Disable()
				}
				if !ok {
					updateOutput2("Nilakantha calculation canceled")
					return
				}
				iters := 30000000 // 30,000,000
				precision := 128
				itersStr = removeCommasAndPeriods(itersStr) // ::: allow user to enter a number with a comma
				val1, err1 := strconv.Atoi(itersStr)
				if err1 != nil {
					fmt.Println("Error converting iterations val1:", err1) // handle error
					iters = 30000000
				} else {
					iters = val1
				}
				val2, err2 := strconv.Atoi(precStr)
				if err2 != nil {
					fmt.Println("Error converting precision val2:", err2) // handle error 
					updateOutput2("setting precision to 128") 
					// fyneFunc(fmt.Sprintf("setting precision to 512")) //  ::: cannot do this instead because ??
					precision = 128
				} else {
					precision = val2
				}
				go NilakanthaBig(updateOutput2, iters, precision)
				calculating = false
				for _, btn := range buttons2 {
					btn.Enable()
				}
			})
	})
	GregLeibnizBtn2 := NewColoredButton("Gregory Leibniz\n-- circa 1676\n quick - 4s, 9 digits of pi\n pi", color.RGBA{100, 255, 100, 215}, // Greenish for variety
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons2 {
				btn.Disable()
			}
			updateOutput2("\nRunning Gregory Leibniz...\n\n")
			go func() {
				GregoryLeibniz(updateOutput2)
				calculating = false
				for _, btn := range buttons2 {
					btn.Enable()
				}
			}()
		},
	)
	GottfriedWilhelmLeibnizBtn2 := NewColoredButton("Gottfried Wilhelm Leibniz -- runs 20sec\ngives 10 digits of pi", color.RGBA{100, 255, 100, 225}, 
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons2 {
				btn.Disable()
			}
			updateOutput2("\nRunning Gregory Leibniz...\n\n")
			go func() {
				GottfriedWilhelmLeibniz(updateOutput2)
				calculating = false
				for _, btn := range buttons2 {
					btn.Enable()
				}
			}()
		},
	)
	// ::: temp, Bailey concur goes here
	ChudnovskyBtn2 := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235}, func() {
		getSingleInput2("Input Required", "Enter the number of digits for the chudnovsky calculation (e.g., 46):", "46",
			func(digitsStr string, ok bool) {
				var chudDigits int 
				if !ok {
					updateOutput2("chudnovsky calculation canceled")
					return
				}
				chudDigits = 46
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err) // handel error 
					updateOutput2("Invalid input, using default 46 digits")
				} else if val <= 0 {
					updateOutput2("Input must be positive, using default 46 digits")
				} else if val > 10000 {
					updateOutput2("Input must be less than 10,001, using default 46 digits")
				} else {
					chudDigits = val
				}
				go func() {
					chudnovskyBig(updateOutput2, chudDigits)
					calculating = false
					for _, btn := range buttons2 {
						btn.Enable()
					}
				}()
			})
	})
	buttons2 = []*ColoredButton{archimedesBtn2, NilakanthaBtn2, GregLeibnizBtn2, GottfriedWilhelmLeibnizBtn2, ChudnovskyBtn2} // array used only for range btn.Enable()

	// ::: Layout
	content2 := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),

		container.NewGridWithColumns(4, archimedesBtn2, NilakanthaBtn2, GregLeibnizBtn2, GottfriedWilhelmLeibnizBtn2,
			ChudnovskyBtn2),

		scrollContainer2, // ::: was and probably should be scrollContainer1 ?????
	)
	window2.SetContent(content2)
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
	
	window2.SetContent(content2)
	return window2
} // end of createWindow2


// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow3(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window3 := myApp.NewWindow("Odd Pi calculators")
	window3.Resize(fyne.NewSize(1900, 1600))
	outputLabel3 := widget.NewLabel("Odd Pi calculators, make a selection")
	outputLabel3.Wrapping = fyne.TextWrapWord
	scrollContainer3 := container.NewScroll(outputLabel3)
	scrollContainer3.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer3 := container.NewGridWithColumns(4,
		widget.NewButton("Button 9", func() {}),
		widget.NewButton("Button 10", func() {}),
		widget.NewButton("Button 11", func() {}),
		widget.NewButton("Button 12", func() {}),
		widget.NewButton("Button 13", func() {}),
		widget.NewButton("Button 14", func() {}),
		widget.NewButton("Button 15", func() {}),
		widget.NewButton("Button 16", func() {}),
	)
	content3 := container.NewVBox(buttonContainer3, scrollContainer3)
	window3.SetContent(content3)
	return window3
}

// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow4(myApp fyne.App) fyne.Window {
	// Similar structure to createWindow2
	window4 := myApp.NewWindow("Misc Maths")
	window4.Resize(fyne.NewSize(1900, 1600))
	outputLabel4 := widget.NewLabel("Misc Maths, make a selection")
	outputLabel4.Wrapping = fyne.TextWrapWord
	scrollContainer4 := container.NewScroll(outputLabel4)
	scrollContainer4.SetMinSize(fyne.NewSize(1900, 1300))
	buttonContainer4 := container.NewGridWithColumns(4,
		widget.NewButton("Button 17", func() {}), widget.NewButton("Button 18", func() {}), widget.NewButton("Button 19", func() {}), widget.NewButton("Button 20", func() {}),
		widget.NewButton("Button 21", func() {}), widget.NewButton("Button 22", func() {}), widget.NewButton("Button 23", func() {}), widget.NewButton("Button 24", func() {}),
	)
	content4 := container.NewVBox(buttonContainer4, scrollContainer4)
	window4.SetContent(content4)
	return window4
}
