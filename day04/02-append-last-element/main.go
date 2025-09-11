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

	//2. append 1 element
	numbers = append(numbers, 15)
	showBackingArray("numbers", &numbers)

	//3. append 2 element
	numbers = append(numbers, 25, 30)
	showBackingArray("numbers", &numbers)

	//4. append 1 element
	numbers = append(numbers, 55)
	showBackingArray("numbers", &numbers)

	//5. remove last element, :6, numbers akan keep element
	// mulai dari index 0 s/d k 5
	numbers = numbers[:6]
	showBackingArray("numbers", &numbers)

	//6. remove last element
	// mulai dari index 0 s/d idx len(number)-1
	numbers = numbers[:len(numbers)-1]
	showBackingArray("numbers", &numbers)

}
