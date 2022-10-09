package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	current := first

	var a, b, remainder, result int

	for {
		a = 0
		b = 0

		if l1 != nil {
			a = l1.Val
		}

		if l2 != nil {
			b = l2.Val
		}

		result = a + b + remainder
		current.Val = result % 10
		remainder = result / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break
		}

		current.Next = &ListNode{}
		current = current.Next
	}

	if remainder > 0 {
		current.Next = &ListNode{Val: remainder}
	}

	return first
}

// BELOW IS JUST FOR TESTING LOCALLY

func (self *ListNode) ToSlice() []int {
	slice := make([]int, 1)
	slice[0] = self.Val
	current := self.Next

	for {
		if current == nil {
			break
		}

		slice = append(slice, current.Val)
		current = current.Next
	}

	return slice
}

func sliceToListNode(ints []int) *ListNode {
	var previous, current *ListNode

	for i := len(ints) - 1; i >= 0; i-- {
		current = &ListNode{
			Val:  ints[i],
			Next: previous,
		}
		previous = current
	}

	return current
}

func execTestCaseAddTwoNumbers(t *testing.T, l1 []int, l2 []int, expectation []int) {
	sl1 := sliceToListNode(l1)
	sl2 := sliceToListNode(l2)
	exp := sliceToListNode(expectation)

	t.Run(fmt.Sprintf("%v + %v", l1, l2), func(t *testing.T) {
		if res := addTwoNumbers(sl1, sl2); !reflect.DeepEqual(res, exp) {
			t.Errorf("Expected %v but got %v", expectation, res.ToSlice())
		}
	})
}

func TestAddTwoNumbers(t *testing.T) {
	execTestCaseAddTwoNumbers(t, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8})
	execTestCaseAddTwoNumbers(t, []int{0}, []int{0}, []int{0})
	execTestCaseAddTwoNumbers(t, []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1})
}
