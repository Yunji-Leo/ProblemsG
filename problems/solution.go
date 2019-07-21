package problems

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"math"
	"sort"
	"strings"
)

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	check_list := make(map[int]int)
	for i, n := range nums {
		complement := target - nums[i]
		if val, ok := check_list[complement]; ok {
			ans[0] = val
			ans[1] = i
			break
		} else {
			check_list[n] = i
		}
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummuyNode := new(ListNode)
	previousNode := dummuyNode

	carry := 0

	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		} else {
			val1 = 0
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		} else {
			val2 = 0
		}

		sum := val1 + val2 + carry
		val := sum % 10
		carry = sum / 10

		newNode := ListNode{Val: val}
		previousNode.Next = &newNode
		previousNode = &newNode
	}

	if carry != 0 {
		newNode := ListNode{Val: carry}
		previousNode.Next = &newNode
	}

	return dummuyNode.Next
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	left, right := 0, 0
	set := make(map[rune]bool)

	for right < len(s) {
		if set[rune(s[right])] == true {
			delete(set, rune(s[left]))
			left++
		} else {
			set[rune(s[right])] = true
			right++
			if right-left > result {
				result = right - left
			}
		}
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if (m+n)%2 == 0 {
		return float64(getKth(nums1, nums2, (m+n)/2+1)+getKth(nums1, nums2, (m+n)/2)) * 0.5
	} else {
		return float64(getKth(nums1, nums2, (m+n)/2+1)) * 1.0
	}
}

func getKth(A []int, B []int, k int) int {
	m := len(A)
	n := len(B)

	if m > n {
		return getKth(B, A, k)
	}
	if m == 0 {
		return B[k-1]
	}
	if k == 1 {
		return min(A[0], B[0])
	}

	pa := min(k/2, m)
	pb := k - pa
	if A[pa-1] <= B[pb-1] {
		return getKth(A[pa:], B, pb)
	} else {
		return getKth(A, B[pb:], pa)
	}

}

func longestPalindrome(s string) string {
	max := -1
	result := ""
	for i := 0; i < len(s); i++ {
		length, value := expandCheck(s, i, i)
		if length > max {
			result = value
			max = length
		}
		length, value = expandCheck(s, i, i+1)
		if length > max {
			result = value
			max = length
		}
	}
	return result
}

func expandCheck(s string, i, j int) (length int, result string) {
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	i++
	j--
	if i <= j {
		return j - i + 1, s[i : j+1]
	}
	return 0, ""
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, 0)
	for i := 0; i < min(len(s), numRows); i++ {
		rows = append(rows, strings.Builder{})
	}

	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow].WriteString(string(c))
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}
	return result.String()
}

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	result := 0
	if str == "" {
		return result
	}
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			result = result*10 + int(str[i]-'0')
		} else {
			break
		}

		if result > math.MaxInt32 && sign == 1 {
			return math.MaxInt32
		} else if result > math.MaxInt32+1 && sign == -1 {
			result = math.MinInt32
			return math.MinInt32
		}
	}

	result = sign * result

	return result
}

func isPalindrome(x int) bool {
	origin := x
	if x < 0 {
		return false
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result == origin
}

func isMatch1(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	var first_match bool
	first_match = len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch1(s, p[2:]) || (first_match && isMatch1(s[1:], p))
	} else {
		return first_match && isMatch1(s[1:], p[1:])
	}
}

var memo [][]int

const (
	TRUE  = 1
	FALSE = 2
)

func isMatch2(s string, p string) bool {
	memo = make([][]int, len(s)+1)
	for i := range memo {
		memo[i] = make([]int, len(p)+1)
	}
	return isMatch2DP(0, 0, s, p)
}

func isMatch2DP(i, j int, text, pattern string) bool {
	if memo[i][j] != 0 {
		return memo[i][j] == TRUE
	}
	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		first_match := (i < len(text) && (pattern[j] == text[i] || pattern[j] == '.'))
		if j+1 < len(pattern) && pattern[j+1] == '*' {
			ans = (isMatch2DP(i, j+2, text, pattern) || first_match && isMatch2DP(i+1, j, text, pattern))
		} else {
			ans = first_match && isMatch2DP(i+1, j+1, text, pattern)
		}
	}
	if ans {
		memo[i][j] = TRUE
	} else {
		memo[i][j] = FALSE
	}
	return ans
}

