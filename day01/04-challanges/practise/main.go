package main

import "fmt"

func main() {
	/* 	boxHollow(5)
	   	fmt.Println("")
	   	triangle01(5)
	   	fmt.Println("")
	   	triangle02(5)
	   	fmt.Println("")
	   	triangle03(5)
	   	fmt.Println("")
	   	triangle04(5)
	   	fmt.Println("")
	   	diagonal(5) */

	elPattern(5)

	var count = countDigit(3452)
	fmt.Printf("count digit : %d", count)
	fmt.Printf("count digit2 : %d", countDigit(345678))
	fmt.Println("")

	fmt.Printf("is Prime : %t", isPrime(5))

	fmt.Println("")
	fmt.Printf("is Prime : %t", isPrime(9))

	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\nBOX PRIME")

	boxPrime(150)
}

func boxPrime(n int) {
	count := 1
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			if count%6 == 0 {
				fmt.Println("")
			} else {

				fmt.Printf("%d ", i)
			}
			count++

		}
	}
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countDigit(n int) int {
	count := 0
	sisa := n

	for {
		sisa = sisa / 10
		count++
		if sisa == 0 {
			break
		}
	}

	return count
}

func elPattern(n int) {
	for i := 1; i <= n; i++ { // baris
		for j := 1; j <= n; j++ { //kolom
			if j == 1 {
				fmt.Printf("%d ", i)
			} else if i == n {
				fmt.Printf("%d ", i+j-1)
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func diagonal(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func triangle04(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if j <= (n+1)-i {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}
		}
		fmt.Println("")
	}
}

func triangle03(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= n-i {
				fmt.Printf("%s ", " ")
			} else {
				fmt.Printf("%s ", "*")
			}
		}
		fmt.Println("")
	}
}

func triangle02(n int) {
	for i := 1; i <= n; i++ {
		for j := (n + 1) - i; j >= 1; j-- {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func triangle01(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%s ", "*")
		}
		fmt.Println("")
	}
}

func boxHollow(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")
			}

		}
		fmt.Println("")
	}
}
