package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}

	ans := subdomainVisits(cpdomains)

	fmt.Printf("ans : %v \n", ans)
}

func subdomainVisits(cpdomains []string) []string {
	// 申明map，存储域名：访问次数
	count := make(map[string]int, 0)
	// 开启遍历
	for _, value := range cpdomains {
		// 按空格切割
		splits := strings.Split(value, " ")
		nums, _ := strconv.Atoi(splits[0])
		// 循环遍历域名splits[1]，域顶级域名跳出
		for {
			count[splits[1]] += nums
			// 当前域名第一个.的位置
			dotIndex := strings.Index(splits[1], ".")
			// 小于0，说明为顶级域名
			if dotIndex < 0 {
				break
			}
			// 获取更高一级域名
			splits[1] = splits[1][dotIndex+1:]
		}
	}

	ans := make([]string, 0, len(count))
	for key, value := range count {
		ans = append(ans, fmt.Sprintf("%d %s", value, key))
	}
	return ans
}