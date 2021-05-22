package string

import (
	"fmt"
	"strings"
	"testing"
)

/*
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
相关标签
双指针 字符串

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	toLower := strings.ToLower(s)
	arr := []byte(toLower)
	l := len(arr)
	for i, j := 0, l-1; i < l; {
		if i >= j {
			break
		}
		if !((arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= '0' && arr[i] <= '9')) {
			i++
			continue
		}
		if !((arr[j] >= 'a' && arr[j] <= 'z') || (arr[j] >= '0' && arr[j] <= '9')) {
			j--
			continue
		}
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	reverse := IsPalindrome(s)
	fmt.Println(reverse)
}
