package string

import (
	"fmt"
	"testing"
)

/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//st := make(map[int32]int,0)
	//for _, tt := range s {
	//	st[tt]++
	//}
	//
	//tt := make(map[int32]int,0)
	//for _, ts := range t {
	//	tt[ts]++
	//}
	//
	//if len(st) != len(tt) {
	//	return false
	//}
	//
	//for k, v := range st {
	//	if vv ,ok := tt[k];!ok || vv != v {
	//		return false
	//	}
	//}

	ss := []byte(s)
	tt := []byte(t)

	for i := 0; i < len(ss); i++ {
		j := 0
		ttl := len(tt)
		for ; j < ttl; j++ {
			if tt[j] == ss[i] {
				tt = append(tt[:j], tt[j+1:]...)
				break
			}
		}
		if j >= ttl {
			return false
		}
	}

	if len(tt) != 0 {
		return false
	}

	return true

}

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	s1 := "nagaram"
	reverse := IsAnagram(s, s1)
	fmt.Println(reverse)
}
