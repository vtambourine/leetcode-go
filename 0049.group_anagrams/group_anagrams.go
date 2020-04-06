package group_anagrams

import (
	"sort"
)

type chars []rune

func (c chars) Len() int {
	return len(c)
}

func (c chars) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c chars) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func getHash(s string) string {
	c := chars(s)
	sort.Sort(c)
	return string(c)
}

func groupAnagrams(strs []string) [][]string {
	index := make(map[string][]string)

	for _, s := range strs {
		hash := getHash(s)
		index[hash] = append(index[hash], s)
	}

	result := make([][]string, 0)
	for _, v := range index {
		result = append(result, v)
	}

	return result
}
