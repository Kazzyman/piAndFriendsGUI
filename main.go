package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

func main() {
	calculating = false
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	window := myApp.NewWindow("Pi Estimation Demo")
	window.Resize(fyne.NewSize(1900, 1600))

	// UI setup
	outputLabel = widget.NewLabel("\nPress a button to estimate π...\n\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer = container.NewVScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1100))

	
	// ::: Single input dialog, deprecated 
	getSingleInput := func(title, prompt, defaultValue string, callback func(string, bool)) {
		confirmed := false // Track if OK was clicked
		d := dialog.NewEntryDialog(title, prompt, func(value string) { // change to dialog.NewForm() with a widget.Entry inside instead.
			confirmed = true
			callback(value, true)
		}, window)
		d.SetText(defaultValue)
		d.SetOnClosed(func() {
			if !confirmed { // Only trigger cancel if OK wasn’t clicked
				callback("", false)
			}
		})
		d.Show()
	}

	// ::: getSingleInput function, new way
	getSingleInputNew := func(title, prompt, defaultValue string, callback func(string, bool)) {
		// Create an entry widget
		entry := widget.NewEntry()
		entry.SetText(defaultValue)
		entry.SetPlaceHolder(prompt)

		// Create form item
		formItems := []*widget.FormItem{
			widget.NewFormItem("", entry), // Empty label since prompt is in placeholder
		}

		// Track if confirmed
		confirmed := false

		// Create the form dialog
		d := dialog.NewForm(
			title,          // Title
			"OK",           // Confirm button text
			"Cancel",       // Cancel button text
			formItems,      // Form fields
			func(conf bool) {
				confirmed = conf
				if conf {
					callback(entry.Text, true)
				}
			},
			window, // Parent window
		)

		// Handle cancel case
		d.SetOnClosed(func() {
			if !confirmed {
				callback("", false)
			}
		})

		// Show the dialog
		d.Show()
	}


	// Custom colored ::: Buttons
	archimedesBtn := NewColoredButton(
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
	JohnWallisBtn := NewColoredButton("John Wallis RUNS LONG -- does billions of calculations\n-- by Rick Woolley\n three\n four", color.RGBA{25, 200, 100, 215}, 
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons {
				btn.Disable()
			}
			updateOutput("\nRunning John Wallis...\n\n")
			go func() {
				JohnWallis(updateOutput)
				calculating = false
				for _, btn := range buttons {
					btn.Enable()
				}
			}()
		},
	)
	BBPfast46Btn := NewColoredButton("BBP super-fast digits, up to 10,000\nIt only takes like 4s to do 10,000 digits of pi\nsays Rick Woolley", color.RGBA{100, 100, 255, 185}, 
	func() {
		if calculating {
			return
		}
		calculating = true
		for _, btn := range buttons {
			btn.Disable()
		}
		updateOutput("\nRunning BBP-fast-46...\n\n")
		
		getSingleInput("Input Required", "Enter the number of digits for BBP calculation (e.g., 46):", "46", 
			func(digitsStr string, ok bool) {
				if !ok {
					updateOutput("BBP calculation canceled, make another selection")
					for _, btn := range buttons {
						btn.Enable()
					}
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
					bbpFast46(updateOutput, digits)
					calculating = false
					for _, btn := range buttons {
						btn.Enable()
					}
				}()
			})
	})
				// ::: spigot actually goes here in place of Gregory Leibniz
				leibnizBtn := NewColoredButton("Gregory Leibniz\n-- circa 1676\n quick\n pi", color.RGBA{100, 255, 100, 215}, // Greenish for variety
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
	ChudnovskyBtn := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235}, func() {
		getSingleInputNew("Input Required", "Enter the number of digits for the chudnovsky calculation (e.g., 46):", "499555",
			func(digitsStr string, ok bool) {
				if !ok {
					updateOutput("chudnovsky calculation canceled, make another selection")
					for _, btn := range buttons {
						btn.Enable()
					}
					return
				}
				digits = 499888
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err)
					updateOutput("Invalid input, using default 46 digits")
				} else if val <= 0 {
					updateOutput("Input must be positive, using default 46 digits")
				} else if val > 500000 {
					updateOutput("Input must be less than 500,001 -- using default of 46 digits")
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

	buttons = []*ColoredButton{archimedesBtn, JohnWallisBtn, BBPfast46Btn, leibnizBtn, ChudnovskyBtn}

	// ::: Layout
	content := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),

		container.NewGridWithColumns(4, archimedesBtn, JohnWallisBtn, BBPfast46Btn, leibnizBtn,
			ChudnovskyBtn),

		scrollContainer,
	)
	// ::: drop-down menus
	logFilesMenu := fyne.NewMenu("Log Files",
		fyne.NewMenuItem("View Log 1", func() { dialog.ShowInformation("Log Files", "Viewing Log 1", window) }),
		fyne.NewMenuItem("View Log 2", func() { dialog.ShowInformation("Log Files", "Viewing Log 2", window) }),
	)
	windowsMenu := fyne.NewMenu("Collections",
		fyne.NewMenuItem("Fast Pi calculators", func() { window.Show() }),
		fyne.NewMenuItem("Classic Pi calculators", func() { createWindow2(myApp).Show() }),
		fyne.NewMenuItem("Odd Pi calculators", func() { createWindow3(myApp).Show() }),
		fyne.NewMenuItem("Misc Maths", func() { createWindow4(myApp).Show() }),
	)
		informationMenu := fyne.NewMenu("Information",
		fyne.NewMenuItem("Help", func() {
			dialog.ShowInformation("Information", "Help...", window)
		}),
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("Information", "About...", window)
		}),
	)
	mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
	window.SetMainMenu(mainMenu)
	
	window.SetContent(content)
	window.ShowAndRun()
}
