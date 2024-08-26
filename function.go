package main

import (
	"fmt"
)

func increment(x *int) {
	*x = *x + 1
}

func increment_array(y *[3]int) {
	(*y)[0] = (*y)[0] + 1
}

func main() {
	var x int = 25
	increment(&x)
	fmt.Println(x)

	a := [3]int{1, 2, 3}
	increment_array(&a)
	fmt.Println(a[0])

}
