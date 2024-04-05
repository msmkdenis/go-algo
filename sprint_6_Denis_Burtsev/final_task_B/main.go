package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25070/run-report/111327952/

/*
-- ПРИНЦИП РАБОТЫ --
  Исходя из условий задачи мы можем представить карту железных дорог в виде ориентированного графа.
  При этом дороги вида R и B можно представить, как разнонаправленные.
  В таком случае задача сведется к двум подзадачам:
  - корректное построение графа
  - подтверждение наличия или отсутствия циклов в графе
  Поиск цикла будем осуществлять с помощью обхода графа в глубину (DFS алгоритма).

-- РЕАЛИЗАЦИЯ --
  Создадим граф в виде map, где ключами являются вершины (города), а значениями - списки смежных вершин (городов).
  Обработаем входящие данные с учетом условия выше: R и B дороги являются разнонаправленными ребрами.
  Затем для каждой вершины в графе запускается алгоритм поиска в глубину (DFS) с использованием стека, чтобы проверить наличие циклов.
  По общему условию если мы встречаем второй раз серую вершину - значит мы обнаружили цикл.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Алгоритм поиска в глубину (DFS) позволяет обнаружить циклы в ориентированном графе.
  При заходе в вершину мы красим ее в серый цвет.
  Дальше мы определяем инцидентные вершины для текущей вершины.
  При этом вершины не должны быть серыми, т.к. это означает что мы встретили родительскую вершину - встретили цикл.
  При выходе из вершины мы красим ее в черный.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма DFS: O(v + E), где v - количество вершин в графе, E - количество рёбер.
  Количество вершин = количество городов, то есть V = N
  Количество ребер = количество дорог между городами.
  По условию задачи из первого города есть дорога во все остальные, из второго - во все, кроме первого, из третьего - во все кроме первого и второго
  Это арифметическая прогрессия с разницей -1
  Сумма этой арифметической прогрессии - общее количество дорог.
  Общее кол-во дорог между городами (ребер в графе) S = (a1+aN)/2*N = (N-1+0)/2*N = (N^2-N)/2 - квадратичная зависимость от вершин
  Для плотных графов, в которых число рёбер квадратично зависит от числа вершин сложность DFS составит O(v^2)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность эквивалентна временной и составит О(v^2)
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
