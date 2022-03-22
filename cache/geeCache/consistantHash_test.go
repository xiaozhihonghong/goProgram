package geeCache

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T)  {
	hash := NewMap(3, func(key []byte) uint32 {
		i, _ :=  strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("2", "4", "6")

	testCase := map[string]string {
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCase {
		if hash.Get(k) != v {
			fmt.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	hash.Add("8")
	testCase["27"] = "8"

	for k, v := range testCase {
		if hash.Get(k) != v {
			fmt.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
