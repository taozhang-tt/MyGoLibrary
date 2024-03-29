package LeetCode

var listCode = new(ListCode)

type ListCode struct {
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func (o *ListCode) PrintList(l *ListNode) {
	for {
		if l == nil {
			break
		}
		print(l.Val, ", ")
		l = l.Next
	}
	println()
}

func (o *ListCode) GenerateList(data ...int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	ret := new(ListNode)
	current := ret
	for _, v := range data{
		current.Next = new(ListNode)
		current.Next.Val = v
		current = current.Next
	}
	return ret.Next
}
//Q2：两数相加
func (o *ListCode) AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	current := ret
	p := 0
	q := 0
	carry := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			p = l1.Val
			l1 = l1.Next
		} else {
			p = 0
		}
		if l2 != nil {
			q = l2.Val
			l2 = l2.Next
		} else {
			q = 0
		}
		sum := p + q + carry
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	if carry > 0 {
		current.Next = new(ListNode)
		current.Next.Val = carry
	}
	return ret.Next
}

func (o *ListCode) AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	l3 := ret
	carry := 0
	p := 0
	q := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			p = 0
		} else {
			p = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			q = 0
		} else {
			q = l2.Val
			l2 = l2.Next
		}
		l3.Val = (p + q + carry) % 10
		carry = (p + q + carry) / 10
		if l1 != nil || l2 != nil {
			l3.Next = new(ListNode)
			l3 = l3.Next
		}
	}
	if carry > 0 {
		l3.Next = new(ListNode)
		l3.Next.Val = carry
	}
	return ret
}

//Q21 合并两个有序链表
//这是自己的思路
func (o *ListCode) MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//不要忽略边界问题
	if l1 == nil && l2 == nil {
		return nil
	}
	ret := new(ListNode)
	l3 := ret
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if (l1 != nil && l2 != nil && l1.Val < l2.Val) || l2 == nil {
			l3.Val = l1.Val
			l1 = l1.Next
		} else {
			l3.Val = l2.Val
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			l3.Next = new(ListNode)
			l3 = l3.Next
		}
	}

	return ret
}

//这是网上看到的递归思路
func (o *ListCode) MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		listCode.PrintList(l2)
		return l2
	}
	if l2 ==  nil {
		listCode.PrintList(l1)
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = listCode.MergeTwoLists2(l1.Next, l2)
		listCode.PrintList(l1)
		return l1
	} else {
		l2.Next = listCode.MergeTwoLists2(l1, l2.Next)
		listCode.PrintList(l2)
		return l2
	}
}

//Q23 合并K个排序链表
//借助 Q21 两两合并
func (o *ListCode) MergeKLists(lists []*ListNode) *ListNode {
	var ret *ListNode
	for _, item := range lists {
		ret = listCode.MergeTwoLists2(ret, item)
	}
	return ret
}

//Q83 删除排序链表中的重复元素(兼容链表是无序的）
func (o *ListCode) DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ret := head
	collection := make(map[int]int)
	collection[head.Val] = 1
	for {
		if head.Next == nil {
			break
		}
		if _, ok := collection[head.Next.Val]; ok {
			if head.Next.Next != nil {
				head.Next = head.Next.Next
			} else {
				head.Next = nil
				break
			}
		} else {
			collection[head.Next.Val] = 1
			head = head.Next
		}
	}
	return ret
}
//Q83 删除排序链表中的重复元素（要求链表是有序的）
func (o *ListCode) DeleteDuplicates2(head *ListNode) *ListNode {
	ret := head
	for {
		if head != nil && head.Next != nil {
			if head.Val == head.Next.Val {
				head.Next = head.Next.Next
			} else {
				head = head.Next
			}
		} else {
			break
		}
	}
	return ret
}

//Q82 删除排序链表中的重复元素 II
func (o *ListCode) DeleteDuplicates3(head *ListNode) *ListNode {
	pre := new(ListNode)	//一个空的头节点
	pre.Next = head
	ret := pre
	sign := false	//标志当前节点是否应该被删除
	for {
		if head != nil && head.Next != nil {
			if head.Val == head.Next.Val {
				head.Next = head.Next.Next
				sign = true
			} else {
				if sign {
					pre.Next = head.Next	//删除当前节点
					sign = false
				} else {
					pre = head
				}
				head = head.Next
			}
		} else {
			break
		}
	}
	if sign {
		pre.Next = head.Next
	}
	return ret.Next
}

//Q19  删除链表的倒数第N个节点
func (o *ListCode) RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	ret := new(ListNode)
	ret.Next = head
	pre := ret
	for {
		if head == nil {
			break
		}
		if n == 0 {
			pre = pre.Next
		} else {
			n --
		}
		head = head.Next
	}
	pre.Next = pre.Next.Next
	return ret.Next
}

//Q24 两两交换链表中的节点
func (o *ListCode) SwapPairs(head *ListNode) *ListNode {
	ret := new(ListNode)
	ret.Next = head
	pre := ret
	next := new(ListNode)
	for  {
		if head == nil {
			break
		}
		next = head.Next
		if next == nil {	//已到达最后一个单节点
			break
		}
		//交换两个节点
		pre.Next = next
		head.Next = next.Next
		next.Next = head
		if next.Next == nil {
			break
		}
		//向后移动两个节点
		pre = pre.Next.Next
		head = pre.Next
	}
	return ret.Next
}

//Q206 反转链表
//pre, head, next 一个一个反转
func (o *ListCode) ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	head = head.Next
	pre.Next = nil
	next := new(ListNode)
	for {
		if head == nil {
			break
		}
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//Q206
//递归做法
func (o *ListCode) ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := listCode.ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}