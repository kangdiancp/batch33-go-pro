package main

import (
	"fmt"
	"sort"
)

/*
[1, 2, 3, 4, 5] => [2 3 4 5 1]
*/
func rotateArrayLeft(arr []int) []int {
	arr = append(arr[1:], arr[:1]...)
	return arr
}

/*
[10, 20, 30, 40, 50] =>[50 10 20 30 40]
*/
func rotateArrayRight(arr []int) []int {
	arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
	return arr
}

func displayArray(arr []int) {
	fmt.Printf("%v\n", arr)
}

/*
[2, 7, 4, 10, 6, 9, 3, 5, 1, 8] => [2 4 6 8 10 1 3 5 7 9]
*/

func evenBeforeOdd(nums []int) []int {
	var (
		even []int
		odd  []int
	)

	sort.Ints(nums)

	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	nums = append(even, odd...)
	return nums

}

// 10, 20, 30, 40, 50 => 40
func getSecondLargest(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-2]
}

/*
[2, 3, 3, 2] => true
[1, 2, 3, 4, 5] => false
*/
func isPalindrome(arr []int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}

	return true
}

/*
addoneplus : last element +1
[1 3 4 5] => [1 3 4 6]
[1 3 4 8] => [1 3 4 9]
[1 3 4 9] => [1 3 5 0]
[1 9 9 9] => [2 0 0 0]
[9 9 9]   => [1 0 0 0]
*/

func addOnePlus(nums []int) []int {
	overflow := true
	pos := len(nums) - 1
	for pos >= 0 && overflow {
		currentValue := nums[pos]
		if overflow {
			currentValue += 1
		}

		nums[pos] = currentValue % 10
		overflow = currentValue >= 10
		pos--
	}
	if overflow {
		nums = append([]int{1}, nums...)
	}

	return nums
}

func main() {
	/* numsLeft := []int{1, 2, 3, 4, 5}
	numsRight := []int{10, 20, 30, 40, 50}
	displayArray(rotateArrayLeft(numsLeft))
	displayArray(rotateArrayRight(numsRight))

	numsEvenOdd := []int{2, 7, 4, 10, 6, 9, 3, 5, 1, 8}
	displayArray(evenBeforeOdd(numsEvenOdd))

	numLargest := []int{12, 234, 323, 765, 23, 657, 21}
	fmt.Printf("2nd max : %d\n", getSecondLargest(numLargest))

	fmt.Printf("isPalindorme:%t\n", isPalindrome([]int{2, 3, 3, 2}))
	fmt.Printf("isPalindorme:%t\n", isPalindrome([]int{2, 1, 3, 2})) */

	displayArray(addOnePlus([]int{1, 3, 4, 5}))
	displayArray(addOnePlus([]int{1, 3, 4, 9}))
	displayArray(addOnePlus([]int{9, 9}))
	//displayArray(addOnePlus([]int{9, 9, 9, 9}))

}
