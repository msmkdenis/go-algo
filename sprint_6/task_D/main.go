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
Алгоритм поиска в ширину (BFS) используется для обхода графа от заданной начальной вершины и поиска всех достижимых из неё вершин.
Он работает путем пошагового расширения области поиска, начиная с начальной вершины и постепенного просмотра её соседей.
Алгоритм работает на основе принципа "первым пришел, первым обслужен", то есть вершины посещаются в порядке их отдаленности от начальной вершины.

-- РЕАЛИЗАЦИЯ --
Функция `BFS` принимает количество вершин `n`, количество рёбер `m`, список смежности графа `graph` и стартовую вершину `start`.
Создаются массивы `visited` для отслеживания посещенных вершин и `queue` для хранения вершин, ожидающих обработки.
Алгоритм продолжается, пока очередь не пуста.
На каждом шаге извлекается вершина из начала очереди, добавляется в список обхода `traversal`, а затем её соседи добавляются в очередь для последующей обработки.
Соседи сортируются перед добавлением в очередь, чтобы обеспечить порядок возрастания при обходе.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Временная сложность алгоритма поиска в ширину: O(N + M), где N - количество вершин, M - количество рёбер в графе.
Это объясняется тем, что каждая вершина и каждое ребро посещаются ровно один раз.
Поскольку для каждой вершины выполняется поиск её соседей, а затем каждый сосед добавляется в очередь,
каждый ребро будет посещено ровно один раз, а каждая вершина будет добавлена в очередь ровно один раз.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность алгоритма: O(N + M).
Используются массивы `visited` и `queue`, каждый размером O(N), а также список смежности `graph`, который может занимать до O(M) памяти.
Массив `visited` отслеживает посещенные вершины, а массив `queue` хранит вершины, ожидающие обработки.
Список смежности `graph` содержит информацию о связях между вершинами графа.
*/

func main() {
	// Чтение данных из файла
	n, m, edges, start := readGraphFromFile("input.txt")

	// Построение графа в виде списка смежности
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Поиск в ширину
	traversal := BFS(n, m, graph, start)

	// Вывод результата
	for _, v := range traversal {
		fmt.Print(v, " ")
	}
}

func readGraphFromFile(filename string) (int, int, [][]int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Считываем количество вершин и рёбер
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	// Считываем рёбра
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		edges[i] = []int{u, v}
	}

	// Считываем стартовую вершину
	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())

	return n, m, edges, start
}

func BFS(n, m int, graph map[int][]int, start int) []int {
	visited := make([]bool, n+1)
	queue := make([]int, 0)
	queue = append(queue, start)
	visited[start] = true
	traversal := make([]int, 0)

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		traversal = append(traversal, vertex)

		sort.Ints(graph[vertex]) // Сортируем соседей в порядке возрастания
		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return traversal
}
