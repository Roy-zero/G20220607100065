package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 开辟保护节点
	pre := &ListNode{}
	// 定义一个游标，初始状态下，指向保护节点
	ret := pre
	// 开启遍历
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	// 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向合并完的链表即可
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return ret.Next
}