func isMatch3(text string, pattern string) bool {
	dp := make([][]bool, len(text)+1)
	for i := range dp {
		dp[i] = make([]bool, len(pattern)+1)
	}
	dp[len(text)][len(pattern)] = true

	for i := len(text); i >= 0; i-- {
		for j := len(pattern) - 1; j >= 0; j-- {
			first_match := (i < len(text) &&
				(pattern[j] == text[i] || pattern[j] == '.'))
			if j+1 < len(pattern) && pattern[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || first_match && dp[i+1][j]
			} else {
				dp[i][j] = first_match && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	tmplist := make([]int, 0)
	backtrackSubsetsRef(&result, &tmplist, nums, 0)
	return result
}

func backtrackSubsetsRef(result *[][]int, tmplist *[]int, nums []int, start int) {
	clone := make([]int, len(*tmplist))
	copy(clone, *tmplist)
	*result = append(*result, clone)

	for i := start; i < len(nums); i++ {
		*tmplist = append(*tmplist, nums[i])
		backtrackSubsetsRef(result, tmplist, nums, i+1)
		*tmplist = (*tmplist)[:len(*tmplist)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	tmplist := make([]int, 0)
	sort.Ints(nums)
	backtrackSubsetsWithDupRef(&result, &tmplist, nums, 0)
	return result
}

func backtrackSubsetsWithDupRef(result *[][]int, tmplist *[]int, nums []int, start int) {
	tmplistClone := append([]int{}, *tmplist...)
	*result = append(*result, tmplistClone)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*tmplist = append(*tmplist, nums[i])
		backtrackSubsetsWithDupRef(result, tmplist, nums, i+1)
		*tmplist = (*tmplist)[:len(*tmplist)-1]
	}
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	if 2*k > len(prices) {
		result := 0
		for i := 0; i < len(prices)-1; i++ {
			if prices[i+1] > prices[i] {
				result += prices[i+1] - prices[i]
			}
		}
		return result
	}

	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][1] = math.MinInt32

	for i := 1; i < len(prices)+1; i++ {
		dp[i][0][1] = max(dp[i-1][0][1], -prices[i-1])
	}

	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < len(prices)+1; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i-1])
		}
	}

	return dp[len(prices)][k][0]
}

func recoverFromPreorder(S string) *TreeNode {
	nodeQueue := queue.New()
	depthQueue := queue.New()
	generateNodeQueues(S, nodeQueue, depthQueue)
	if nodeQueue.Len() == 0 {
		return nil
	}

	nodeStack := stack.New()
	depthStack := stack.New()

	root := nodeQueue.Dequeue().(*TreeNode)
	depth := depthQueue.Dequeue().(int)
	nodeStack.Push(root)
	depthStack.Push(depth)

	for nodeQueue.Len() > 0 {
		node := nodeQueue.Dequeue().(*TreeNode)
		depth = depthQueue.Dequeue().(int)

		var parent *TreeNode
		parentDepth := -1
		for true {
			parent = nodeStack.Pop().(*TreeNode)
			parentDepth = depthStack.Pop().(int)
			if parentDepth == depth-1 {
				break
			}
		}
		if parent.Left == nil {
			parent.Left = node
			nodeStack.Push(parent)
			depthStack.Push(parentDepth)
		} else {
			parent.Right = node
		}
		nodeStack.Push(node)
		depthStack.Push(depth)
	}
	return root

}

func generateNodeQueues(S string, nodeQueue, depthQueue *queue.Queue) {
	if len(S) == 0 {
		return
	}
	isDigit := false
	depth := 0
	value := 0
	for i := range S {
		if '-' != S[i] {
			if isDigit {
				value = value*10 + int(S[i]-'0')
			} else {
				isDigit = true
				value = int(S[i] - '0')
			}
		} else {
			if !isDigit {
				depth++
			} else {
				nodeQueue.Enqueue(&TreeNode{Val: value})
				depthQueue.Enqueue(depth)
				isDigit = false
				depth = 1
			}
		}
	}
	nodeQueue.Enqueue(&TreeNode{Val: value})
	depthQueue.Enqueue(depth)
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	for left < right {
		minH := -1
		width := right - left
		if height[left] <= height[right] {
			minH = height[left]
			left++
		} else {
			minH = height[right]
			right--
		}
		result = max(result, minH*width)
	}
	return result
}

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

func romanToInt(s string) int {
	roman := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	result := 0
	for i := 0; i < len(s)-1; i++ {
		if roman[rune(s[i])] < roman[rune(s[i+1])] {
			result -= roman[rune(s[i])]
		} else {
			result += roman[rune(s[i])]
		}
	}
	return result + roman[rune(s[len(s)-1])]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefixDC(strs, 0, len(strs)-1)
}

func longestCommonPrefixDC(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}

	mid := (l + r) / 2
	left := longestCommonPrefixDC(strs, l, mid)
	right := longestCommonPrefixDC(strs, mid+1, r)
	return commonPrefix(left, right)
}

func commonPrefix(left, right string) string {
	minLength := min(len(left), len(right))
	for i := 0; i < minLength; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLength]
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[left]+nums[right] > 0 {
					right--
				} else if nums[i]+nums[left]+nums[right] < 0 {
					left++
				} else {
					ans := []int{nums[i], nums[left], nums[right]}
					result = append(result, ans)
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return result
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	distance := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < distance {
				result = sum
				distance = abs(sum - target)
			}
			if sum == target {
				return sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	phone := map[uint8][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	letterCombinationsRecursive(&result, phone, digits, 0, "")
	return result
}

func letterCombinationsRecursive(result *[]string, phone map[uint8][]string, digits string, pos int, temp string) {
	if pos == len(digits) {
		*result = append(*result, temp)
		return
	}

	for _, s := range phone[digits[pos]] {
		temp += s
		letterCombinationsRecursive(result, phone, digits, pos+1, temp)
		temp = temp[:len(temp)-1]
	}
}
