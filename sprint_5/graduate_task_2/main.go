package main

// https://contest.yandex.ru/contest/24810/run-report/108428174/

/*
-- ПРИНЦИП РАБОТЫ --
  Для удаления вершины бинарного дерева поиска по ключу необходимо выполнить две операции:
  1. Найти вершину
  2. Корректно удалить вершину, сохраняя целостность дерева

-- РЕАЛИЗАЦИЯ --
  Сначала найдем удаляемую вершину, а также предок этой вершины.
  	- если искомое значение меньше значения вершины - рекурсивно ищем в левом поддереве
  	- если искомое значение больше значения вершины - рекурсивно ищем в правом поддереве
	- одновременно сохраняем предок и значение вершины

  После того как мы найдем вершину для удаления возможны варианты:
  	- у удаляемой вершины нет левого или правого потомка:
		если нет левого потомка - во временную вершину запишем правого потомка удаляемой вершины
		если нет правого потомка - во временную вершину запишем левого потомка удаляемой вершины
		при этом:
		если у удаляемой вершины нет родителя, значит удаляемая вершина корень, сразу вернем имеющуюся временную вершину
		если удаляемая вершина - это правый потомок родителя, запишем вместо правого потомка имеющуюся временную вершину
		если удаляемая вершина - это левый потомок родителя, запишем вместо левого потомка имеющуюся временную вершину

  	- у удаляемой вершины есть и правый и левый потомки:
		найдем вершину с наибольшим значением в левом поддереве, а также родителя такой вершины
		благодаря устройству бинарного дерева поиска значение в данной вершине:
			однозначно больше любого значения из левого поддерева относительно удаляемой вершины
			однозначно меньше любого значения из правого поддерева относительно удаляемой вершины
		заменим значение в удаляемой вершине на найденное наибольшее значение из левого поддерева
		если у найденной вершины есть левые потомки - присоединяем их к родителю такой вершины в качестве правого потомка
		если у найденной вершины нет левых потомков - обнуляем правых потомков родителя такой вершины

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Корректность обеспечивается устройством бинарного дерева поиска:
  - для каждой вершины Х бинарного дерева поиска обязательно выполняется условие:
		в левом поддереве вершины со значениями меньше Х
		в правом поддереве вершины со значениями больше Х
  - вершина в левом поддереве относительно любого Х с максимальным значением:
		всегда имеет значение менее Х
		всем имеет значение больше любого значения вершин в левом поддереве относительно Х
		никогда не имеет потомков в правом поддереве

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Алгоритм выполняется за O(H), где H - высота дерева.
  При поиске удаляемого значения мы рекурсивно спускаемся по дереву максимально на К шагов.
  При необходимости мы ищем в левом поддереве вершину с максимальным значением, также рекурсивно максимально на Т шагов.
  При этом сумма К и Т не может быть больше Н - высоты дерева.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Дополнительная память потребуется на хранения стэка вызовов, который может составить до Н.
*/

type Node struct {
	value int
	left  *Node
	right *Node
}

func remove(node *Node, key int) *Node {

	// если дерево пустое
	if node == nil {
		return nil
	}

	deletedNode, prev := findNode(node, nil, key)

	if deletedNode == nil {
		return node
	}

	// если у удаляемой вершины отсутствует левый или правый потомок
	if deletedNode.right == nil || deletedNode.left == nil {
		var newNode *Node

		if deletedNode.left == nil {
			newNode = deletedNode.right
		} else {
			newNode = deletedNode.left
		}

		// если удаляемая вершина - корень дерева
		if prev == nil {
			return newNode
		}

		// присоединяем к предыдущей вершине (относительно удаляемой)
		// выпадающее дерево
		if deletedNode == prev.left {
			prev.left = newNode
		} else {
			prev.right = newNode
		}

		deletedNode = nil
	} else {
		// если у удаляемой вершины есть оба потомка
		// находим наибольшее значение левого поддерева, а также его родитель
		leftMaxNode, prevLeftMaxNode := findLeftMaxNode(deletedNode.left, deletedNode)
		// меняем значение у удаляемой вершины
		deletedNode.value = leftMaxNode.value
		// если у вершины с наибольшим значением левого поддерева есть потомки
		// присоединяем их к правому потомку родителя этой вершины
		if leftMaxNode.left != nil {
			prevLeftMaxNode.right = leftMaxNode.left
			// иначе обнуляем ссылку на праву ветвь этой вершины
		} else {
			prevLeftMaxNode.right = nil
		}
		leftMaxNode = nil
	}

	return node
}

func findNode(root *Node, prev *Node, value int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if value < root.value {
		return findNode(root.left, root, value)
	}

	if value > root.value {
		return findNode(root.right, root, value)
	}

	var node *Node
	if value == root.value {
		node = root
	}

	return node, prev
}

func findLeftMaxNode(node, prev *Node) (*Node, *Node) {
	root := node
	for root.right != nil {
		prev = root
		root = root.right
	}
	return root, prev
}
