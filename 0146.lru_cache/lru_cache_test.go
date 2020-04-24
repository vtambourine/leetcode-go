package lru_cache

import (
	"strconv"
	"strings"
	"testing"
)

type test struct {
	methods []string
	args    [][]int
	expect  string
}

func TestLRUCache(t *testing.T) {
	tests := []test{
		{
			[]string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			[][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			"null,null,null,1,null,-1,null,-1,3,4",
		},
		{
			[]string{"LRUCache", "put", "put", "get", "put", "put", "get"},
			[][]int{{2}, {2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}},
			"null,null,null,2,null,null,-1",
		},
	}

	for _, c := range tests {
		output := make([]string, len(c.methods))
		var cache LRUCache
		for i := range c.methods {
			switch c.methods[i] {
			case "LRUCache":
				cache = Constructor(c.args[i][0])
				output[i] = "null"
			case "put":
				cache.Put(c.args[i][0], c.args[i][1])
				output[i] = "null"
			case "get":
				output[i] = strconv.Itoa(cache.Get(c.args[i][0]))
			}
		}
		result := strings.Join(output, ",")
		if result != c.expect {
			t.Fatalf("LRUCache failed.\nInput\n%v\n%v\nExpected %v\nReceived %v", c.methods, c.args, c.expect, result)
		}
	}
}
