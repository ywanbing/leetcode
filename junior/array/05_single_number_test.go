package array

import (
	"fmt"
	"testing"
)

/*
只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例2:

输入: [4,1,2,1,2]
输出: 4

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func SingleNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			delete(numMap, v)
		} else {
			numMap[v] = true
		}
	}
	for v := range numMap {
		return v
	}
	return 0
}

func SingleNumberN(nums []int) int {

	num := 0
	for _, v := range nums {
		num ^= v
	}

	return num
}

func TestSingleNumber(t *testing.T) {
	//nums := []int{0, 1, 1}
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	has := SingleNumber(nums)
	fmt.Println(has)
	has = SingleNumberN(nums)
	fmt.Println(has)
}
