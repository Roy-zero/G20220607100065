package main

import "fmt"

func main()  {
	nums := []int{1,2,2,3,1,4,2}
	ans := findShortestSubArray(nums)

	fmt.Printf("ans : %v \n", ans)
}

type Appear struct {
	AppearTime int  // 第一次出现的位置
	FirstIndex int   // 最后一次出现的位置
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*Appear)
	// 初始化数组的度
	maxAppearTime := 0
	minLength := len(nums)
	// 开启遍历
	for i, x := range nums {
		t, ok := m[x]
		// 判断是否第一次出现
		if !ok {
			t = &Appear{FirstIndex: i}
		}
		t.AppearTime++

		m[x] = t

		if maxAppearTime < t.AppearTime {
			// 判断是否超过当前最大次数
			maxAppearTime = t.AppearTime
			minLength = i - t.FirstIndex + 1
		} else if maxAppearTime == t.AppearTime && minLength > i-t.FirstIndex + 1 {
			// 与某元素的出现频数相同，取最小
			minLength = i - t.FirstIndex + 1
		}
	}
	return minLength
}