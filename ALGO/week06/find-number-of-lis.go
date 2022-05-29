package main

func main() {

}

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence
func findNumberOfLIS(nums []int) int {
	type info struct {
		maxLen int //最大递增子序列长度
		cnt    int //最大递增子序列个数
	}
	var dp [2000]info
	for i := 0; i < len(nums); i++ { //初始化，长度为1，个数为1
		dp[i] = info{1, 1}
	}

	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] { //当第i数大于第j数的值时
				if 1+dp[j].maxLen > dp[i].maxLen { //当第j数的最大递增子序列长度+1 大于 当前第i数的最大长度
					dp[i].maxLen = 1 + dp[j].maxLen //更新第i数的最大长度
					dp[i].cnt = dp[j].cnt           //更新第i数的最大长度个数
				} else if 1+dp[j].maxLen == dp[i].maxLen { //当第j数的最大递增子序列长度+1 等于 当前第i数的最大长度
					dp[i].cnt += dp[j].cnt //则累加个数
				}
			}
		}
	}

	maxLen, cnt := 0, 0 //数组內最大递增子序列长度和个数
	for i := 0; i < len(nums); i++ {
		if dp[i].maxLen > maxLen { //大于
			maxLen = dp[i].maxLen //更新长度
			cnt = dp[i].cnt       //更新个数
		} else if dp[i].maxLen == maxLen { //等于
			cnt += dp[i].cnt //累加个数
		}
	}
	return cnt
}
