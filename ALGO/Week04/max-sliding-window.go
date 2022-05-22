package main

func main()  {

}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		queue []int
		ans []int
	)
	for f := 0; f < len(nums); f++ {
		//维持队列递减, 将 k 插入合适的位置, queue中 <=k 的 元素都不可能是窗口中的最大值, 直接弹出
		for len(queue) > 0 && nums[f] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		// 等大的后来者也应入队
		if len(queue) == 0 || nums[f] <= nums[queue[len(queue)-1]] {
			queue = append(queue, f)
		}

		if f >= k - 1 {
			ans = append(ans, nums[queue[0]])
			//弹出离开窗口的队首
			if f - k + 1 == queue[0] {
				queue = queue[1:]
			}
		}
	}

	return ans
}