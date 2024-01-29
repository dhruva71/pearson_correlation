package statistics

import "math"

func Mean(x []float64) float64 {
	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}

func StandardDeviation(x []float64) float64 {
	var sum float64
	mean := Mean(x)
	for _, v := range x {
		sum += (v - mean) * (v - mean)
	}
	variance := sum / float64(len(x))
	std_dev := math.Sqrt(variance)
	return std_dev
}

// Covariance is a measure of the relationship between two random variables and to what extent, they change together.
func Covariance(x []float64, y []float64) float64 {
	var sum float64
	xMean := Mean(x)
	yMean := Mean(y)

	for i := 0; i < len(x); i++ {
		sum += (x[i] - xMean) * (y[i] - yMean)
	}

	return sum / float64(len(x))
}
