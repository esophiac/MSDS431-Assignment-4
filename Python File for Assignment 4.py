## Python File for Assignment 4 - testing Linear Regression

## import statements
from scipy.stats import linregress
import numpy as np
import pandas as pd
import time

start_time = time.time()

## set up the demonstration data
## demonstration data from
## Anscombe, F. J. 1973, February. Graphs in statistical analysis. 
##  The American Statistician 27: 17–21.

## set up the data as a dataframe
anscombe = pd.DataFrame({'x1' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x2' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x3' : [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
    'x4' : [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
    'y1' : [8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68],
    'y2' : [9.14, 8.14,  8.74, 8.77, 9.26, 8.1, 6.13, 3.1,  9.13, 7.26, 4.74],
    'y3' : [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73],
    'y4' : [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]})

## run the linear regression tests
slope1, intercept1, r_value1, p_value1, std_err1 = linregress(anscombe['x1'], anscombe['y1'])
print(round(slope1, 2), round(intercept1, 2))

slope2, intercept2, r_value2, p_value2, std_err2 = linregress(anscombe['x2'], anscombe['y2'])
print(round(slope2, 2), round(intercept2, 2))

slope3, intercept3, r_value3, p_value3, std_err3 = linregress(anscombe['x3'], anscombe['y3'])
print(round(slope3, 2), round(intercept3, 2))

slope4, intercept4, r_value4, p_value4, std_err4 = linregress(anscombe['x4'], anscombe['y4'])
print(round(slope4, 2), round(intercept4, 2))

print("%.2f seconds" % (time.time() - start_time))


