package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Гоша любит играть в игру «Подпоследовательность»: даны 2 строки, и нужно понять, является ли первая из них подпоследовательностью второй.
Когда строки достаточно длинные, очень трудно получить ответ на этот вопрос, просто посмотрев на них.
Помогите Гоше написать функцию, которая решает эту задачу.
*/

func main() {

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}

	first := []rune(s[0])
	second := []rune(s[1])
	answer := isSubSequence(first, second)
	if answer {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func isSubSequence(first, second []rune) bool {
	if len(first) == 0 {
		return true
	}
	if len(second) == 0 {
		return false
	}
	if first[0] == second[0] {
		return isSubSequence(first[1:], second[1:])
	}
	return isSubSequence(first, second[1:])
}
