package main

import "fmt"
import "math"
import "errors"

func quad(a float64, b float64, c float64, x float64) float64 {
	return a * math.Pow(x, 2) + b * x + c
}

func solve(a float64, b float64, c float64) (float64, float64, error) {
	delta := math.Pow(b, 2) - 4 * a * c
	if delta < 0 {
		return 0., 0., errors.New("Negative delta")
	} else {
		return (-b + math.Sqrt(delta))/2/a, (-b - math.Sqrt(delta))/2/a, nil
	}
}

func main() {
	// fmt.Println(quad(1., 1., 1., 1.))
	fmt.Println(solve(1., 4., +16.))
}