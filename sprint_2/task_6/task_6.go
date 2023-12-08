package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/102337928/

/*
Нужно реализовать класс StackMax, который поддерживает операцию определения максимума среди всех элементов в стеке.
Класс должен поддерживать операции push(x), где x – целое число, pop() и get_max().

Формат ввода
В первой строке записано одно число n — количество команд, которое не превосходит 10000.
В следующих n строках идут команды. Команды могут быть следующих видов:

push(x) — добавить число x в стек. Число x не превышает 105;
pop() — удалить число с вершины стека;
get_max() — напечатать максимальное число в стеке;
Если стек пуст, при вызове команды get_max() нужно напечатать «None», для команды pop() — «error»
*/

type StackMax struct {
	stack   []int
	max     int
	isEmpty bool
}

func NewStackMax() *StackMax {
	return &StackMax{
		stack:   make([]int, 0),
		isEmpty: true,
	}
}

func (s *StackMax) Push(x int) {
	s.stack = append(s.stack, x)
	if s.isEmpty {
		s.max = x
		s.isEmpty = false
	} else {
		if x > s.max {
			s.max = x
		}
	}
}

func (s *StackMax) Pop() {
	if s.isEmpty {
		fmt.Println("error")
		return
	} else {
		s.stack = s.stack[:(len(s.stack) - 1)]
		s.isEmpty = len(s.stack) == 0
		if !s.isEmpty {
			if len(s.stack) == 1 {
				s.max = s.stack[0]
			} else {
				s.max = s.stack[0]
				for _, v := range s.stack {
					if v > s.max {
						s.max = v
					}
				}
			}

		}
	}
}

func (s *StackMax) GetMax() {
	if s.isEmpty {
		fmt.Println("None")
	} else {
		fmt.Println(s.max)
	}
}

func main() {
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var inputData []string
	for scanner.Scan() {
		line := scanner.Text()
		inputData = append(inputData, line)
		if line == "" {
			break
		}
	}
	stack := NewStackMax()
	processCommands(inputData[1:], stack)
}

func processCommands(inputData []string, stack *StackMax) {
	for _, v := range inputData {
		command := strings.Split(v, " ")
		if command[0] == "push" {
			x, _ := strconv.Atoi(command[1])
			stack.Push(x)
		}
		if command[0] == "pop" {
			stack.Pop()
		}
		if command[0] == "get_max" {
			stack.GetMax()
		}
	}
}
