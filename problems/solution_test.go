package problems

import (
	"fmt"
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

func TestReverseKGroup(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	head := reverseKGroup(&n1, 2)
	fmt.Println(head.Val)
}

func TestFindSubString(t *testing.T) {
	findSubstring("barfoothefoobarfoo", []string{"foo", "bar"})
}
