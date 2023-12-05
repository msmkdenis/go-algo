package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Метеорологическая служба вашего города решила исследовать погоду новым способом.

Под температурой воздуха в конкретный день будем понимать максимальную температуру в этот день.
Под хаотичностью погоды за n дней служба понимает количество дней, в которые температура строго больше, чем в день до (если такой существует) и в день после текущего (если такой существует).
Например, если за 5 дней максимальная температура воздуха составляла [1, 2, 5, 4, 8] градусов, то хаотичность за этот период равна 2: в 3-й и 5-й дни выполнялись описанные условия.
Определите по ежедневным показаниям температуры хаотичность погоды за этот период.

Заметим, что если число показаний n=1, то единственный день будет хаотичным.
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

	count, _ := strconv.Atoi(s[0])
	days := makeIntSlice(s[1], count)
	fmt.Println(countDays(days))
}

func makeIntSlice(s string, len int) []int {
	result := make([]int, 0, len)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func countDays(days []int) int {
	if len(days) == 1 {
		return 1
	}

	if len(days) == 2 {
		if days[1] > days[0] || days[0] > days[1] {
			return 1
		} else {
			return 0
		}
	}

	var counter int
	for i := 1; i < len(days)-1; i++ {
		if days[i] > days[i-1] && days[i] > days[i+1] {
			counter++
		}
	}
	if days[0] > days[1] {
		counter++
	}
	if days[len(days)-1] > days[len(days)-2] {
		counter++
	}
	return counter
}
