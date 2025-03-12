package main

import (
	"fmt"
	"math/big"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// @formatter:off
// ; fyneFunc(fmt.Sprintf(""))

func ArchimedesBig(fyneFunc func(string)) {
	fyneFunc(fmt.Sprintf("\nYou selected Rick's improved version of Archimedes' method\n"))
	
	r := big.NewFloat(1); fyneFunc(fmt.Sprintf("r := big.NewFloat(1)  // The radius of our hexigon will always be == 1"))
	s1 := big.NewFloat(1); fyneFunc(fmt.Sprintf("s1 := big.NewFloat(1)  // s1 will initially be == 1"))
	numberOfSides := big.NewFloat(6); fyneFunc(fmt.Sprintf("numberOfSides := big.NewFloat(6)  // The initial number of sides for our polygon"))
				fyneFunc(fmt.Sprintf("\n"))
			
	a := new(big.Float); fyneFunc(fmt.Sprintf("a := new(big.Float)  // variable for height of bisected triangle : side \"a\""))
	b := new(big.Float); fyneFunc(fmt.Sprintf("b := new(big.Float)  // variable for new short side \"b\""))
	p := new(big.Float); fyneFunc(fmt.Sprintf("p := new(big.Float)  // variable for Perimeter of the triangle: \"p\""))
	s2 := new(big.Float); fyneFunc(fmt.Sprintf("s2 := new(big.Float)   // variable for new hypotenuse : new side \"s2\""))
	p_d := new(big.Float); fyneFunc(fmt.Sprintf("p_d := new(big.Float)  // variable for calculated pi thus far"))
	s1_2 := new(big.Float); fyneFunc(fmt.Sprintf("s1_2 := new(big.Float) // s1_2 variable will be s1/2 or half of s1"))

	precision := 55000
	p_d.SetPrec(uint(precision))
	a.SetPrec(uint(precision))
	s1_2.SetPrec(uint(precision))
	s2.SetPrec(uint(precision))
	b.SetPrec(uint(precision))
	p.SetPrec(uint(precision))
	r.SetPrec(uint(precision))
	s1.SetPrec(uint(precision))
	numberOfSides.SetPrec(uint(precision))

	numberOfSides.Mul(numberOfSides, big.NewFloat(2))
	s1_2.Quo(s1, big.NewFloat(2))
	a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))
	b.Sub(r, a)
	s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2)))
	s1.Set(s2)
	p.Mul(numberOfSides, s1)
	p_d.Set(p)

	for i := 0; i < 5001; i++ {
		numberOfSides.Mul(numberOfSides, big.NewFloat(2))
		s1_2.Quo(s1, big.NewFloat(2))
		a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))
		b.Sub(r, a)
		s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2)))
		s1.Set(s2)
		p.Mul(numberOfSides, s1)
		p_d.Set(p)
		p_d.Quo(p_d, big.NewFloat(2))

		if i == 24 {
			fyneFunc(fmt.Sprintf("    %.20f is the big.Float of what we have calculated per Archimedes' at 24 iters, 20f", p_d))
			fyneFunc("    3.141592653589793238 is the value of π from the web")
			fyneFunc(fmt.Sprintf(" the above was estimated from a %.0f sided polygon", numberOfSides))
			fyneFunc(fmt.Sprintf("%.0f as parsed against ...", numberOfSides))
			fyneFunc("100000000 which is one-hundred-million, for comparison to the above line")
			fyneFunc("... Which is to say a 201,326,592 sided polygon, Mr. A. would have wept!\n")
		}

		if i == 50 {
			fyneFunc(fmt.Sprintf("    %.33f is the big.Float of what we have calculated per Archimedes' at 50 iters, 33f", p_d))
			fyneFunc("    3.141592653589793238462643383279502 is the value of π from the web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 31 correct digits of π", i))
			fyneFunc(fmt.Sprintf("the above was estimated from a %.0f sided polygon (formatted as a .0f) \n\n", numberOfSides))
		}

		if i == 150 {
			fyneFunc(fmt.Sprintf("  %.95f per Archimedes'", p_d))
			fyneFunc("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211 is from web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 92 correct digits of π", i))
			fyneFunc(fmt.Sprintf("Calculated from a %.0f sided polygon\n\n", numberOfSides))
		}

		if i == 200 {
			fyneFunc(fmt.Sprintf("  %.122f per Archimedes'", p_d))
			fyneFunc("  3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709 is from web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 121 correct digits of π", i))
			fyneFunc(fmt.Sprintf("Calculated from a %.0f sided polygon\n\n\n", numberOfSides))
			fyneFunc(" ... working ...\n\n")
		}

		if i == 1200 || i == 2200 || i == 3200 || i == 4200 {
			fyneFunc(fmt.Sprintf("... still working, %d iterations completed ...\n", i))
		}

		if i == 5000 { // was 5500
			fyneFunc(fmt.Sprintf(" A peek at the result formatted 1500f is: %.1500f \nper Archimedes'\n", p_d)) // show the first 1,500 digits of calculated pi
			fyneFunc(fmt.Sprintf("%d iterations were completed, \n ... which generated a %.999f sided polygon!!\n", i, numberOfSides))
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 2,712 correct digits of π!!!\n", i))
			fyneFunc(fmt.Sprintf("Go's math/big objects were set to a precision value of:%d", precision))
			fyneFunc(fmt.Sprintf("  %.2800f per Archimedes'", p_d))
		}

		time.Sleep(1 * time.Millisecond) // Slow it down slightly for visibility
	}
}


func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Archimedes Pi")
	myWindow.Resize(fyne.NewSize(1900, 1600)) // Adjust to 2065x1350 for your app

	// Button to trigger the function
	button := widget.NewButton("Calculate Pi", nil)

	// Label for output
	outputLabel := widget.NewLabel("Press the button to start...\n")
	outputLabel.Wrapping = fyne.TextWrapWord

	// Scrollable container for the label
	scrollContainer := container.NewScroll(outputLabel)
	scrollContainer.SetMinSize(fyne.NewSize(1900, 1300)) // Set a fixed height, adjust as needed

	// Vertical layout: button on top, scrollable output below
	content := container.New(layout.NewVBoxLayout(), button, scrollContainer)

	// Update function to append output
	var outputText string
	callBkPrn2canvas := func(oneLineSansCR string) {
		outputText += oneLineSansCR + "\n"
		outputLabel.SetText(outputText)
		// fmt.Println(line)                // Keep CLI output
		scrollContainer.ScrollToBottom() // Auto-scroll to the latest output
	}

	// Button click handler
	button.OnTapped = func() {
		outputText = "" // Reset output
		outputLabel.SetText(outputText)
		go ArchimedesBig(callBkPrn2canvas) // Run with selection 14
	}

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
