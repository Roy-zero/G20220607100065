package main

import (
	"math/rand"
	"time"

	"G20220607100065/Go-009/week05/pkg"
)

func main() {
	size := 100
	r := pkg.NewRollingWindow(size)
	r.Launch()

	// 随机生成值，计入结果
	for i := 0; i < size; i++ {
		randInt := setReqResult()
		if randInt%2 == 0 {
			r.RecordReqResult(false)
		} else {
			r.RecordReqResult(true)
		}
	}

	r.ShowAllBucket()
}

func Wrapper(size int) {
}

func setReqResult() int {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	return rand.Intn(100)
}
