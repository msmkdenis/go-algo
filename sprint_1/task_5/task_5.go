package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Чтобы подготовиться к семинару, Гоше надо прочитать статью по эффективному менеджменту.
Так как Гоша хочет спланировать день заранее, ему необходимо оценить сложность статьи.

Он придумал такой метод оценки: берётся случайное предложение из текста и в нём ищется самое длинное слово. Его длина и будет условной сложностью статьи.

Помогите Гоше справиться с этой задачей.
*/

func main() {
	const maxCapacity = 1024 * 1024
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

	words := strings.Split(s[1], " ")

	var index int
	for i, v := range words {
		if len(v) > len(words[index]) {
			index = i
		}
	}

	fmt.Println(words[index])
	fmt.Println(len(words[index]))
}
