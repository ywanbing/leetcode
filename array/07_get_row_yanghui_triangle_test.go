package array

import (
	"fmt"
	"runtime"
	"testing"
)

/*
119. 杨辉三角 II
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？
*/
func GetRowGenerateYanghuiTriangle(rowIndex int) []int {
	m := make(map[int]int)
	arr := make([]int, rowIndex+1)
	for i := 0; i < rowIndex/2+1; i++ {
		arr[i] = GetNum(m, rowIndex, i)
		arr[rowIndex-i] = arr[i]
	}
	m = nil
	runtime.GC()
	return arr
}

// 采用空间换时间，并使用递归
func GetNum(m map[int]int, i, j int) int {
	if j == 0 || j == i {
		return 1
	}
	num1 := 0
	if v, ok := m[(i-1)*1000+j-1]; ok {
		num1 = v
	} else {
		num1 = GetNum(m, i-1, j-1)
		m[(i-1)*1000+j-1] = num1
	}

	num2 := 0
	if v, ok := m[(i-1)*1000+j]; ok {
		num2 = v
	} else {
		num2 = GetNum(m, i-1, j)
		m[(i-1)*1000+j] = num2
	}

	return num1 + num2
}

func TestGetRowGenerateYanghuiTriangle(t *testing.T) {
	triangle := GetRowGenerateYanghuiTriangle(4)
	fmt.Println(triangle)
}
