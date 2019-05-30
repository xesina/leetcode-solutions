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

// To examine if tt is a rearrangement of ss, we can count
// occurrences of each letter in the two strings and compare them.
func isAnagram2(s string, t string) bool {
	s = strings.Join(strings.Fields(s), "")
	t = strings.Join(strings.Fields(t), "")

	if len(s) != len(t) {
		return false
	}

	sm := make(map[rune]int)
	for _, v := range s {
		sm[v]++
	}

	for _, v := range t {
		if _, ok := sm[v]; ok {
			sm[v]--
		}
	}

	for _, v := range sm {
		if v != 0 {
			return false
		}
	}

	return true
}
