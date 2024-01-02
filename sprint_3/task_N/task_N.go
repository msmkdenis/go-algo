package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Алла захотела, чтобы у неё под окном были узкие клумбы с тюльпанам.
На схеме земельного участка клумбы обозначаются просто горизонтальными отрезками, лежащими на одной прямой.
Для ландшафтных работ было нанято n садовников. Каждый из них обрабатывал какой-то отрезок на схеме.
Процесс был организован не очень хорошо, иногда один и тот же отрезок или его часть могли быть обработаны сразу несколькими садовниками.
Таким образом, отрезки, обрабатываемые двумя разными садовниками, сливаются в один. Непрерывный обработанный отрезок затем станет клумбой.
Нужно определить границы будущих клумб.
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
	lineCount, _ := strconv.Atoi(s[0])
	var lines [][]int
	for i := 1; i <= lineCount; i++ {
		lineNumbers := makeIntSlice(s[i], 2)
		lines = append(lines, lineNumbers)
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})

	var answer [][]int
	var subLine []int
	for i := 1; i < len(lines); i++ {
		if len(subLine) == 0 {
			if lines[i][0] > lines[i-1][0] && lines[i][0] > lines[i-1][1] {
				answer = append(answer, lines[i-1])
				if i == len(lines)-1 {
					answer = append(answer, lines[i])
				}
			}
			if (lines[i][0] >= lines[i-1][0] && lines[i][0] <= lines[i-1][1]) && (lines[i][1] >= lines[i-1][1]) {
				subLine = []int{lines[i-1][0], lines[i][1]}
			}
			if (lines[i][0] >= lines[i-1][0] && lines[i][0] <= lines[i-1][1]) && (lines[i][1] <= lines[i-1][1]) {
				subLine = []int{lines[i-1][0], lines[i-1][1]}
			}
		} else {
			if lines[i][0] > subLine[1] && lines[i][1] > subLine[1] {
				answer = append(answer, subLine)
				subLine = []int{lines[i][0], lines[i][1]}
			}
			if (lines[i][0] >= subLine[0] && lines[i][0] <= subLine[1]) && (lines[i][1] >= subLine[1]) {
				subLine = []int{subLine[0], lines[i][1]}
			}
		}
	}
	if len(subLine) != 0 {
		answer = append(answer, subLine)
	}
	if len(lines) == 1 {
		answer = append(answer, lines[0])
	}
	for _, v := range answer {
		fmt.Println(v[0], v[1])
	}
}

func makeIntSlice(s string, cap int) []int {
	result := make([]int, 0, cap)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
