package anagram

// https://leetcode.com/problems/valid-anagram

import (
	"strings"

	"github.com/xesina/leetcode-solutions/utils"
)

func isAnagram(s string, t string) bool {
	s = strings.Join(strings.Fields(s), "")
	t = strings.Join(strings.Fields(t), "")

	if len(s) != len(t) {
		return false
	}

	s = utils.SortString(s)
	t = utils.SortString(t)

	if s != t {
		return false
	}

	return true
}
