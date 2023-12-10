package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://contest.yandex.ru/contest/22779/problems/H/

/*
Вот какую задачу Тимофей предложил на собеседовании одному из кандидатов.
Если вы с ней ещё не сталкивались, то наверняка столкнётесь –— она довольно популярная.

Дана скобочная последовательность. Нужно определить, правильная ли она.

Будем придерживаться такого определения:

пустая строка —– правильная скобочная последовательность;
правильная скобочная последовательность, взятая в скобки одного типа, –— правильная скобочная последовательность;
правильная скобочная последовательность с приписанной слева или справа правильной скобочной последовательностью —– тоже правильная.
На вход подаётся последовательность из скобок трёх видов: [], (), {}.
Напишите функцию is_correct_bracket_seq, которая принимает на вход скобочную последовательность и возвращает True,
если последовательность правильная, а иначе False.
*/

type StackMax struct {
	stack   []string
	isEmpty bool
}

func NewStackMax() *StackMax {
	return &StackMax{
		stack:   make([]string, 0),
		isEmpty: true,
	}
}

func (s *StackMax) Push(x string) {
	s.stack = append(s.stack, x)
	s.isEmpty = false
}

func (s *StackMax) Pop() string {
	if s.isEmpty {
		return "error"
	} else {
		x := s.stack[len(s.stack)-1]
		s.stack = s.stack[:(len(s.stack) - 1)]
		s.isEmpty = len(s.stack) == 0
		return x
	}
}

func main() {
	const maxCapacity = 150 * 1024 * 1024
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

	if inputData[0] == "" {
		fmt.Println("True")
		return
	}

	brackets := strings.Split(inputData[0], "")

	bracketMap := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}

	stack := NewStackMax()

	for _, bracket := range brackets {
		if bracket == "[" || bracket == "(" || bracket == "{" {
			stack.Push(bracket)
		} else if bracket == "]" || bracket == ")" || bracket == "}" {
			previousBracket := stack.Pop()
			if previousBracket == "error" {
				fmt.Println("False")
				return
			}
			if bracketMap[bracket] != previousBracket {
				fmt.Println("False")
				return
			}
		}
	}

	if !stack.isEmpty {
		fmt.Println("False")
		return
	}

	fmt.Println("True")
	return
}
