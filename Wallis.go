package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func JohnWallis(fyneFunc func(string)) { // case 8: // -- AMFJohnWallisA
	usingBigFloats = false
	fyneFunc(fmt.Sprintf("\n   You selected A Go language exercize which can be used to test the speed of your hardware."))
	fyneFunc(fmt.Sprintf("   We will calculate π to a maximum of ten digits of accuracy using an infinite series by John Wallis circa 1655"))
	fyneFunc(fmt.Sprintf("   Up to 40 Billion iterations of the following formula will be executed "))
	fyneFunc(fmt.Sprintf("   π = 2 * ((2/1)*(2/3)) * ((4/3)*(4/5)) * ((6/5)*(6/7)) ..."))
	start := time.Now()
	iterFloat64 = 0
	var numerators float64
	numerators = 2
	var firstDenom float64
	firstDenom = 1
	var secondDenom float64
	secondDenom = 3
	var cumulativeProduct float64
	cumulativeProduct = (numerators / firstDenom) * (numerators / secondDenom)
	iterInt64 = 0
	for iterInt64 < 1000000000 {
		iterInt64++
		iterFloat64++
		numerators = numerators + 2
		firstDenom = firstDenom + 2
		secondDenom = secondDenom + 2
		cumulativeProduct = cumulativeProduct * (numerators / firstDenom) * (numerators / secondDenom)
		π = cumulativeProduct * 2
		/*            if iterInt64 == 100 {
		                  fyneFunc(fmt.Sprintf("%0.9f calculated using an infinite series by John Wallis circa 1655", π))
		                  fyneFunc(fmt.Sprintf("    3.1,415926535897933.14159265358979323846264338327950288419716939937510  is, again, the value of π from the web"))
		                  t := time.Now()
		                  elapsed := t.Sub(start)
		                  fyneFunc(fmt.Sprintf(iterInt64, " iterations were completed in %s yielding 2 digits of π\n", RunTimeAsString))
		              }
		              if iterInt64 == 500 {
		                  fyneFunc(fmt.Sprintf("%0.9f calculated using an infinite series by John Wallis circa 1655", π))
		                  fyneFunc(fmt.Sprintf("    3.14,15926535897933.14159265358979323846264338327950288419716939937510  is, again, the value of π from the web"))
		                  t := time.Now()
		                  elapsed := t.Sub(start)
		                  fyneFunc(fmt.Sprintf(iterInt64, " iterations were completed in %s yielding 3 digits of π\n", RunTimeAsString))
		              }
		*/
		if iterInt64 == 2000 {
			fyneFunc(fmt.Sprintf("%0.5f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("%d iterations were completed in %s yielding 4 digits of π\n", iterInt64, RunTimeAsString))
		}
		if iterInt64 == 10000 {
			fyneFunc(fmt.Sprintf("%0.6f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("10,000 iterations were completed in %s yielding 5 digits of π\n", RunTimeAsString))
		}
		if iterInt64 == 50000 { // 50,000
			fyneFunc(fmt.Sprintf("%0.7f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("50,000 iterations were completed in %s yielding 5 digits of π\n", RunTimeAsString))
		}
		if iterInt64 == 500000 { // 500,000 done
			fyneFunc(fmt.Sprintf("%0.8f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("500,000 iterations were completed in %s yielding 6 digits of π\n", RunTimeAsString))
		}
		if iterInt64 == 2000000 { // 2M done
			fyneFunc(fmt.Sprintf("%0.9f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("2,000,000 iterations were completed in %s yielding 7 digits of π\n", RunTimeAsString))
		}
		if iterInt64 == 40000000 { // 40M done
			fyneFunc(fmt.Sprintf("%0.10f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()

			piAsAString := strconv.FormatFloat(π, 'g', -1, 64)
			copyOfLastPosition = checkPi(piAsAString)
			fyneFunc(fmt.Sprintf("40,000,000 iterations were completed in %s yielding %d confirmed digits of π\n\n", RunTimeAsString, copyOfLastPosition))
			fyneFunc(fmt.Sprintf("  .. working .. on another factor-of-ten iterations\n"))
		}
		if iterInt64 == 400000000 { // 400M done
			fyneFunc(fmt.Sprintf("%0.11f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is, again, the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()

			piAsAString := strconv.FormatFloat(π, 'g', -1, 64)
			copyOfLastPosition = checkPi(piAsAString)

			fyneFunc(fmt.Sprintf("400,000,000 iterations were completed in %s yielding %d confirmed digits of π\n\n", RunTimeAsString, copyOfLastPosition))

			LinesPerIter = 36 // an estimate
			fyneFunc(fmt.Sprintf("at aprox %0.1f lines of code per iteration ...", LinesPerIter))
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds()
			formattedLinesPerSecond := formatInt64WithThousandSeparators(int64(LinesPerSecond)) // .Seconds() returns a float64
			fyneFunc(fmt.Sprintf("Aprox %s lines of code were executed per second \n", formattedLinesPerSecond))
			// a brief Red notification follows :
			fyneFunc(fmt.Sprintf(" ... will be working on doing Billions more iterations ...\n\n"))
		}
		//
		if iterInt64 == 600000000 { // 600M done
			fyneFunc(fmt.Sprintf("  600M done, still working on another Two-Hundred-Thousand iterations ... working ...\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("%s \n", RunTimeAsString))
			fyneFunc(fmt.Sprintf("Calculating the next digit of pi may require 40B iterations, which takes a few minutes \n"))
			fyneFunc(fmt.Sprintf("- Ctrl-C to End/Exit without saving results\n"))
			LinesPerIter = 36 // an estimate
			fyneFunc(fmt.Sprintf("at aprox %0.1f lines of code per iteration ...", LinesPerIter))
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds()
			formattedLinesPerSecond := formatInt64WithThousandSeparators(int64(LinesPerSecond)) // .Seconds() returns a float64
			fyneFunc(fmt.Sprintf("Aprox %s lines of code were executed per second \n", formattedLinesPerSecond))
			fyneFunc(fmt.Sprintf(" ... still working ..."))
		}
		if iterInt64 == 800000000 { // 800M done
			fyneFunc(fmt.Sprintf("  800M done, still working on yet another Two Hundred Thousand iterations ... working ...\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()
			fyneFunc(fmt.Sprintf("%s \n", RunTimeAsString))
		}
		if iterInt64 == 1000000000 { // 1B done
			fyneFunc(fmt.Sprintf("%0.11f calculated using an infinite series by John Wallis circa 1655", π))
			fyneFunc(fmt.Sprintf("3.14159265358  is the value of π from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			RunTimeAsString := elapsed.String()

			piAsAString := strconv.FormatFloat(π, 'g', -1, 64)
			copyOfLastPosition = checkPi(piAsAString)
			fyneFunc(fmt.Sprintf("\nOne Billion iterations were completed in %s still only yielding π to %d confirmed digits\n", RunTimeAsString, copyOfLastPosition))
			fyneFunc(fmt.Sprintf(" per --  an infinite series by John Wallis circa 1655\n")) // ----------------------

			LinesPerIter = 36 // an estimate
			fyneFunc(fmt.Sprintf("at aprox %0.1f lines of code per iteration ...", LinesPerIter))
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds()
			formattedLinesPerSecond := formatInt64WithThousandSeparators(int64(LinesPerSecond)) // .Seconds() returns a float64
			fyneFunc(fmt.Sprintf("Aprox %s lines of code were executed per second \n", formattedLinesPerSecond))

			// store reults in a log file which can be displayed from within the program by selecting option #12
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis -- on %s \n", Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			_, err2 := fmt.Fprintf(fileHandle, "%s was Lines/Second  \n", formattedLinesPerSecond)
			check(err2)
			_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
			check(err4)
			_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
			check(err5)
			TotalRun := elapsed.String()                                         // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) // add total runtime of this calculation
			check(err7)
		}
	} // end of first for loop

	// :::fyneFunc(fmt.Sprintf("Enter any positive digit to continue with an additional 39 billion iterations, 0 to exit"))
	option39 := 1
	// fmt.Scan(&option39)
	if option39 > 0 {

		fyneFunc(fmt.Sprintf("\n\nYou elected to continue the infinite series by John Wallis"))
		fyneFunc(fmt.Sprintf("\n    an additionl 39 billion iterations will be executed \n\n   ... working ...\n"))

		fyneFunc(fmt.Sprintf(" ... still working ... on Billions of iterations, 39 to go ...\n"))

		fyneFunc(fmt.Sprintf("\n ... 39 Billion additional loops now ensue, just to get maybe one additional digit of pi"))

		start := time.Now()

		for iterInt64 < 40000000000 {
			iterInt64++
			iterFloat64++
			numerators = numerators + 2
			firstDenom = firstDenom + 2
			secondDenom = secondDenom + 2
			cumulativeProduct = cumulativeProduct * (numerators / firstDenom) * (numerators / secondDenom)
			π = cumulativeProduct * 2

			if iterInt64 == 2000000000 { // 2B completed
				fyneFunc(fmt.Sprintf("  2B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 3000000000 { // 3B completed
				fyneFunc(fmt.Sprintf("  3B done, still working ... on another Billion iterations ... working ... Ctrl-C to End/Exit without saving stats"))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 4000000000 { // 4B completed
				fyneFunc(fmt.Sprintf("  4B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 5000000000 { // 5B completed
				fyneFunc(fmt.Sprintf("  5B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 6000000000 { // 6B completed
				fyneFunc(fmt.Sprintf("  6B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 7000000000 { // 7B completed
				fyneFunc(fmt.Sprintf("  7B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 8000000000 { // 8B completed
				fyneFunc(fmt.Sprintf("  8B done, still working ... on another Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 9000000000 { // 9B completed
				fyneFunc(fmt.Sprintf("  9B done, still working ... on another five Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 14000000000 { // 14B completed
				fyneFunc(fmt.Sprintf("  14B done, still working ... on another five Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 19000000000 { // 19B completed
				fyneFunc(fmt.Sprintf("  19B done, still working ... on another five Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 24000000000 { // 24B completed
				fyneFunc(fmt.Sprintf("  24B done, still working ... on another five Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 29000000000 { // 29B completed
				fyneFunc(fmt.Sprintf("  29B done, still working ... on another five Billion iterations ... working ..."))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 34000000000 { // 34B completed
				fyneFunc(fmt.Sprintf("  34B done, still working ... just another six Billion iterations to go! ... "))
				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()
				fyneFunc(fmt.Sprintf("%s", RunTimeAsString))
			}
			if iterInt64 == 40000000000 { // 40B completed
				fyneFunc(fmt.Sprintf("%0.12f calculated using an infinite series by John Wallis circa 1655", π))
				fyneFunc(fmt.Sprintf("3.14159265358  is the value of π from the web"))

				t := time.Now()
				elapsed := t.Sub(start)
				RunTimeAsString := elapsed.String()

				piAsAString := strconv.FormatFloat(π, 'g', -1, 64)
				copyOfLastPosition = checkPi(piAsAString)
				fyneFunc(fmt.Sprintf("Forty Billion iterations were completed in %s yielding π to %d confirmed digits\n", RunTimeAsString, copyOfLastPosition))
				fyneFunc(fmt.Sprintf(" per --  an infinite series by John Wallis circa 1655\n")) // ----------------------
				LinesPerIter = 36                                                                // an estimate
				fyneFunc(fmt.Sprintf("at aprox %0.1f lines of code per iteration ...", LinesPerIter))
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds()
				formattedLinesPerSecond := formatInt64WithThousandSeparators(int64(LinesPerSecond)) // .Seconds() returns a float64
				fyneFunc(fmt.Sprintf("Aprox %s lines of code were executed per second \n", formattedLinesPerSecond))

				// store reults in a log file which can be displayed from within the program by selecting option #12
				fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
				defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
				Hostname, _ := os.Hostname()
				_, err0 := fmt.Fprintf(fileHandle, "\n  -- John Wallis (cont.) -- on %s \n", Hostname)
				check(err0)
				current_time := time.Now()
				_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
				check(err6)
				_, err2 := fmt.Fprintf(fileHandle, "%s was Lines/Second  \n", formattedLinesPerSecond)
				check(err2)
				_, err4 := fmt.Fprintf(fileHandle, "%.02f was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
				check(err4)
				_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
				check(err5)
				TotalRun := elapsed.String()                                         // cast time durations to a String type for Fprintf "formatted print"
				_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun) // add total runtime of this calculation
				check(err7)
			}
		} // end of second for loop
	} // end of 40B continuation if
	// written entirely by Richard Woolley
} // end of JohnWallis()
