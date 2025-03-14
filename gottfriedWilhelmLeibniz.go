package main

import (
	"fmt"
	"os"
	"time"
)

// @formatter:off

func GottfriedWilhelmLeibniz(fyneFunc func(string)){
usingBigFloats = false
fyneFunc(fmt.Sprintf("\n\nYou selected Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... "))
fyneFunc(fmt.Sprintf("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton"))
fyneFunc(fmt.Sprintf("    ... and Gottfried Wilhelm Leibniz"))
fyneFunc(fmt.Sprintf("   4 Billion iterations will be executed ... "))
fyneFunc(fmt.Sprintf(""))
fyneFunc(fmt.Sprintf(" ... working ...\n"))
start := time.Now()
iterFloat64 = 0
var denom float64
denom = 3
var sum float64
sum = 1 - (1 / denom)
iterInt64 = 1
for iterInt64 < 4000000000 {
iterFloat64++
iterInt64++
denom = denom + 2
if iterInt64%2 == 0 {
sum = sum + 1/denom
} else {
sum = sum - 1/denom
}
π = 4 * sum
if iterInt64 == 100000000 {
fyneFunc(fmt.Sprintf("... 100,000,000 completed iterations ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.1415926,53589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  100,000,000 iterations in %s yields 8 digits of π\n\n", elapsed))
}
if iterInt64 == 200000000 {
fyneFunc(fmt.Sprintf("... 200,000,000 gets another digit ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.14159265,3589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  200,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
}
if iterInt64 == 400000000 {
fyneFunc(fmt.Sprintf("... 400,000,000 iterations completed, still at nine ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.14159265,3589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  400,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
}
if iterInt64 == 600000000 {
fyneFunc(fmt.Sprintf("... 600,000,000 iterations, still at nine ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.14159265,3589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  600,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
}
if iterInt64 == 1000000000 {
fyneFunc(fmt.Sprintf("... 1 Billion iterations completed, still nine ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.14159265,3589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  1,000,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
}
if iterInt64 == 2000000000 {
fyneFunc(fmt.Sprintf("... 2 Billion, and still just nine ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.14159265,3589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  2,000,000,000 iterations in %s yields 9 digits of π\n\n", elapsed))
}
if iterInt64 == 4000000000 { // last one
fyneFunc(fmt.Sprintf("\n... 4 Billion, gets us ten digits  ..."))
fyneFunc(fmt.Sprintf("   %0.5f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.141592653,589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  4,000,000,000 iterations in %s yields 10 digits of π\n\n", elapsed))
fyneFunc(fmt.Sprintf(" per option  --  the Gottfried Wilhelm Leibniz formula\n"))

LinesPerIter = 14
fyneFunc(fmt.Sprintf("at aprox %0.2f lines of code per iteration ...", LinesPerIter))
LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)

// store reults in a log file which can be displayed from within the program by selecting option #12
fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
Hostname, _ := os.Hostname()
_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz --  on %s \n", Hostname)
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
TotalRun := elapsed.String()                                            // cast time duration to a String type for Fprintf "formatted print"
_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
check(err7)
}
} // end of first for loop

fyneFunc(fmt.Sprintf( "\n\nWe continue the Gottfried Wilhelm Leibniz formula  :  π = 4 * ( 1 - 1/3 + 1/5 - 1/7 + 1/9 ... "))
fyneFunc(fmt.Sprintf("    π = 3 + 4/(2*3*4) - 4/(4*5*6) + 4/(6*7*8) - 4/(8*9*10) + 4/(10*11*12) ..."))

fyneFunc(fmt.Sprintf("   Infinitesimal calculus was developed independently in the late 17th century by Isaac Newton"))
fyneFunc(fmt.Sprintf("    ... and Gottfried Wilhelm Leibniz"))
fyneFunc(fmt.Sprintf("    9 billion iterations will be executed \n\n   ... working ...\n"))

start = time.Now()

/*        iterFloat64 = 0
          var denom float64
              denom = 3
          var sum float64
          iterInt64 = 1
*/
//      sum = 1-(1/denom)

for iterInt64 < 9000000000 {
iterFloat64++
iterInt64++
denom = denom + 2
if iterInt64%2 == 0 {
sum = sum + 1/denom
} else {
sum = sum - 1/denom
}
π = 4 * sum

if iterInt64 == 6000000000 {
fyneFunc(fmt.Sprintf("... 6 Billion completed ... \n"))
fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.141592653,589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  6,000,000,000 iterations in %s still yields 10 digits of π\n", elapsed))
fyneFunc(fmt.Sprintf( "  ... working ...\n"))
}
if iterInt64 == 8000000000 {
fyneFunc(fmt.Sprintf("... 8 Billion completed. still ten ...\n"))
fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.141592653,589793 is from the web"))
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("  8,000,000,000 iterations in %s still yields 10 digits of π\n", elapsed))
fyneFunc(fmt.Sprintf( "  ... working ...\n"))
}
if iterInt64 == 9000000000 {
fyneFunc(fmt.Sprintf("   %0.13f was calculated by the Gottfried Wilhelm Leibniz formula", π))
fyneFunc(fmt.Sprintf("    3.141592653,589793 is from the web"))
// fyneFunc(fmt.Sprintf("   ", iter)
t := time.Now()
elapsed := t.Sub(start)
fyneFunc(fmt.Sprintf("\n... 9B iterations in %s, but to get 10 digits we only needed 4B iterations\n\n", elapsed))
fyneFunc(fmt.Sprintf(" per  --  the Gottfried Wilhelm Leibniz formula\n"))

t = time.Now()
elapsed = t.Sub(start)

LinesPerIter = 14
fyneFunc(fmt.Sprintf("at aprox %0.2f lines of code per iteration ...", LinesPerIter))
LinesPerSecond = (LinesPerIter * iterFloat64) / elapsed.Seconds() // .Seconds() returns a float64
fmt.Printf("Aprox %.0f lines of code were executed per second \n", LinesPerSecond)

// store reults in a log file which can be displayed from within the program by selecting option #12
fileHandle, err1 := os.OpenFile("dataLog-From_calculate-pi-and-friends.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
check(err1)                                                                                                             // ... gets a file handle to dataLog-From_calculate-pi-and-friends.txt
defer fileHandle.Close()                                                                                                // It’s idiomatic to defer a Close immediately after opening a file.
Hostname, _ := os.Hostname()
_, err0 := fmt.Fprintf(fileHandle, "\n  -- Gottfried Wilhelm Leibniz (cont.) -- on %s \n", Hostname)
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
TotalRun := elapsed.String()                                            // cast time duration to a String type for Fprintf "formatted print"
_, err7 := fmt.Fprintf(fileHandle, "Total runTime was %s \n", TotalRun) // add total runtime of this calculation
check(err7)
}
} 
} // written entirely by Richard Woolley
