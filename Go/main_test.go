package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// creating a Coordinate object to run all of the tests with
var testData = []stats.Coordinate{{1, 2}, {2, 3}, {3, 4}, {4, 5}}

// test the varMeans function
func TestVarMeans(t *testing.T) {

	out1, out2 := 2.5, 3.5

	result1, result2 := varMeans(testData)

	if out1 != result1 && out2 != result2 {
		t.Errorf("Both values were incorrect. Expected %f, got %f instead, and expected %f, got %f instead", out1, result1, out2, result2)
	} else if out1 != result1 {
		t.Errorf("X mean was incorrect. Expected %f, got %f instead", out1, result1)
	} else if out2 != result2 {
		t.Errorf("Y mean was incorrect. Expected %f, got %f instead", out2, result2)
	}

}

// test the gradient calculation function
func TestGradient(t *testing.T) {
	out := 1.0

	result := gradient(testData)

	if out != result {
		t.Errorf("Gradient was incorrect. Expected %f, got %f instead", out, result)
	}
}

// test the y-intercept calculation function
func TestYIntercept(t *testing.T) {
	out := 1.0

	result := yIntercept(testData)

	if out != result {
		t.Errorf("Gradient was incorrect. Expected %f, got %f instead", out, result)
	}
}

func BenchmarkFinalAnswer(b *testing.B) {
	// set the the points in the Anscombe Quartet
	num1 := []stats.Coordinate{{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33}, {14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68}}
	num2 := []stats.Coordinate{{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26}, {14, 8.1}, {6, 6.13}, {4, 3.1}, {12, 9.13}, {7, 7.26}, {5, 4.74}}
	num3 := []stats.Coordinate{{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81}, {14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73}}
	num4 := []stats.Coordinate{{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47}, {8, 7.04}, {8, 5.25}, {19, 12.5}, {8, 5.56}, {8, 7.91}, {8, 6.89}}

	//find the linear regression coefficients each dataset
	stats.LinearRegression(num1)
	// linreg1, _ := stats.LinearRegression(num1)
	//fmt.Println(gradient(linreg1), yIntercept(linreg1))

	stats.LinearRegression(num2)
	//linreg2, _ := stats.LinearRegression(num2)
	//fmt.Println(gradient(linreg2), yIntercept(linreg2))

	stats.LinearRegression(num3)
	//linreg3, _ := stats.LinearRegression(num3)
	//fmt.Println(gradient(linreg3), yIntercept(linreg3))

	stats.LinearRegression(num4)
	//linreg4, _ := stats.LinearRegression(num4)
	//fmt.Println(gradient(linreg4), yIntercept(linreg4))
}
