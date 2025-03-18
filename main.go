package main // limits scope (main is for all executable, others are for libraries or non executable packages)

import ( // packages to include in the build 
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"sync"
	"time"
)

// @formatter:off // ::: causes GoLand to not touch your formatting 

type highContrastTheme struct { // a user-defined type, based on fyne.Theme -- pairs with the color method below ::: - -
	fyne.Theme // since there is only this one statement in the struct, highContrastTheme is really just a synonym for fyne.Theme -- needed to make a modifiable instance. 
}
// The following method returns a color based on input parameters. (the struct allows the method to override or extend fyne.Theme)
func (h *highContrastTheme) color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color { // parentheses after func defines this as a method -- attached to a type.  ::: - -
	switch name { // *highC... is a dereferenced address/pointer to a custom type; func can modify instance. * is dereference operator; used to access the value stored at a memory address.
	case theme.ColorNameBackground:				// ColorNameBackground is a const string defined in image/color (package: image, file: color) -- image/color/Theme.Color
		return color.White						// color refers to the image/color package; White is the exported/exposed element: Color ... i.e. ...
	case theme.ColorNameInputBackground:		// Color(ThemeColorName, ThemeVariant) color.Color // from: type Theme interface {} which also contains 2 resources and a float 32
		return color.White						// returns white: (typically RGBA{255, 255, 255, 255})
	case theme.ColorNameForeground:				// ColorNameForeground was defined thus: ColorNameForeground fyne.ThemeColorName = "foreground" -- in package theme/color
		return color.Black						// 
	}										// .Theme.Color [below] means, or refers to: the exposed part of: image/color/Theme.Color
	return h.Theme.Color(name, variant) // "else" return name fyne.ThemeColorName, variant fyne.ThemeVariant -- h is local var name for instance aka: receiver, self, this, handler.
} // hence forward we could call highContrastTheme.color which would resolve to either Black or White -- as: RGBA() (r, g, b, a uint32)

type updateData struct { // ::: - -
	text      string
	clearText bool
}

