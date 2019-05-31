package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var input []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		t = strings.ToLower(t)
		input = append(input, t)
	}

	n := len(input)
	for i := 0; i < n; i++ {
		var result []string
		result = append(result, input[i])
		delMark := make(map[int]int)
		for j := 0; j < n; j++ {
			if isAnagram(input[i], input[j]) && i != j {
				result = append(result, input[j])
				delMark[j] = j // mark for delete
			}
		}

		sort.Strings(result)
		fmt.Println(strings.Join(result, ","))
		for k, _ := range delMark {
			if k == n {
				input = input[:n-1]
			} else {
				input = append(input[:k], input[k+1:]...)
			}
			n--
		}
	}
}

func isAnagram(s string, t string) bool {
	s = strings.Join(strings.Fields(s), "")
	t = strings.Join(strings.Fields(t), "")

	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, v := range s {
		if string(v) != " " {
			m[v]++
		}
	}

	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v]--
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
