package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
-- ПРИНЦИП РАБОТЫ --
Алгоритм проверяет, может ли данная ориентированная сеть представлять из себя дороги между городами так,
чтобы из каждого города можно было добраться в любой другой, не проходя по одной и той же дороге дважды.
Для этого используется алгоритм поиска в глубину (DFS), который проверяет наличие циклов в графе.

-- РЕАЛИЗАЦИЯ --
1. Сначала считываются входные данные о количестве городов и графе дорог из файла input.txt.
2. Создаётся пустой граф, представленный в виде карты, где ключами являются вершины (города), а значениями - списки смежных вершин (городов).
3. Для каждой строки ввода, представляющей дороги между городами, формируется соответствующее ребро графа.
4. Затем для каждой вершины в графе запускается алгоритм поиска в глубину (DFS) с использованием стека, чтобы проверить наличие циклов.
5. Если цикл обнаружен, программа выводит "NO" и завершает работу, иначе выводится "YES".

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм поиска в глубину (DFS) позволяет обнаружить циклы в ориентированном графе.
Если обнаружен цикл, значит, существует путь из одной вершины графа в эту же вершину, что противоречит условиям задачи.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Временная сложность алгоритма DFS: O(V + E), где V - количество вершин в графе, E - количество рёбер.
Считывание входных данных и создание графа: O(n^2), где n - количество городов.
Общая временная сложность: O(n^2 + V + E).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность алгоритма: O(V + E), где V - количество вершин в графе, E - количество рёбер.
Используются структуры данных для хранения графа и цветов вершин.
*/

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	input := getInputData()
	towns, _ := strconv.Atoi(strings.Split(input[0], " ")[0])
	graph := make(map[int][]int, towns)

	for i := 1; i < len(input); i++ {
		s := strings.Split(input[i], "")
		for j, v := range s {
			from := i
			to := i + j + 1
			if v == "R" {
				from, to = to, from
			}
			if _, ok := graph[from]; !ok {
				graph[from] = make([]int, 0)
			}
			graph[from] = append(graph[from], to)
		}
	}

	color := make([]string, 1, towns+1)
	for i := 0; i < towns; i++ {
		color = append(color, "white")
	}

	for i := 1; i <= towns; i++ {
		if color[i] != "white" {
			continue
		}
		if dfs(i, color, graph) {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}

func dfs(startVertex int, color []string, graph map[int][]int) bool {
	stack := NewStack()
	stack.Push(startVertex)

	for stack.Size() > 0 {
		v, _ := stack.Pop()
		if color[v] == "white" {
			color[v] = "grey"
			stack.Push(v)
			tops := graph[v]
			for _, w := range tops {
				if color[w] == "white" {
					stack.Push(w)
				}
				if color[w] == "grey" {
					return true
				}
			}
		} else if color[v] == "grey" {
			color[v] = "black"
		}
	}
	return false
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
