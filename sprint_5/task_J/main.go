package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func insert(root *Node, key int) *Node {
	if root.value <= key {
		if root.right == nil {
			root.right = &Node{value: key}
		} else {
			insert(root.right, key)
		}
		return root
	} else {
		if root.left == nil {
			root.left = &Node{value: key}
		} else {
			insert(root.left, key)
		}
		return root
	}
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		panic("WA")
	}
	if newHead.left.value != 6 {
		panic("WA")
	}
}
