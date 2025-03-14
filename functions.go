package main

import (
	"fyne.io/fyne/v2"
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


func check(e error) { // create a func named check which takes one parameter "e" of type error
	if e != nil {
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

// RicksMouseDomain extends CanvasObject to catch and handle taps (mouse clicks). // RicksMouseDomain extends CanvasObject to snag and handle taps (mouse clicks galore)!
type RicksMouseDomain struct {
	fyne.CanvasObject // // Embeds the drawable base, size, and position (the first step towards extending it. 
	/* per grok:
	   Embedding fyne.CanvasObject (no field name, just type) gives RicksMouseDomain all its methods (Resize(), Move(), Refresh()) and
	   fields. It’s Go’s inheritance trick — CanvasObject is Fyne’s drawable foundation (rectangles, lines, circles). This makes
	   RicksMouseDomain a CanvasObject, ready for containers like staffContainer.Add(staffAreaTapped). Next, we add tap powers!
	*/
	/* said my way:
	this is a hybrid of inheritance-like behavior (via embedding) and event handling. fyne.CanvasObject is an embedded field (no name, just the type). Embedding lets
	RicksMouseDomain inherit all methods and fields of fyne.CanvasObject . CanvasObject is Fyne’s base interface for anything drawable, such as rectangles, lines, or
	circles. It includes methods like Resize(), Move(), and Refresh() . RicksMouseDomain thereby becomes a CanvasObject : and, it can therefore be added to containers
	(e.g., staffContainer.Add(staffAreaTapped) . Next, we extend it.
	*/
	OnTapped func(*fyne.PointEvent) // A named field of the struct, OnTapped, is a function type (a callback func : for tap action; which is where the magic happens!). 
	// It takes a *fyne.PointEvent (a pointer to a struct with X/Y coordinates and other event data) and returns nothing. This is the handler you’ll define later : 
	// it defines what happens when the player clicks on the canvas.
	/* per grok:
	OnTapped func(*fyne.PointEvent) // Named field — a callback function for tap action, where the magic unfolds!
	Takes *fyne.PointEvent (X/Y coords and event data), returns nada. Set this later to define player tap behavior!
	*/
}

// Tapped is a declared method, (t *RicksMouseDomain) is the method receiver; or method receiver declaration. It creates a local var 't' as a pointer to an instance of RicksMouseDomain ...
// t functions as the instance (below in t.OnTapped(e) 
/* per grok:
Tapped, a method with receiver (t *RicksMouseDomain)—declares ‘t’ as a pointer to our RicksMouseDomain instance!
*/
func (t *RicksMouseDomain) Tapped(e *fyne.PointEvent) { // Implements Fyne’s Tappable interface; tap-ready! Notice that RicksMouseDomain became an extended CanvasObject -- above.
	/* And; that other guy ---- ^ e will be used above as: OnTapped: func(e *fyne.PointEvent) { // e is the argument passed to OnTapped
	This is a method on RicksMouseDomain with a receiver t (the instance being tapped). Named Tapped : this is key! It implements Fyne’s Tappable interface, which
	requires a Tapped(*fyne.PointEvent) method. Notice that it takes the same *fyne.PointEvent as the OnTapped field/(callback func) : the coordinates of the tap or PointEvent.
	*/
	/* per grok:
	   Method on RicksMouseDomain — receiver ‘t’ is the tapped instance. ‘Tapped’ name is key: satisfies Fyne’s Tappable interface
	   with Tapped(*fyne.PointEvent). Matches OnTapped’s *fyne.PointEvent for tap coords!
	*/
	// t below is a local variable: and contains a pointer to a type (in this case the user-defined struct RicksMouseDomain). Which is an extended version of CanvasObject.
	/* per grok:
	t points to a RicksMouseDomain instance -- ‘t’ is the receiver — a pointer to this RicksMouseDomain instance, extended from CanvasObject.
	*/
	if t.OnTapped != nil { // This checks to assure that OnTapped was set (is not nil). In Go, function types default to nil if unassigned, preventing 
		// a panic crash that would result from calling a null/unset function. Safety is hereby assured; we'll have no nil crashes here! Proceed only if not nil.
		t.OnTapped(e) // This calls "back" to the stored OnTapped tap-handling function of the preceding struct, passing it the tap event (e). This delegates the actual ...
		// ... tap logic to whatever you plugged in (e.g., your note-placing code). It thereby fires off our custom tap handling logic.
		/* per grok:
		t.OnTapped(e) // Fires the stored OnTapped callback with tap event ‘e’ — unleashes your custom logic!
		*/
	}
}