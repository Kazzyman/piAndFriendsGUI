package main

import (
	"math/big"
	"strings"
)

// @formatter:off

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