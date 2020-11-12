package leetcode

import (
	"fmt"
	"testing"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSumTwoNumber(t *testing.T) {
	var l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	var l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := addTwoNumbers(l1, l2)
	t.Log(res.Val)
	t.Log(res.Next.Val)
	t.Log(res.Next.Next.Val)
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	next := &result
	num := 0
	for l1 != nil || l2 != nil || num > 0 {
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		*next = &ListNode{
			Val:  num % 10,
			Next: nil,
		}
		num = num / 10
		next = &((*next).Next)
	}
	return result
}

// 递归方式
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a, b []int
	for p := l1; p != nil; p = p.Next {
		a = append(a, p.Val)
	}
	for p := l2; p != nil; p = p.Next {
		b = append(b, p.Val)
	}
	fmt.Println("a =", a)
	fmt.Println("a =", b)
	var res *ListNode
	carry := 0
	aLen, bLen := len(a), len(b)
	for i, j := 0, 0; i < aLen || j < bLen; i, j = i+1, j+1 {
		sum := carry
		if i < aLen {
			sum += a[aLen-i-1]
		}
		if j < bLen {
			sum += b[bLen-j-1]
		}
		node := ListNode{Val: sum % 10, Next: nil}
		if res == nil {
			res = &node
		} else {
			node.Next = res
			res = &node
		}
		carry = sum / 10
	}
	if carry != 0 {
		res = &ListNode{Val: carry, Next: res}
	}
	return res
}
