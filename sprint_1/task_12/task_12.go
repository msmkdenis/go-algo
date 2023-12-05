package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
Васе очень нравятся задачи про строки, поэтому он придумал свою.
Есть 2 строки s и t, состоящие только из строчных букв.
Строка t получена перемешиванием букв строки s и добавлением 1 буквы в случайную позицию.
Нужно найти добавленную букву.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}

	firstString := strings.Split(s[0], "")
	sort.Strings(firstString)
	secondString := strings.Split(s[1], "")
	sort.Strings(secondString)

	for i, v := range secondString {
		if i == len(secondString)-1 {
			fmt.Println(v)
			break
		}
		if v != firstString[i] {
			fmt.Println(v)
			break
		}
	}
}
