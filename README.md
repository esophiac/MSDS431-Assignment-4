# Assignment 4 - Go for Statistics
The goal of this assignment is compare the performance of Python and R to Go (Golang) for use in statistics, in a hypothetical corporate setting. Using the Anscombe Quartet, this assignment tested how quickly each language performed a linear regression, and then returned a gradient and y-intercept for each set of points. Each language found the correct gradient(0.5) and the correct y-intercept (3). The time it took to the run the program for each language was the following:
- R: 0.02 sec 
- Python: 0.43 sec
- Go: 0.39 sec
The execution time for R and Python was determined with a timing packages in the main body of the program. The timing for the Go program was done with a benchmark in the main_test.go file.

This test found that Go could be used for statistics within the company, given some adaptation.

## Background
The purpose of the Anscombe Quartet is to demonstrate how it is important to visualize datasets instead of relying on pure data. Each of the datasets in the quartet returns identifical regression coefficients, despite being very different. 

For more information, see:
Anscombe, F. J. 1973, February. "Graphs in Statistical Analysis." *The American Statistician 27(1)*: 17–21. Available online at [https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf).

## Recommendation to Management
This test found that Go would be suitable for enterprise use, particularly given that the Go program produced identical results for the coefficients. While R performed faster than both Go and Python, we recommend Go additionally because of the potential to tailor functions to business need. This test used the [stats](https://pkg.go.dev/github.com/montanaflynn/stats) package. However, it was simple to create additional functions that extended the function of the selected library. Moving forward, the team also recommends that the statistics that we use be packaged into a library for use by the Go community.

## Roles of Programs and Data
All Go programs used in this test are found in the Go folder in this repository. The data used can be found in the .pdf linked above in this README, but the data was created with a suitable data structure in each format.
- Linear Regression with R.R: the file to test the Anscombe quartet with the R programming language. The standard library in R was used.
- Python File for Assignment 4: the file to test the Anscombe quartet with the Python programming language. The scipy.stats library was used for the linear regression, in addition to other standard data analysis packages.
- README.md: the readme file for the repository
- go.mod: defines the module's properties
- go.sum: record of the library the project depends on
- main_test.go: tests and benchmarks the fuctions in the main.go file
- main.go: the file to test the Anscombe quartet with the Go programming language. The stats library was used. 
- m.exe: the executable for this project. Assembled on Windows.

## Application
An executable for this project was created using Windows. To create your own executable, run **go build** in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

## Use of AI
AI was not used for this assignment.