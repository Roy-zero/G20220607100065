package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		// 链表为空或长度为0，直接返回空指针
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

/**
 * 辅助函数，在还没有达到最基本情况前，不断递归调用自己
 * 输入是链表数组，和当前要处理的在链表中开始和结束的下标
 * 输出是合并后的链表
 */
func merge(lists []*ListNode, start, end int) *ListNode {
	// 开始下标等于结束下标时
	if start == end {
		// 说明当前要处理的只有一个链表，直接返回它即可
		return lists[start]
	}
	// 开始下标大于结束下标时
	if start > end {
		// 无效，直接返回空指针nil
		return nil
	}
	// 分而治之；先找到当前子数组的中间下标
	mid := start + (end-start)/2
	// 然后分别递归处理前一半和后一半链表
	left := merge(lists, start, mid)
	right := merge(lists, mid+1, end)
	// 得到的结果是两条合并后的有序链表; 最后再把这两条链表也合并即可
	return mergeTwoSortedLists(left, right)
}

/**
 * 辅助函数，用于合并两个有序链表
 */
func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	// 连上l1剩余的链,连上l2剩余的链
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}
