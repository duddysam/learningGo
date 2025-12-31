package main

import "fmt"

func main() {
	var i int = 20

	var f = float64(i)

	fmt.Printf("printing var i int = %d\n", i)
	fmt.Printf("printing var f = %f\n", f)

	const value = 10

	var y int = value
	var z float64 = value

	fmt.Printf("checking that var y = %d and var z = %f are valid", y, z)
}
