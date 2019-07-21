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
