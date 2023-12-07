package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, elem string) int {
	idx := 0
	for head != nil {
		if head.data == elem {
			break
		}
		head = head.next
		idx++
	}

	if head == nil {
		return -1
	}

	return idx
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
