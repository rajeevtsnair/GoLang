/*Given an integer x, return true if x is a palindrome , and false otherwise.*/

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	num := x
	rev_num := 0

	for num > 0 {
		digit := num % 10
		rev_num = rev_num*10 + digit
		num /= 10
	}

	if rev_num == x {
		return true
	} else {
		return false
	}
}

func main() {

	result := isPalindrome(-12321)
	fmt.Println(result)

}
