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
	numbers := make([]int, 3, 6)
	showBackingArray("numbers", &numbers)

	//2. append 3 element in first element
	numbers = append(numbers[:0], 15, 35, 40)
	showBackingArray("numbers", &numbers)

	//3. append 3 element in first element
	numbers = append(numbers, 50, 55, 60)
	showBackingArray("numbers", &numbers)

	//4. append 1 element in first element
	numbers = append(numbers, 65)
	showBackingArray("numbers", &numbers)

}
