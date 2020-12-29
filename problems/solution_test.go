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

func TestCreateTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	fmt.Println(result)
}

func TestMoveCount(t *testing.T) {
	//movingCount(7, 2, 3)
	verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10})
}

func TestCopyLink(t *testing.T) {
	node0 := Node{
		Val: 0,
	}
	node1 := Node{
		Val: 1,
	}
	node2 := Node{
		Val: 2,
	}
	node0.Next = &node1
	node1.Next = &node2
	node0.Random = &node2
	node1.Random = &node0
	node2.Random = &node1
	copyRandomList(&node0)
}

/*
func TestFoo(t *testing.T) {
	lengthOfLongestSubstring2("abcabcbb")

	mf := Constructor()
	mf.AddNum(-1)
	mf.AddNum(-2)
	mf.AddNum(-3)

	reversePairs([]int{7, 5, 6, 4})
	search([]int{5, 7, 7, 8, 8, 10}, 6)
}

*/

func TestOddEvenLink(t *testing.T) {
	node0 := ListNode{
		Val: 0,
	}
	node1 := ListNode{
		Val: 1,
	}
	node2 := ListNode{
		Val: 2,
	}
	node3 := ListNode{
		Val: 3,
	}
	node4 := ListNode{
		Val: 4,
	}
	node0.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	result := oddEvenList(&node0)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
