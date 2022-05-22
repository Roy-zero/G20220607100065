package main

import "fmt"

func main() {
	digits := []int{9, 9, 9}
	ret := plusOne(digits)
	fmt.Printf("ret: %v \n", ret)
}

/*
 * 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一
 * digits = [1,2,3] 输出[1,2,4]
 * digits = [1,2,8,9] 输出[1,2,9,0]
 */
func plusOne(digits []int) []int {
	n := len(digits)
	/*
	 * 从最后开始便利，查看倒数第几位元素不等于9
	 * 等于9，赋值为0
	 * 不等于9，则+1
	 */
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	// 所有元素都为9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
