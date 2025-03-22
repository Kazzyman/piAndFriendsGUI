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
var outputLabel1 = widget.NewLabel("\nPress a button to estimate π...\n\n")
var scrollContainer1 = container.NewVScroll(outputLabel1)

func main() {
	calculating = false
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	window1 := myApp.NewWindow("Pi Estimation Demo")
	window1.Resize(fyne.NewSize(1900, 1600))
	outputLabel1.Wrapping = fyne.TextWrapWord
	scrollContainer1.SetMinSize(fyne.NewSize(1900, 1050))
	
	getSingleInputBpp1 := func(title, prompt, defaultValue string, callback func(string, bool)) {
		if calculating {
			return
		}
		confirmed := false // Track if OK was clicked
		d := dialog.NewEntryDialog(title, prompt, func(value string) {
			confirmed = true
			callback(value, true)
		}, window1)
		d.SetText(defaultValue)
		d.SetOnClosed(func() {
			if !confirmed { // Only trigger cancel if OK wasn’t clicked
				callback("", false)
			}
		})
		d.Show()
	}

	getSingleInputChud1 := func(title, prompt, defaultValue string, callback func(string, bool)) {
		if calculating {return}
		confirmed := false // Track if OK was clicked
		d := dialog.NewEntryDialog(title, prompt, func(value string) {
			confirmed = true
			callback(value, true)
		}, window1)
		d.SetText(defaultValue)
		d.SetOnClosed(func() {
			if !confirmed { // Only trigger cancel if OK wasn’t clicked
				callback("", false)
			}
		})
		d.Show()
	}


	// Custom colored ::: Buttons1
	archimedesBtn1 := NewColoredButton(
		"modified Archimedes \n-- by Rick Woolley\n three\n four", color.RGBA{255, 100, 100, 215},
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			updateOutput1("\nRunning ArchimedesBig...\n\n")
			go func() {
				ArchimedesBig(updateOutput1) 
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	JohnWallisBtn1 := NewColoredButton("John Wallis RUNS LONG -- does billions of calculations\n-- by Rick Woolley\n three\n four", color.RGBA{110, 110, 255, 185}, 
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			updateOutput1("\nRunning John Wallis...\n\n")
			go func() {
				JohnWallis(updateOutput1)
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	BBPfast44Btn1 := NewColoredButton("BBP super-fast digits, up to 10,000\nIt only takes like 4s to do 10,000 digits of pi\nsays Rick Woolley", color.RGBA{25, 200, 100, 215}, 
	func() {
		if calculating {
			return
		}
		updateOutput1("\nRunning BBP-fast-190 up to here...\n\n") 
		
		getSingleInputBpp1("Input Required", "Enter the number of digits for BBP calculation (e.g., 190):", "190", 
			func(digitsStr string, ok bool) {
				// calculating = true
				for _, btn := range buttons1 {
					btn.Disable()
				}
				for _, btn := range BPPbut {
					calculating = true // keep it from being restarted in parallel 
					btn.Enable() // even though the button is enabled 
				}
				if !ok {
					updateOutput1("BBP calculation canceled, make another selection")
					for _, btn := range buttons1 {
						btn.Enable()
					}
					calculating = false // ::: this is the trick to allow others to run after the dialog is canceled. 
					return
				}
				digits = 190
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err)
					updateOutput1("Invalid input, using default 190 digits")
				} else if val <= 0 {
					fmt.Println("here in val <= 0")
					updateOutput1("Input must be positive, using default 190 digits")
				} else if val > 10000 {
					fmt.Println("here in val > 10000")
					updateOutput1("Input must be less than 10,001, using default 190 digits")
				} else {
					digits = val 
				}
				go func() {
					bbpFast44(updateOutput1, digits)
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
				}()
			})
	})
	SpigotBtn1 := NewColoredButton("Spigot (magic)\nInstantly spits out unlimited digits of pi\nsays Rick Woolley", color.RGBA{110, 110, 255, 185},
		func() {
			if calculating {
				return
			}
			updateOutput1("\nRunning The Spigot...\n\n")

			getSingleInputBpp1("Input Required", "Enter the number of digits for Spigot calculation (e.g., 160):", "160",
				func(digitsStr string, ok bool) {
					// calculating = true
					for _, btn := range buttons1 {
						btn.Disable()
					}
					for _, btn := range BPPbut {
						calculating = true // keep it from being restarted in parallel 
						btn.Enable() // even though the button is enabled 
					}
					if !ok {
						updateOutput1("Spigot calculation canceled, make another selection")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled. 
						return
					}
					digits = 160
					val, err := strconv.Atoi(digitsStr)
					if err != nil {
						fmt.Println("Error converting input:", err)
						updateOutput1("Invalid input, using default 160 digits")
					} else if val <= 0 {
						fmt.Println("here in val <= 0")
						updateOutput1("Input must be positive, using default 160 digits")
					} else if val > 10000 {
						fmt.Println("here in val > 10000")
						updateOutput1("Input must be less than 10,001, using default 160 digits")
					} else {
						digits = val
					}
					go func() {
						TheSpigot(updateOutput1, digits)
						calculating = false
						for _, btn := range buttons1 {
							btn.Enable()
						}
					}()
				})
		})
	ChudnovskyBtn1 := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235}, func() {
		getSingleInputChud1("Input Required", "Enter the number of digits for the chudnovsky calculation (e.g., 499888):", "499888",
			func(digitsStr string, ok bool) {
				if calculating {
					return
				}
				// calculating = true
				for _, btn := range buttons1 {
					btn.Disable()
				}
				for _, btn := range chudBut {
					calculating = true // keep it from being restarted in parallel 
					btn.Enable() // even though the button is enabled 
				}
				if !ok {
					updateOutput1("chudnovsky calculation canceled, make another selection") 
					for _, btn := range buttons1 {
						btn.Enable()
					}
					calculating = false // ::: this is the trick to allow others to run after the dialog is canceled. 
					return
				}
				digits = 499888
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err)
					updateOutput1("Invalid input, using default 499888 digits")
				} else if val <= 0 {
					updateOutput1("Input must be positive, using default 499888 digits")
				} else if val > 500000 {
					updateOutput1("Input must be less than 500,001 -- using default of 499888 digits")
				} else {
					digits = val
				}
				go func() {
					chudnovskyBig(updateOutput1, digits)
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
				}()
			})
	})

	chudBut = []*ColoredButton{ChudnovskyBtn1} // used as bug fixes 
	BPPbut = []*ColoredButton{BBPfast44Btn1}
	BPPbut = []*ColoredButton{SpigotBtn1}
	
	buttons1 = []*ColoredButton{archimedesBtn1, JohnWallisBtn1, BBPfast44Btn1, SpigotBtn1, ChudnovskyBtn1} // array used only for range btn.Enable() // will have 7-8

	// ::: Layout
	content1 := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),

		container.NewGridWithColumns(4, archimedesBtn1, JohnWallisBtn1, BBPfast44Btn1, SpigotBtn1,
			ChudnovskyBtn1),

		scrollContainer1,
	)
	// ::: drop-down menus
	logFilesMenu := fyne.NewMenu("Log Files",
		fyne.NewMenuItem("View Log 1", func() { dialog.ShowInformation("Log Files", "Viewing Log 1", window1) }),
		fyne.NewMenuItem("View Log 2", func() { dialog.ShowInformation("Log Files", "Viewing Log 2", window1) }),
	)
	windowsMenu := fyne.NewMenu("Collections",
		fyne.NewMenuItem("Fast Pi calculators", func() { window1.Show() }),
		fyne.NewMenuItem("Classic Pi calculators", func() { createWindow2(myApp).Show() }),
		fyne.NewMenuItem("Odd Pi calculators", func() { createWindow3(myApp).Show() }),
		fyne.NewMenuItem("Misc Maths", func() { createWindow4(myApp).Show() }),
	)
		informationMenu := fyne.NewMenu("Information",
		fyne.NewMenuItem("Help", func() {
			dialog.ShowInformation("Information", "Help...", window1)
		}),
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("Information", "About...", window1)
		}),
	)
	mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
	window1.SetMainMenu(mainMenu)
	
	window1.SetContent(content1) 
	window1.ShowAndRun()
}
