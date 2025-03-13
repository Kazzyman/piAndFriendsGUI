package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"math/big"
	"strings"
	"time"
	"fyne.io/fyne/v2/theme"
)

// @formatter:off

func ArchimedesBig(fyneFunc func(string)) {
	fyneFunc(fmt.Sprintf("\nYou've selected a demonstration of Rick's improved version of Archimedes' method for aproximating the value of Pi : 3.14159...\n"))

	fyneFunc(fmt.Sprintf("The goal is to accurately calculate over 2,700 correct digits of Pi. We'll need to use floating-point numbers with thousands of decimal places."))
	fyneFunc(fmt.Sprintf("This can be done using Rick's most-favoured language: go.lang, or simply Go. GoLand (by JetBrains) wil be our IDE."))
	fyneFunc(fmt.Sprintf("We'll also be using the Fyne.io package to create a graphical (windowed) user interface.\n"))

	fyneFunc(fmt.Sprintf("Let's begin with creating our variables (all must be big.Floats)\n"))
	
	r := big.NewFloat(1); fyneFunc(fmt.Sprintf("r := big.NewFloat(1)\t\t\t// The radius of our hexigon will always be == 1"))
	s1 := big.NewFloat(1); fyneFunc(fmt.Sprintf("s1 := big.NewFloat(1)\t\t\t// s1 will initially be == 1"))
	numberOfSides := big.NewFloat(6); fyneFunc(fmt.Sprintf("numberOfSides := big.NewFloat(6)\t\t// The initial number of sides for our polygon"))
				fyneFunc(fmt.Sprintf("\n"))
			
	a := new(big.Float); fyneFunc(fmt.Sprintf("a := new(big.Float)\t\t\t// variable for height of bisected triangle : side \"a\""))
	b := new(big.Float); fyneFunc(fmt.Sprintf("b := new(big.Float)\t\t// variable for new short side \"b\""))
	p := new(big.Float); fyneFunc(fmt.Sprintf("p := new(big.Float)\t\t\t// variable for Perimeter of the triangle: \"p\""))
	s2 := new(big.Float); fyneFunc(fmt.Sprintf("s2 := new(big.Float)\t\t\t// variable for new hypotenuse : new side \"s2\""))
	p_d := new(big.Float); fyneFunc(fmt.Sprintf("p_d := new(big.Float)\t\t\t// variable for calculated pi thus far"))
	s1_2 := new(big.Float); fyneFunc(fmt.Sprintf("s1_2 := new(big.Float)\t\t\t// s1_2 variable will be s1/2 or half of s1"))
				fyneFunc(fmt.Sprintf("\nNext we'll set go's precision to 55000 on all of our variables.\n"))

	precision := 55000; fyneFunc(fmt.Sprintf("precision := 55000\t\t\t// Initial valeue of Go's math/big objects precision"))
	p_d.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("p_d.SetPrec(uint(precision))\t\t// Set the precision of all the above Big Floats"))
	a.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("a.SetPrec(uint(precision))\t\t//  \""))
	s1_2.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("s1_2.SetPrec(uint(precision))\t\t//  \""))
	s2.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("s2.SetPrec(uint(precision))\t\t//  \""))
	b.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("b.SetPrec(uint(precision))\t\t//  \""))
	p.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("p.SetPrec(uint(precision))\t\t//  \""))
	r.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("r.SetPrec(uint(precision))\t\t//  \""))
	s1.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("s1.SetPrec(uint(precision))\t\t//  \""))
	numberOfSides.SetPrec(uint(precision)); fyneFunc(fmt.Sprintf("numberOfSides.SetPrec(uint(precision))\t\t//  \""))
				fyneFunc(fmt.Sprintf("\nNext, we do some initial asignments and calculations:\n"))

	numberOfSides.Mul(numberOfSides, big.NewFloat(2)); fyneFunc(fmt.Sprintf("numberOfSides.Mul(numberOfSides, big.NewFloat(2))\t\t// Double the number of sides"))
	s1_2.Quo(s1, big.NewFloat(2)); fyneFunc(fmt.Sprintf("s1_2.Quo(s1, big.NewFloat(2))\t\t\t// Set s1_2 = s1/2"))
	a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) // a = Sqrt(r-(s1_2)^2) where r is always 1  [all values must be big floats, including the const 1]
				fyneFunc(fmt.Sprintf("a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2)))\t\t// a = Sqrt(r-(s1_2)^2)"))
	
	fyneFunc(fmt.Sprintf("\n\t\tTo calculate the height (a) of a right triangle formed by bisecting a side of a polygon inscribed in a unit circle (radius r = 1). \n\t\t\tThe polygon’s side length (s1) is halved (s1_2 = s1 / 2), and this computation helps refine the polygon’s perimeter to approximate \n\t\t\tπ as the number of sides increases.")) // print to canvas a rune of the above comments 
	
	b.Sub(r, a); fyneFunc(fmt.Sprintf("\nb.Sub(r, a)\t\t// b = r-a  where a: is the Height of the bisected triangle (from the prior step)."))
	s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2)))
				fyneFunc(fmt.Sprintf("s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2)))\t\t// s2 = Sqrt(b*b + s1_2*s1_2)"))
	/* The above statement is explained below:
			Inputs:
			  b: short side from midpoint to circle edge (a big float)
			  s1_2: half the current side length (s1 / 2, a big float)
			Output:
			  s2: new side length of the polygon (a big float)
		
			1. Compute b^2
			temp1 = b * b
		
			2. Compute (s1_2)^2
			temp2 = s1_2 * s1_2
		
			3. Add the two squares
			temp3 = temp1 + temp2
		
			4. Take the square root to get the new side length
			s2 = square_root(temp3)
	 */
	fyneFunc(fmt.Sprintf("\n\t\t\tHere is some seudo code for the algorithm:\n\n\t\t\tInputs:\n\t\t\t\t  b: short side from midpoint to circle edge (a big float)\n\t\t\t\t  s1_2: half the current side length (s1 / 2, a big float)\n\t\t\tOutput:\n\t\t\t\t  s2: new side length of the polygon (a big float)\n\t\t\n\t\t\t1. Compute b^2\n\t\t\ttemp1 = b * b\n\t\t\n\t\t\t2. Compute (s1_2)^2\n\t\t\ttemp2 = s1_2 * s1_2\n\t\t\n\t\t\t3. Add the two squares\n\t\t\ttemp3 = temp1 + temp2\n\t\t\n\t\t\t4. Take the square root to get the new side length\n\t\t\ts2 = square_root(temp3)"))
	
	s1.Set(s2); fyneFunc(fmt.Sprintf("\ns1.Set(s2)\t\t\t// Use big.Float method Set, to asign s2 to s1"))
	p.Mul(numberOfSides, s1); fyneFunc(fmt.Sprintf("p.Mul(numberOfSides, s1)\t\t// p = numberOfSides * s1"))
	p_d.Set(p); fyneFunc(fmt.Sprintf("p_d.Set(p)\t\t\t// Use big.Float method Set, to asign p to p_d  'read p sub d'"))
	fyneFunc(fmt.Sprintf("\n"))
	
	fyneFunc(fmt.Sprintf("for i := 0; i < 5001; i++ {"))
	fyneFunc(fmt.Sprintf("\tnumberOfSides.Mul(numberOfSides, big.NewFloat(2)) \t\t// Double the number of sides"))
	fyneFunc(fmt.Sprintf("\ts1_2.Quo(s1, big.NewFloat(2)) \t\t\t\t// s1_2 = s1/2"))
	fyneFunc(fmt.Sprintf("\ta.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) \t// a = Sqrt(r-(s1_2)^2)"))
	fyneFunc(fmt.Sprintf("\tb.Sub(r, a) \t\t\t// b = r-a"))
	fyneFunc(fmt.Sprintf("\ts2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) \t// a = Sqrt(r-(s1_2)^2)"))
	fyneFunc(fmt.Sprintf("\ts1.Set(s2) \t\t\t// Use big.Float method Set, to asign s2 to s1"))
	fyneFunc(fmt.Sprintf("\tp.Mul(numberOfSides, s1) \t\t\t// p = numberOfSides * s1"))
	fyneFunc(fmt.Sprintf("\tp_d.Set(p) \t\t\t// Use big.Float method Set, to asign p to p_d  'read p sub d'"))
	fyneFunc(fmt.Sprintf("\tp_d.Quo(p_d, big.NewFloat(2)) \t\t\t// p_d = p_d / 2"))
				fyneFunc(fmt.Sprintf("}\n"))

	fyneFunc(fmt.Sprintf("if i == 24 { "))
	fyneFunc(fmt.Sprintf("\tfyneFunc(fmt.Sprintf(\"    '.20f' is the big.Float of what we have calculated  ----- per Archimedes' at 24 iters, 20f\", p_d))"))
	fyneFunc(fmt.Sprintf("}"))
	fyneFunc(fmt.Sprintf("\nEtcetera ..."))

	fyneFunc(fmt.Sprintf("\n\nThe results of our iterative calculations follow:"))
		
	for i := 0; i < 5001; i++ { 
		numberOfSides.Mul(numberOfSides, big.NewFloat(2)) // ; 
		s1_2.Quo(s1, big.NewFloat(2)) // ; 
		a.Sqrt(new(big.Float).Sub(r, new(big.Float).Mul(s1_2, s1_2))) // ; 
		b.Sub(r, a) // ; 
		s2.Sqrt(new(big.Float).Add(new(big.Float).Mul(b, b), new(big.Float).Mul(s1_2, s1_2))) // ; 
		s1.Set(s2) // ; 
		p.Mul(numberOfSides, s1) // ; 
		p_d.Set(p) // ; 
		p_d.Quo(p_d, big.NewFloat(2)) // ; 

		if i == 24 { 
			fyneFunc(fmt.Sprintf("    %.20f is the big.Float of what we have calculated  ----- per Archimedes' at 24 iters, 20f", p_d))
			fyneFunc("    3.141592653589793238  vs the value of π from the web")

			number := numberOfSides
		
			// Convert to big.Int
			numInt, _ := number.Int(nil)
		
			// Get the string representation
			numStr := numInt.String()
		
			// Manually format with thousand separators
			formattedNum := formatWithThousandSeparators(numStr)
		
			fmt.Println("big.Int formatted:", formattedNum)
					fyneFunc(fmt.Sprintf("rick: the above was estimated from a %s  --- sided polygon\n", formattedNum))
					fyneFunc("... Mister A. would have wept!\n\n")
		}

		if i == 50 {
			fyneFunc(fmt.Sprintf("    %.33f is the big.Float of what we have calculated  ----- per Archimedes' at 50 iters, 33f", p_d))
			fyneFunc("    3.141592653589793238462643383279502  ----- is the value of π from the web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 31 correct digits of π", i))

			// Convert to big.Int
			numInt, _ := numberOfSides.Int(nil)

			// Get the string representation
			numStr := numInt.String()

			// Manually format with thousand separators
			formattedNum := formatWithThousandSeparators(numStr)
			fyneFunc(fmt.Sprintf(" the above was estimated from a %s  --- sided polygon\n\n", formattedNum))
		}

		if i == 150 {
			fyneFunc(fmt.Sprintf("   %.95f   ----- per Rick's modified Archimedes' method", p_d))
			fyneFunc("   3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211  ----- is from web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 92 correct digits of π", i))

			// Convert to big.Int
			numInt, _ := numberOfSides.Int(nil)

			// Get the string representation
			numStr := numInt.String()

			// Manually format with thousand separators
			formattedNum := formatWithThousandSeparators(numStr)
			fyneFunc(fmt.Sprintf(" the above was estimated from a %s  --- sided polygon\n\n", formattedNum))
		}

		if i == 200 {
			fyneFunc(fmt.Sprintf("   %.122f   ----- per Rick's modified Archimedes' method", p_d))
			fyneFunc("   3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709  ----- is from web")
			fyneFunc(fmt.Sprintf("%d iterations were completed yielding 121 correct digits of π", i))

			// Convert to big.Int
			numInt, _ := numberOfSides.Int(nil)

			// Get the string representation
			numStr := numInt.String()

			// Manually format with thousand separators
			formattedNum := formatWithThousandSeparators(numStr)
			fyneFunc(fmt.Sprintf(" the above was estimated from a %s  --- sided polygon\n\n", formattedNum))
			fyneFunc(" ... working ...\n\n")
		}

		if i == 1200 || i == 2200 || i == 3200 || i == 4200 {
			fyneFunc(fmt.Sprintf("... still working, %d iterations completed ...\n", i))
		}

		if i == 5000 { // was 5500
			// fyneFunc(fmt.Sprintf(" A peek at the result formatted 1500f is: %.1500f \nper Archimedes'\n", p_d)) // show the first 1,500 digits of calculated pi
			
			// fyneFunc(fmt.Sprintf("%d iterations were completed, \n ... which generated a %.999f sided polygon!!\n", i, numberOfSides))

			// Convert to big.Int
			numInt, _ := numberOfSides.Int(nil)

			// Get the string representation
			numStr := numInt.String()

			// Manually format with thousand separators
			formattedNum := formatWithThousandSeparators(numStr)
			fyneFunc(fmt.Sprintf("Done! So, how many sides does our polygon have now? A lot:\n\nA staggering:\n\n%s\n\nSIDED POLYGON !!!\n\n", formattedNum))
			fyneFunc(fmt.Sprintf("%d iterations were completed to yeild 2,712 correct digits of π!!!\n", i))
			fyneFunc(fmt.Sprintf("Go's math/big objects were set to a precision value of: %d  --- here is your GIANT slice of pie:\n", precision))
			fyneFunc(fmt.Sprintf("  %.2800f \n\nper Rick's modified Archimedes' method\n\n", p_d))
		}
		time.Sleep(1 * time.Millisecond) // Slow it down slightly for visibility
	}
}


// Define a custom theme with larger text
type myTheme struct {
	fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 20 // Larger default text size
	}
	return theme.DefaultTheme().Size(name)
}

func main() {
	myApp := app.New()
	// myApp.Settings().SetTheme(&myTheme{theme: theme.DefaultTheme()})
	myApp.Settings().SetTheme(&myTheme{Theme: theme.DefaultTheme()})
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
		fmt.Println(oneLineSansCR)                // Keep CLI output
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


func formatWithThousandSeparators(numStr string) string {
	// Handle negative numbers
	prefix := ""
	if strings.HasPrefix(numStr, "-") {
		prefix = "-"
		numStr = numStr[1:]
	}
	// Insert commas every three digits from the right
	result := ""
	for i, char := range numStr {
		if i > 0 && (len(numStr)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return prefix + result
}