/*
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/

package main

import (
	"fmt"
)

func climbStairs(n int) int {
	x, y := 0, 1
    
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return y
}


func main() {
	fmt.Println(climbStairs(20))
}
