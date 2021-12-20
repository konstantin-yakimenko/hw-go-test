package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Word struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	mapCounter := make(map[string]int)
	for _, v := range strings.Fields(s) {
		mapCounter[v]++
	}
	structList := make([]Word, 0)
	for key, value := range mapCounter {
		structList = append(structList, Word{
			Word:  key,
			Count: value,
		})
	}

	sort.SliceStable(structList, func(i, j int) bool {
		if structList[i].Count != structList[j].Count {
			return structList[i].Count > structList[j].Count
		}
		return structList[i].Word < structList[j].Word
	})

	var result []string
	for i := 0; i < 10 && i < len(structList); i++ {
		result = append(result, structList[i].Word)
	}
	return result
}
