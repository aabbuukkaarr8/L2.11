package main

import (
	"fmt"

	"github.com/aabbuukkaarr8/anagrams"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	res := anagrams.FindAnagrams(words)
	for k, group := range res {
		fmt.Println(k, "->", group)
	}
}
