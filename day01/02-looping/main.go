package main

import "fmt"

func main() {
	/* 	loopHorizontal(8)
	   	loopVertical(8) */

	fizzBuzz(15)
	sumFizz(15)

	// call countfizz
	count := countFizz(15)
	fmt.Println("count: ", count)

	fmt.Println("count fizz : ", countFizz(15))
}

func countFizz(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i%4 != 0 {
			count++
			// sum = sum +i
		}
	}
	return count
}

func sumFizz(n int) {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i%4 != 0 {
			sum += i
			// sum = sum +i
		}
	}
	fmt.Println(sum)
}

func fizzBuzz(n int) {
	for i := 0; i <= n; i++ {
		if i%3 == 0 && i%4 == 0 {
			fmt.Println(i, "=FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Println(i, "=Fizz ")
		} else if i%4 == 0 {
			fmt.Println(i, "=Buzz ")
		} else {
			fmt.Println(i)
		}
	}
}

func loopHorizontal(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
	}
}

func loopVertical(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
