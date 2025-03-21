package main

import (
	"math/big"
	"strconv"
	"strings"
)

// @formatter:off

// Placeholder Pi calculation functions

func MonteCarloBig(print func(string)) {
	print("Starting Monte Carlo calculation...")
	print("Pi approximation: 3.14159 (stub)")
	print("Finished Monte Carlo.")
}


// fyneFunc func(string)
func formatWithThousandSeparators(num *big.Float) string {
	// Convert to big.Int
	numInt, _ := num.Int(nil)

	// Get the string representation
	numStr := numInt.String()

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
// fyneFunc func(string)
func formatInt64WithThousandSeparators(num int64) string {

	numStr := strconv.FormatInt(num, 10)

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

func formatFloat64WithThousandSeparators(numFloat64 float64) string {

	numStr := strconv.FormatFloat(numFloat64, 'f', -1, 64)

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

func check(e error) { // create a func named check which takes one parameter "e" of type error ::: - -
	if e != nil {
		// ::: that should do it ???	
		panic(e) // use panic() to display error code
	}
}


func checkPi(stringOfSum string) int {
	posInPi := 0               // to be the incremented offset : piChar = piAs59766chars[posInPi]
	var piChar byte            // one byte (character) of pi as string, e.g. piChar = piAs59766chars[posInPi]
	var stringVerOfCorrectDigits = []string{}
	for positionInString, charAtRangePos := range stringOfSum {
		piChar = bigPieAs255Chars[posInPi] // go.lang has limits on the size of constants, but this case is rather small
		if charAtRangePos == rune(piChar) {
			stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
			copyOfLastPosition  = positionInString // save an external global copy, of the last position found to have matched pi, as an int
		} else {
			break // to print result and info below
		}
		posInPi++
	}
	return copyOfLastPosition
}


// Updated output function with auto-scrolling
func updateOutput1(text string) {
	current := outputLabel1.Text
	if len(current) > 99500 {
		current = current[len(current)-10:]
	}
	outputLabel1.SetText(current + text)
	outputLabel1.Refresh()
	scrollContainer1.ScrollToBottom()
}
// fyneFunc(fmt.Sprintf("Send a message to the scroll area via fyne func"))
// Updated output function with auto-scrolling
func updateOutput2(text string) {
	current := outputLabel2.Text
	if len(current) > 99500 {
		current = current[len(current)-10:]
	}
	outputLabel2.SetText(current + text)
	outputLabel2.Refresh()
	scrollContainer2.ScrollToBottom()
}
// fyneFunc(fmt.Sprintf("Send a message to the scroll area via fyne func"))
