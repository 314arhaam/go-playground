package main

import "fmt"
import "math"
import "os"
import "strconv"

type Root struct {
	r1 float64
	r2 float64
}

type Result struct {
    Value Root
    Err   error
}

func solve(a float64, b float64, c float64) (float64, float64) {
	delta := math.Pow(b, 2) - 4 * a * c
	return (-b + math.Sqrt(delta))/2/a, (-b - math.Sqrt(delta))/2/a
}

func main() {
	// fmt.Println(quad(1., 1., 1., 1.))
	a, _ := strconv.ParseFloat(os.Args[1], 64)
    b, _ := strconv.ParseFloat(os.Args[2], 64)
	c, _ := strconv.ParseFloat(os.Args[3], 64)
	fmt.Println(solve(a, b, c))
}