package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
-- ПРИНЦИП РАБОТЫ --
Алгоритм Крускала используется для поиска минимального остовного дерева в связном взвешенном графе.
Однако в данном случае мы модифицируем его для поиска максимального остовного дерева.
Принцип работы алгоритма состоит в том, чтобы постепенно добавлять рёбра с наибольшим весом в остовное дерево, при этом проверяя, не создаются ли циклы.

-- РЕАЛИЗАЦИЯ --
1. Сначала рёбра графа сортируются по убыванию веса.
2. Затем создаётся массив parent, который будет использоваться для представления множеств вершин.
3. Для каждого ребра в порядке убывания веса проверяется, принадлежат ли его концы разным множествам.
Если да, то ребро добавляется к остовному дереву, и множества концов ребра объединяются.
Если концы ребра уже принадлежат одному множеству, оно игнорируется, чтобы избежать образования циклов в дереве.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм Крускала гарантирует, что на каждом шаге добавляется ребро с наибольшим весом, которое не создаёт цикл в остовном дереве.
Это обеспечивает корректность алгоритма и гарантирует получение максимального остовного дерева.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Временная сложность алгоритма Крускала:
- Сортировка рёбер: O(E * log E)
- Обход рёбер и проверка наличия циклов: O(E * log V) (где V - число вершин)
Общая временная сложность: O(E * log E)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность алгоритма: O(V + E), где V - число вершин, E - число рёбер.
Необходимо хранить список рёбер графа и массив parent.
*/

type Edge struct {
	src, dest, weight int
}

type Graph struct {
	V, E int
	edge []Edge
}

func NewGraph(V, E int) *Graph {
	graph := &Graph{V: V, E: E}
	graph.edge = make([]Edge, E)
	return graph
}

func (g *Graph) AddEdge(src, dest, weight int) {
	g.edge = append(g.edge, Edge{src: src, dest: dest, weight: weight})
}

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func Union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	parent[xSet] = ySet
}

func MaxSpanningTreeWeight(graph *Graph) (int, bool) {
	V := graph.V
	E := graph.E

	sort.Slice(graph.edge, func(i, j int) bool {
		return graph.edge[i].weight > graph.edge[j].weight
	})

	parent := make([]int, V)
	for i := range parent {
		parent[i] = -1
	}

	maxWeight := 0
	edgeCount := 0
	for i := 0; i < E; i++ {
		x := find(parent, graph.edge[i].src)
		y := find(parent, graph.edge[i].dest)

		if x != y {
			Union(parent, x, y)
			maxWeight += graph.edge[i].weight
			edgeCount++
		}
	}

	if edgeCount == V-1 {
		return maxWeight, true
	}

	return 0, false
}

func main() {

	input := getInputData()
	n, _ := strconv.Atoi(strings.Split(input[0], " ")[0])
	m, _ := strconv.Atoi(strings.Split(input[0], " ")[1])

	graph := NewGraph(n, m)

	// Чтение рёбер и их весов
	for i := 0; i < m; i++ {
		parts := strings.Split(input[i+1], " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		w, _ := strconv.Atoi(parts[2])
		graph.AddEdge(u-1, v-1, w)
	}

	maxWeight, exists := MaxSpanningTreeWeight(graph)
	if exists {
		fmt.Println(maxWeight)
	} else {
		fmt.Println("Oops! I did it again")
	}
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	const maxCapacity = 10240 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
