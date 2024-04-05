package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
-- РЕАЛИЗАЦИЯ --
1. Считывание данных из файла и парсинг: O(E), где E - количество рёбер.
2. Построение карты смежности графа: O(E), так как каждое ребро обрабатывается ровно один раз.
3. Поиск компонент связности без использования рекурсии:
   - Создается пустой стек для отслеживания текущих вершин и массив посещенных вершин.
   - Для каждой вершины, которая еще не была посещена, выполняется следующее:
     - Вершина помещается в стек.
     - Пока стек не пуст:
       - Извлекается вершина из стека.
       - Помечается как посещенная.
       - Добавляются все непосещенные смежные вершины в стек.
     - Компонента связности завершена, и вершины, которые были добавлены в стек во время этого процесса, формируют ее.
   - Данный подход не использует рекурсию, поскольку он использует стек вместо вызова функций.
4. Сортировка компонент связности: O(K * log K), где K - количество компонент связности.
5. Общая временная сложность: O(E + K * log K).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
1. Дополнительная память для хранения данных о графе: O(v + E), где v - количество вершин, E - количество рёбер.
2. Дополнительная память для хранения компонент связности: O(K * max(|C|)), где K - количество компонент связности, |C| - средний размер компоненты связности.
3. Общая пространственная сложность: O(v + E + K * max(|C|)).

В данной реализации алгоритма поиска компонент связности не используется рекурсия. Вместо этого используется стек для отслеживания текущих вершин и обхода графа в глубину.
*/

func main() {
	// Считываем данные из файла
	input := getInputData()

	// Парсим количество вершин и рёбер
	parts := strings.Split(input[0], " ")
	n, _ := strconv.Atoi(parts[0])

	// Создаём карту для хранения смежности графа
	graph := make(map[int][]int)

	// Заполняем карту смежности графа
	for i := 2; i <= len(input); i++ {
		edge := strings.Split(input[i-1], " ")
		u, _ := strconv.Atoi(edge[0])
		if len(edge) > 1 {
			v, _ := strconv.Atoi(edge[1])
			graph[u] = append(graph[u], v)
			graph[v] = append(graph[v], u)
		}
	}

	// Массив для отслеживания посещенных вершин
	visited := make(map[int]bool)

	// Массив для хранения компонент связности
	var components [][]int

	// Функция для поиска компонент связности с использованием DFS
	var dfs func(start int) []int
	dfs = func(start int) []int {
		visited[start] = true
		stack := []int{start}
		component := []int{}

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			component = append(component, node)

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					stack = append(stack, neighbor)
				}
			}
		}

		return component
	}

	// Обходим все вершины графа
	for i := 1; i <= n; i++ {
		if !visited[i] {
			components = append(components, dfs(i))
		}
	}

	// Сортируем компоненты и выводим результат
	sort.Slice(components, func(i, j int) bool {
		return components[i][0] < components[j][0]
	})

	// Выводим результат
	fmt.Println(len(components))
	var answer strings.Builder
	for _, component := range components {
		slices.Sort(component)
		for _, vertex := range component {
			answer.WriteString(strconv.Itoa(vertex) + " ")
		}
		answer.WriteString("\n")
	}
	fmt.Println(answer.String())
}

// Функция для чтения данных из файла
func getInputData() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
