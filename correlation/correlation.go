package correlation

import (
	"pearson_correlation/correlation/statistics"
)

func CalculateCorrelation(x []float64, y []float64) float64 {
	xStdDev := statistics.StandardDeviation(x)
	yStdDev := statistics.StandardDeviation(y)

	// Calculate the covariance of x and y
	covariance := statistics.Covariance(x, y)

	// Calculate the correlation of x and y
	correlation := covariance / (xStdDev * yStdDev)
	return correlation
}
