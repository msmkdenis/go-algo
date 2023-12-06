package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/102128756/

/*
Игра «Тренажёр для скоростной печати» представляет собой поле 4x4 из клавиш, на которых — либо точка, либо цифра от одного до девяти.
Суть игры следующая: каждый раунд на поле появляется комбинация цифр и точек.
В момент времени t игрок должен одновременно нажать на все клавиши, где есть цифра t.

Если в момент t нажаты все нужные клавиши, игроки получают один балл.
Если клавиш с цифрой t на поле нет, победное очко не начисляется

Два игрока в один момент могут нажать на k клавиш каждый.
Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём.
Рассмотрим пример 1, в котором k=3.

Допустим, t=1. В таком случае один игрок должен нажать на две кнопки с цифрой 1.
Чтобы узнать, сколько клавиш нажмут два игрока, воспользуемся формулой: k*2.
Получается, что вместе мальчики нажмут на шесть клавиш и получат победное очко.

Когда t=2, двум игрокам необходимо нажать на семь кнопок одновременно.
Но это не под силу ребятам: каждый может нажать только по три кнопки. Победное очко не начисляется.

При t=3, каждому игроку нужно нажать на одну кнопку.
Успех! Теперь на счету Гоши и Тимофея целых два победных очка.

Других цифр на поле нет. Поэтому в следующих раундах, где t=4...t=9, победные очки начисляться не будут.
Таким образом, Гоша и Тимофей заработают два балла.

Найдите число баллов, которое смогут заработать Гоша и Тимофей, если будут нажимать на клавиши вдвоём.
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
	k, _ := strconv.Atoi(s[0])

	data := makeIntSlice(s[1:], 16)
	fmt.Println(countScores(k, data))
}

func makeIntSlice(s []string, len int) []int {
	result := make([]int, 0, len)
	for _, v := range s {
		subString := strings.Split(v, "")
		for _, vs := range subString {
			if vs != "." {
				i, _ := strconv.Atoi(vs)
				result = append(result, i)
			} else {
				result = append(result, 0)
			}
		}
	}
	return result
}

func countScores(k int, data []int) int {
	scores := 0
	countNumbers := 0
	for i := 1; i <= 9; i++ {
		for j := 0; j < len(data); j++ {
			if i == data[j] {
				countNumbers++
			}
		}
		if countNumbers != 0 && countNumbers <= k*2 {
			scores++
		}
		countNumbers = 0
	}

	return scores
}
