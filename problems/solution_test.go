package problems

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethod(t *testing.T) {
	isValid("()")
	assert.True(t, isPalindrome(121))
	recoverFromPreorder("1-2--3---4-5--6---7")
}

func TestMergeSortedList(t *testing.T) {
	a1 := ListNode{Val: 0}
	a2 := ListNode{Val: 2}
	b1 := ListNode{Val: 1}
	b2 := ListNode{Val: 4}
	c1 := ListNode{Val: 3}
	c2 := ListNode{Val: 5}

	a1.Next = &a2
	b1.Next = &b2
	c1.Next = &c2

	mergeKLists([]*ListNode{&a1, &b1, &c1})
}
