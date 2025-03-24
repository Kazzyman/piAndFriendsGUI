package main

import (
	"fmt"
	"os"
	"time"
)

// @formatter:off

func GregoryLeibniz(fyneFunc func(string)) {
	usingBigFloats = false
	
	codeSnippet := `
What follows is the actual source code for this method:

    usingBigFloats = false
	fyneFunc(fmt.Sprintf("\n\nYou selected the Gregory-Leibniz series ... this will be quick!\n\n"))
	fyneFunc(fmt.Sprintf("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...\n"))
	fyneFunc(fmt.Sprintf("Three-hundred-million iterations will be executed ... working ...\n"))
	start := time.Now()
	iterFloat64 = 0
	var nextOdd float64
	nextOdd = 1
	four = 4
	var tally float64
	tally = (four / nextOdd)
	iterInt64 = 0
	for iterInt64 < 300000000 {
		iterInt64++
		iterFloat64++
		nextOdd = nextOdd + 2
		tally = tally - (tally / nextOdd)
		tally = tally + (tally / nextOdd) // pi (tally) is set equl to the sum of a subtraction and an addition, alternatively

		if iterInt64 == 10000000 {
			fyneFunc(fmt.Sprintf("... 10,000,000 of three hundred million completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("   %0.6f was calculated by the Gregory-Leibniz series\n", tally))
			fyneFunc(fmt.Sprintf("   3.141592,653589793 is from the web\n\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  10,000,000 iterations in %s yields 7 digits of π\n\n", elapsed))
		}
		// 7 digits of Pi were found per the above code
		// the next two ifs give eight digits of Pi
		if iterInt64 == 50000000 {
			fyneFunc(fmt.Sprintf("... 50,000,000 of three hundred million completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("      %0.8f was calculated by the Gregory-Leibniz series\n", tally))
			fyneFunc(fmt.Sprintf("     3.1415926,53589793 is from the web\n\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  50,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
			fyneFunc(fmt.Sprintf(" "))
		}
		if iterInt64 == 100000000 {
			fyneFunc(fmt.Sprintf("... 100,000,000 of three hundred million completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("      %0.9f was calculated by the Gregory-Leibniz series\n", tally))
			fyneFunc(fmt.Sprintf("     3.1415926,53589793 is from the web\n\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  100,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
		}
		// 9 digits of Pi are found below
		if iterInt64 == 200000000 {
			fyneFunc(fmt.Sprintf("... 200,000,000 of three hundred million completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("      %0.10f was calculated by the Gregory-Leibniz series\n", tally))
			fyneFunc(fmt.Sprintf("     3.14159265,3589793 is from the web"))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  200,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
		}
		if iterInt64 == 300000000 { // last one, still 9 digits
			fyneFunc(fmt.Sprintf("       %0.11f was calculated by the Gregory-Leibniz series\n", tally))
			fyneFunc(fmt.Sprintf("\n      3.141592653589793 is from the web\n\n"))
			t := time.Now()
			elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  300 million iterations still yields 9 digits, ")) // no Println here
			fyneFunc(fmt.Sprintf("in %s\n\n", elapsed))
			fyneFunc(fmt.Sprintf(" per option  --  the Gregory-Leibniz series, circa 1676\n\n"))

			LinesPerIter = 11 // an estimate of the number of lines per iteration
			fyneFunc(fmt.Sprintf("at aprox %0.2f lines of code per iteration ...\n", LinesPerIter))
			LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fyneFunc(fmt.Sprintf("       %.0f lines of code were executed per second \n\n", LinesPerSecond))

			fyneFunc(fmt.Sprintf("\n That was the Gregory-Leibniz series:\n\n"))
			fyneFunc(fmt.Sprintf("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...\n\n"))
			fyneFunc(fmt.Sprintf("Three-hundred-million iterations were executed. This section was written entirely by Richard Woolley\n\n"))

			
			// store results in a log file which can be displayed from within the program by selecting option #12
			fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
			Hostname, _ := os.Hostname()
			_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- on %s \n", Hostname)
			check(err0)
			current_time := time.Now()
			_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
			check(err6)
			_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
			check(err2)
			_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
			check(err4)
			_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
			check(err5)
			TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
			check(err7)
		}
	}
	output := fmt.Sprintf("\n%s\nIs my test code snip\n", codeSnippet)
	fyneFunc(output)

	// written entirely by Richard Woolley
`
	fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err1)                     // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
			defer fileHandle.Close()   // It’s idiomatic to defer a Close immediately after opening a file.
	
	output := fmt.Sprintf("\n%s\nThat, was the actual code.\n", codeSnippet)
	fyneFunc(output)
	
	fyneFunc(fmt.Sprintf("\n\nYou selected the Gregory-Leibniz series ... this will be quick!\n\n"))
	fyneFunc(fmt.Sprintf("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...\n"))
	fyneFunc(fmt.Sprintf("Three-hundred-million iterations will be executed\n\n ... working ...\n\n\n"))

	var nextOdd float64
	var tally float64
	
			start := time.Now()
			
	iterFloat64 = 0
	nextOdd = 1
	four = 4
	tally = (four / nextOdd)
	iterInt64 = 0
	
	for iterInt64 < 300000000 {
		iterInt64++
		iterFloat64++
		nextOdd = nextOdd + 2
		tally = tally - (tally / nextOdd)
		tally = tally + (tally / nextOdd) // pi (tally) is set equl to the sum of a subtraction and an addition, alternatively

		if iterInt64 == 10000000 {
			fyneFunc(fmt.Sprintf("... 10,000,000 of three hundred million iterations already completed. \n\nstill working, but ...\n\n"))
			fyneFunc(fmt.Sprintf("   %0.6f was calculated thus far via the Gregory-Leibniz series\n", tally))
				t := time.Now()
				elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  10,000,000 iterations in %s yields 7 digits of π\n\n", elapsed))
		}
		// 7 digits of Pi were found per the above code
		// the next two ifs give eight digits of Pi
		if iterInt64 == 50000000 {
			fyneFunc(fmt.Sprintf("... 50,000,000 of three hundred million completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("      %0.8f was calculated by the Gregory-Leibniz series, so far\n", tally))
				t := time.Now()
				elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  50,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
			fyneFunc(fmt.Sprintf(" "))
		}
		if iterInt64 == 100000000 {
			fyneFunc(fmt.Sprintf("... 100,000,000 of three hundred million completed. still working, and ...\n"))
			fyneFunc(fmt.Sprintf("      %0.9f was calculated by the Gregory-Leibniz series\n", tally))
				t := time.Now()
				elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  100,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
		}
		// 9 digits of Pi are found below
		if iterInt64 == 200000000 {
			fyneFunc(fmt.Sprintf("... 200,000,000 of three hundred million now completed. still working, but ...\n"))
			fyneFunc(fmt.Sprintf("      %0.10f was calculated thus far by the Gregory-Leibniz series\n", tally))
				t := time.Now()
				elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  200,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
		}
		if iterInt64 == 300000000 { // last one, still 9 digits
			fyneFunc(fmt.Sprintf("       %0.11f was calculated by the Gregory-Leibniz series \n", tally))
				t := time.Now()
				elapsed := t.Sub(start)
			fyneFunc(fmt.Sprintf("  300 million iterations have finally finished; still yielding only 9 digits of pi, ")) // no Println here
			fyneFunc(fmt.Sprintf("in %s\n\n", elapsed))
			fyneFunc(fmt.Sprintf(" per the Gregory-Leibniz series, circa 1676\n\n"))

				LinesPerIter = 11 // an estimate of the number of lines per iteration
				linePerApp := LinesPerIter * 300000000
				stringOfTotal := formatFloat64WithThousandSeparators(linePerApp)
				LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
			fyneFunc(fmt.Sprintf("at aprox %0.0f lines of code per iteration ... SLOC executed was aprox. %s \n", LinesPerIter, stringOfTotal))
			fyneFunc(fmt.Sprintf("       %.0f lines of code were executed per second \n\n", LinesPerSecond))

					Hostname, _ := os.Hostname()
					current_time := time.Now()
					TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
			
			fyneFunc(fmt.Sprintf("\n That was the Gregory-Leibniz series:\n\n"))
			fyneFunc(fmt.Sprintf("π = (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...\n\n"))
			fyneFunc(fmt.Sprintf("Runtime was: %s\n", TotalRun))
			fyneFunc(fmt.Sprintf("Three-hundred-million iterations were executed. This section was written entirely by Richard Woolley\n\n"))

			
			// store results in a log ::: file

						
				_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gregory-Leibniz -- on %s \n", Hostname)
					check(err0)
				_, err6 := fmt.Fprint(fileHandle, "was run on: ", current_time.Format(time.ANSIC), "\n")
					check(err6)
				_, err2 := fmt.Fprintf(fileHandle, "%.0f was Lines/Second  \n", LinesPerSecond)
					check(err2)
				_, err4 := fmt.Fprintf(fileHandle, "%e was Iterations/Seconds  \n", iterFloat64/elapsed.Seconds())
					check(err4)
				_, err5 := fmt.Fprintf(fileHandle, "%e was total Iterations  \n", iterFloat64)
					check(err5)
				_, err7 := fmt.Fprintf(fileHandle, "Total run was %s \n ", TotalRun)
					check(err7)
		}
	}
	// ::: Prepare to exit the Gregory Leibniz method function
	calculating = false // Allow another method to be selected.
	for _, btn := range buttons2 { // ok to only Enable buttons2, because I expect to only ever execute this from window2
		btn.Enable() // ::: Enable
	}
	// written entirely by Richard Woolley
}