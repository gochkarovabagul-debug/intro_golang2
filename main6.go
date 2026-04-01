package main

import (
	"fmt"
	"math"
)

func pow(x, y, z float64) float64 {
	if v := math.Pow(x, y); v < z {
		return v
	}
	return z
}
func main() {
	var x float64
	var y float64
	var z float64
	fmt.Printf("enter your numbers: \n")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	fmt.Println(pow(x, y, z))
}
