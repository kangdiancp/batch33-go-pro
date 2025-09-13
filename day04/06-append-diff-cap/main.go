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

	slice1 := append(words[2:4:4], "CODING", "GO", "JAVA", "C#")
	showBackingArray("slice1", &slice1)

	slice2 := make([]string, len(slice1)-1)
	copy(slice2, slice1)
	slice2 = append(slice2, "PHP", "PYTHON")
	showBackingArray("slice2", &slice2)

	slice3 := append(slice1, slice2[5:]...)
	showBackingArray("slice3", &slice3)

}
