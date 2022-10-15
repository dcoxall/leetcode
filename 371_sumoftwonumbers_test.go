package leetcode

import (
	"fmt"
	"testing"
)

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	return getSum(a^b, (a&b)<<1)
}

// BELOW IS JUST FOR TESTING LOCALLY

func execTestGetSum(t *testing.T, a int, b int, expectation int) {
	t.Run(fmt.Sprintf("%d + %d = %d", a, b, expectation), func(t *testing.T) {
		if res := getSum(a, b); res != expectation {
			t.Errorf("Expected %d but got %d", expectation, res)
		}
	})
}

func TestGetSum(t *testing.T) {
	execTestGetSum(t, 1, 2, 3)
	execTestGetSum(t, 2, 3, 5)
}
