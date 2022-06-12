package pkg

import (
	"fmt"
	"sync"
	"time"
)

type RollingWindow struct {
	sync.RWMutex
	broken bool
	// 滑动窗口大小
	size int
	// 桶队列
	buckets []*Bucket
}

// NewRollingWindow 新建滑动窗口
func NewRollingWindow(
	size int,
) *RollingWindow {
	return &RollingWindow{
		size:    size,
		buckets: make([]*Bucket, 0, size),
	}
}

// AppendBucket 追加一个新桶
func (r *RollingWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}
}

// GetBucket 获取当前队列末端的桶
func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

// RecordReqResult 在桶中记录当次结果
func (r *RollingWindow) RecordReqResult(result bool) {
	r.GetBucket().Record(result)
}

// ShowAllBucket 展示当前滑动窗口的所有桶状态
func (r *RollingWindow) ShowAllBucket() {
	for _, v := range r.buckets {
		fmt.Printf("id: [%v] | total: [%d] | failed: [%d]\n", v.Timestamp, v.Total, v.Failed)
	}
}

// Launch 启动滑动窗口
func (r *RollingWindow) Launch() {
	go func() {
		for {
			r.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}
