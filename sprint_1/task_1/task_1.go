package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вася делает тест по математике: вычисляет значение функций в различных точках.
Стоит отличная погода, и друзья зовут Васю гулять.
Но мальчик решил сначала закончить тест и только после этого идти к друзьям.
К сожалению, Вася пока не умеет программировать. Зато вы умеете.
Помогите Васе написать код функции, вычисляющей y = ax2 + bx + c.

Напишите программу, которая будет по коэффициентам a, b, c и числу x выводить значение функции в точке x.
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
	data := makeIntSlice(s[0])

	a := data[0]
	x := data[1]
	b := data[2]
	c := data[3]

	fmt.Println(a*x*x + b*x + c)
}

func makeIntSlice(s string) []int {
	var result []int
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
