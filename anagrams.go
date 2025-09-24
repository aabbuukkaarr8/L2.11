package anagrams

import (
	"sort"
	"strings"
)

func FindAnagrams(words []string) map[string][]string {

	anagramGroups := make(map[string][]string)

	for _, w := range words {
		word := strings.ToLower(w)
		letters := []rune(word)
		sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
		key := string(letters)
		anagramGroups[key] = append(anagramGroups[key], word)
	}

	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) < 2 {
			continue
		}
		sort.Strings(group)
		first := group[0]
		result[first] = group
	}

	return result
}

//принцип - сортируем руны в слове - кладем как ключ в мапу
