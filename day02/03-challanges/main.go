package main

import "fmt"

func main() {
	arr := [5]int{5, 3, 2, 1, 8}
	arr2 := rotateArray(arr)
	displayArray(arr2)

}

func displayArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
	}
}

func rotateArray(arr [5]int) [5]int {
	//1. isi variable temp dengan nilai arr[0]=5
	temp := arr[0]

	/* 	arr[0] = arr[1]
	   	arr[1] = arr[2]
	   	arr[2] = arr[3]
	   	arr[3] = arr[4] */

	//2. looping, geser ke kiri
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}

	//3. isi array index terakhir dengna temp
	arr[len(arr)-1] = temp

	return arr
}
