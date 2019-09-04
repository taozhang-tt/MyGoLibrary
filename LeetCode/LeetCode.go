package LeetCode

//Q2：两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
