package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	v, _ := strconv.Atoi(strings.Split(input[0], " ")[0])
	graph := makeGraph(input, v)

	color := make([]string, 1, v)
	for i := 0; i < v; i++ {
		color = append(color, "white")
	}
	entry := make([]int, v+1)
	leave := make([]int, v+1)
	dfs(1, color, graph, entry, leave)
}

func makeGraph(input []string, count int) map[int][]int {
	graph := make(map[int][]int, count)
	for i := 1; i <= count; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 1; i < len(input); i++ {
		fields := strings.Fields(input[i])
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		graph[v1] = append(graph[v1], v2)
	}
	return graph
}

func dfs(startVertex int, color []string, graph map[int][]int, entry, leave []int) string {
	stack := NewStack()
	stack.Push(startVertex)
	time := -1

	var visited []int
	for stack.Size() > 0 {
		v, _ := stack.Pop()
		if color[v] == "white" {
			time = time + 1
			entry[v] = time
			color[v] = "grey"
			visited = append(visited, v)
			stack.Push(v)
			tops := graph[v]
			sort.Slice(tops, func(i, j int) bool { return tops[i] > tops[j] })
			for _, w := range tops {
				if color[w] == "white" {
					stack.Push(w)
				}
			}
		} else if color[v] == "grey" {
			time = time + 1
			leave[v] = time
			color[v] = "black"
		}
	}
	for i := 1; i < len(leave); i++ {
		fmt.Println(entry[i], leave[i])
	}
	return strings.Join(strings.Fields(fmt.Sprint(visited)), " ")
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
