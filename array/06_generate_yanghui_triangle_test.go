package array

import (
	"fmt"
	"testing"
)

/*
118. 杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
func GenerateYanghuiTriangle(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		length := i + 1
		arr := make([]int, length)
		// 提前判断第一个和最后一个
		arr[0] = 1
		arr[length-1] = 1

		for j := 1; j < length-1; j++ {
			// 第2个数是前面的 第一个加第二个数
			arr[j] = res[i-1][j] + res[i-1][j-1]
		}
		res[i] = arr
	}
	return res
}

func TestGenerateYanghuiTriangle(t *testing.T) {
	triangle := GenerateYanghuiTriangle(10)

	for _, ints := range triangle {
		fmt.Println(ints)
	}
}
