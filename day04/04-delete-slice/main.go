package main

import "fmt"

func showBackingArray(varName string, slice *[]int) {
	fmt.Printf("[%s]\taddr[%p] len[%d] cap[%d]\n", varName, slice, len(*slice), cap(*slice))

	for i, s := range *slice {
		fmt.Printf("[%d] addr:[%p]	value:[%#v]\n", i, &(*slice)[i], s)
	}
}

func main() {
	//initial value
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	showBackingArray("numbers", &numbers)

	//delete index ke-0 or first elment
	//[1:10] => [low:high] or [start:end]
	// [low:high] high selalu lebih besar dari low
	//[1:10]=> muali dari index ke-0 s/d index-10, tapi 10 ga includ
	numbers = numbers[1:10]
	showBackingArray("numbers", &numbers)

	//deelete last element
	numbers = numbers[0:8]
	showBackingArray("numbers", &numbers)

	//delete di middle element
	numbers = numbers[2:6]
	showBackingArray("numbers", &numbers)

	// delete di middle but with capacity
	numbers = numbers[:3]
	showBackingArray("numbers", &numbers)

}
