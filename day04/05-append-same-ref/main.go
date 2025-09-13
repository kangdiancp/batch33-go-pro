package main

import "fmt"

func showBackingArray(varName string, slice *[]string) {
	fmt.Printf("[%s]\taddr[%p] len[%d] cap[%d]\n", varName, slice, len(*slice), cap(*slice))

	for i, s := range *slice {
		fmt.Printf("[%d] addr:[%p]	value:[%#v]\n", i, &(*slice)[i], s)
	}
}

func main() {
	//init word slice
	words := []string{"A", "B", "C", "D", "E", "F"}
	showBackingArray("words", &words)

	slice1 := append(words[2:4], "CODING")
	showBackingArray("slice1", &slice1)

	slice1 = append(slice1, "GOLANG")
	showBackingArray("slice1", &slice1)
	//showBackingArray("words", &words)

	slice1 = append(slice1, "JAVA")
	showBackingArray("slice1", &slice1)
	showBackingArray("words", &words)

}
