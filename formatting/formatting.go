package main

import (
	"fmt"
	"math"
)

// Definition of constant with specific type
const two float64 = 2

// Definition of constant without type
const five = 5

func main() {
	result := math.Exp2(two)
	println(result)
	fmt.Printf("%f\n", result)
	fmt.Printf("%.2f\n", result)

}
