package main

import "fmt"

func main() {
	//1. declare array int
	var arr [4]int
	arr[0] = 12
	arr[1] = 15
	arr[2] = 10
	arr[3] = 45

	//1.1 display array
	for idx, v := range arr {
		fmt.Println(idx, ":", v)
	}

	//2. declare array int with initial value
	numbers := [5]int{1, 2, 3, 4, 5}
	numbers[2] = 10

	for idx, v := range numbers {
		fmt.Println(idx, ":", v)
	}

	//3 declare array using 3 period
	anime := [...]string{"demon", "spirit away", "inu", "slayar"}
	for idx, v := range anime {
		fmt.Println(idx, ":", v)
	}
}
