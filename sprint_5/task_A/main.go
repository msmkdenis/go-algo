package main

/*
	Гоша повесил на стену гирлянду в виде бинарного дерева, в узлах которого находятся лампочки.
	У каждой лампочки есть своя яркость.
	Уровень яркости лампочки соответствует числу, расположенному в узле дерева.
	Помогите Гоше найти самую яркую лампочку в гирлянде, то есть такую, у которой яркость наибольшая.
*/

var maxValue = 0

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {
	maxValue = 0
	recurse(root)

	return maxValue
}

func recurse(node *Node) {
	if node == nil {
		return
	}

	if node.value > maxValue {
		maxValue = node.value
	}

	if node.left != nil {
		recurse(node.left)
	}
	if node.right != nil {
		recurse(node.right)
	}
}
