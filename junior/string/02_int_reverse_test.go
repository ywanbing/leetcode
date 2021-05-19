package string

import (
	"fmt"
	"math"
	"testing"
)

/*
整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func IntReverse(x int) int {
	t := false
	if x < 0 {
		x = -x
		t = true
	}
	a := 0
	for ; x > 0; x = x / 10 {
		a = a*10 + x%10
	}
	if t {
		a = -a
	}
	if a > math.MaxInt32 || a < -math.MaxInt32-1 {
		return 0
	}
	return a
}

func TestIntReverse(t *testing.T) {
	num := 1534236469
	reverse := IntReverse(num)
	fmt.Println(reverse)
}
