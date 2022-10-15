package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func wordBreak(s string, wordDict []string) bool {
	paths := make([]string, 1)
	paths[0] = s
	visited := make(map[string]bool)

	var word string

	for {
		if len(paths) == 0 {
			break
		}

		word = paths[len(paths)-1]
		paths = paths[:len(paths)-1]

		if _, alreadyChecked := visited[word]; alreadyChecked {
			continue
		} else {
			visited[word] = true
		}

		if word == "" {
			return true
		}

		for _, dictWord := range wordDict {
			if strings.HasPrefix(word, dictWord) {
				paths = append(paths, word[len(dictWord):])
			}
		}
	}

	return false
}

// BELOW IS JUST FOR TESTING LOCALLY

func execTestWordBreak(t *testing.T, s string, dict []string, expectation bool) {
	t.Run(fmt.Sprintf("%s %v is %t", s, dict, expectation), func(t *testing.T) {
		if res := wordBreak(s, dict); res != expectation {
			t.Errorf("Expected %t but got %t", expectation, res)
		}
	})
}

func TestWordBreak(t *testing.T) {
	execTestWordBreak(t, "leetcode", []string{"leet", "code"}, true)
	execTestWordBreak(t, "applepenapple", []string{"apple", "pen"}, true)
	execTestWordBreak(t, "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false)
}
