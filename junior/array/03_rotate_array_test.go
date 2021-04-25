package array

import (
	"fmt"
	"testing"
)

/*
给定一个数组，将数组中的元素向右移动 k个位置，其中 k 是非负数。

进阶：

	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
	输出: [5,6,7,1,2,3,4]
解释:
	向右旋转 1 步: [7,1,2,3,4,5,6]
	向右旋转 2 步: [6,7,1,2,3,4,5]
	向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
	输出：[3,99,-1,-100]
解释:
	向右旋转 1 步: [99,-1,-100,3]
	向右旋转 2 步: [3,99,-1,-100]

提示：

	1 <= nums.length <= 2 * 104
	-231 <= nums[i] <= 231 - 1
	0 <= k <= 105

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Rotate2N(nums []int, k int) {
	length := len(nums)
	rotateNum := k % length
	if rotateNum == 0 || length <= 1 {
		return
	}
	temp := make([]int, length, length)
	copy(temp, nums)
	for i := 0; i < length; i++ {
		nums[(i+k)%length] = temp[i]
	}
}

// Rotate 实现 O（1）空间  O(n) 时间
func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func TestRotate(t *testing.T) {
	prices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//prices := []int{1, 1, 2, 3, 2,1}

	//Rotate2N(prices, 4)
	Rotate(prices, 4)
	fmt.Println(prices)
}
