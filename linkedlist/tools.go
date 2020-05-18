package linkedlist

func (list *List) HasCircle() bool {
	if list.size != 0 {
		fast := list.head
		slow := list.head

		for fast != nil && fast.next != nil {
			fast = fast.next.next
			slow = slow.next
			if fast == slow {
				return true
			}
		}
	}

	return false
}
