package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

// calculate the mean of the variables in a stats.Coordinate object
// the first return is X and the second return is y
func varMeans(inputData []stats.Coordinate) (xMean float64, yMean float64) {

	// calculate the mean of the x values
	var xSum float64 = 0
	for i := 0; i < len(inputData); i++ {
		xSum = xSum + inputData[i].X
	}

	xMean = xSum / float64(len(inputData))

	// calculate the mean of the y values
	var ySum float64 = 0
	for i := 0; i < len(inputData); i++ {
		ySum = ySum + inputData[i].Y
	}

	yMean = ySum / float64(len(inputData))

	return xMean, yMean
}

// calculates the gradient of a stats.Coordinate object
func gradient(inputData []stats.Coordinate) (gradient float64) {
	// calculate gradient using Sxy / Sxx

	xMean, yMean := varMeans(inputData)
	var numerator float64
	var denominator float64

	for i := 0; i < len(inputData); i++ {
		numerator = numerator + ((inputData[i].X - xMean) * (inputData[i].Y - yMean))
		denominator = denominator + ((inputData[i].X - xMean) * (inputData[i].X - xMean))
	}

	gradient = numerator / denominator

	return gradient
}

// calculates the intersection of a stats.Coordinate object
func yIntercept(inputData []stats.Coordinate) (intercept float64) {

	lineGradient := gradient(inputData)

	xMean, yMean := varMeans(inputData)

	intercept = yMean - lineGradient*xMean

	return intercept
}

func main() {

	start := time.Now()

	// set the the points in the Anscombe Quartet
	num1 := []stats.Coordinate{{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33}, {14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68}}
	num2 := []stats.Coordinate{{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26}, {14, 8.1}, {6, 6.13}, {4, 3.1}, {12, 9.13}, {7, 7.26}, {5, 4.74}}
	num3 := []stats.Coordinate{{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81}, {14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73}}
	num4 := []stats.Coordinate{{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47}, {8, 7.04}, {8, 5.25}, {19, 12.5}, {8, 5.56}, {8, 7.91}, {8, 6.89}}

	//find the linear regression coefficients each dataset
	linreg1, _ := stats.LinearRegression(num1)
	fmt.Println(gradient(linreg1), yIntercept(linreg1))

	linreg2, _ := stats.LinearRegression(num2)
	fmt.Println(gradient(linreg2), yIntercept(linreg2))

	linreg3, _ := stats.LinearRegression(num3)
	fmt.Println(gradient(linreg3), yIntercept(linreg3))

	linreg4, _ := stats.LinearRegression(num4)
	fmt.Println(gradient(linreg4), yIntercept(linreg4))

	fmt.Printf("Took %v", time.Since(start))

}
