package group_anagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

type test struct {
	input  []string
	expect [][]string
}

func copyAnswer(answer [][]string) [][]string {
	result := make([][]string, 0)
	for _, str := range answer {
		newStr := make([]string, len(str))
		copy(newStr, str)
		result = append(result, newStr)
	}
	return result
}

func sortAnswer(answer [][]string) {
	for _, s := range answer {
		sort.Strings(s)
	}

	sort.Slice(answer, func(i, j int) bool {
		return strings.Join(answer[i], "") < strings.Join(answer[j], "")
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []test{
		{[]string{"boo", "bob"},
			[][]string{
				{"bob"},
				{"boo"},
			},
		},
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"ate", "eat", "tea"},
				{"bat"},
				{"nat", "tan"},
			},
		},
		{[]string{"hos", "boo", "nay", "deb", "wow", "bop", "bob", "brr", "hey", "rye", "eve", "elf", "pup", "bum", "iva", "lyx", "yap", "ugh", "hem", "rod", "aha", "nam", "gap", "yea", "doc", "pen", "job", "dis", "max", "oho", "jed", "lye", "ram", "pup", "qua", "ugh", "mir", "nap", "deb", "hog", "let", "gym", "bye", "lon", "aft", "eel", "sol", "jab"},
			[][]string{
				{"sol"}, {"wow"}, {"gap"}, {"hem"}, {"yap"}, {"bum"}, {"ugh", "ugh"}, {"aha"}, {"jab"}, {"eve"}, {"bop"}, {"lyx"}, {"jed"}, {"iva"}, {"rod"}, {"boo"}, {"brr"}, {"hog"}, {"nay"}, {"mir"}, {"deb", "deb"}, {"aft"}, {"dis"}, {"yea"}, {"hos"}, {"rye"}, {"hey"}, {"doc"}, {"bob"}, {"eel"}, {"pen"}, {"job"}, {"max"}, {"oho"}, {"lye"}, {"ram"}, {"nap"}, {"elf"}, {"qua"}, {"pup", "pup"}, {"let"}, {"gym"}, {"nam"}, {"bye"}, {"lon"},
			},
		},
	}

	for _, c := range tests {
		result := copyAnswer(groupAnagrams(c.input))
		sortAnswer(result)
		sortAnswer(c.expect)
		if !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("groupAnagrams(%v) fails.\nExpected %v\nReceived %v", c.input, c.expect, result)
		}
	}
}
