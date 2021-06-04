package array

import (
	"fmt"
	"sort"
	"testing"
)

/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]

提示：

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[i] <= 109
*/
func MergeArray(nums1 []int, m int, nums2 []int, n int) {
	l := m + n
	arr := make([]int, l)
	idx := 0
	for i, j := 0, 0; i < m || j < n; {
		if i == m {
			arr[idx] = nums2[j]
			idx++
			j++
			continue
		}
		if j == n {
			arr[idx] = nums1[i]
			idx++
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			arr[idx] = nums2[j]
			idx++
			j++
			continue
		}
		arr[idx] = nums1[i]
		idx++
		i++
	}

	for i := 0; i < l; i++ {
		nums1[i] = arr[i]
	}
}

/*
	666
*/
func MergeArrayT1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func TestMergeArray(t *testing.T) {
	nums1 := []int{1, 2, 0, 0}
	nums2 := []int{-2, 4}
	MergeArray(nums1, 2, nums2, 2)

	fmt.Println(nums1)
	//fmt.Println(nums[:duplicates])

}
