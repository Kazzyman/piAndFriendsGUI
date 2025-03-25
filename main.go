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

	done := make(chan bool) // local, kill channel for all listening goroutines::: only Archimedes, and Wallis fon window1
	
	// Custom colored ::: Buttons1
	archimedesBtn1 := NewColoredButton(
		"Archimedes method for finding π, modified by Richard Woolley\n" +
			"a purely geometric method which is easy to understand\n" +
			"produces 3,012 digits of delicious Pi in under a minute\n" +
			"             -*-*-*- Rick's personal favorite -*-*-*-          ",
			color.RGBA{255, 100, 100, 215},
		func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			// ::: We want to cause the button that corresponds to the currently executing method to remain bright, while the other buttons remain dimmed...
			for _, btn := range archiBut { // This trick accomplishes that because the archiBut array comes after the creation of archimedesBtn1
				calculating = true // This keeps archimedesBtn1 from being restarted in parallel with itself...
				btn.Enable() // ... even though we herewith enable archimedesBtn1  ::: note that simply doing: archimedesBtn1.Enable() would not work...
			} // ::: ... because, we are inside of the creation of archimedesBtn1 [ it is a timing and scoping issue ]
			updateOutput1("\nRunning ArchimedesBig...\n\n")
			go func() {
				ArchimedesBig(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	/*
	.
	.
	 */
	JohnWallisBtn1 := NewColoredButton(
		"John Wallis infinite series -- 40 billion iterations -- runs 5m30s\n" +
		"π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ...\n" +
		"only manages to do 10 digits of Pi in well-over five minutes\n" +
		"an infinite series circa 1655    --- served here by Rick Woolley ---",		color.RGBA{110, 110, 255, 185}, 
			func() {
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range walisBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			updateOutput1("\nRunning John Wallis...\n\n")
			go func() { // made this the goroutine as per your example 
				JohnWallis(updateOutput1, done) // ::: func < - - - - - - - - - - - - - < -
				calculating = false
				for _, btn := range buttons1 {
					btn.Enable()
				}
			}()
		},
	)
	/*
	.
	.
	 */
	BBPfast44Btn1 := NewColoredButton(
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
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range BPPbut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}			
		updateOutput1("\nRunning BBP-fast-190 up to here...\n\n")
			
		showCustomEntryDialog(
			"Input Desired number of digits",
			"Any number less than 190",
			func(input string) {
				if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
					input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
					val, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Error converting input:", err)
						updateOutput1("Invalid input, using default 190 digits")
					} else if val <= 0 {
						updateOutput1("Input must be positive, using default 190 digits")
					} else if val > 10000 {
						updateOutput1("Input must be less than 191 -- using default of 190 digits")
					} else {
						BppDigits = val
					}
					go func() {
						bbpFast44(updateOutput1, BppDigits) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
						calculating = false
						for _, btn := range buttons1 {
							btn.Enable()
						}
					}()
				} else {
					// dialog canceled 
					updateOutput1("spigot calculation canceled, make another selection")
					for _, btn := range buttons1 {
						btn.Enable()
					}
					calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					// return // don't think I need this, don't know how it got here ?
				}
			},
		)
	})
	/*
	.
	.
	 */
	SpigotBtn1 := NewColoredButton(
		"The Spigot Algorithm, a Leibniz series. Served hot, bite by byte\n" +
			"spits out a nearly-unlimited, continuous stream of Pi goodness\n" +
			"This trick made possible by a bit of code mooched off of GitHub\n" +
			"bakes π without using any floating-point arithmetic",
		color.RGBA{110, 110, 255, 185},
		func() {
			var spigotDigits int
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range spigotBut { // Refer to the comments in the initial assignment and creation of archimedesBtn1
				calculating = true
				btn.Enable()
			}
			updateOutput1("\nRunning The Spigot...\n\n")
			
			showCustomEntryDialog(
				"Input Desired number of digits",
				"Any number less than 461",
				func(input string) {
					if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
						input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
						val, err := strconv.Atoi(input)
						if err != nil {
							fmt.Println("Error converting input:", err)
							updateOutput1("Invalid input, using default 460 digits")
						} else if val <= 0 {
							updateOutput1("Input must be positive, using default 460 digits")
						} else if val > 460 {
							updateOutput1("Input must be less than 461 -- using default of 460 digits")
						} else {
							spigotDigits = val
						}
						go func() {
							TheSpigot(updateOutput1, spigotDigits) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons1 {
								btn.Enable()
							}
						}()
					} else {
						// dialog canceled 
						updateOutput1("spigot calculation canceled, make another selection")
						for _, btn := range buttons1 {
							btn.Enable()
						}
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
						// return // don't think I need this, don't know how it got here ?
					}
				},
			)
		})
	/*
	.
	.
	pi = \frac{1}{12} \left[ \sum_{n=0}^{\infty} \frac{(-1)^n (6n)! (13591409 + 545140134n)}{(3n)! (n!)^3 (640320^{3n + 3/2})} \right]^{-1}      */ 
	ChudnovskyBtn1 := NewColoredButton(
		"Chudnovsky -- by David & Gregory Chudnovsky -- late 1980s\n" +
			"extremely efficient, quickly bakes world-record quantities of Pi\n" +
			"this algorithm is a rapidly converging infinite series which\n" +
			"leverages properties of j-invariant from elliptic function theory",
		color.RGBA{255, 255, 100, 235}, 
	func() {
		// 
		var chudDigits int
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range chudBut { // chudBut is an array with only one member
				calculating = true // keep it from being restarted in parallel
				btn.Enable() // even though the button is enabled
			}
		updateOutput1("\nRunning Chudnovsky...\n\n")

		showCustomEntryDialog(
			"Input Desired number of digits",
			"Any number less than 49,999",
			func(input string) {
				if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
					input = removeCommasAndPeriods(input) // allow user to enter a number with a comma
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
						chudnovskyBig(updateOutput1, chudDigits) // ::: func < - - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
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
						calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					// return // don't think I need this, don't know how it got here ?
				}
			},
		)
	})
	/*
	.
	.
	 */
	MontyBtn1 := NewColoredButton(
		"Monte Carlo method for converging on π  --  big floats, & float64\n" +
			"Flavor: no fancy equations are used, only Go's pure randomness\n" +
			"4 digits of pi in 21s ; 7 digits possible in 1h30m w/ 119k grid\n" +
			"                   -*-*- Rick's second-favorite method -*-*-     ",
		color.RGBA{255, 255, 100, 235}, 
	func() {
		var MontDigits string
			if calculating {
				return
			}
			calculating = true
			for _, btn := range buttons1 {
				btn.Disable()
			}
			for _, btn := range montBut { // montBut is an array with only one member
				calculating = true // keep it from being restarted in parallel
				btn.Enable() // even though the button is enabled
			}
		updateOutput1("\nRunning Monte Carlo ...\n\n")
			
		showCustomEntryDialog(
			"Input Desired number of grid elements",
			"max 120,000; 10,000 will produce 4 pi digits, 110,00 may get you 5 digits",
			func(input string) {
				if input != "" { // This if-else is part of the magic that allows us to dismiss a dialog and allow others to run after the dialog is canceled/dismissed.
					input = removeCommasAndPeriods(input) // ::: allow user to enter a number with a comma
					val, err := strconv.Atoi(input) // val is now an int and input is a string
					if err != nil {
						fmt.Println("Error converting input:", err)
						updateOutput1("Invalid input, using default 10,000 digits")
					} else if val <= 1 {
						updateOutput1("Input must be greater than 1, using default 10,000 digits")
					} else if val > 120000 {
						updateOutput1("Input must be less than 120,001 -- using default of 10,000 digits")
					} else {
						MontDigits = strconv.Itoa(val) // val here is a number, an int to be precise. So, we use strconv.Itoa to convert the int to a string and assign it to MontDigits. 
					}
						go func() {
							Monty(updateOutput1, MontDigits) // ::: func < - - - - - - - - - - - - < -  NOT AMENABLE TO KILLING VIA A DONE CHANNEL 
							calculating = false
							for _, btn := range buttons1 {
								btn.Enable()
							}
						}()
				} else {
					// dialog canceled 
					updateOutput1("Monte Carlo calculation canceled, make another selection")
					for _, btn := range buttons1 {
						btn.Enable()
					}
					calculating = false // ::: this is the trick to allow others to run after the dialog is canceled/dismissed.
					// return // don't think I need this, don't know how it got here ?
				}
			},
		)
	})

	archiBut = []*ColoredButton{archimedesBtn1} // 1
	walisBut = []*ColoredButton{JohnWallisBtn1} // 1
	BPPbut = []*ColoredButton{BBPfast44Btn1} // 2
	spigotBut = []*ColoredButton{SpigotBtn1} // 2
	chudBut = []*ColoredButton{ChudnovskyBtn1} // 3      used as bug preventions // keep methods from being started or restarted in parallel (over-lapping) 
	montBut = []*ColoredButton{MontyBtn1} // 3
	
	buttons1 = []*ColoredButton{archimedesBtn1, JohnWallisBtn1, BBPfast44Btn1, SpigotBtn1, ChudnovskyBtn1, MontyBtn1} // used only for range btn.Enable()

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
			select { // select is a concurrency-specific channel-only construct used to handle multiple channel operations, see explanation in second comment-block below. 
			// // Check if the done channel is already closed (chan receive [<-] succeeds on a closed chan and false is returned in the case of chan type bool)
			case <-done: // chan syntax for receive on chan "done"
				updateOutput1("\nGoroutines already notified to terminate\n")
			default: // chan was open but empty, receive has "failed" (nothing to receive: "blocks"), case has "failed" (does not trigger), chan has blocked until a value is sent on the chan; default ensues 
				close(done) // "else" close the done chan, which will be interpreted as a termination signal by all listening processes
				// Assume chan initialization as: done := make(chan bool) // understanding that "bools are false upon creation, and chans nil till initialized"
				updateOutput1("\nTermination signals were sent to all current processes that may be listening\n")
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
	window1.SetMainMenu(mainMenu)
	
	window1.SetContent(content1) 
	window1.ShowAndRun()
}
