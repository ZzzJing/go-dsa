package linkedlist

func (list *List) IsPalindrome() bool {
	slow := list.head.next
	fast := list.head.next
	var prev *Node = nil
	var temp *Node = nil

	if list.head == nil || list.head.next == nil {
		return true
	}
	//快指针跑两步,慢指针跑一步,并反转慢指针跑过的节点
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		temp = slow.next
		slow.next = prev
		prev = slow
		slow = temp
	}

	if fast != nil {
		slow = slow.next
	}

	l1 := prev
	l2 := slow

	for l1 != nil && l2 != nil && l1.GetValue() == l2.GetValue() {
		l1 = l1.next
		l2 = l2.next
	}

	result := (l1 == nil && l2 == nil)

	//恢复链表
	var pp *Node = nil

	for prev != nil {
		pp = prev.next
		prev.next = temp
		temp = prev
		prev = pp
	}

	return result
}
