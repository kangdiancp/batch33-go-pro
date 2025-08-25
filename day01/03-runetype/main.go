package main

import "fmt"

func main() {
	vocal, consonant := countVowelConsonant("hello codeid")
	fmt.Println("vocal : ", vocal)
	fmt.Println("consonant : ", consonant)

	vocal, consonant = countVowelConsonantRange("hello codeid bootcamp")
	fmt.Println("vocal : ", vocal)
	fmt.Println("consonant : ", consonant)

	loopBox(5)

	loopBoxHollow(9)

}

// input param word : hello codeid
// string kumpulan dari rune/char
// count vokal : 5
// count consonant : 6
func countVowelConsonant(word string) (int, int) {
	vocal := 0
	consonant := 0

	for i := 0; i < len(word); i++ {
		if word[i] == 'a' || word[i] == 'i' || word[i] == 'u' || word[i] == 'e' || word[i] == 'o' {
			vocal++
		} else if word[i] != ' ' {
			consonant++
		}
	}

	return vocal, consonant
}

func countVowelConsonantRange(word string) (int, int) {
	vocal := 0
	consonant := 0

	for _, v := range word {
		if v == 'a' || v == 'i' || v == 'u' || v == 'e' || v == 'o' {
			vocal++
		} else if v != ' ' {
			consonant++
		}
	}

	return vocal, consonant
}

func loopBox(n int) {
	// outer loop pindah baris only
	for i := 0; i < n; i++ {
		//print *
		for i := 0; i < n; i++ {
			fmt.Print("*", " ")
		}
		fmt.Println("")
	}
}

func loopBoxHollow(n int) {

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || row == n-1 || col == 0 || col == n-1 {
				fmt.Printf("%s", "*")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println("")
	}
}
