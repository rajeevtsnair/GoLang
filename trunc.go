package main

import "fmt"

func main() {

	var floatNum float64
	var IntNum int

	fmt.Println("Enter a floating point number..")
	fmt.Scan(&floatNum)

	IntNum = int(floatNum)

	// Printing the truncated number
	fmt.Println("Truncated number : ", IntNum)
}
