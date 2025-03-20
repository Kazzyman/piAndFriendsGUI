package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

// @formatter:off

	func NilakanthaBig(fyneFunc func(string), iters int, precision int) { // Changed signature ::: - -

	if iters > 36111222 {
		// fyneFunc(fmt.Sprintf(" ... "))
		fyneFunc(fmt.Sprintf(" ... working ... Nilakantha using big floats")) // Send to channel
	}
	if iters > 42000000 {
		// fyneFunc(fmt.Sprintf("... werkin ..."))
		fyneFunc(fmt.Sprintf("... werkin ...")) // Send to channel
	}
	if iters > 55111222 {
		fyneFunc(fmt.Sprintf("... working for a while ..."))
	}
	if iters > 69111222 {
		fyneFunc(fmt.Sprintf("... will be working for quite a while ..."))
	}
	if iters > 80111222 {
		fyneFunc(fmt.Sprintf("... a very long while ... working ..."))
	}

	start := time.Now()
	var iterBig int

	// big.Float constants:

	twoBig := big.NewFloat(2)
	threeBig := big.NewFloat(3)
	fourBig := big.NewFloat(4)

	// big.Float variables:

	digitoneBig := new(big.Float)
	*digitoneBig = *twoBig

	digittwoBig := new(big.Float)
	*digittwoBig = *threeBig

	digitthreeBig := new(big.Float)
	*digitthreeBig = *fourBig

	sumBig := new(big.Float)
	nexttermBig := new(big.Float)

	// set precision to a user-specified value
	sumBig.SetPrec(uint(precision))
	twoBig.SetPrec(uint(precision))
	threeBig.SetPrec(uint(precision))
	fourBig.SetPrec(uint(precision))
	digitoneBig.SetPrec(uint(precision))
	digittwoBig.SetPrec(uint(precision))
	digitthreeBig.SetPrec(uint(precision))
	nexttermBig.SetPrec(uint(precision))

	sumBig.Add(threeBig, new(big.Float).Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig))))

	iterBig = 1
	for iterBig < iters { // 1,000,000,000 yeilds 19 digits in 13 sec

		// 256p 100,000,000 : 56s 25 digits digits June 26 2023
		// 1,280p and 1Bil : 23 min without ending
		// 128p and 100,000,000 49s gave 25 digits June 26 2023
		// Total run with SetPrec at: 128 and iters of 1,000,000,000 was 7m57.3179415s
		// got 31 digits in 1 hour and 26 min using this algorithm with one billion iters at 128 prec
		// 1,000,000,002 and 64 bits prec yielded :: 17 digits in 5m41s

		iterBig++

		digitoneBig.Add(digitoneBig, twoBig)
		digittwoBig.Add(digittwoBig, twoBig)
		digitthreeBig.Add(digitthreeBig, twoBig)

		nexttermBig.Quo(fourBig, new(big.Float).Mul(digitoneBig, new(big.Float).Mul(digittwoBig, digitthreeBig)))

		if iterBig%2 == 0 { // % is modulus operator
			sumBig.Sub(sumBig, nexttermBig)
		} else {
			sumBig.Add(sumBig, nexttermBig)
		}

		if iterBig == 20111222 {
			// fyneFunc(fmt.Sprintf(" ... doin some ... "))
			fyneFunc(fmt.Sprintf(" ... doin some ... ")) // Send to channel
		}
		if iterBig == 36111222 {
			fyneFunc(fmt.Sprintf(" ... werkin ... "))
		}
		if iterBig == 42000000 {
			fyneFunc(fmt.Sprintf("... still werkin ... Nilakantha Somayaji method using big.Float types -- with some patience one can generate 31 correct digits of pi this way."))
		}
		if iterBig == 55111222 {
			fyneFunc(fmt.Sprintf("... been working for a while ..."))
		}
		if iterBig == 69111222 {
			fyneFunc(fmt.Sprintf("... been working for quite a while ..."))
		}
		if iterBig == 80111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while ... but still working ..."))
		}
		if iterBig == 180111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, 180,111,222 down, ... and still working ..."))
		}
		if iterBig == 280111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, 280,111,222 down, ... and still working ..."))
		}
		if iterBig == 480111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, 480,111,222 down, ... still working ..."))
		}
		if iterBig == 680111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, 680,111,222 down, ...  working ..."))
		}
		if iterBig == 880111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, down, 880,111,222, down ... still, working ..."))
		}
		if iterBig == 977111222 {
			fyneFunc(fmt.Sprintf("... it's been a very long while, 977,111,222 already ... why am I still working? ..."))
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String()

	var piAs150chars string
	//  59,766 is the limit of the size of token Go can handle, so the following example of a maximally-long sequence of digits of pi will have to suffice.
	piAs150chars = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284"
	stringOfSum := sumBig.Text('f', 999999) // create a string version of a big result, up to a million digits can be handled, if not verified, here-with
	shortStringOfSumbig := stringOfSum[:30]
	/*
	shortStringOfSumbig and copyOfLastPosition

	These variables are used in fmt.Sprintf calls to create strings that are sent to the updateChan. It is important to verify that the value of these variables are correctly calculated before they are used.
	 */

	posInPi := 0    // to be the incremented offset : piChar = piAs59766chars[posInPi]
	var piChar byte // one byte (character) of pi as string, e.g. piChar = piAs59766chars[posInPi]
	// var copyOfLastPosition int // an external (to the loop) copy of positionInString ::: already in globals.go
	var stringVerOfCorrectDigits = []string{}
	for positionInString, charAtRangePos := range stringOfSum {
		piChar = piAs150chars[posInPi]
		if charAtRangePos == rune(piChar) {
			stringVerOfCorrectDigits = append(stringVerOfCorrectDigits, string(charAtRangePos))
			copyOfLastPosition = positionInString // save an external copy, of the last position found to have matched pi, as an int
		} else {
			break // to print result and info below
		}
		posInPi++
	}

	if copyOfLastPosition > 55000 { // if length of pi is > 55,000 digits we have something really big
		// print (log) to a special file
		fyneFunc(fmt.Sprintf("\n\n\nWe have been tasked with making a lot of pie and it was sooo big it needed its own file ...\n"))
		fyneFunc(fmt.Sprintf("\n\n  After allowing this process to finish (you may have to continue prodding this thing along for a while) ... \n"))
		fyneFunc(fmt.Sprintf("... Go have a look in /.big_pie_is_in_here.txt to find all the digits of π you had requested. \n\n"))

		fileHandleBig, err1prslc2c := os.OpenFile("big_pie_is_in_here.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1prslc2c)                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandleBig.Close()                                                                                    // It’s idiomatic to defer a Close immediately after opening a file.
		_, err2prslc2c := fmt.Fprintf(fileHandleBig, "\nThese are the %d verified digits we have calculated  :: \n", copyOfLastPosition)
		check(err2prslc2c)
		for _, oneChar := range stringVerOfCorrectDigits {
			// fmt.Print(oneChar) // to the console // the whole point of using an alternate file is to not clutter up the console or the default file
			// *************************************** this is the one and only logging loop ******************************************************************************
			_, err8prslc2c := fmt.Fprint(fileHandleBig, oneChar) // to a file
			check(err8prslc2c)
		}
		_, err9prslc2c := fmt.Fprintf(fileHandleBig, "\n...the preceeding was logged one char at a time \n")
		check(err9prslc2c)
		fileHandleBig.Close()
	} else if true {
			// regularExpression := regexp.MustCompile(`^3.1.........................................................................................`)
			// firstSectionOfPiFromWeb := regularExpression.FindStringSubmatch(piAs59766chars)
			fyneFunc(fmt.Sprintf("\npi from the web begins thusly: 3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534"))

		fyneFunc(fmt.Sprintf("\npi as calculated herein is: %s", shortStringOfSumbig))



		fyneFunc(fmt.Sprintf("\n.... we have matched %d digits: ", copyOfLastPosition))
			
			

			fileHandleNilakan, err1prslc2c := os.OpenFile("dataLog-From_Nilakantha_Method_lengthy_prints.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1prslc2c)                                                                                                                            // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandleNilakan.Close()                                                                                                               // It’s idiomatic to defer a Close immediately after opening a file.
			_, err2prslc2c := fmt.Fprintf(fileHandleNilakan,
				"\n\nBelow are the %d verified digits we have calculated via Nilakantha using precision of %d and iterations of %d: \n",
				copyOfLastPosition, precision, iterBig)
			check(err2prslc2c)
			for _, oneChar := range stringVerOfCorrectDigits {
				// fmt.Print(oneChar) // to the console we log pi, one digit at a time
				_, err8prslc2c := fmt.Fprint(fileHandleNilakan, oneChar) // to a file we log pi one digit at a time
				check(err8prslc2c)
			}
			fileHandleNilakan.Close()
	} else {
		fileHandleDefault, err1prslc2d := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1prslc2d)                                                                                                                    // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandleDefault.Close()                                                                                                       // It’s idiomatic to defer a Close immediately after opening a file.
		// to file
		_, err2prslc2d := fmt.Fprintf(fileHandleDefault, "\nThese are the %d verified digits we have calculated  :: \n", copyOfLastPosition)
		check(err2prslc2d)
		// to screen
		// fyneFunc(fmt.Sprintf("\n These are the %d verified digits we have calculated: \n", copyOfLastPosition))
		fyneFunc(fmt.Sprintf("\n These are the %d verified digits we have calculated: \n", copyOfLastPosition))

		for _, oneChar := range stringVerOfCorrectDigits {
			// to screen
			fmt.Print(oneChar)
			// to file
			_, err8prslc2d := fmt.Fprint(fileHandleDefault, oneChar)
			check(err8prslc2d)
		}
		fileHandleDefault.Close()
	}
		fyneFunc(fmt.Sprintf("\n"))


		fileHandleDefault, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1)                                                                                                                    // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
		defer fileHandleDefault.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
		Hostname, _ := os.Hostname()
		_, err0 := fmt.Fprintf(fileHandleDefault, "\n  -- Nilakantha Somayaji -- on %s \n", Hostname)
		check(err0)
		current_time := time.Now()
		_, err6 := fmt.Fprint(fileHandleDefault, "was run on: ", current_time.Format(time.ANSIC), "\n")
		check(err6)
		_, err5 := fmt.Fprintf(fileHandleDefault, "%d was total Iterations; %d was precision setting for the big.Float types \n", iterBig, precision)
		check(err5)
		// TotalRun := elapsed.String() // e.g., it was cast to a String type in another func and passed to this func as TotalRun
		_, err7 := fmt.Fprintf(fileHandleDefault, "Total run was %s \n ", TotalRun)
		check(err7)
		_, err2prslc2da := fmt.Fprint(fileHandleDefault, "\nResults from running Nilakantha (case 5:) can be viewed in dataLog-From_Nilakantha_Method_lengthy_prints.txt\n")
		check(err2prslc2da)
		fileHandleDefault.Close()
	

	// print to screen:
	fyneFunc(fmt.Sprintf(" via Nilakantha with big floats. Written entirely by Richard Woolley\n"))
	fmt.Printf("Total run with SetPrec at: %d and iters of %d was %s \n\n ", precision, iterBig, TotalRun)
	// written entirely by Richard Woolley
} // end of Nilakantha_Somayaji_with_big_Float_types() // -- AMFNilakantha_Somayaji_with_big_Float_typesB
