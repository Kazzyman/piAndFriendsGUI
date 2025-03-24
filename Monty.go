package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"
)

func Monty(fyneFunc func(string), gridSize int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the grid size to be used (10000 will give 4 digits in about 21s) : ")
	input, _ := reader.ReadString('\n')
	gridSize, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid input, please enter an integer.")
		return
	}
	fmt.Printf("Size of the grid has been set to: %d\n", gridSize)
	if gridSize > 3000 {
		fmt.Println("working ...")
	}
	// ::: calculate pi
	piApprox := GridPi(gridSize) // ::: GridPi
	fmt.Printf("Approximated Pi as big float: %s\n", piApprox.Text('f', 30))
	piApproxFloat64, _ := piApprox.Float64() // Convert piApprox to a float64 type
	fmt.Printf("Approximated Pi as float64:   %f\n", piApproxFloat64)
	piFromMathLib := math.Pi                       // Obtain Pi from math library
	piFromMathLibBF := big.NewFloat(piFromMathLib) // Create a big float object version of Pi from math library
	fmt.Printf("Pi from Math Library:         %s\n", piFromMathLibBF.Text('f', 30))
	fmt.Printf("Difference: %f\n", math.Abs(piApproxFloat64-math.Pi))
}

/*
.
.
.
*/
func GridPi(gridSize int) *big.Float {
	start := time.Now()
	insideCircle := big.NewInt(0)
	totalPoints := big.NewInt(int64(gridSize * gridSize))
	increment := big.NewFloat(1.0 / float64(gridSize))

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			x := new(big.Float).Mul(increment, big.NewFloat(float64(i)))
			y := new(big.Float).Mul(increment, big.NewFloat(float64(j)))

			// x*x + y*y
			xSquared := new(big.Float).Mul(x, x)
			ySquared := new(big.Float).Mul(y, y)
			sum := new(big.Float).Add(xSquared, ySquared)

			// Check if x^2 + y^2 <= 1
			if sum.Cmp(big.NewFloat(1.0)) <= 0 {
				insideCircle.Add(insideCircle, big.NewInt(1))
			}
			iterationsForMonte16j = j
		}
		iterationsForMonte16i = i
	}
	iterationsForMonteTotal = iterationsForMonte16j * iterationsForMonte16i
	// Calculate Pi approximation
	four := big.NewFloat(4.0)
	insideCircleF := new(big.Float).SetInt(insideCircle)
	totalPointsF := new(big.Float).SetInt(totalPoints)
	piApprox := new(big.Float).Quo(insideCircleF, totalPointsF)
	piApprox.Mul(piApprox, four)

	t := time.Now()
	elapsed := t.Sub(start)
	TotalRun := elapsed.String() // cast time durations to a String type for Fprintf "formatted print"
	// print results: (piApprox, 0, "MonteCarlo", iterationsForMonteTotal, TotalRun, selection)

	fmt.Printf("Total iterations: %d, Elapsed time: %s \n\n", iterationsForMonteTotal, TotalRun)
	return piApprox
}
