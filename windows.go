package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

// @formatter:off

var bgsc2 = canvas.NewRectangle(color.NRGBA{R: 130, G: 160, B: 250, A: 140}) // Light blue // was: 130, 160, 250, 160
var bgwc2 = canvas.NewRectangle(color.NRGBA{R: 110, G: 255, B: 160, A: 150}) // Light green

var outputLabel2 = widget.NewLabel("Classic Pi calculators, make a selection")
var scrollContainer2 = container.NewScroll(outputLabel2)
var window2 = myApp.NewWindow("Rick's Pi calculation Demo, set #2")

// Three Additional Windows: 
// ::: ------------------------------------------------------------------------------------------------------------------------------------------------------------
func createWindow2(myApp fyne.App) fyne.Window {
	window2.Resize(fyne.NewSize(1900, 1600))
	outputLabel2.Wrapping = fyne.TextWrapWord
	scrollContainer2.SetMinSize(fyne.NewSize(1900, 1050))

	coloredScroll2 := container.NewMax(bgsc2, scrollContainer2) // Light blue-ish scroll bg

	// ::: Get single input dialog < - - - - - - - - - - - - - - - - - - - - - - - - < -
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
	// ::: Dual input dialog < - - - - - - - - - - - - - - - - - - - - - - - - < -
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
	done := make(chan bool) // local, kill channel for all goroutines that are listening: 
	
	// ::: Bailey chan -- will go here
	BBPfast44Btn2 := NewColoredButton(
		"BBP, the Bailey–Borwein–Plouffe formula for π, circa 1995\n" +
			"FAST -- only runs 4s to produce 10,000 digits of Pi" +
			"uses channels: GOMAXPROCS(numCPU), and using Go's big floats\n" +
			"                     --- done here by Rick Woolley ---          ",
		color.RGBA{25, 200, 100, 215},
		func() {
			var BppDigits int
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons2 {
				btn.Disable()
			}
			for _, btn := range BPPbut2 { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			updateOutput2("\nRunning BBP-fast-190 up to here...\n\n")

			showCustomEntryDialog2(
				"Input Desired number of digits",
				"Any number less than 190",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
						val, err := strconv.Atoi(input)
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput2("Invalid input, using default 190 digits")
						} else if val <= 0 {
							updateOutput2("Input must be positive, using default 190 digits")
						} else if val > 10000 {
							updateOutput2("Input must be less than 191 -- using default of 190 digits")
						} else {
							BppDigits = val
						}
						go func() {
							bbpFast44(updateOutput2, BppDigits, done) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons2 {
								btn.Enable()
							}
						}()
					} else {
						// dialog canceled 
						updateOutput2("spigot calculation canceled, make another selection")
						for _, btn := range buttons2 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					}
				},
			)
		})
	/*
	.
	.
	 */
	
	// ::: nila 3 goes here ??
	NilakanthaBtn2 := NewColoredButton(
		"Nilakantha -- input iterations\n" +
		"output up to 26 digits of pi",
		color.RGBA{255, 255, 100, 235},
	func() {
		if calculating {
			return
		}
		calculating = true
		for _, btn := range buttons2 {
			btn.Disable()
		}
		for _, btn := range nilaBut2 { // Refer to the comments in the initial assignment and creation of archimedesBtn1
			calculating = true
			btn.Enable()
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
				go NilakanthaBig(updateOutput2, iters, precision, done) // ::: probably want to add a done channel to this one
				calculating = false
				for _, btn := range buttons2 {
					btn.Enable()
				}
			})
	})

	// ::: temp, Bailey concur goes here
	ChudnovskyBtn2 := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 235},
	func() {
		if calculating {
			return
		}
		calculating = true
		for _, btn := range buttons2 {
			btn.Disable()
		}
		for _, btn := range chudBut2 { // Refer to the comments in the initial assignment and creation of archimedesBtn1
			calculating = true
			btn.Enable()
		}
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
					chudnovskyBig(updateOutput2, chudDigits, done)
					calculating = false
					for _, btn := range buttons2 {
						btn.Enable()
					}
				}()
			})
	})

	BPPbut2 = []*ColoredButton{BBPfast44Btn2}
	chudBut2 = []*ColoredButton{ChudnovskyBtn2}
	nilaBut2 = []*ColoredButton{}

	buttons2 = []*ColoredButton{BBPfast44Btn2, NilakanthaBtn2, ChudnovskyBtn2} // array used only for range btn.Enable()

	// ::: Layout
		content2 := container.NewVBox(
			widget.NewLabel("\nSelect a method to estimate π:\n"),
			container.NewGridWithColumns(4, BBPfast44Btn2, NilakanthaBtn2, ChudnovskyBtn2),
			coloredScroll2, // Use coloredScroll2 directly or windowContent2 if you want an extra layer
		)
		windowContent2 := container.NewMax(bgwc2, content2) // Light green window bg
	
	window2.Canvas().SetOnTypedRune(func(r rune) { // Main-thread update loop using Fyne's lifecycle
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
	informationMenu := fyne.NewMenu("Actions and Information",
		fyne.NewMenuItem("Help", func() {
			dialog.ShowInformation("Information", "Help...", window2)
		}),
		fyne.NewMenuItem("Abort current method", func() {
			select { // select is a concurrency-specific channel-only construct used to handle multiple channel operations, see explanation in second comment-block below. 
			// // Check if the done channel is already closed (chan receive [<-] succeeds on a closed chan and false is returned in the case of chan type bool)
			case <-done: // chan syntax for receive on chan "done"
				updateOutput2("\nGoroutines already notified to terminate\n")
			default: // chan was open but empty, receive has "failed" (nothing to receive: "blocks"), case has "failed" (does not trigger), chan has blocked until a value is sent on the chan; default ensues 
				close(done) // "else" close the done chan, which will be interpreted as a termination signal by all listening processes
				// Assume chan initialization as: done := make(chan bool) // understanding that "bools are false upon creation, and chans nil till initialized"
				updateOutput2("\nTermination signals were sent to all current processes that may be listening\n")
			}
			/*
				operation (<-ch) on a closed channel:
				    Succeeds immediately (no blocking/waiting).
				    Returns the zero value of the channel’s type (false for chan bool, 0 for chan int, "" for chan string, etc.).
				When you try <-ch on an empty, open channel, it doesn’t fail — it blocks. Blocking means the operation pauses (waits) until something is put into the pipe
				... but in the context of a select, waiting is not succeeding, hence the default case is run.
			*/
			/*
				Switch: Like picking a door based on a number you’re holding — door 1, 2, or 3 opens depending on your number. Your num matches no doors? You get the default door.
					vs
				Select: Like waiting at a row of mailboxes for a letter to arrive — you grab the first one you see, or immediately walk away if you see none (default).
			*/
		}),
	)
		mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
		window2.SetMainMenu(mainMenu)

		window2.SetContent(windowContent2) // Set once with the full layout
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
