package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for r, x := range nums {
		if l, ok := cache[target-x]; ok {
			return []int{l, r}
		}
		cache[x] = r
	}
	return []int{}
}

// BELOW IS JUST FOR TESTING LOCALLY

func execTestCaseTwoSum(t *testing.T, nums []int, target int, expectation []int) {
	t.Run(fmt.Sprintf("%v %d", nums, target), func(t *testing.T) {
		if res := twoSum(nums, target); !reflect.DeepEqual(res, expectation) {
			t.Errorf("Expected %v but got %v", expectation, res)
		}
	})
}

func TestTwoSum(t *testing.T) {
	execTestCaseTwoSum(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
	execTestCaseTwoSum(t, []int{3, 2, 4}, 6, []int{1, 2})
	execTestCaseTwoSum(t, []int{3, 3}, 6, []int{0, 1})
}
