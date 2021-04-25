package array

import (
	"fmt"
	"testing"
)

/*
买卖股票的最佳时机 II
给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: prices = [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例3:

输入: prices = [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

提示：

1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func MaxProfit(prices []int) int {
	count := 0
	pay := false
	payIndex := 0
	for i := 0; i < len(prices)-1; i++ {
		// 只需要寻找上升的区间，遇到下降，如果存在已经购买，那么就出售，否则就继续找
		if prices[i] < prices[i+1] {
			if !pay {
				pay = true
				payIndex = i
			}
			if pay && len(prices)-1 == i+1 {
				count += prices[i+1] - prices[payIndex]
				pay = false
			}
		} else if pay {
			count += prices[i] - prices[payIndex]
			pay = false
		}
	}
	return count
}

func TestMaxProfit(t *testing.T) {
	//prices := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	prices := []int{1, 1, 2, 3, 2,1}

	n := MaxProfit(prices)
	fmt.Println(n)
}