package main

import "fmt"

func main() {
	var x[5] int //Define array of size 5
	x[0] = 3 //Initialize index 0
	fmt.Println(x[1])

	var y[6]int = [6]int{1, 2, 3, 4, 5, 6} //Initialize array of size 6
	//y := [6]int{1, 2, 3, 4, 5, 6} //Alternate way
	fmt.Println(y[2])

    z := [...]int{1, 2, 3, 4, 5, 6,7} //Initialize array of arbitrary length

	for i, v := range z {
		fmt.Printf("Array z Index: %d Value: %d\n", i,v)
	}

	//Slices example. Uses the array z defined above.
	fmt.Println("*********Slices Example****************")

	s1 := z[1:4]
	s2 := z[3:5]
	
	for i, v := range s1 {
		fmt.Printf("Slice s1 Index: %d Value: %d\n", i,v)
	}

	for i, v := range s2 {
		fmt.Printf("Slice s2 Index: %d Value: %d\n", i,v)
	}

	fmt.Println(s1[2])
	fmt.Println(s2[0])

	//Slice literal
	fmt.Println("**********Slice Literal******************")

	sli_1 := []int{1,2,3}	

	fmt.Println(len(sli_1))
	fmt.Println(cap(sli_1))

	sli_2 := make([]int, 0, 3)
	sli_2 = append(sli_2, 10)
	fmt.Println(sli_2[0])

	sli_2 = append(sli_2, 15)
	fmt.Println(sli_2[1])

}
