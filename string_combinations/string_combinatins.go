package stringcombinations

import (
	"math"
	"fmt"
)

func combinations(s string) [][]string {
	n := len(s)
	pSize := int(math.Pow(2, float64(n)))
	powerset := make([][]string, 0, pSize)

	for counter := 1; counter < pSize; counter++ {
		var subset []string
		for j := 0; j < n; j++ {
			if counter&(1<<uint(j)) > 0 {
				char := string(fmt.Sprintf("%c", s[j]))
				subset = append(subset, char)
			}
		}

		powerset = append(powerset, subset)
	}

	return powerset
}