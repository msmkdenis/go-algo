package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		newHead := head.next
		head.next = nil
		return newHead
	}

	previousNode := getNodeByIndex(head, idx-1)
	currentNode := previousNode.next
	nextNode := currentNode.next

	previousNode.next = nextNode
	currentNode.next = nil

	return head
}

func getNodeByIndex(node *ListNode, index int) *ListNode {
	for index > 0 {
		node = node.next
		index--
	}
	return node
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
