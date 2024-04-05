package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25070/run-report/111324850/

/*
-- ПРИНЦИП РАБОТЫ --
  По условию задачи мы имеем взвешенный неориентированный граф.
  Существуют несколько алгоритмов для поиска min остовного дерева в указанном графе, например алгоритм Прима.
  Мы можем модифицировать алгоритм Прима для поиска max остовного дерева.

-- РЕАЛИЗАЦИЯ --
  Создадим структуру для хранения графа с использованием двумерного слайса.
  Т.к. граф неориентированный каждое ребро требуется отобразить два раза.

  Создадим очередь с приоритетом для хранения вершины окончания ребра и веса ребра в качестве приоритета.
  Создадим слайс, в котором будем хранить посещенные вершины.
  Создадим также слайс для хранения вершины (окончания) ребра и веса ребра для доп. проверки перед пушем в очередь.
  Инициализируем такой слайс минимальными значениями.

  Для алгоритма Прима неважно, с какой вершины будет начинаться обход графа, поэтому начнем с вершины 0.

  Добавим в очередь вершину 0 с приоритетом 0 (инициализация).
  Далее в цикле будем перебирать вершины приоритетной очереди, пока она не будет пуста:
    Извлечем вершину из очереди (всегда имеет max значение веса ребра - гарантия heap структуры).
    Если вершина не была посещена, то добавим её в слайс посещенных вершин.
    Обновим вес остовного дерева.
    Получим все инцидентные вершины с вершиной, извлеченной из очереди и для каждой из них проверим:
	  - Не была ли посещена вершина до этого
      - Вес ребра к этой вершине больше или меньше записанного в слайс весов ребер.
		Если оба утверждения верны: добавим в очередь вершину с приоритетом веса ребра.

  Цикл продолжается пока не переберем все вершины в очереди.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Построение графа: перебираем входные данные O(E)
  Алгоритм Прима: O((v + E) log v), где v - количество вершин, E - количество рёбер в графе

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность зависит от количества вершин и рёбер в графе:
  O(2v) для хранения двух массивов: посещенных вершин и весов ребер.
  O(E) для хранения очереди
  В общем случае потребуется ~ O(v + E) дополнительной памяти.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Мы перебираем все вершины графа, на что делается соответствующая проверка.
  При этом каждый раз мы выбираем ребро с максимальным весом в графе.
*/

type Edge struct {
	dest, weight int
}

type Graph struct {
	v     int
	edges [][]Edge
}

func NewGraph(v int) *Graph {
	graph := &Graph{v: v}
	graph.edges = make([][]Edge, v)
	return graph
}

func (g *Graph) AddEdge(src, dest, weight int) {
	g.edges[src] = append(g.edges[src], Edge{dest: dest, weight: weight})
	g.edges[dest] = append(g.edges[dest], Edge{dest: src, weight: weight}) // граф неориентированный
}

func PrimMaxST(graph *Graph) (int, bool) {
	v := graph.v
	visited := make([]bool, v)
	key := make([]int, v)
	for i := 0; i < v; i++ {
		key[i] = -(1 << 31) // инициализируем веса ребер min значениями
	}
	key[0] = 0

	maxWeight := 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: 0, priority: 0})

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.value

		if visited[u] {
			continue
		}

		visited[u] = true
		maxWeight += item.priority

		for _, edge := range graph.edges[u] {
			v := edge.dest
			wt := edge.weight
			if !visited[v] && wt > key[v] { // key[v] - вес ребра с вершины u до вершины v
				key[v] = wt
				heap.Push(&pq, &Item{value: v, priority: wt})
			}
		}
	}

	// проверка на то, что мы посетили все вершины графа
	for i := 0; i < v; i++ {
		if !visited[i] {
			return 0, false
		}
	}
	return maxWeight, true
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	input := getInputData()
	n, _ := strconv.Atoi(strings.Split(input[0], " ")[0])
	m, _ := strconv.Atoi(strings.Split(input[0], " ")[1])

	graph := NewGraph(n)

	// Заполняем граф
	for i := 0; i < m; i++ {
		parts := strings.Split(input[i+1], " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		w, _ := strconv.Atoi(parts[2])
		graph.AddEdge(u-1, v-1, w)
	}

	maxWeight, exists := PrimMaxST(graph)
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
