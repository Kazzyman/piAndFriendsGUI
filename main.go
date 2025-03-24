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
var myApp = app.New()
var window1 = myApp.NewWindow("Pi Estimation Demo")

func main() {
	calculating = false
	myApp.Settings().SetTheme(theme.LightTheme())
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

	// identical to the above except in name, i.e., Bpp1 vs Chud1; the latter being "deprecated" by Rick
	/* Chud now using: showCustomEntryDialog(
	getSingleInputChud1 := func(title, prompt, defaultValue string, callback func(string, bool)) {
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
	 */


	done := make(chan bool) // kill channel for all goroutines 
	done2 := make(chan bool) // kill channel for all goroutines 
	done3 := make(chan bool) // kill channel for all goroutines // ::: remove some of these ??
	
	// Custom colored ::: Buttons1
	archimedesBtn1 := NewColoredButton(
		"modified Archimedes \n-- by Rick Woolley, and Rick's personal favorite\n 3012 digits of pi in under a minute\n four", color.RGBA{255, 100, 100, 215},
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
				ArchimedesBig(updateOutput1, done) 
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	JohnWallisBtn1 := NewColoredButton("John Wallis 5m30s -- does 40 billion calculations\n-- by Rick Woolley\n just 10 digits of pi\n four", color.RGBA{110, 110, 255, 185}, 
			func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			updateOutput1("\nRunning John Wallis...\n\n")
			go func() { // made this the goroutine as per your example 
				JohnWallis(updateOutput1, done) // made this a normal func call, per your example
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	BBPfast44Btn1 := NewColoredButton("BBP super-fast digits, up to 10,000\nIt only takes like 4s to do 10,000 digits of pi\nsays Rick Woolley", color.RGBA{25, 200, 100, 215}, 
	func() {
		var BppDigits int 
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
				BppDigits = 190
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
					BppDigits = val 
				}
				go func() {
					bbpFast44(updateOutput1, BppDigits) 
					calculating = false
					for _, btn := range buttons1 {
						btn.Enable()
					}
				}()
			})
	})
	SpigotBtn1 := NewColoredButton("Spigot (magic)\nInstantly spits out unlimited digits of pi\nsays Rick Woolley", color.RGBA{110, 110, 255, 185},
		func() {
			var spigotDigits int 
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
					for _, btn := range spigotBut {
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
					spigotDigits = 160
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
						spigotDigits = val
					}
					go func() {
						TheSpigot(updateOutput1, spigotDigits) // ::: func
						calculating = false
						for _, btn := range buttons1 {
							btn.Enable()
						}
					}()
				})
		})
	ChudnovskyBtn1 := NewColoredButton("chudnovsky -- 23,000 digits of pi\nin less than 8s", color.RGBA{255, 255, 100, 235}, 
	func() {
		var chudDigits int 
		if calculating {
			return
		}
		for _, btn := range buttons1 {
			btn.Disable()
		}
		for _, btn := range chudBut { // chudBut is an array with only one member
			calculating = true // keep it from being restarted in parallel
			btn.Enable() // even though the button is enabled
		}

		showCustomEntryDialog(
			"Input Desired number of digits",
			"Any number less than 49,999",
			func(input string) {
				if input != "" {
					input = removeCommasAndPeriods(input) // ::: allow user to enter a number with a comma
					val, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Error converting input:", err)
						updateOutput1("Invalid input, using default 49,000 digits")
					} else if val <= 0 {
						updateOutput1("Input must be positive, using default 49000 digits")
					} else if val > 50000 {
						updateOutput1("Input must be less than 50,000 -- using default of 49,000 digits")
					} else {
						chudDigits = val
					}
					go func() {
						chudnovskyBig(updateOutput1, chudDigits) // ::: func
						calculating = false
						for _, btn := range buttons1 {
							btn.Enable()
						}
					}()
				} else {
					// dialog canceled 
						updateOutput1("chudnovsky calculation canceled, make another selection")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled.
						return
				}
			},
		)
		
		
		/*
				getSingleInputChud1("Input Required", "Number of digits desired from the chudnovsky calculation (max 49,000):", "49000",
			func(digitsStr string, ok bool) {
				
				digits = 49000
				digitsStr = removeCommasAndPeriods(digitsStr) // allow user to enter a number with a comma
				val, err := strconv.Atoi(digitsStr)
				if err != nil {
					fmt.Println("Error converting input:", err)
					updateOutput1("Invalid input, using default 49,000 digits")
				} else if val <= 0 {
					updateOutput1("Input must be positive, using default 49000 digits")
				} else if val > 50000 {
					updateOutput1("Input must be less than 50,000 -- using default of 49,000 digits")
				} else {
					digits = val
				}
		 */
		
		

				
		
	})
	/*
	.
	.
	 */
	MontyBtn1 := NewColoredButton("Montycarlo using big floats, and float64-- 4 digits of pi\nin 21s\nRick's second favorite", color.RGBA{255, 255, 100, 235}, 
	func() {
		var MontDigits int 
		if calculating {
			return
		}
		for _, btn := range buttons1 {
			btn.Disable()
		}
		for _, btn := range montBut { // chudBut is an array with only one member
			calculating = true // keep it from being restarted in parallel
			btn.Enable() // even though the button is enabled
		}

		showCustomEntryDialog(
			"Input Desired number of grid elements",
			"max 5k; 10,000 will produce 4 pi digits",
			func(input string) {
				if input != "" {
					input = removeCommasAndPeriods(input) // ::: allow user to enter a number with a comma
					val, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Error converting input:", err)
						updateOutput1("Invalid input, using default 10,000 digits")
					} else if val <= 0 {
						updateOutput1("Input must be positive, using default 10,000 digits")
					} else if val > 50000 {
						updateOutput1("Input must be less than 50,000 -- using default of 10,000 digits")
					} else {
						MontDigits = val
					}
					go func() {
						Monty(updateOutput1, MontDigits) // ::: func 
						calculating = false
						for _, btn := range buttons1 {
							btn.Enable()
						}
					}()
				} else {
					// dialog canceled 
					updateOutput1("MontyCarlo calculation canceled, make another selection")
					for _, btn := range buttons1 {
						btn.Enable()
					}
					calculating = false // ::: this is the trick to allow others to run after the dialog is canceled.
					return
				}
			},
		)
	})

	chudBut = []*ColoredButton{ChudnovskyBtn1} // used as bug fixes 
	BPPbut = []*ColoredButton{BBPfast44Btn1}
	spigotBut = []*ColoredButton{SpigotBtn1}
	montBut = []*ColoredButton{MontyBtn1}
	
	buttons1 = []*ColoredButton{archimedesBtn1, JohnWallisBtn1, BBPfast44Btn1, SpigotBtn1, ChudnovskyBtn1, MontyBtn1} // array used only for range btn.Enable() // will have 7-8

	// ::: Layout
	content1 := container.NewVBox(widget.NewLabel("\nSelect a method to estimate π:\n"),
		container.NewGridWithColumns(4, archimedesBtn1, JohnWallisBtn1, BBPfast44Btn1, SpigotBtn1,
			ChudnovskyBtn1, MontyBtn1), 
		scrollContainer1,
	)
	// ::: drop-down menus
	logFilesMenu := fyne.NewMenu("Log Files",
		fyne.NewMenuItem("View Log 1", func() { dialog.ShowInformation("Log Files", "Viewing Log 1", window1) }),
		fyne.NewMenuItem("View Log 2", func() { dialog.ShowInformation("Log Files", "Viewing Log 2", window1) }),
	)
	windowsMenu := fyne.NewMenu("Collections/functions",
		fyne.NewMenuItem("Fast Pi calculators", func() { window1.Show() }),
		fyne.NewMenuItem("Classic Pi calculators", func() { createWindow2(myApp).Show() }),
		fyne.NewMenuItem("Odd Pi calculators", func() { createWindow3(myApp).Show() }),
		fyne.NewMenuItem("Misc Maths", func() { createWindow4(myApp).Show() }),
	)
	informationMenu := fyne.NewMenu("Actions and Information",
		fyne.NewMenuItem("Help", func() {
			dialog.ShowInformation("Information", "Help...", window1)
		}),
		fyne.NewMenuItem("Abort current method", func() {
			select {
			case <-done: // Check if already closed
				updateOutput1("Goroutine1 already terminated\n")
			case <-done2: // 
				updateOutput1("Goroutine2 already terminated\n")
			case <-done3:
				updateOutput1("Goroutine3 already terminated\n")
			default:
				close(done) // Signal termination
				close(done2)
				close(done3)
				updateOutput1("Termination signals sent to all current processes\n")
			}			
			// dialog.ShowInformation("Information", "About...", window1)
		}),
	)
	mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
	window1.SetMainMenu(mainMenu)
	
	window1.SetContent(content1) 
	window1.ShowAndRun()
}
