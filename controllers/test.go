package controllers

import beego "github.com/beego/beego/v2/server/web"

type TestController struct {
	beego.Controller
}

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func (c *TestController) Test() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ReverseList(&head)
}

// ReverseList 反转列表
//指定一个列表 类似 1->2->3->4->5
//输出 5->4->3->2->1
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