func main() {
	countSLOC()
	myApp := app.New()
	myApp.Settings().SetTheme(&highContrastTheme{Theme: theme.LightTheme()}) // ??? points to a custom theme, defined above as a method called color [non-exported due to lower case name] ...
		fmt.Printf("color.Color is: %s\n", "what could I put here to see what is/was returned by the above color method?")
	myWindow := myApp.NewWindow("Fast Pi calculators")					// '&' is the address-of operator. It is used to get the memory address of a variable (see examples below).
	myWindow.Resize(fyne.NewSize(1900, 1600))
		
			// Declare a variable of type int and assign it 42
			var Value int = 42
			fmt.Printf("value is: %d\n", Value) // prints: value is: 42
		
			// Use & to get the memory address of value and store it in ptr as type int 
			var ptr *int = &Value
			fmt.Println(*ptr, "is value accessed via a pointer of type int") // prints: 42 is value accessed via a pointer of type int
		
			// Print the value of Value and its memory address
			fmt.Printf("Value: %d, Address of var Value: %p\n", Value, &Value) // prints: Value: 42, Address of var Value: 0x140002a99a0
		
			// Use * to dereference the pointer and access the value at the memory address of Value
			fmt.Printf("Dereferenced pointer yeilds value of Value: %d\n", *ptr) // prints: Dereferenced pointer yeilds value of Value: 42
		
			// Modify the value called Value through the pointer
			*ptr = 100 
			// &ptr = 100 // would give two errors: Cannot assign to &ptr ; and: '100' (type untyped int) cannot be represented by the type **int
			// ptr = 100 // would give error: '100' (type untyped int) cannot be represented by the type *int
		
			// Print the new value of Value
			fmt.Printf("New value of Value: %d\n", Value) // prints: New value of Value: 100

	outputLabel := widget.NewLabel("Press a button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1100)) // was 1300

	promptLabel := widget.NewLabel("")
	inputContainer := container.NewVBox()
	inputContainer.Hide() // for Nilakantha's two input fields

	// var outputText string
	updateChan := make(chan updateData, 100) // Changed to struct, from the previous var outputText string
		// Note::: updateChan is used for BOTH  slow and fast-scrolling methods, i.e., 
		//  ::: updateChan <- updateData{text: with normal function calls,    and ,     callBkPrn2canvas + Fyne with go routines instead of normal function calls 
		
		// NilakanthaBig(updateChan chan updateData, iters int, precision int) // ::: slower scrolling method -- is OK for Nilakantha since it does no scrolling  
		//  ::: updateChan <- updateData{text: ,   with   , normal function calls
			
	var mu sync.Mutex

	callBkPrn2canvas := func(oneLineSansCR string) {
		updateChan <- updateData{text: oneLineSansCR}
	}
	
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
		entryFields[0].Move(fyne.NewPos(0, 0)) // ::: what are these doing ????
		entryFields[1].Move(fyne.NewPos(230, 0))
		submitBtn.Move(fyne.NewPos(390, 0))
		hbox.Resize(fyne.NewSize(500, 40))

		inputContainer.Add(container.NewBorder(nil, nil, nil, nil, hbox))
		inputContainer.Resize(fyne.NewSize(500, 60))
		inputContainer.Show()

		return inputChan
	}
	
	// ::: ===== get input value (one int)  ========== = = = = = = = = = = = = = = = = = = = = = = = = =  
	getInputValue := func(digits int) chan string {
		inputContainer.Objects = nil
	
		entryField := widget.NewEntry()
		entryField.SetPlaceHolder("e.g., 9256")
		entryField.Resize(fyne.NewSize(150, 40))

		inputChan := make(chan string)

		submitBtn := widget.NewButton("Submit", func() {
				value := entryField.Text
				fmt.Println("Input value:", value)
			inputContainer.Hide()
			promptLabel.SetText("")
			fmt.Println("Submit button clicked")
			inputChan <- value
			close(inputChan) // Close the channel
		})

		submitBtn.Resize(fyne.NewSize(95, 40))
		submitBtn.Importance = widget.HighImportance

		hbox := container.NewWithoutLayout(
			entryField,
			submitBtn,
		)
		// entryFields[0].Move(fyne.NewPos(0, 0)) // ::: what ??? ^^^
		entryField.Move(fyne.NewPos(230, 0))
		submitBtn.Move(fyne.NewPos(390, 0))
		hbox.Resize(fyne.NewSize(500, 40))

		inputContainer.Add(container.NewBorder(nil, nil, nil, nil, hbox))
		inputContainer.Resize(fyne.NewSize(500, 60))
		inputContainer.Show()

		return inputChan
	}
	
	// ::: Buttons ========== = = = = = = = = = = = = = = = = = = = = = = = = =  
	buttonArchimedes := NewColoredButton("modified Archimedes \n-- by Rick Woolley\n three\n four", color.RGBA{255, 100, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go ArchimedesBig(callBkPrn2canvas)
	})
	buttonLeibniz := NewColoredButton("Gottfried Wilhelm Leibniz -- runs long", color.RGBA{100, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GottfriedWilhelmLeibniz(callBkPrn2canvas)
	})
	
	// ::: complex buttons - - - - - - -
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

	// ::: button - - - - - - - - -
	buttonChudnovsky := NewColoredButton("chudnovsky -- takes input", color.RGBA{255, 255, 100, 255}, func() {
		updateChan <- updateData{clearText: true}
		go func() {
			printAprompt(callBkPrn2canvas)
			inputChan := getInputValue(9)
			stringVerOfReqDigitsOfPi := <-inputChan 
					
			ReqDigitsOfPi := 999

			val1, err1 := strconv.Atoi(stringVerOfReqDigitsOfPi)
			if err1 != nil {
				fmt.Println("Error converting input1:", err1)
				fmt.Println("setting iters to 40,000")
				ReqDigitsOfPi = 40000
			} else {
				fmt.Println("Value of input1:", val1)
				ReqDigitsOfPi = val1
			}
			updateChan <- updateData{clearText: true}
			go chudnovskyBig(callBkPrn2canvas, updateChan, callBkPrn2canvas, ReqDigitsOfPi) // ::: two identical callbacks :
			// chudnovskyBig(fyneFunc func(string), updateChan chan updateData, callBkPrn2canvas func(oneLineSansCR string), digits int)
			// ::: first will be named: fyneFunc,  while the second gets named: callBkPrn2canvas
			// the first being customary; while the second is for passing-on ::: (I could have just made a copy at the destination)
		}()
	})

	// more simple buttons:
	buttonGregory := NewColoredButton("Gregory-Leibniz, is quick", color.RGBA{100, 100, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go GregoryLeibniz(callBkPrn2canvas)
	})
	buttonMonteCarlo := NewColoredButton("Monte Carlo", color.RGBA{100, 255, 255, 255}, func() {
		updateChan <- updateData{clearText: true}
		go MonteCarloBig(callBkPrn2canvas)
	})
	
	// stub buttons:
	buttonExtra1 := NewColoredButton("Extra 1", color.RGBA{200, 200, 200, 255}, func() {
		updateChan <- updateData{text: "Extra 1 clicked"}
	})

	buttonExtra2 := NewColoredButton("Extra 2", color.RGBA{150, 150, 150, 255}, func() {
			updateChan <- updateData{text: "Extra 2 clicked"}
	})

	// load our eight buttons: 
	buttonContainer := container.NewGridWithColumns(4,
		buttonArchimedes, buttonLeibniz, buttonGregory, buttonNilakantha,
		buttonChudnovsky, buttonMonteCarlo, buttonExtra1, buttonExtra2,
	)
	content := container.NewVBox(buttonContainer, promptLabel, inputContainer, scrollContainer)
	myWindow.SetContent(content)

	// ::: Main-thread update loop using Fyne's lifecycle = = = = = = = = = =
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
	// ::: end of Main-thread update loop using Fyne's lifecycle = = = = = = = = 

	// Drop-Down Menus
	logFilesMenu := fyne.NewMenu("Log Files",
		fyne.NewMenuItem("View Log 1", func() {
			// Implement log file viewing here
			dialog.ShowInformation("Log Files", "Viewing Log 1", myWindow)
		}),
		fyne.NewMenuItem("View Log 2", func() {
			// Implement log file viewing here
			dialog.ShowInformation("Log Files", "Viewing Log 2", myWindow)
		}),
	)

	windowsMenu := fyne.NewMenu("Collections",
		fyne.NewMenuItem("Fast Pi calculators", func() {
			myWindow.Show()
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
			dialog.ShowInformation("Information", "Help...", myWindow)
		}),
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("Information", "About...", myWindow)
		}),
	)

	mainMenu := fyne.NewMainMenu(logFilesMenu, windowsMenu, informationMenu)
	myWindow.SetMainMenu(mainMenu)

	myWindow.ShowAndRun()
} // ::: end of main

// helper func for chudnovsky method
func printAprompt(fyneFunc func(string)){
	fyneFunc(fmt.Sprintf("Enter the number of digits of pi to calculate per the chudnovsky method; \nThe sky is the limit with this method, so don't be shy."))
}